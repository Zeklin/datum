// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/entitlementhistory"
	"github.com/datumforge/enthistory"
)

// EntitlementHistoryCreate is the builder for creating a EntitlementHistory entity.
type EntitlementHistoryCreate struct {
	config
	mutation *EntitlementHistoryMutation
	hooks    []Hook
}

// SetHistoryTime sets the "history_time" field.
func (ehc *EntitlementHistoryCreate) SetHistoryTime(t time.Time) *EntitlementHistoryCreate {
	ehc.mutation.SetHistoryTime(t)
	return ehc
}

// SetNillableHistoryTime sets the "history_time" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableHistoryTime(t *time.Time) *EntitlementHistoryCreate {
	if t != nil {
		ehc.SetHistoryTime(*t)
	}
	return ehc
}

// SetOperation sets the "operation" field.
func (ehc *EntitlementHistoryCreate) SetOperation(et enthistory.OpType) *EntitlementHistoryCreate {
	ehc.mutation.SetOperation(et)
	return ehc
}

// SetRef sets the "ref" field.
func (ehc *EntitlementHistoryCreate) SetRef(s string) *EntitlementHistoryCreate {
	ehc.mutation.SetRef(s)
	return ehc
}

// SetNillableRef sets the "ref" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableRef(s *string) *EntitlementHistoryCreate {
	if s != nil {
		ehc.SetRef(*s)
	}
	return ehc
}

// SetCreatedAt sets the "created_at" field.
func (ehc *EntitlementHistoryCreate) SetCreatedAt(t time.Time) *EntitlementHistoryCreate {
	ehc.mutation.SetCreatedAt(t)
	return ehc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableCreatedAt(t *time.Time) *EntitlementHistoryCreate {
	if t != nil {
		ehc.SetCreatedAt(*t)
	}
	return ehc
}

// SetUpdatedAt sets the "updated_at" field.
func (ehc *EntitlementHistoryCreate) SetUpdatedAt(t time.Time) *EntitlementHistoryCreate {
	ehc.mutation.SetUpdatedAt(t)
	return ehc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableUpdatedAt(t *time.Time) *EntitlementHistoryCreate {
	if t != nil {
		ehc.SetUpdatedAt(*t)
	}
	return ehc
}

// SetCreatedBy sets the "created_by" field.
func (ehc *EntitlementHistoryCreate) SetCreatedBy(s string) *EntitlementHistoryCreate {
	ehc.mutation.SetCreatedBy(s)
	return ehc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableCreatedBy(s *string) *EntitlementHistoryCreate {
	if s != nil {
		ehc.SetCreatedBy(*s)
	}
	return ehc
}

// SetUpdatedBy sets the "updated_by" field.
func (ehc *EntitlementHistoryCreate) SetUpdatedBy(s string) *EntitlementHistoryCreate {
	ehc.mutation.SetUpdatedBy(s)
	return ehc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableUpdatedBy(s *string) *EntitlementHistoryCreate {
	if s != nil {
		ehc.SetUpdatedBy(*s)
	}
	return ehc
}

// SetDeletedAt sets the "deleted_at" field.
func (ehc *EntitlementHistoryCreate) SetDeletedAt(t time.Time) *EntitlementHistoryCreate {
	ehc.mutation.SetDeletedAt(t)
	return ehc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableDeletedAt(t *time.Time) *EntitlementHistoryCreate {
	if t != nil {
		ehc.SetDeletedAt(*t)
	}
	return ehc
}

// SetDeletedBy sets the "deleted_by" field.
func (ehc *EntitlementHistoryCreate) SetDeletedBy(s string) *EntitlementHistoryCreate {
	ehc.mutation.SetDeletedBy(s)
	return ehc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableDeletedBy(s *string) *EntitlementHistoryCreate {
	if s != nil {
		ehc.SetDeletedBy(*s)
	}
	return ehc
}

// SetOwnerID sets the "owner_id" field.
func (ehc *EntitlementHistoryCreate) SetOwnerID(s string) *EntitlementHistoryCreate {
	ehc.mutation.SetOwnerID(s)
	return ehc
}

// SetTier sets the "tier" field.
func (ehc *EntitlementHistoryCreate) SetTier(e enums.Tier) *EntitlementHistoryCreate {
	ehc.mutation.SetTier(e)
	return ehc
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableTier(e *enums.Tier) *EntitlementHistoryCreate {
	if e != nil {
		ehc.SetTier(*e)
	}
	return ehc
}

// SetExternalCustomerID sets the "external_customer_id" field.
func (ehc *EntitlementHistoryCreate) SetExternalCustomerID(s string) *EntitlementHistoryCreate {
	ehc.mutation.SetExternalCustomerID(s)
	return ehc
}

// SetNillableExternalCustomerID sets the "external_customer_id" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableExternalCustomerID(s *string) *EntitlementHistoryCreate {
	if s != nil {
		ehc.SetExternalCustomerID(*s)
	}
	return ehc
}

// SetExternalSubscriptionID sets the "external_subscription_id" field.
func (ehc *EntitlementHistoryCreate) SetExternalSubscriptionID(s string) *EntitlementHistoryCreate {
	ehc.mutation.SetExternalSubscriptionID(s)
	return ehc
}

// SetNillableExternalSubscriptionID sets the "external_subscription_id" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableExternalSubscriptionID(s *string) *EntitlementHistoryCreate {
	if s != nil {
		ehc.SetExternalSubscriptionID(*s)
	}
	return ehc
}

// SetExpires sets the "expires" field.
func (ehc *EntitlementHistoryCreate) SetExpires(b bool) *EntitlementHistoryCreate {
	ehc.mutation.SetExpires(b)
	return ehc
}

// SetNillableExpires sets the "expires" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableExpires(b *bool) *EntitlementHistoryCreate {
	if b != nil {
		ehc.SetExpires(*b)
	}
	return ehc
}

// SetExpiresAt sets the "expires_at" field.
func (ehc *EntitlementHistoryCreate) SetExpiresAt(t time.Time) *EntitlementHistoryCreate {
	ehc.mutation.SetExpiresAt(t)
	return ehc
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableExpiresAt(t *time.Time) *EntitlementHistoryCreate {
	if t != nil {
		ehc.SetExpiresAt(*t)
	}
	return ehc
}

// SetCancelled sets the "cancelled" field.
func (ehc *EntitlementHistoryCreate) SetCancelled(b bool) *EntitlementHistoryCreate {
	ehc.mutation.SetCancelled(b)
	return ehc
}

// SetNillableCancelled sets the "cancelled" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableCancelled(b *bool) *EntitlementHistoryCreate {
	if b != nil {
		ehc.SetCancelled(*b)
	}
	return ehc
}

// SetID sets the "id" field.
func (ehc *EntitlementHistoryCreate) SetID(s string) *EntitlementHistoryCreate {
	ehc.mutation.SetID(s)
	return ehc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ehc *EntitlementHistoryCreate) SetNillableID(s *string) *EntitlementHistoryCreate {
	if s != nil {
		ehc.SetID(*s)
	}
	return ehc
}

// Mutation returns the EntitlementHistoryMutation object of the builder.
func (ehc *EntitlementHistoryCreate) Mutation() *EntitlementHistoryMutation {
	return ehc.mutation
}

// Save creates the EntitlementHistory in the database.
func (ehc *EntitlementHistoryCreate) Save(ctx context.Context) (*EntitlementHistory, error) {
	ehc.defaults()
	return withHooks(ctx, ehc.sqlSave, ehc.mutation, ehc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ehc *EntitlementHistoryCreate) SaveX(ctx context.Context) *EntitlementHistory {
	v, err := ehc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ehc *EntitlementHistoryCreate) Exec(ctx context.Context) error {
	_, err := ehc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehc *EntitlementHistoryCreate) ExecX(ctx context.Context) {
	if err := ehc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ehc *EntitlementHistoryCreate) defaults() {
	if _, ok := ehc.mutation.HistoryTime(); !ok {
		v := entitlementhistory.DefaultHistoryTime()
		ehc.mutation.SetHistoryTime(v)
	}
	if _, ok := ehc.mutation.CreatedAt(); !ok {
		v := entitlementhistory.DefaultCreatedAt()
		ehc.mutation.SetCreatedAt(v)
	}
	if _, ok := ehc.mutation.UpdatedAt(); !ok {
		v := entitlementhistory.DefaultUpdatedAt()
		ehc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ehc.mutation.Tier(); !ok {
		v := entitlementhistory.DefaultTier
		ehc.mutation.SetTier(v)
	}
	if _, ok := ehc.mutation.Expires(); !ok {
		v := entitlementhistory.DefaultExpires
		ehc.mutation.SetExpires(v)
	}
	if _, ok := ehc.mutation.Cancelled(); !ok {
		v := entitlementhistory.DefaultCancelled
		ehc.mutation.SetCancelled(v)
	}
	if _, ok := ehc.mutation.ID(); !ok {
		v := entitlementhistory.DefaultID()
		ehc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ehc *EntitlementHistoryCreate) check() error {
	if _, ok := ehc.mutation.HistoryTime(); !ok {
		return &ValidationError{Name: "history_time", err: errors.New(`generated: missing required field "EntitlementHistory.history_time"`)}
	}
	if _, ok := ehc.mutation.Operation(); !ok {
		return &ValidationError{Name: "operation", err: errors.New(`generated: missing required field "EntitlementHistory.operation"`)}
	}
	if v, ok := ehc.mutation.Operation(); ok {
		if err := entitlementhistory.OperationValidator(v); err != nil {
			return &ValidationError{Name: "operation", err: fmt.Errorf(`generated: validator failed for field "EntitlementHistory.operation": %w`, err)}
		}
	}
	if _, ok := ehc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`generated: missing required field "EntitlementHistory.owner_id"`)}
	}
	if _, ok := ehc.mutation.Tier(); !ok {
		return &ValidationError{Name: "tier", err: errors.New(`generated: missing required field "EntitlementHistory.tier"`)}
	}
	if v, ok := ehc.mutation.Tier(); ok {
		if err := entitlementhistory.TierValidator(v); err != nil {
			return &ValidationError{Name: "tier", err: fmt.Errorf(`generated: validator failed for field "EntitlementHistory.tier": %w`, err)}
		}
	}
	if _, ok := ehc.mutation.Expires(); !ok {
		return &ValidationError{Name: "expires", err: errors.New(`generated: missing required field "EntitlementHistory.expires"`)}
	}
	if _, ok := ehc.mutation.Cancelled(); !ok {
		return &ValidationError{Name: "cancelled", err: errors.New(`generated: missing required field "EntitlementHistory.cancelled"`)}
	}
	return nil
}

func (ehc *EntitlementHistoryCreate) sqlSave(ctx context.Context) (*EntitlementHistory, error) {
	if err := ehc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ehc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ehc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected EntitlementHistory.ID type: %T", _spec.ID.Value)
		}
	}
	ehc.mutation.id = &_node.ID
	ehc.mutation.done = true
	return _node, nil
}

func (ehc *EntitlementHistoryCreate) createSpec() (*EntitlementHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &EntitlementHistory{config: ehc.config}
		_spec = sqlgraph.NewCreateSpec(entitlementhistory.Table, sqlgraph.NewFieldSpec(entitlementhistory.FieldID, field.TypeString))
	)
	_spec.Schema = ehc.schemaConfig.EntitlementHistory
	if id, ok := ehc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ehc.mutation.HistoryTime(); ok {
		_spec.SetField(entitlementhistory.FieldHistoryTime, field.TypeTime, value)
		_node.HistoryTime = value
	}
	if value, ok := ehc.mutation.Operation(); ok {
		_spec.SetField(entitlementhistory.FieldOperation, field.TypeEnum, value)
		_node.Operation = value
	}
	if value, ok := ehc.mutation.Ref(); ok {
		_spec.SetField(entitlementhistory.FieldRef, field.TypeString, value)
		_node.Ref = value
	}
	if value, ok := ehc.mutation.CreatedAt(); ok {
		_spec.SetField(entitlementhistory.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ehc.mutation.UpdatedAt(); ok {
		_spec.SetField(entitlementhistory.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ehc.mutation.CreatedBy(); ok {
		_spec.SetField(entitlementhistory.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ehc.mutation.UpdatedBy(); ok {
		_spec.SetField(entitlementhistory.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ehc.mutation.DeletedAt(); ok {
		_spec.SetField(entitlementhistory.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := ehc.mutation.DeletedBy(); ok {
		_spec.SetField(entitlementhistory.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := ehc.mutation.OwnerID(); ok {
		_spec.SetField(entitlementhistory.FieldOwnerID, field.TypeString, value)
		_node.OwnerID = value
	}
	if value, ok := ehc.mutation.Tier(); ok {
		_spec.SetField(entitlementhistory.FieldTier, field.TypeEnum, value)
		_node.Tier = value
	}
	if value, ok := ehc.mutation.ExternalCustomerID(); ok {
		_spec.SetField(entitlementhistory.FieldExternalCustomerID, field.TypeString, value)
		_node.ExternalCustomerID = value
	}
	if value, ok := ehc.mutation.ExternalSubscriptionID(); ok {
		_spec.SetField(entitlementhistory.FieldExternalSubscriptionID, field.TypeString, value)
		_node.ExternalSubscriptionID = value
	}
	if value, ok := ehc.mutation.Expires(); ok {
		_spec.SetField(entitlementhistory.FieldExpires, field.TypeBool, value)
		_node.Expires = value
	}
	if value, ok := ehc.mutation.ExpiresAt(); ok {
		_spec.SetField(entitlementhistory.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = &value
	}
	if value, ok := ehc.mutation.Cancelled(); ok {
		_spec.SetField(entitlementhistory.FieldCancelled, field.TypeBool, value)
		_node.Cancelled = value
	}
	return _node, _spec
}

// EntitlementHistoryCreateBulk is the builder for creating many EntitlementHistory entities in bulk.
type EntitlementHistoryCreateBulk struct {
	config
	err      error
	builders []*EntitlementHistoryCreate
}

// Save creates the EntitlementHistory entities in the database.
func (ehcb *EntitlementHistoryCreateBulk) Save(ctx context.Context) ([]*EntitlementHistory, error) {
	if ehcb.err != nil {
		return nil, ehcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ehcb.builders))
	nodes := make([]*EntitlementHistory, len(ehcb.builders))
	mutators := make([]Mutator, len(ehcb.builders))
	for i := range ehcb.builders {
		func(i int, root context.Context) {
			builder := ehcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EntitlementHistoryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ehcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ehcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ehcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ehcb *EntitlementHistoryCreateBulk) SaveX(ctx context.Context) []*EntitlementHistory {
	v, err := ehcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ehcb *EntitlementHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ehcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ehcb *EntitlementHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ehcb.Exec(ctx); err != nil {
		panic(err)
	}
}