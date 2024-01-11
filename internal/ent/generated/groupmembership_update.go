// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/groupmembership"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// GroupMembershipUpdate is the builder for updating GroupMembership entities.
type GroupMembershipUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMembershipMutation
}

// Where appends a list predicates to the GroupMembershipUpdate builder.
func (gmu *GroupMembershipUpdate) Where(ps ...predicate.GroupMembership) *GroupMembershipUpdate {
	gmu.mutation.Where(ps...)
	return gmu
}

// SetUpdatedAt sets the "updated_at" field.
func (gmu *GroupMembershipUpdate) SetUpdatedAt(t time.Time) *GroupMembershipUpdate {
	gmu.mutation.SetUpdatedAt(t)
	return gmu
}

// SetUpdatedBy sets the "updated_by" field.
func (gmu *GroupMembershipUpdate) SetUpdatedBy(s string) *GroupMembershipUpdate {
	gmu.mutation.SetUpdatedBy(s)
	return gmu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gmu *GroupMembershipUpdate) SetNillableUpdatedBy(s *string) *GroupMembershipUpdate {
	if s != nil {
		gmu.SetUpdatedBy(*s)
	}
	return gmu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (gmu *GroupMembershipUpdate) ClearUpdatedBy() *GroupMembershipUpdate {
	gmu.mutation.ClearUpdatedBy()
	return gmu
}

// SetDeletedAt sets the "deleted_at" field.
func (gmu *GroupMembershipUpdate) SetDeletedAt(t time.Time) *GroupMembershipUpdate {
	gmu.mutation.SetDeletedAt(t)
	return gmu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gmu *GroupMembershipUpdate) SetNillableDeletedAt(t *time.Time) *GroupMembershipUpdate {
	if t != nil {
		gmu.SetDeletedAt(*t)
	}
	return gmu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gmu *GroupMembershipUpdate) ClearDeletedAt() *GroupMembershipUpdate {
	gmu.mutation.ClearDeletedAt()
	return gmu
}

// SetDeletedBy sets the "deleted_by" field.
func (gmu *GroupMembershipUpdate) SetDeletedBy(s string) *GroupMembershipUpdate {
	gmu.mutation.SetDeletedBy(s)
	return gmu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gmu *GroupMembershipUpdate) SetNillableDeletedBy(s *string) *GroupMembershipUpdate {
	if s != nil {
		gmu.SetDeletedBy(*s)
	}
	return gmu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (gmu *GroupMembershipUpdate) ClearDeletedBy() *GroupMembershipUpdate {
	gmu.mutation.ClearDeletedBy()
	return gmu
}

// SetRole sets the "role" field.
func (gmu *GroupMembershipUpdate) SetRole(e enums.Role) *GroupMembershipUpdate {
	gmu.mutation.SetRole(e)
	return gmu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (gmu *GroupMembershipUpdate) SetNillableRole(e *enums.Role) *GroupMembershipUpdate {
	if e != nil {
		gmu.SetRole(*e)
	}
	return gmu
}

// Mutation returns the GroupMembershipMutation object of the builder.
func (gmu *GroupMembershipUpdate) Mutation() *GroupMembershipMutation {
	return gmu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gmu *GroupMembershipUpdate) Save(ctx context.Context) (int, error) {
	if err := gmu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, gmu.sqlSave, gmu.mutation, gmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmu *GroupMembershipUpdate) SaveX(ctx context.Context) int {
	affected, err := gmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gmu *GroupMembershipUpdate) Exec(ctx context.Context) error {
	_, err := gmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmu *GroupMembershipUpdate) ExecX(ctx context.Context) {
	if err := gmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmu *GroupMembershipUpdate) defaults() error {
	if _, ok := gmu.mutation.UpdatedAt(); !ok {
		if groupmembership.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupmembership.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := groupmembership.UpdateDefaultUpdatedAt()
		gmu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gmu *GroupMembershipUpdate) check() error {
	if v, ok := gmu.mutation.Role(); ok {
		if err := groupmembership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "GroupMembership.role": %w`, err)}
		}
	}
	if _, ok := gmu.mutation.GroupID(); gmu.mutation.GroupCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "GroupMembership.group"`)
	}
	if _, ok := gmu.mutation.UserID(); gmu.mutation.UserCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "GroupMembership.user"`)
	}
	return nil
}

func (gmu *GroupMembershipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupmembership.Table, groupmembership.Columns, sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString))
	if ps := gmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmu.mutation.UpdatedAt(); ok {
		_spec.SetField(groupmembership.FieldUpdatedAt, field.TypeTime, value)
	}
	if gmu.mutation.CreatedByCleared() {
		_spec.ClearField(groupmembership.FieldCreatedBy, field.TypeString)
	}
	if value, ok := gmu.mutation.UpdatedBy(); ok {
		_spec.SetField(groupmembership.FieldUpdatedBy, field.TypeString, value)
	}
	if gmu.mutation.UpdatedByCleared() {
		_spec.ClearField(groupmembership.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := gmu.mutation.DeletedAt(); ok {
		_spec.SetField(groupmembership.FieldDeletedAt, field.TypeTime, value)
	}
	if gmu.mutation.DeletedAtCleared() {
		_spec.ClearField(groupmembership.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gmu.mutation.DeletedBy(); ok {
		_spec.SetField(groupmembership.FieldDeletedBy, field.TypeString, value)
	}
	if gmu.mutation.DeletedByCleared() {
		_spec.ClearField(groupmembership.FieldDeletedBy, field.TypeString)
	}
	if value, ok := gmu.mutation.Role(); ok {
		_spec.SetField(groupmembership.FieldRole, field.TypeEnum, value)
	}
	_spec.Node.Schema = gmu.schemaConfig.GroupMembership
	ctx = internal.NewSchemaConfigContext(ctx, gmu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, gmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmembership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gmu.mutation.done = true
	return n, nil
}

// GroupMembershipUpdateOne is the builder for updating a single GroupMembership entity.
type GroupMembershipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMembershipMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (gmuo *GroupMembershipUpdateOne) SetUpdatedAt(t time.Time) *GroupMembershipUpdateOne {
	gmuo.mutation.SetUpdatedAt(t)
	return gmuo
}

// SetUpdatedBy sets the "updated_by" field.
func (gmuo *GroupMembershipUpdateOne) SetUpdatedBy(s string) *GroupMembershipUpdateOne {
	gmuo.mutation.SetUpdatedBy(s)
	return gmuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gmuo *GroupMembershipUpdateOne) SetNillableUpdatedBy(s *string) *GroupMembershipUpdateOne {
	if s != nil {
		gmuo.SetUpdatedBy(*s)
	}
	return gmuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (gmuo *GroupMembershipUpdateOne) ClearUpdatedBy() *GroupMembershipUpdateOne {
	gmuo.mutation.ClearUpdatedBy()
	return gmuo
}

// SetDeletedAt sets the "deleted_at" field.
func (gmuo *GroupMembershipUpdateOne) SetDeletedAt(t time.Time) *GroupMembershipUpdateOne {
	gmuo.mutation.SetDeletedAt(t)
	return gmuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (gmuo *GroupMembershipUpdateOne) SetNillableDeletedAt(t *time.Time) *GroupMembershipUpdateOne {
	if t != nil {
		gmuo.SetDeletedAt(*t)
	}
	return gmuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (gmuo *GroupMembershipUpdateOne) ClearDeletedAt() *GroupMembershipUpdateOne {
	gmuo.mutation.ClearDeletedAt()
	return gmuo
}

// SetDeletedBy sets the "deleted_by" field.
func (gmuo *GroupMembershipUpdateOne) SetDeletedBy(s string) *GroupMembershipUpdateOne {
	gmuo.mutation.SetDeletedBy(s)
	return gmuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (gmuo *GroupMembershipUpdateOne) SetNillableDeletedBy(s *string) *GroupMembershipUpdateOne {
	if s != nil {
		gmuo.SetDeletedBy(*s)
	}
	return gmuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (gmuo *GroupMembershipUpdateOne) ClearDeletedBy() *GroupMembershipUpdateOne {
	gmuo.mutation.ClearDeletedBy()
	return gmuo
}

// SetRole sets the "role" field.
func (gmuo *GroupMembershipUpdateOne) SetRole(e enums.Role) *GroupMembershipUpdateOne {
	gmuo.mutation.SetRole(e)
	return gmuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (gmuo *GroupMembershipUpdateOne) SetNillableRole(e *enums.Role) *GroupMembershipUpdateOne {
	if e != nil {
		gmuo.SetRole(*e)
	}
	return gmuo
}

// Mutation returns the GroupMembershipMutation object of the builder.
func (gmuo *GroupMembershipUpdateOne) Mutation() *GroupMembershipMutation {
	return gmuo.mutation
}

// Where appends a list predicates to the GroupMembershipUpdate builder.
func (gmuo *GroupMembershipUpdateOne) Where(ps ...predicate.GroupMembership) *GroupMembershipUpdateOne {
	gmuo.mutation.Where(ps...)
	return gmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gmuo *GroupMembershipUpdateOne) Select(field string, fields ...string) *GroupMembershipUpdateOne {
	gmuo.fields = append([]string{field}, fields...)
	return gmuo
}

// Save executes the query and returns the updated GroupMembership entity.
func (gmuo *GroupMembershipUpdateOne) Save(ctx context.Context) (*GroupMembership, error) {
	if err := gmuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gmuo.sqlSave, gmuo.mutation, gmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gmuo *GroupMembershipUpdateOne) SaveX(ctx context.Context) *GroupMembership {
	node, err := gmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gmuo *GroupMembershipUpdateOne) Exec(ctx context.Context) error {
	_, err := gmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gmuo *GroupMembershipUpdateOne) ExecX(ctx context.Context) {
	if err := gmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gmuo *GroupMembershipUpdateOne) defaults() error {
	if _, ok := gmuo.mutation.UpdatedAt(); !ok {
		if groupmembership.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupmembership.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := groupmembership.UpdateDefaultUpdatedAt()
		gmuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gmuo *GroupMembershipUpdateOne) check() error {
	if v, ok := gmuo.mutation.Role(); ok {
		if err := groupmembership.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`generated: validator failed for field "GroupMembership.role": %w`, err)}
		}
	}
	if _, ok := gmuo.mutation.GroupID(); gmuo.mutation.GroupCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "GroupMembership.group"`)
	}
	if _, ok := gmuo.mutation.UserID(); gmuo.mutation.UserCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "GroupMembership.user"`)
	}
	return nil
}

func (gmuo *GroupMembershipUpdateOne) sqlSave(ctx context.Context) (_node *GroupMembership, err error) {
	if err := gmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupmembership.Table, groupmembership.Columns, sqlgraph.NewFieldSpec(groupmembership.FieldID, field.TypeString))
	id, ok := gmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "GroupMembership.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupmembership.FieldID)
		for _, f := range fields {
			if !groupmembership.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != groupmembership.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gmuo.mutation.UpdatedAt(); ok {
		_spec.SetField(groupmembership.FieldUpdatedAt, field.TypeTime, value)
	}
	if gmuo.mutation.CreatedByCleared() {
		_spec.ClearField(groupmembership.FieldCreatedBy, field.TypeString)
	}
	if value, ok := gmuo.mutation.UpdatedBy(); ok {
		_spec.SetField(groupmembership.FieldUpdatedBy, field.TypeString, value)
	}
	if gmuo.mutation.UpdatedByCleared() {
		_spec.ClearField(groupmembership.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := gmuo.mutation.DeletedAt(); ok {
		_spec.SetField(groupmembership.FieldDeletedAt, field.TypeTime, value)
	}
	if gmuo.mutation.DeletedAtCleared() {
		_spec.ClearField(groupmembership.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := gmuo.mutation.DeletedBy(); ok {
		_spec.SetField(groupmembership.FieldDeletedBy, field.TypeString, value)
	}
	if gmuo.mutation.DeletedByCleared() {
		_spec.ClearField(groupmembership.FieldDeletedBy, field.TypeString)
	}
	if value, ok := gmuo.mutation.Role(); ok {
		_spec.SetField(groupmembership.FieldRole, field.TypeEnum, value)
	}
	_spec.Node.Schema = gmuo.schemaConfig.GroupMembership
	ctx = internal.NewSchemaConfigContext(ctx, gmuo.schemaConfig)
	_node = &GroupMembership{config: gmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupmembership.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gmuo.mutation.done = true
	return _node, nil
}