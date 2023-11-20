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
	"github.com/datumforge/datum/internal/ent/generated/usersettings"
)

// UserSettings is the model entity for the UserSettings schema.
type UserSettings struct {
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
	// user account is locked if unconfirmed or explicitly locked
	Locked bool `json:"locked,omitempty"`
	// The time notifications regarding the user were silenced
	SilencedAt *time.Time `json:"silenced_at,omitempty"`
	// The time the user was suspended
	SuspendedAt *time.Time `json:"suspended_at,omitempty"`
	// local user password recovery code generated during account creation - does not exist for oauth'd users
	RecoveryCode *string `json:"-"`
	// Status holds the value of the "status" field.
	Status usersettings.Status `json:"status,omitempty"`
	// Role holds the value of the "role" field.
	Role usersettings.Role `json:"role,omitempty"`
	// Permissions holds the value of the "permissions" field.
	Permissions []string `json:"permissions,omitempty"`
	// EmailConfirmed holds the value of the "email_confirmed" field.
	EmailConfirmed bool `json:"email_confirmed,omitempty"`
	// tags associated with the object
	Tags []string `json:"tags,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserSettingsQuery when eager-loading is set.
	Edges        UserSettingsEdges `json:"edges"`
	user_setting *string
	selectValues sql.SelectValues
}

// UserSettingsEdges holds the relations/edges for other nodes in the graph.
type UserSettingsEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserSettingsEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserSettings) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usersettings.FieldPermissions, usersettings.FieldTags:
			values[i] = new([]byte)
		case usersettings.FieldLocked, usersettings.FieldEmailConfirmed:
			values[i] = new(sql.NullBool)
		case usersettings.FieldID, usersettings.FieldCreatedBy, usersettings.FieldUpdatedBy, usersettings.FieldRecoveryCode, usersettings.FieldStatus, usersettings.FieldRole:
			values[i] = new(sql.NullString)
		case usersettings.FieldCreatedAt, usersettings.FieldUpdatedAt, usersettings.FieldSilencedAt, usersettings.FieldSuspendedAt:
			values[i] = new(sql.NullTime)
		case usersettings.ForeignKeys[0]: // user_setting
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserSettings fields.
func (us *UserSettings) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usersettings.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				us.ID = value.String
			}
		case usersettings.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				us.CreatedAt = value.Time
			}
		case usersettings.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				us.UpdatedAt = value.Time
			}
		case usersettings.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				us.CreatedBy = value.String
			}
		case usersettings.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				us.UpdatedBy = value.String
			}
		case usersettings.FieldLocked:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field locked", values[i])
			} else if value.Valid {
				us.Locked = value.Bool
			}
		case usersettings.FieldSilencedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field silenced_at", values[i])
			} else if value.Valid {
				us.SilencedAt = new(time.Time)
				*us.SilencedAt = value.Time
			}
		case usersettings.FieldSuspendedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field suspended_at", values[i])
			} else if value.Valid {
				us.SuspendedAt = new(time.Time)
				*us.SuspendedAt = value.Time
			}
		case usersettings.FieldRecoveryCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field recovery_code", values[i])
			} else if value.Valid {
				us.RecoveryCode = new(string)
				*us.RecoveryCode = value.String
			}
		case usersettings.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				us.Status = usersettings.Status(value.String)
			}
		case usersettings.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				us.Role = usersettings.Role(value.String)
			}
		case usersettings.FieldPermissions:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field permissions", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &us.Permissions); err != nil {
					return fmt.Errorf("unmarshal field permissions: %w", err)
				}
			}
		case usersettings.FieldEmailConfirmed:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field email_confirmed", values[i])
			} else if value.Valid {
				us.EmailConfirmed = value.Bool
			}
		case usersettings.FieldTags:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field tags", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &us.Tags); err != nil {
					return fmt.Errorf("unmarshal field tags: %w", err)
				}
			}
		case usersettings.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_setting", values[i])
			} else if value.Valid {
				us.user_setting = new(string)
				*us.user_setting = value.String
			}
		default:
			us.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the UserSettings.
// This includes values selected through modifiers, order, etc.
func (us *UserSettings) Value(name string) (ent.Value, error) {
	return us.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the UserSettings entity.
func (us *UserSettings) QueryUser() *UserQuery {
	return NewUserSettingsClient(us.config).QueryUser(us)
}

// Update returns a builder for updating this UserSettings.
// Note that you need to call UserSettings.Unwrap() before calling this method if this UserSettings
// was returned from a transaction, and the transaction was committed or rolled back.
func (us *UserSettings) Update() *UserSettingsUpdateOne {
	return NewUserSettingsClient(us.config).UpdateOne(us)
}

// Unwrap unwraps the UserSettings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (us *UserSettings) Unwrap() *UserSettings {
	_tx, ok := us.config.driver.(*txDriver)
	if !ok {
		panic("generated: UserSettings is not a transactional entity")
	}
	us.config.driver = _tx.drv
	return us
}

// String implements the fmt.Stringer.
func (us *UserSettings) String() string {
	var builder strings.Builder
	builder.WriteString("UserSettings(")
	builder.WriteString(fmt.Sprintf("id=%v, ", us.ID))
	builder.WriteString("created_at=")
	builder.WriteString(us.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(us.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(us.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(us.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("locked=")
	builder.WriteString(fmt.Sprintf("%v", us.Locked))
	builder.WriteString(", ")
	if v := us.SilencedAt; v != nil {
		builder.WriteString("silenced_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := us.SuspendedAt; v != nil {
		builder.WriteString("suspended_at=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("recovery_code=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", us.Status))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", us.Role))
	builder.WriteString(", ")
	builder.WriteString("permissions=")
	builder.WriteString(fmt.Sprintf("%v", us.Permissions))
	builder.WriteString(", ")
	builder.WriteString("email_confirmed=")
	builder.WriteString(fmt.Sprintf("%v", us.EmailConfirmed))
	builder.WriteString(", ")
	builder.WriteString("tags=")
	builder.WriteString(fmt.Sprintf("%v", us.Tags))
	builder.WriteByte(')')
	return builder.String()
}

// UserSettingsSlice is a parsable slice of UserSettings.
type UserSettingsSlice []*UserSettings