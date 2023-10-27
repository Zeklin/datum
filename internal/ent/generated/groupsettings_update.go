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
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/groupsettings"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/google/uuid"
)

// GroupSettingsUpdate is the builder for updating GroupSettings entities.
type GroupSettingsUpdate struct {
	config
	hooks    []Hook
	mutation *GroupSettingsMutation
}

// Where appends a list predicates to the GroupSettingsUpdate builder.
func (gsu *GroupSettingsUpdate) Where(ps ...predicate.GroupSettings) *GroupSettingsUpdate {
	gsu.mutation.Where(ps...)
	return gsu
}

// SetUpdatedAt sets the "updated_at" field.
func (gsu *GroupSettingsUpdate) SetUpdatedAt(t time.Time) *GroupSettingsUpdate {
	gsu.mutation.SetUpdatedAt(t)
	return gsu
}

// SetCreatedBy sets the "created_by" field.
func (gsu *GroupSettingsUpdate) SetCreatedBy(i int) *GroupSettingsUpdate {
	gsu.mutation.ResetCreatedBy()
	gsu.mutation.SetCreatedBy(i)
	return gsu
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gsu *GroupSettingsUpdate) SetNillableCreatedBy(i *int) *GroupSettingsUpdate {
	if i != nil {
		gsu.SetCreatedBy(*i)
	}
	return gsu
}

// AddCreatedBy adds i to the "created_by" field.
func (gsu *GroupSettingsUpdate) AddCreatedBy(i int) *GroupSettingsUpdate {
	gsu.mutation.AddCreatedBy(i)
	return gsu
}

// ClearCreatedBy clears the value of the "created_by" field.
func (gsu *GroupSettingsUpdate) ClearCreatedBy() *GroupSettingsUpdate {
	gsu.mutation.ClearCreatedBy()
	return gsu
}

// SetUpdatedBy sets the "updated_by" field.
func (gsu *GroupSettingsUpdate) SetUpdatedBy(i int) *GroupSettingsUpdate {
	gsu.mutation.ResetUpdatedBy()
	gsu.mutation.SetUpdatedBy(i)
	return gsu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gsu *GroupSettingsUpdate) SetNillableUpdatedBy(i *int) *GroupSettingsUpdate {
	if i != nil {
		gsu.SetUpdatedBy(*i)
	}
	return gsu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (gsu *GroupSettingsUpdate) AddUpdatedBy(i int) *GroupSettingsUpdate {
	gsu.mutation.AddUpdatedBy(i)
	return gsu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (gsu *GroupSettingsUpdate) ClearUpdatedBy() *GroupSettingsUpdate {
	gsu.mutation.ClearUpdatedBy()
	return gsu
}

// SetVisibility sets the "visibility" field.
func (gsu *GroupSettingsUpdate) SetVisibility(gr groupsettings.Visibility) *GroupSettingsUpdate {
	gsu.mutation.SetVisibility(gr)
	return gsu
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (gsu *GroupSettingsUpdate) SetNillableVisibility(gr *groupsettings.Visibility) *GroupSettingsUpdate {
	if gr != nil {
		gsu.SetVisibility(*gr)
	}
	return gsu
}

// SetJoinPolicy sets the "join_policy" field.
func (gsu *GroupSettingsUpdate) SetJoinPolicy(gp groupsettings.JoinPolicy) *GroupSettingsUpdate {
	gsu.mutation.SetJoinPolicy(gp)
	return gsu
}

// SetNillableJoinPolicy sets the "join_policy" field if the given value is not nil.
func (gsu *GroupSettingsUpdate) SetNillableJoinPolicy(gp *groupsettings.JoinPolicy) *GroupSettingsUpdate {
	if gp != nil {
		gsu.SetJoinPolicy(*gp)
	}
	return gsu
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gsu *GroupSettingsUpdate) SetGroupID(id uuid.UUID) *GroupSettingsUpdate {
	gsu.mutation.SetGroupID(id)
	return gsu
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (gsu *GroupSettingsUpdate) SetNillableGroupID(id *uuid.UUID) *GroupSettingsUpdate {
	if id != nil {
		gsu = gsu.SetGroupID(*id)
	}
	return gsu
}

// SetGroup sets the "group" edge to the Group entity.
func (gsu *GroupSettingsUpdate) SetGroup(g *Group) *GroupSettingsUpdate {
	return gsu.SetGroupID(g.ID)
}

// Mutation returns the GroupSettingsMutation object of the builder.
func (gsu *GroupSettingsUpdate) Mutation() *GroupSettingsMutation {
	return gsu.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (gsu *GroupSettingsUpdate) ClearGroup() *GroupSettingsUpdate {
	gsu.mutation.ClearGroup()
	return gsu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gsu *GroupSettingsUpdate) Save(ctx context.Context) (int, error) {
	if err := gsu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, gsu.sqlSave, gsu.mutation, gsu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gsu *GroupSettingsUpdate) SaveX(ctx context.Context) int {
	affected, err := gsu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gsu *GroupSettingsUpdate) Exec(ctx context.Context) error {
	_, err := gsu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsu *GroupSettingsUpdate) ExecX(ctx context.Context) {
	if err := gsu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsu *GroupSettingsUpdate) defaults() error {
	if _, ok := gsu.mutation.UpdatedAt(); !ok {
		if groupsettings.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupsettings.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := groupsettings.UpdateDefaultUpdatedAt()
		gsu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gsu *GroupSettingsUpdate) check() error {
	if v, ok := gsu.mutation.Visibility(); ok {
		if err := groupsettings.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`generated: validator failed for field "GroupSettings.visibility": %w`, err)}
		}
	}
	if v, ok := gsu.mutation.JoinPolicy(); ok {
		if err := groupsettings.JoinPolicyValidator(v); err != nil {
			return &ValidationError{Name: "join_policy", err: fmt.Errorf(`generated: validator failed for field "GroupSettings.join_policy": %w`, err)}
		}
	}
	return nil
}

func (gsu *GroupSettingsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gsu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupsettings.Table, groupsettings.Columns, sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeUUID))
	if ps := gsu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gsu.mutation.UpdatedAt(); ok {
		_spec.SetField(groupsettings.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gsu.mutation.CreatedBy(); ok {
		_spec.SetField(groupsettings.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := gsu.mutation.AddedCreatedBy(); ok {
		_spec.AddField(groupsettings.FieldCreatedBy, field.TypeInt, value)
	}
	if gsu.mutation.CreatedByCleared() {
		_spec.ClearField(groupsettings.FieldCreatedBy, field.TypeInt)
	}
	if value, ok := gsu.mutation.UpdatedBy(); ok {
		_spec.SetField(groupsettings.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := gsu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(groupsettings.FieldUpdatedBy, field.TypeInt, value)
	}
	if gsu.mutation.UpdatedByCleared() {
		_spec.ClearField(groupsettings.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := gsu.mutation.Visibility(); ok {
		_spec.SetField(groupsettings.FieldVisibility, field.TypeEnum, value)
	}
	if value, ok := gsu.mutation.JoinPolicy(); ok {
		_spec.SetField(groupsettings.FieldJoinPolicy, field.TypeEnum, value)
	}
	if gsu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   groupsettings.GroupTable,
			Columns: []string{groupsettings.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   groupsettings.GroupTable,
			Columns: []string{groupsettings.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gsu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupsettings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gsu.mutation.done = true
	return n, nil
}

// GroupSettingsUpdateOne is the builder for updating a single GroupSettings entity.
type GroupSettingsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupSettingsMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (gsuo *GroupSettingsUpdateOne) SetUpdatedAt(t time.Time) *GroupSettingsUpdateOne {
	gsuo.mutation.SetUpdatedAt(t)
	return gsuo
}

// SetCreatedBy sets the "created_by" field.
func (gsuo *GroupSettingsUpdateOne) SetCreatedBy(i int) *GroupSettingsUpdateOne {
	gsuo.mutation.ResetCreatedBy()
	gsuo.mutation.SetCreatedBy(i)
	return gsuo
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (gsuo *GroupSettingsUpdateOne) SetNillableCreatedBy(i *int) *GroupSettingsUpdateOne {
	if i != nil {
		gsuo.SetCreatedBy(*i)
	}
	return gsuo
}

// AddCreatedBy adds i to the "created_by" field.
func (gsuo *GroupSettingsUpdateOne) AddCreatedBy(i int) *GroupSettingsUpdateOne {
	gsuo.mutation.AddCreatedBy(i)
	return gsuo
}

// ClearCreatedBy clears the value of the "created_by" field.
func (gsuo *GroupSettingsUpdateOne) ClearCreatedBy() *GroupSettingsUpdateOne {
	gsuo.mutation.ClearCreatedBy()
	return gsuo
}

// SetUpdatedBy sets the "updated_by" field.
func (gsuo *GroupSettingsUpdateOne) SetUpdatedBy(i int) *GroupSettingsUpdateOne {
	gsuo.mutation.ResetUpdatedBy()
	gsuo.mutation.SetUpdatedBy(i)
	return gsuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (gsuo *GroupSettingsUpdateOne) SetNillableUpdatedBy(i *int) *GroupSettingsUpdateOne {
	if i != nil {
		gsuo.SetUpdatedBy(*i)
	}
	return gsuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (gsuo *GroupSettingsUpdateOne) AddUpdatedBy(i int) *GroupSettingsUpdateOne {
	gsuo.mutation.AddUpdatedBy(i)
	return gsuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (gsuo *GroupSettingsUpdateOne) ClearUpdatedBy() *GroupSettingsUpdateOne {
	gsuo.mutation.ClearUpdatedBy()
	return gsuo
}

// SetVisibility sets the "visibility" field.
func (gsuo *GroupSettingsUpdateOne) SetVisibility(gr groupsettings.Visibility) *GroupSettingsUpdateOne {
	gsuo.mutation.SetVisibility(gr)
	return gsuo
}

// SetNillableVisibility sets the "visibility" field if the given value is not nil.
func (gsuo *GroupSettingsUpdateOne) SetNillableVisibility(gr *groupsettings.Visibility) *GroupSettingsUpdateOne {
	if gr != nil {
		gsuo.SetVisibility(*gr)
	}
	return gsuo
}

// SetJoinPolicy sets the "join_policy" field.
func (gsuo *GroupSettingsUpdateOne) SetJoinPolicy(gp groupsettings.JoinPolicy) *GroupSettingsUpdateOne {
	gsuo.mutation.SetJoinPolicy(gp)
	return gsuo
}

// SetNillableJoinPolicy sets the "join_policy" field if the given value is not nil.
func (gsuo *GroupSettingsUpdateOne) SetNillableJoinPolicy(gp *groupsettings.JoinPolicy) *GroupSettingsUpdateOne {
	if gp != nil {
		gsuo.SetJoinPolicy(*gp)
	}
	return gsuo
}

// SetGroupID sets the "group" edge to the Group entity by ID.
func (gsuo *GroupSettingsUpdateOne) SetGroupID(id uuid.UUID) *GroupSettingsUpdateOne {
	gsuo.mutation.SetGroupID(id)
	return gsuo
}

// SetNillableGroupID sets the "group" edge to the Group entity by ID if the given value is not nil.
func (gsuo *GroupSettingsUpdateOne) SetNillableGroupID(id *uuid.UUID) *GroupSettingsUpdateOne {
	if id != nil {
		gsuo = gsuo.SetGroupID(*id)
	}
	return gsuo
}

// SetGroup sets the "group" edge to the Group entity.
func (gsuo *GroupSettingsUpdateOne) SetGroup(g *Group) *GroupSettingsUpdateOne {
	return gsuo.SetGroupID(g.ID)
}

// Mutation returns the GroupSettingsMutation object of the builder.
func (gsuo *GroupSettingsUpdateOne) Mutation() *GroupSettingsMutation {
	return gsuo.mutation
}

// ClearGroup clears the "group" edge to the Group entity.
func (gsuo *GroupSettingsUpdateOne) ClearGroup() *GroupSettingsUpdateOne {
	gsuo.mutation.ClearGroup()
	return gsuo
}

// Where appends a list predicates to the GroupSettingsUpdate builder.
func (gsuo *GroupSettingsUpdateOne) Where(ps ...predicate.GroupSettings) *GroupSettingsUpdateOne {
	gsuo.mutation.Where(ps...)
	return gsuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gsuo *GroupSettingsUpdateOne) Select(field string, fields ...string) *GroupSettingsUpdateOne {
	gsuo.fields = append([]string{field}, fields...)
	return gsuo
}

// Save executes the query and returns the updated GroupSettings entity.
func (gsuo *GroupSettingsUpdateOne) Save(ctx context.Context) (*GroupSettings, error) {
	if err := gsuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, gsuo.sqlSave, gsuo.mutation, gsuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gsuo *GroupSettingsUpdateOne) SaveX(ctx context.Context) *GroupSettings {
	node, err := gsuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gsuo *GroupSettingsUpdateOne) Exec(ctx context.Context) error {
	_, err := gsuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gsuo *GroupSettingsUpdateOne) ExecX(ctx context.Context) {
	if err := gsuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gsuo *GroupSettingsUpdateOne) defaults() error {
	if _, ok := gsuo.mutation.UpdatedAt(); !ok {
		if groupsettings.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized groupsettings.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := groupsettings.UpdateDefaultUpdatedAt()
		gsuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (gsuo *GroupSettingsUpdateOne) check() error {
	if v, ok := gsuo.mutation.Visibility(); ok {
		if err := groupsettings.VisibilityValidator(v); err != nil {
			return &ValidationError{Name: "visibility", err: fmt.Errorf(`generated: validator failed for field "GroupSettings.visibility": %w`, err)}
		}
	}
	if v, ok := gsuo.mutation.JoinPolicy(); ok {
		if err := groupsettings.JoinPolicyValidator(v); err != nil {
			return &ValidationError{Name: "join_policy", err: fmt.Errorf(`generated: validator failed for field "GroupSettings.join_policy": %w`, err)}
		}
	}
	return nil
}

func (gsuo *GroupSettingsUpdateOne) sqlSave(ctx context.Context) (_node *GroupSettings, err error) {
	if err := gsuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(groupsettings.Table, groupsettings.Columns, sqlgraph.NewFieldSpec(groupsettings.FieldID, field.TypeUUID))
	id, ok := gsuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "GroupSettings.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gsuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, groupsettings.FieldID)
		for _, f := range fields {
			if !groupsettings.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != groupsettings.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gsuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := gsuo.mutation.UpdatedAt(); ok {
		_spec.SetField(groupsettings.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := gsuo.mutation.CreatedBy(); ok {
		_spec.SetField(groupsettings.FieldCreatedBy, field.TypeInt, value)
	}
	if value, ok := gsuo.mutation.AddedCreatedBy(); ok {
		_spec.AddField(groupsettings.FieldCreatedBy, field.TypeInt, value)
	}
	if gsuo.mutation.CreatedByCleared() {
		_spec.ClearField(groupsettings.FieldCreatedBy, field.TypeInt)
	}
	if value, ok := gsuo.mutation.UpdatedBy(); ok {
		_spec.SetField(groupsettings.FieldUpdatedBy, field.TypeInt, value)
	}
	if value, ok := gsuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(groupsettings.FieldUpdatedBy, field.TypeInt, value)
	}
	if gsuo.mutation.UpdatedByCleared() {
		_spec.ClearField(groupsettings.FieldUpdatedBy, field.TypeInt)
	}
	if value, ok := gsuo.mutation.Visibility(); ok {
		_spec.SetField(groupsettings.FieldVisibility, field.TypeEnum, value)
	}
	if value, ok := gsuo.mutation.JoinPolicy(); ok {
		_spec.SetField(groupsettings.FieldJoinPolicy, field.TypeEnum, value)
	}
	if gsuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   groupsettings.GroupTable,
			Columns: []string{groupsettings.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gsuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   groupsettings.GroupTable,
			Columns: []string{groupsettings.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GroupSettings{config: gsuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gsuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{groupsettings.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gsuo.mutation.done = true
	return _node, nil
}