// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/user"
	"github.com/datumforge/datum/internal/ent/generated/webauthn"
)

// Webauthn is the model entity for the Webauthn schema.
type Webauthn struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// CreatedBy holds the value of the "created_by" field.
	CreatedBy string `json:"created_by,omitempty"`
	// UpdatedBy holds the value of the "updated_by" field.
	UpdatedBy string `json:"updated_by,omitempty"`
	// MappingID holds the value of the "mapping_id" field.
	MappingID string `json:"mapping_id,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID string `json:"owner_id,omitempty"`
	// A probabilistically-unique byte sequence identifying a public key credential source and its authentication assertions
	CredentialID []byte `json:"credential_id,omitempty"`
	// The public key portion of a Relying Party-specific credential key pair, generated by an authenticator and returned to a Relying Party at registration time
	PublicKey []byte `json:"public_key,omitempty"`
	// The attestation format used (if any) by the authenticator when creating the credential
	AttestationType string `json:"attestation_type,omitempty"`
	// The AAGUID of the authenticator; AAGUID is defined as an array containing the globally unique identifier of the authenticator model being sought
	Aaguid []byte `json:"aaguid,omitempty"`
	// SignCount -Upon a new login operation, the Relying Party compares the stored signature counter value with the new signCount value returned in the assertions authenticator data. If this new signCount value is less than or equal to the stored value, a cloned authenticator may exist, or the authenticator may be malfunctioning
	SignCount int32 `json:"sign_count,omitempty"`
	// transport
	Transports []string `json:"transports,omitempty"`
	// Flag backup eligible indicates the credential is able to be backed up and/or sync'd between devices. This should NEVER change
	BackupEligible bool `json:"backup_eligible,omitempty"`
	// Flag backup state indicates the credential has been backed up and/or sync'd
	BackupState bool `json:"backup_state,omitempty"`
	// Flag user present indicates the users presence
	UserPresent bool `json:"user_present,omitempty"`
	// Flag user verified indicates the user performed verification
	UserVerified bool `json:"user_verified,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WebauthnQuery when eager-loading is set.
	Edges        WebauthnEdges `json:"edges"`
	selectValues sql.SelectValues
}

// WebauthnEdges holds the relations/edges for other nodes in the graph.
type WebauthnEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WebauthnEdges) OwnerOrErr() (*User, error) {
	if e.Owner != nil {
		return e.Owner, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Webauthn) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case webauthn.FieldTags, webauthn.FieldCredentialID, webauthn.FieldPublicKey, webauthn.FieldAaguid, webauthn.FieldTransports:
			values[i] = new([]byte)
		case webauthn.FieldBackupEligible, webauthn.FieldBackupState, webauthn.FieldUserPresent, webauthn.FieldUserVerified:
			values[i] = new(sql.NullBool)
		case webauthn.FieldSignCount:
			values[i] = new(sql.NullInt64)
		case webauthn.FieldID, webauthn.FieldCreatedBy, webauthn.FieldUpdatedBy, webauthn.FieldMappingID, webauthn.FieldOwnerID, webauthn.FieldAttestationType:
			values[i] = new(sql.NullString)
		case webauthn.FieldCreatedAt, webauthn.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Webauthn fields.
func (w *Webauthn) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case webauthn.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				w.ID = value.String
			}
		case webauthn.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				w.CreatedAt = value.Time
			}
		case webauthn.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				w.UpdatedAt = value.Time
			}
		case webauthn.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				w.CreatedBy = value.String
			}
		case webauthn.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				w.UpdatedBy = value.String
			}
		case webauthn.FieldMappingID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mapping_id", values[i])
			} else if value.Valid {
				w.MappingID = value.String
			}
		case webauthn.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case webauthn.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				w.OwnerID = value.String
			}
		case webauthn.FieldCredentialID:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field credential_id", values[i])
			} else if value != nil {
				w.CredentialID = *value
			}
		case webauthn.FieldPublicKey:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field public_key", values[i])
			} else if value != nil {
				w.PublicKey = *value
			}
		case webauthn.FieldAttestationType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field attestation_type", values[i])
			} else if value.Valid {
				w.AttestationType = value.String
			}
		case webauthn.FieldAaguid:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field aaguid", values[i])
			} else if value != nil {
				w.Aaguid = *value
			}
		case webauthn.FieldSignCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sign_count", values[i])
			} else if value.Valid {
				w.SignCount = int32(value.Int64)
			}
		case webauthn.FieldTransports:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field transports", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &w.Transports); err != nil {
					return fmt.Errorf("unmarshal field transports: %w", err)
				}
			}
		case webauthn.FieldBackupEligible:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field backup_eligible", values[i])
			} else if value.Valid {
				w.BackupEligible = value.Bool
			}
		case webauthn.FieldBackupState:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field backup_state", values[i])
			} else if value.Valid {
				w.BackupState = value.Bool
			}
		case webauthn.FieldUserPresent:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field user_present", values[i])
			} else if value.Valid {
				w.UserPresent = value.Bool
			}
		case webauthn.FieldUserVerified:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field user_verified", values[i])
			} else if value.Valid {
				w.UserVerified = value.Bool
			}
		default:
			w.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Webauthn.
// This includes values selected through modifiers, order, etc.
func (w *Webauthn) Value(name string) (ent.Value, error) {
	return w.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Webauthn entity.
func (w *Webauthn) QueryOwner() *UserQuery {
	return NewWebauthnClient(w.config).QueryOwner(w)
}

// Update returns a builder for updating this Webauthn.
// Note that you need to call Webauthn.Unwrap() before calling this method if this Webauthn
// was returned from a transaction, and the transaction was committed or rolled back.
func (w *Webauthn) Update() *WebauthnUpdateOne {
	return NewWebauthnClient(w.config).UpdateOne(w)
}

// Unwrap unwraps the Webauthn entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (w *Webauthn) Unwrap() *Webauthn {
	_tx, ok := w.config.driver.(*txDriver)
	if !ok {
		panic("generated: Webauthn is not a transactional entity")
	}
	w.config.driver = _tx.drv
	return w
}

// String implements the fmt.Stringer.
func (w *Webauthn) String() string {
	var builder strings.Builder
	builder.WriteString("Webauthn(")
	builder.WriteString(fmt.Sprintf("id=%v, ", w.ID))
	builder.WriteString("created_at=")
	builder.WriteString(w.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(w.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(w.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(w.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("mapping_id=")
	builder.WriteString(w.MappingID)
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", w.Tags))
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(w.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("credential_id=")
	builder.WriteString(fmt.Sprintf("%v", w.CredentialID))
	builder.WriteString(", ")
	builder.WriteString("public_key=")
	builder.WriteString(fmt.Sprintf("%v", w.PublicKey))
	builder.WriteString(", ")
	builder.WriteString("attestation_type=")
	builder.WriteString(w.AttestationType)
	builder.WriteString(", ")
	builder.WriteString("aaguid=")
	builder.WriteString(fmt.Sprintf("%v", w.Aaguid))
	builder.WriteString(", ")
	builder.WriteString("sign_count=")
	builder.WriteString(fmt.Sprintf("%v", w.SignCount))
	builder.WriteString(", ")
	builder.WriteString("transports=")
	builder.WriteString(fmt.Sprintf("%v", w.Transports))
	builder.WriteString(", ")
	builder.WriteString("backup_eligible=")
	builder.WriteString(fmt.Sprintf("%v", w.BackupEligible))
	builder.WriteString(", ")
	builder.WriteString("backup_state=")
	builder.WriteString(fmt.Sprintf("%v", w.BackupState))
	builder.WriteString(", ")
	builder.WriteString("user_present=")
	builder.WriteString(fmt.Sprintf("%v", w.UserPresent))
	builder.WriteString(", ")
	builder.WriteString("user_verified=")
	builder.WriteString(fmt.Sprintf("%v", w.UserVerified))
	builder.WriteByte(')')
	return builder.String()
}

// Webauthns is a parsable slice of Webauthn.
type Webauthns []*Webauthn
