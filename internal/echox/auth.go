package echox

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"

	ent "github.com/datumforge/datum/internal/ent/generated"
	entUser "github.com/datumforge/datum/internal/ent/generated/user"
	database "github.com/datumforge/datum/internal/entdb"
)

type Service struct {
	Store *database.Database
	OAuth struct {
		Provider *oidc.Provider
		Config   oauth2.Config
	}
}

func NewService(store *database.Database) *Service {
	oAuthEnabled := viper.GetBool("oauth_enabled")
	if oAuthEnabled {
		// Fetch environment variables
		providerURL := os.Getenv("OAUTH_PROVIDER_URL")
		oauthClientID := os.Getenv("OAUTH_CLIENT_ID")
		oauthClientSecret := os.Getenv("OAUTH_CLIENT_SECRET")
		oauthRedirectURL := os.Getenv("OAUTH_REDIRECT_URL")
		if providerURL == "" || oauthClientID == "" || oauthClientSecret == "" || oauthRedirectURL == "" {
			log.Fatal().Msg("missing environment variables for oauth authentication")
		}
		provider, err := oidc.NewProvider(context.Background(), providerURL)
		if err != nil {
			log.Fatal().Err(err).Msg("error creating oauth provider")
		}

		config := oauth2.Config{
			ClientID:     oauthClientID,
			ClientSecret: oauthClientSecret,
			RedirectURL:  oauthRedirectURL,
			Endpoint:     provider.Endpoint(),
			Scopes:       []string{oidc.ScopeOpenID, "profile"},
		}

		err = FetchJWKS()
		if err != nil {
			log.Fatal().Err(err).Msg("error fetching jwks")
		}

		return &Service{
			Store: store,
			OAuth: struct {
				Provider *oidc.Provider
				Config   oauth2.Config
			}{
				Provider: provider,
				Config:   config,
			},
		}
	} else {
		return &Service{Store: store}
	}
}

type ChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func (s *Service) Register(c echo.Context, user user.User) (*ent.User, error) {
	// hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.hashedPassword), 14)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %v", err)
	}

	u, err := s.Store.Client.User.Create().SetUsername(user.Username).SetPassword(string(hashedPassword)).Save(c.Request().Context())
	if err != nil {
		if _, ok := err.(*ent.ConstraintError); ok {
			return nil, fmt.Errorf("user already exists")
		}
		return nil, fmt.Errorf("error creating user: %v", err)
	}
	return u, nil
}

func (s *Service) Login(c echo.Context, uDto user.User) (*ent.User, error) {
	u, err := s.Store.Client.User.Query().Where(entUser.Username(uDto.Username)).Only(c.Request().Context())
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(uDto.Password))
	if err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	uDto = user.User{
		ID:       u.ID,
		Username: u.Username,
		Role:     u.Role,
	}

	// Generate JWT and set cookie
	err = GenerateTokensAndSetCookies(&uDto, c)
	if err != nil {
		return nil, fmt.Errorf("error generating tokens: %v", err)
	}

	return u, nil
}

func (s *Service) Refresh(c echo.Context, refreshToken string) error {

	tkn, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetJWTRefreshSecret()), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return fmt.Errorf("invalid refresh token")
		}
		return fmt.Errorf("error parsing refresh token: %v", err)
	}

	if claims, ok := tkn.Claims.(jwt.MapClaims); ok && tkn.Valid {
		uID := claims["user_id"].(string)
		if err != nil {
			return fmt.Errorf("error parsing user id: %v", err)
		}
		u, err := s.Store.Client.User.Query().Where(entUser.ID(uID)).Only(c.Request().Context())
		if err != nil {
			return fmt.Errorf("error getting user: %v", err)
		}

		// Generate JWT and set cookie
		err = GenerateTokensAndSetCookies(&user.User{ID: u.ID, Username: u.Username, Role: u.Role}, c)
		if err != nil {
			return fmt.Errorf("error generating tokens: %v", err)
		}
		return nil
	}

	return err
}

func (s *Service) Me(c *CustomContext) (*ent.User, error) {
	return c.User, nil
}

func (s *Service) ChangePassword(c *CustomContext, passwordDto ChangePassword) error {
	u, err := s.Store.Client.User.Query().Where(entUser.ID(c.User.ID)).Only(c.Request().Context())
	if err != nil {
		return fmt.Errorf("error getting user: %v", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(passwordDto.OldPassword))
	if err != nil {
		return fmt.Errorf("invalid credentials")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(passwordDto.NewPassword), 14)
	if err != nil {
		return fmt.Errorf("error hashing password: %v", err)
	}

	_, err = s.Store.Client.User.Update().Where(entUser.ID(c.User.ID)).SetPassword(string(hashedPassword)).Save(c.Request().Context())
	if err != nil {
		return fmt.Errorf("error changing password: %v", err)
	}

	return nil
}
