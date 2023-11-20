// Code generated by ent, DO NOT EDIT.

package generated

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/datumforge/datum/internal/ent/generated/entitlement"
	"github.com/datumforge/datum/internal/ent/generated/organization"
)

// Entitlement is the model entity for the Entitlement schema.
type Entitlement struct {
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
	// Tier holds the value of the "tier" field.
	Tier entitlement.Tier `json:"tier,omitempty"`
	// used to store references to external systems, e.g. Stripe
	ExternalCustomerID string `json:"external_customer_id,omitempty"`
	// used to store references to external systems, e.g. Stripe
	ExternalSubscriptionID string `json:"external_subscription_id,omitempty"`
	// ExpiresAt holds the value of the "expires_at" field.
	ExpiresAt time.Time `json:"expires_at,omitempty"`
	// UpgradedAt holds the value of the "upgraded_at" field.
	UpgradedAt time.Time `json:"upgraded_at,omitempty"`
	// the tier the customer upgraded from
	UpgradedTier string `json:"upgraded_tier,omitempty"`
	// DowngradedAt holds the value of the "downgraded_at" field.
	DowngradedAt time.Time `json:"downgraded_at,omitempty"`
	// the tier the customer downgraded from
	DowngradedTier string `json:"downgraded_tier,omitempty"`
	// Cancelled holds the value of the "cancelled" field.
	Cancelled bool `json:"cancelled,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EntitlementQuery when eager-loading is set.
	Edges                     EntitlementEdges `json:"edges"`
	organization_entitlements *string
	selectValues              sql.SelectValues
}

// EntitlementEdges holds the relations/edges for other nodes in the graph.
type EntitlementEdges struct {
	// Owner holds the value of the owner edge.
	Owner *Organization `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EntitlementEdges) OwnerOrErr() (*Organization, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Entitlement) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case entitlement.FieldCancelled:
			values[i] = new(sql.NullBool)
		case entitlement.FieldID, entitlement.FieldCreatedBy, entitlement.FieldUpdatedBy, entitlement.FieldTier, entitlement.FieldExternalCustomerID, entitlement.FieldExternalSubscriptionID, entitlement.FieldUpgradedTier, entitlement.FieldDowngradedTier:
			values[i] = new(sql.NullString)
		case entitlement.FieldCreatedAt, entitlement.FieldUpdatedAt, entitlement.FieldExpiresAt, entitlement.FieldUpgradedAt, entitlement.FieldDowngradedAt:
			values[i] = new(sql.NullTime)
		case entitlement.ForeignKeys[0]: // organization_entitlements
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Entitlement fields.
func (e *Entitlement) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case entitlement.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				e.ID = value.String
			}
		case entitlement.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				e.CreatedAt = value.Time
			}
		case entitlement.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				e.UpdatedAt = value.Time
			}
		case entitlement.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				e.CreatedBy = value.String
			}
		case entitlement.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				e.UpdatedBy = value.String
			}
		case entitlement.FieldTier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tier", values[i])
			} else if value.Valid {
				e.Tier = entitlement.Tier(value.String)
			}
		case entitlement.FieldExternalCustomerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_customer_id", values[i])
			} else if value.Valid {
				e.ExternalCustomerID = value.String
			}
		case entitlement.FieldExternalSubscriptionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field external_subscription_id", values[i])
			} else if value.Valid {
				e.ExternalSubscriptionID = value.String
			}
		case entitlement.FieldExpiresAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expires_at", values[i])
			} else if value.Valid {
				e.ExpiresAt = value.Time
			}
		case entitlement.FieldUpgradedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field upgraded_at", values[i])
			} else if value.Valid {
				e.UpgradedAt = value.Time
			}
		case entitlement.FieldUpgradedTier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field upgraded_tier", values[i])
			} else if value.Valid {
				e.UpgradedTier = value.String
			}
		case entitlement.FieldDowngradedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field downgraded_at", values[i])
			} else if value.Valid {
				e.DowngradedAt = value.Time
			}
		case entitlement.FieldDowngradedTier:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field downgraded_tier", values[i])
			} else if value.Valid {
				e.DowngradedTier = value.String
			}
		case entitlement.FieldCancelled:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field cancelled", values[i])
			} else if value.Valid {
				e.Cancelled = value.Bool
			}
		case entitlement.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization_entitlements", values[i])
			} else if value.Valid {
				e.organization_entitlements = new(string)
				*e.organization_entitlements = value.String
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Entitlement.
// This includes values selected through modifiers, order, etc.
func (e *Entitlement) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Entitlement entity.
func (e *Entitlement) QueryOwner() *OrganizationQuery {
	return NewEntitlementClient(e.config).QueryOwner(e)
}

// Update returns a builder for updating this Entitlement.
// Note that you need to call Entitlement.Unwrap() before calling this method if this Entitlement
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Entitlement) Update() *EntitlementUpdateOne {
	return NewEntitlementClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Entitlement entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Entitlement) Unwrap() *Entitlement {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("generated: Entitlement is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Entitlement) String() string {
	var builder strings.Builder
	builder.WriteString("Entitlement(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("created_at=")
	builder.WriteString(e.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(e.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(e.CreatedBy)
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(e.UpdatedBy)
	builder.WriteString(", ")
	builder.WriteString("tier=")
	builder.WriteString(fmt.Sprintf("%v", e.Tier))
	builder.WriteString(", ")
	builder.WriteString("external_customer_id=")
	builder.WriteString(e.ExternalCustomerID)
	builder.WriteString(", ")
	builder.WriteString("external_subscription_id=")
	builder.WriteString(e.ExternalSubscriptionID)
	builder.WriteString(", ")
	builder.WriteString("expires_at=")
	builder.WriteString(e.ExpiresAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("upgraded_at=")
	builder.WriteString(e.UpgradedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("upgraded_tier=")
	builder.WriteString(e.UpgradedTier)
	builder.WriteString(", ")
	builder.WriteString("downgraded_at=")
	builder.WriteString(e.DowngradedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("downgraded_tier=")
	builder.WriteString(e.DowngradedTier)
	builder.WriteString(", ")
	builder.WriteString("cancelled=")
	builder.WriteString(fmt.Sprintf("%v", e.Cancelled))
	builder.WriteByte(')')
	return builder.String()
}

// Entitlements is a parsable slice of Entitlement.
type Entitlements []*Entitlement