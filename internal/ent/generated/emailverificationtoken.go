// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/emailverificationtoken"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// EmailVerificationToken is the model entity for the EmailVerificationToken schema.
type EmailVerificationToken struct {
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
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// DeletedBy holds the value of the "deleted_by" field.
	DeletedBy string `json:"deleted_by,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID string `json:"owner_id,omitempty"`
	// the verification token sent to the user via email which should only be provided to the /verify endpoint + handler
	Token string `json:"token,omitempty"`
	// the ttl of the verification token which defaults to 7 days
	TTL *time.Time `json:"ttl,omitempty"`
	// the email used as input to generate the verification token; this is used to verify that the token when regenerated within the server matches the token emailed
	Email string `json:"email,omitempty"`
	// the comparison secret to verify the token's signature
	Secret *[]byte `json:"secret,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EmailVerificationTokenQuery when eager-loading is set.
	Edges        EmailVerificationTokenEdges `json:"edges"`
	selectValues sql.SelectValues
}

// EmailVerificationTokenEdges holds the relations/edges for other nodes in the graph.
type EmailVerificationTokenEdges struct {
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
func (e EmailVerificationTokenEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*EmailVerificationToken) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case emailverificationtoken.FieldSecret:
			values[i] = new([]byte)
		case emailverificationtoken.FieldID, emailverificationtoken.FieldCreatedBy, emailverificationtoken.FieldUpdatedBy, emailverificationtoken.FieldDeletedBy, emailverificationtoken.FieldOwnerID, emailverificationtoken.FieldToken, emailverificationtoken.FieldEmail:
			values[i] = new(sql.NullString)
		case emailverificationtoken.FieldCreatedAt, emailverificationtoken.FieldUpdatedAt, emailverificationtoken.FieldDeletedAt, emailverificationtoken.FieldTTL:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the EmailVerificationToken fields.
func (evt *EmailVerificationToken) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case emailverificationtoken.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				evt.ID = value.String
			}
		case emailverificationtoken.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				evt.CreatedAt = value.Time
			}
		case emailverificationtoken.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				evt.UpdatedAt = value.Time
			}
		case emailverificationtoken.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				evt.CreatedBy = value.String
			}
		case emailverificationtoken.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				evt.UpdatedBy = value.String
			}
		case emailverificationtoken.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				evt.DeletedAt = value.Time
			}
		case emailverificationtoken.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				evt.DeletedBy = value.String
			}
		case emailverificationtoken.FieldOwnerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				evt.OwnerID = value.String
			}
		case emailverificationtoken.FieldToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field token", values[i])
			} else if value.Valid {
				evt.Token = value.String
			}
		case emailverificationtoken.FieldTTL:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field ttl", values[i])
			} else if value.Valid {
				evt.TTL = new(time.Time)
				*evt.TTL = value.Time
			}
		case emailverificationtoken.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				evt.Email = value.String
			}
		case emailverificationtoken.FieldSecret:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field secret", values[i])
			} else if value != nil {
				evt.Secret = value
			}
		default:
			evt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the EmailVerificationToken.
// This includes values selected through modifiers, order, etc.
func (evt *EmailVerificationToken) Value(name string) (ent.Value, error) {
	return evt.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the EmailVerificationToken entity.
func (evt *EmailVerificationToken) QueryOwner() *UserQuery {
	return NewEmailVerificationTokenClient(evt.config).QueryOwner(evt)
}

// Update returns a builder for updating this EmailVerificationToken.
// Note that you need to call EmailVerificationToken.Unwrap() before calling this method if this EmailVerificationToken
// was returned from a transaction, and the transaction was committed or rolled back.
func (evt *EmailVerificationToken) Update() *EmailVerificationTokenUpdateOne {
	return NewEmailVerificationTokenClient(evt.config).UpdateOne(evt)
}

// Unwrap unwraps the EmailVerificationToken entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (evt *EmailVerificationToken) Unwrap() *EmailVerificationToken {
	_tx, ok := evt.config.driver.(*txDriver)
	if !ok {
		panic("generated: EmailVerificationToken is not a transactional entity")
	}
	evt.config.driver = _tx.drv
	return evt
}

// String implements the fmt.Stringer.
func (evt *EmailVerificationToken) String() string {
	var builder strings.Builder
	builder.WriteString("EmailVerificationToken(")
	builder.WriteString(fmt.Sprintf("id=%v, ", evt.ID))
	builder.WriteString("created_at=")
	builder.WriteString(evt.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(evt.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(evt.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(evt.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(evt.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(evt.DeletedBy)
	builder.WriteString(", ")
	builder.WriteString("owner_id=")
	builder.WriteString(evt.OwnerID)
	builder.WriteString(", ")
	builder.WriteString("token=")
	builder.WriteString(evt.Token)
	builder.WriteString(", ")
	if v := evt.TTL; v != nil {
		builder.WriteString("ttl=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(evt.Email)
	builder.WriteString(", ")
	if v := evt.Secret; v != nil {
		builder.WriteString("secret=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// EmailVerificationTokens is a parsable slice of EmailVerificationToken.
type EmailVerificationTokens []*EmailVerificationToken
