// Package fga includes client libraries to interact with openfga authorization
package fga

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	openfga "github.com/openfga/go-sdk"
	ofgaclient "github.com/openfga/go-sdk/client"
	"github.com/openfga/go-sdk/credentials"
	language "github.com/openfga/language/pkg/go/transformer"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/datumforge/datum/internal/httpserve/config"
)

const (
	// TODO: allow this to be configurable
	storeModelFile = "fga/model/datum.fga"
)

// Client is an event bus client with some configuration
type Client struct {
	// Ofga is the openFGA client
	Ofga ofgaclient.SdkClient
	// Config is the client configuration
	Config ofgaclient.ClientConfiguration
	// Logger is the provided Logger
	Logger *zap.SugaredLogger
}

// Config configures the openFGA setup
type Config struct {
	// config contains the base authz settings
	config config.Authz
	// logger contains the zap logger
	logger *zap.SugaredLogger
}

// NewAuthzConfig returns a new authorization configuration
func NewAuthzConfig(c config.Authz, l *zap.SugaredLogger) *Config {
	return &Config{
		config: c,
		logger: l,
	}
}

// Option is a functional configuration option for openFGA client
type Option func(c *Client)

// NewClient returns a wrapped OpenFGA API client ensuring all calls are made
// to the provided authorization model (id) and returns what is necessary.
func NewClient(host string, opts ...Option) (*Client, error) {
	if host == "" {
		return nil, ErrFGAMissingHost
	}

	// The api host is the only required field when setting up a new FGA client connection
	client := Client{
		Config: ofgaclient.ClientConfiguration{
			ApiHost: host,
		},
	}

	for _, opt := range opts {
		opt(&client)
	}

	fgaClient, err := ofgaclient.NewSdkClient(&client.Config)
	if err != nil {
		return nil, err
	}

	client.Ofga = fgaClient

	return &client, err
}

func (c *Client) GetModelID() string {
	return *c.Config.AuthorizationModelId
}

// WithScheme sets the open fga scheme, defaults to "https"
func WithScheme(scheme string) Option {
	return func(c *Client) {
		c.Config.ApiScheme = scheme
	}
}

// WithStoreID sets the store IDs, not needed when calling `CreateStore` or `ListStores`
func WithStoreID(storeID string) Option {
	return func(c *Client) {
		c.Config.StoreId = storeID
	}
}

// WithAuthorizationModelID sets the authorization model ID
func WithAuthorizationModelID(authModelID string) Option {
	return func(c *Client) {
		c.Config.AuthorizationModelId = &authModelID
	}
}

// WithToken sets the client credentials
func WithToken(token string) Option {
	return func(c *Client) {
		c.Config.Credentials = &credentials.Credentials{
			Method: credentials.CredentialsMethodApiToken,
			Config: &credentials.Config{
				ApiToken: token,
			},
		}
	}
}

// WithLogger sets logger
func WithLogger(l *zap.SugaredLogger) Option {
	return func(c *Client) {
		c.Logger = l
	}
}

// CreateFGAClientWithStore returns a Client with a store and model configured
func CreateFGAClientWithStore(ctx context.Context, c Config) (*Client, error) {
	c.logger.Infow("setting up fga client", "host", c.config.Host, "scheme", c.config.Scheme)

	// create store if an ID was not configured
	if c.config.StoreID == "" {
		// Create new store
		fgaClient, err := NewClient(
			c.config.Host,
			WithScheme(c.config.Scheme),
			WithLogger(c.logger),
		)
		if err != nil {
			return nil, err
		}

		c.config.StoreID, err = fgaClient.CreateStore(ctx, c.config.StoreName)
		if err != nil {
			return nil, err
		}
	}

	// create model if ID was not configured
	if c.config.ModelID == "" {
		// create fga client with store ID
		fgaClient, err := NewClient(
			c.config.Host,
			WithScheme(c.config.Scheme),
			WithStoreID(c.config.StoreID),
			WithLogger(c.logger),
		)
		if err != nil {
			return nil, err
		}

		// Create model if one does not already exist
		modelID, err := fgaClient.CreateModel(ctx, storeModelFile, c.config.CreateNewModel)
		if err != nil {
			return nil, err
		}

		// Set ModelID in the config
		c.config.ModelID = modelID
	}

	// create fga client with store ID
	return NewClient(
		c.config.Host,
		WithScheme(c.config.Scheme),
		WithStoreID(c.config.StoreID),
		WithAuthorizationModelID(c.config.ModelID),
		WithLogger(c.logger),
	)
}

// CreateStore creates a new fine grained authorization store and returns the store ID
func (c *Client) CreateStore(ctx context.Context, storeName string) (string, error) {
	options := ofgaclient.ClientListStoresOptions{
		ContinuationToken: openfga.PtrString(""),
	}

	stores, err := c.Ofga.ListStores(context.Background()).Options(options).Execute()
	if err != nil {
		return "", err
	}

	// Only create a new test store if one does not exist
	if len(stores.GetStores()) > 0 {
		storeID := *stores.GetStores()[0].Id
		c.Logger.Infow("fga store exists", "store_id", storeID)

		return storeID, nil
	}

	// Create new store
	storeReq := c.Ofga.CreateStore(context.Background())

	resp, err := storeReq.Body(ofgaclient.ClientCreateStoreRequest{
		Name: storeName,
	}).Execute()
	if err != nil {
		return "", err
	}

	storeID := resp.GetId()

	c.Logger.Infow("fga store created", "store_id", storeID)

	return storeID, nil
}

// CreateModel creates a new fine grained authorization model and returns the model ID
func (c *Client) CreateModel(ctx context.Context, fn string, forceCreate bool) (string, error) {
	options := ofgaclient.ClientReadAuthorizationModelsOptions{}

	models, err := c.Ofga.ReadAuthorizationModels(context.Background()).Options(options).Execute()
	if err != nil {
		return "", err
	}

	// Only create a new test model if one does not exist and we aren't forcing a new model to be created
	if !forceCreate {
		if len(*models.AuthorizationModels) > 0 {
			modelID := *models.GetAuthorizationModels()[0].Id
			c.Logger.Infow("fga model exists", "model_id", modelID)

			return modelID, nil
		}
	}

	// Create new model
	dsl, err := os.ReadFile(fn)
	if err != nil {
		return "", err
	}

	// convert to json
	dslJSON, err := dslToJSON(dsl)
	if err != nil {
		return "", err
	}

	var body ofgaclient.ClientWriteAuthorizationModelRequest
	if err := json.Unmarshal(dslJSON, &body); err != nil {
		return "", err
	}

	resp, err := c.Ofga.WriteAuthorizationModel(ctx).Body(body).Execute()
	if err != nil {
		fmt.Println("here 1")
		return "", err
	}

	modelID := resp.GetAuthorizationModelId()

	c.Logger.Infow("fga model created", "model_id", modelID)

	return modelID, nil
}

// dslToJSON converts fga model to JSON
func dslToJSON(dslString []byte) ([]byte, error) {
	parsedAuthModel, err := language.TransformDSLToProto(string(dslString))
	if err != nil {
		return []byte{}, errors.Wrap(err, ErrFailedToTransformModel.Error())
	}

	return protojson.Marshal(parsedAuthModel)
}

// Healthcheck reads the model to check if the connection is working
func Healthcheck(client Client) func(ctx context.Context) error {
	return func(ctx context.Context) error {
		opts := ofgaclient.ClientReadAuthorizationModelOptions{
			AuthorizationModelId: client.Config.AuthorizationModelId,
		}

		_, err := client.Ofga.ReadAuthorizationModel(ctx).Options(opts).Execute()
		if err != nil {
			return err
		}

		return nil
	}
}
