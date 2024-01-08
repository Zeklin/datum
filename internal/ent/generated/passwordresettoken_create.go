// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/passwordresettoken"
	"github.com/datumforge/datum/internal/ent/generated/user"
)

// PasswordResetTokenCreate is the builder for creating a PasswordResetToken entity.
type PasswordResetTokenCreate struct {
	config
	mutation *PasswordResetTokenMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (prtc *PasswordResetTokenCreate) SetCreatedAt(t time.Time) *PasswordResetTokenCreate {
	prtc.mutation.SetCreatedAt(t)
	return prtc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (prtc *PasswordResetTokenCreate) SetNillableCreatedAt(t *time.Time) *PasswordResetTokenCreate {
	if t != nil {
		prtc.SetCreatedAt(*t)
	}
	return prtc
}

// SetUpdatedAt sets the "updated_at" field.
func (prtc *PasswordResetTokenCreate) SetUpdatedAt(t time.Time) *PasswordResetTokenCreate {
	prtc.mutation.SetUpdatedAt(t)
	return prtc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (prtc *PasswordResetTokenCreate) SetNillableUpdatedAt(t *time.Time) *PasswordResetTokenCreate {
	if t != nil {
		prtc.SetUpdatedAt(*t)
	}
	return prtc
}

// SetCreatedBy sets the "created_by" field.
func (prtc *PasswordResetTokenCreate) SetCreatedBy(s string) *PasswordResetTokenCreate {
	prtc.mutation.SetCreatedBy(s)
	return prtc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (prtc *PasswordResetTokenCreate) SetNillableCreatedBy(s *string) *PasswordResetTokenCreate {
	if s != nil {
		prtc.SetCreatedBy(*s)
	}
	return prtc
}

// SetUpdatedBy sets the "updated_by" field.
func (prtc *PasswordResetTokenCreate) SetUpdatedBy(s string) *PasswordResetTokenCreate {
	prtc.mutation.SetUpdatedBy(s)
	return prtc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (prtc *PasswordResetTokenCreate) SetNillableUpdatedBy(s *string) *PasswordResetTokenCreate {
	if s != nil {
		prtc.SetUpdatedBy(*s)
	}
	return prtc
}

// SetDeletedAt sets the "deleted_at" field.
func (prtc *PasswordResetTokenCreate) SetDeletedAt(t time.Time) *PasswordResetTokenCreate {
	prtc.mutation.SetDeletedAt(t)
	return prtc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (prtc *PasswordResetTokenCreate) SetNillableDeletedAt(t *time.Time) *PasswordResetTokenCreate {
	if t != nil {
		prtc.SetDeletedAt(*t)
	}
	return prtc
}

// SetDeletedBy sets the "deleted_by" field.
func (prtc *PasswordResetTokenCreate) SetDeletedBy(s string) *PasswordResetTokenCreate {
	prtc.mutation.SetDeletedBy(s)
	return prtc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (prtc *PasswordResetTokenCreate) SetNillableDeletedBy(s *string) *PasswordResetTokenCreate {
	if s != nil {
		prtc.SetDeletedBy(*s)
	}
	return prtc
}

// SetOwnerID sets the "owner_id" field.
func (prtc *PasswordResetTokenCreate) SetOwnerID(s string) *PasswordResetTokenCreate {
	prtc.mutation.SetOwnerID(s)
	return prtc
}

// SetToken sets the "token" field.
func (prtc *PasswordResetTokenCreate) SetToken(s string) *PasswordResetTokenCreate {
	prtc.mutation.SetToken(s)
	return prtc
}

// SetTTL sets the "ttl" field.
func (prtc *PasswordResetTokenCreate) SetTTL(t time.Time) *PasswordResetTokenCreate {
	prtc.mutation.SetTTL(t)
	return prtc
}

// SetEmail sets the "email" field.
func (prtc *PasswordResetTokenCreate) SetEmail(s string) *PasswordResetTokenCreate {
	prtc.mutation.SetEmail(s)
	return prtc
}

// SetSecret sets the "secret" field.
func (prtc *PasswordResetTokenCreate) SetSecret(b []byte) *PasswordResetTokenCreate {
	prtc.mutation.SetSecret(b)
	return prtc
}

// SetID sets the "id" field.
func (prtc *PasswordResetTokenCreate) SetID(s string) *PasswordResetTokenCreate {
	prtc.mutation.SetID(s)
	return prtc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (prtc *PasswordResetTokenCreate) SetNillableID(s *string) *PasswordResetTokenCreate {
	if s != nil {
		prtc.SetID(*s)
	}
	return prtc
}

// SetOwner sets the "owner" edge to the User entity.
func (prtc *PasswordResetTokenCreate) SetOwner(u *User) *PasswordResetTokenCreate {
	return prtc.SetOwnerID(u.ID)
}

// Mutation returns the PasswordResetTokenMutation object of the builder.
func (prtc *PasswordResetTokenCreate) Mutation() *PasswordResetTokenMutation {
	return prtc.mutation
}

// Save creates the PasswordResetToken in the database.
func (prtc *PasswordResetTokenCreate) Save(ctx context.Context) (*PasswordResetToken, error) {
	if err := prtc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, prtc.sqlSave, prtc.mutation, prtc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (prtc *PasswordResetTokenCreate) SaveX(ctx context.Context) *PasswordResetToken {
	v, err := prtc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (prtc *PasswordResetTokenCreate) Exec(ctx context.Context) error {
	_, err := prtc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prtc *PasswordResetTokenCreate) ExecX(ctx context.Context) {
	if err := prtc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (prtc *PasswordResetTokenCreate) defaults() error {
	if _, ok := prtc.mutation.CreatedAt(); !ok {
		if passwordresettoken.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized passwordresettoken.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := passwordresettoken.DefaultCreatedAt()
		prtc.mutation.SetCreatedAt(v)
	}
	if _, ok := prtc.mutation.UpdatedAt(); !ok {
		if passwordresettoken.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized passwordresettoken.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := passwordresettoken.DefaultUpdatedAt()
		prtc.mutation.SetUpdatedAt(v)
	}
	if _, ok := prtc.mutation.ID(); !ok {
		if passwordresettoken.DefaultID == nil {
			return fmt.Errorf("generated: uninitialized passwordresettoken.DefaultID (forgotten import generated/runtime?)")
		}
		v := passwordresettoken.DefaultID()
		prtc.mutation.SetID(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (prtc *PasswordResetTokenCreate) check() error {
	if _, ok := prtc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "PasswordResetToken.created_at"`)}
	}
	if _, ok := prtc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "PasswordResetToken.updated_at"`)}
	}
	if _, ok := prtc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner_id", err: errors.New(`generated: missing required field "PasswordResetToken.owner_id"`)}
	}
	if _, ok := prtc.mutation.Token(); !ok {
		return &ValidationError{Name: "token", err: errors.New(`generated: missing required field "PasswordResetToken.token"`)}
	}
	if v, ok := prtc.mutation.Token(); ok {
		if err := passwordresettoken.TokenValidator(v); err != nil {
			return &ValidationError{Name: "token", err: fmt.Errorf(`generated: validator failed for field "PasswordResetToken.token": %w`, err)}
		}
	}
	if _, ok := prtc.mutation.TTL(); !ok {
		return &ValidationError{Name: "ttl", err: errors.New(`generated: missing required field "PasswordResetToken.ttl"`)}
	}
	if _, ok := prtc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`generated: missing required field "PasswordResetToken.email"`)}
	}
	if v, ok := prtc.mutation.Email(); ok {
		if err := passwordresettoken.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`generated: validator failed for field "PasswordResetToken.email": %w`, err)}
		}
	}
	if _, ok := prtc.mutation.Secret(); !ok {
		return &ValidationError{Name: "secret", err: errors.New(`generated: missing required field "PasswordResetToken.secret"`)}
	}
	if v, ok := prtc.mutation.Secret(); ok {
		if err := passwordresettoken.SecretValidator(v); err != nil {
			return &ValidationError{Name: "secret", err: fmt.Errorf(`generated: validator failed for field "PasswordResetToken.secret": %w`, err)}
		}
	}
	if _, ok := prtc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`generated: missing required edge "PasswordResetToken.owner"`)}
	}
	return nil
}

func (prtc *PasswordResetTokenCreate) sqlSave(ctx context.Context) (*PasswordResetToken, error) {
	if err := prtc.check(); err != nil {
		return nil, err
	}
	_node, _spec := prtc.createSpec()
	if err := sqlgraph.CreateNode(ctx, prtc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected PasswordResetToken.ID type: %T", _spec.ID.Value)
		}
	}
	prtc.mutation.id = &_node.ID
	prtc.mutation.done = true
	return _node, nil
}

func (prtc *PasswordResetTokenCreate) createSpec() (*PasswordResetToken, *sqlgraph.CreateSpec) {
	var (
		_node = &PasswordResetToken{config: prtc.config}
		_spec = sqlgraph.NewCreateSpec(passwordresettoken.Table, sqlgraph.NewFieldSpec(passwordresettoken.FieldID, field.TypeString))
	)
	_spec.Schema = prtc.schemaConfig.PasswordResetToken
	if id, ok := prtc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := prtc.mutation.CreatedAt(); ok {
		_spec.SetField(passwordresettoken.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := prtc.mutation.UpdatedAt(); ok {
		_spec.SetField(passwordresettoken.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := prtc.mutation.CreatedBy(); ok {
		_spec.SetField(passwordresettoken.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := prtc.mutation.UpdatedBy(); ok {
		_spec.SetField(passwordresettoken.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := prtc.mutation.DeletedAt(); ok {
		_spec.SetField(passwordresettoken.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := prtc.mutation.DeletedBy(); ok {
		_spec.SetField(passwordresettoken.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := prtc.mutation.Token(); ok {
		_spec.SetField(passwordresettoken.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := prtc.mutation.TTL(); ok {
		_spec.SetField(passwordresettoken.FieldTTL, field.TypeTime, value)
		_node.TTL = &value
	}
	if value, ok := prtc.mutation.Email(); ok {
		_spec.SetField(passwordresettoken.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := prtc.mutation.Secret(); ok {
		_spec.SetField(passwordresettoken.FieldSecret, field.TypeBytes, value)
		_node.Secret = &value
	}
	if nodes := prtc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   passwordresettoken.OwnerTable,
			Columns: []string{passwordresettoken.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = prtc.schemaConfig.PasswordResetToken
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.OwnerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PasswordResetTokenCreateBulk is the builder for creating many PasswordResetToken entities in bulk.
type PasswordResetTokenCreateBulk struct {
	config
	err      error
	builders []*PasswordResetTokenCreate
}

// Save creates the PasswordResetToken entities in the database.
func (prtcb *PasswordResetTokenCreateBulk) Save(ctx context.Context) ([]*PasswordResetToken, error) {
	if prtcb.err != nil {
		return nil, prtcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(prtcb.builders))
	nodes := make([]*PasswordResetToken, len(prtcb.builders))
	mutators := make([]Mutator, len(prtcb.builders))
	for i := range prtcb.builders {
		func(i int, root context.Context) {
			builder := prtcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PasswordResetTokenMutation)
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
					_, err = mutators[i+1].Mutate(root, prtcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, prtcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, prtcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (prtcb *PasswordResetTokenCreateBulk) SaveX(ctx context.Context) []*PasswordResetToken {
	v, err := prtcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (prtcb *PasswordResetTokenCreateBulk) Exec(ctx context.Context) error {
	_, err := prtcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (prtcb *PasswordResetTokenCreateBulk) ExecX(ctx context.Context) {
	if err := prtcb.Exec(ctx); err != nil {
		panic(err)
	}
}
