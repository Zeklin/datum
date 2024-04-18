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
	"github.com/datumforge/datum/internal/ent/customtypes"
	"github.com/datumforge/datum/internal/ent/enums"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/templatehistory"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// TemplateHistoryUpdate is the builder for updating TemplateHistory entities.
type TemplateHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *TemplateHistoryMutation
}

// Where appends a list predicates to the TemplateHistoryUpdate builder.
func (thu *TemplateHistoryUpdate) Where(ps ...predicate.TemplateHistory) *TemplateHistoryUpdate {
	thu.mutation.Where(ps...)
	return thu
}

// SetUpdatedAt sets the "updated_at" field.
func (thu *TemplateHistoryUpdate) SetUpdatedAt(t time.Time) *TemplateHistoryUpdate {
	thu.mutation.SetUpdatedAt(t)
	return thu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (thu *TemplateHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *TemplateHistoryUpdate {
	if t != nil {
		thu.SetUpdatedAt(*t)
	}
	return thu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (thu *TemplateHistoryUpdate) ClearUpdatedAt() *TemplateHistoryUpdate {
	thu.mutation.ClearUpdatedAt()
	return thu
}

// SetUpdatedBy sets the "updated_by" field.
func (thu *TemplateHistoryUpdate) SetUpdatedBy(s string) *TemplateHistoryUpdate {
	thu.mutation.SetUpdatedBy(s)
	return thu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (thu *TemplateHistoryUpdate) SetNillableUpdatedBy(s *string) *TemplateHistoryUpdate {
	if s != nil {
		thu.SetUpdatedBy(*s)
	}
	return thu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (thu *TemplateHistoryUpdate) ClearUpdatedBy() *TemplateHistoryUpdate {
	thu.mutation.ClearUpdatedBy()
	return thu
}

// SetDeletedAt sets the "deleted_at" field.
func (thu *TemplateHistoryUpdate) SetDeletedAt(t time.Time) *TemplateHistoryUpdate {
	thu.mutation.SetDeletedAt(t)
	return thu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (thu *TemplateHistoryUpdate) SetNillableDeletedAt(t *time.Time) *TemplateHistoryUpdate {
	if t != nil {
		thu.SetDeletedAt(*t)
	}
	return thu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (thu *TemplateHistoryUpdate) ClearDeletedAt() *TemplateHistoryUpdate {
	thu.mutation.ClearDeletedAt()
	return thu
}

// SetDeletedBy sets the "deleted_by" field.
func (thu *TemplateHistoryUpdate) SetDeletedBy(s string) *TemplateHistoryUpdate {
	thu.mutation.SetDeletedBy(s)
	return thu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (thu *TemplateHistoryUpdate) SetNillableDeletedBy(s *string) *TemplateHistoryUpdate {
	if s != nil {
		thu.SetDeletedBy(*s)
	}
	return thu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (thu *TemplateHistoryUpdate) ClearDeletedBy() *TemplateHistoryUpdate {
	thu.mutation.ClearDeletedBy()
	return thu
}

// SetOwnerID sets the "owner_id" field.
func (thu *TemplateHistoryUpdate) SetOwnerID(s string) *TemplateHistoryUpdate {
	thu.mutation.SetOwnerID(s)
	return thu
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (thu *TemplateHistoryUpdate) SetNillableOwnerID(s *string) *TemplateHistoryUpdate {
	if s != nil {
		thu.SetOwnerID(*s)
	}
	return thu
}

// SetName sets the "name" field.
func (thu *TemplateHistoryUpdate) SetName(s string) *TemplateHistoryUpdate {
	thu.mutation.SetName(s)
	return thu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (thu *TemplateHistoryUpdate) SetNillableName(s *string) *TemplateHistoryUpdate {
	if s != nil {
		thu.SetName(*s)
	}
	return thu
}

// SetTemplateType sets the "template_type" field.
func (thu *TemplateHistoryUpdate) SetTemplateType(et enums.DocumentType) *TemplateHistoryUpdate {
	thu.mutation.SetTemplateType(et)
	return thu
}

// SetNillableTemplateType sets the "template_type" field if the given value is not nil.
func (thu *TemplateHistoryUpdate) SetNillableTemplateType(et *enums.DocumentType) *TemplateHistoryUpdate {
	if et != nil {
		thu.SetTemplateType(*et)
	}
	return thu
}

// SetDescription sets the "description" field.
func (thu *TemplateHistoryUpdate) SetDescription(s string) *TemplateHistoryUpdate {
	thu.mutation.SetDescription(s)
	return thu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (thu *TemplateHistoryUpdate) SetNillableDescription(s *string) *TemplateHistoryUpdate {
	if s != nil {
		thu.SetDescription(*s)
	}
	return thu
}

// ClearDescription clears the value of the "description" field.
func (thu *TemplateHistoryUpdate) ClearDescription() *TemplateHistoryUpdate {
	thu.mutation.ClearDescription()
	return thu
}

// SetJsonconfig sets the "jsonconfig" field.
func (thu *TemplateHistoryUpdate) SetJsonconfig(co customtypes.JSONObject) *TemplateHistoryUpdate {
	thu.mutation.SetJsonconfig(co)
	return thu
}

// SetUischema sets the "uischema" field.
func (thu *TemplateHistoryUpdate) SetUischema(co customtypes.JSONObject) *TemplateHistoryUpdate {
	thu.mutation.SetUischema(co)
	return thu
}

// ClearUischema clears the value of the "uischema" field.
func (thu *TemplateHistoryUpdate) ClearUischema() *TemplateHistoryUpdate {
	thu.mutation.ClearUischema()
	return thu
}

// Mutation returns the TemplateHistoryMutation object of the builder.
func (thu *TemplateHistoryUpdate) Mutation() *TemplateHistoryMutation {
	return thu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (thu *TemplateHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, thu.sqlSave, thu.mutation, thu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thu *TemplateHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := thu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (thu *TemplateHistoryUpdate) Exec(ctx context.Context) error {
	_, err := thu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thu *TemplateHistoryUpdate) ExecX(ctx context.Context) {
	if err := thu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thu *TemplateHistoryUpdate) check() error {
	if v, ok := thu.mutation.TemplateType(); ok {
		if err := templatehistory.TemplateTypeValidator(v); err != nil {
			return &ValidationError{Name: "template_type", err: fmt.Errorf(`generated: validator failed for field "TemplateHistory.template_type": %w`, err)}
		}
	}
	return nil
}

func (thu *TemplateHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := thu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(templatehistory.Table, templatehistory.Columns, sqlgraph.NewFieldSpec(templatehistory.FieldID, field.TypeString))
	if ps := thu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if thu.mutation.RefCleared() {
		_spec.ClearField(templatehistory.FieldRef, field.TypeString)
	}
	if thu.mutation.CreatedAtCleared() {
		_spec.ClearField(templatehistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := thu.mutation.UpdatedAt(); ok {
		_spec.SetField(templatehistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if thu.mutation.UpdatedAtCleared() {
		_spec.ClearField(templatehistory.FieldUpdatedAt, field.TypeTime)
	}
	if thu.mutation.CreatedByCleared() {
		_spec.ClearField(templatehistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := thu.mutation.UpdatedBy(); ok {
		_spec.SetField(templatehistory.FieldUpdatedBy, field.TypeString, value)
	}
	if thu.mutation.UpdatedByCleared() {
		_spec.ClearField(templatehistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := thu.mutation.DeletedAt(); ok {
		_spec.SetField(templatehistory.FieldDeletedAt, field.TypeTime, value)
	}
	if thu.mutation.DeletedAtCleared() {
		_spec.ClearField(templatehistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := thu.mutation.DeletedBy(); ok {
		_spec.SetField(templatehistory.FieldDeletedBy, field.TypeString, value)
	}
	if thu.mutation.DeletedByCleared() {
		_spec.ClearField(templatehistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := thu.mutation.OwnerID(); ok {
		_spec.SetField(templatehistory.FieldOwnerID, field.TypeString, value)
	}
	if value, ok := thu.mutation.Name(); ok {
		_spec.SetField(templatehistory.FieldName, field.TypeString, value)
	}
	if value, ok := thu.mutation.TemplateType(); ok {
		_spec.SetField(templatehistory.FieldTemplateType, field.TypeEnum, value)
	}
	if value, ok := thu.mutation.Description(); ok {
		_spec.SetField(templatehistory.FieldDescription, field.TypeString, value)
	}
	if thu.mutation.DescriptionCleared() {
		_spec.ClearField(templatehistory.FieldDescription, field.TypeString)
	}
	if value, ok := thu.mutation.Jsonconfig(); ok {
		_spec.SetField(templatehistory.FieldJsonconfig, field.TypeJSON, value)
	}
	if value, ok := thu.mutation.Uischema(); ok {
		_spec.SetField(templatehistory.FieldUischema, field.TypeJSON, value)
	}
	if thu.mutation.UischemaCleared() {
		_spec.ClearField(templatehistory.FieldUischema, field.TypeJSON)
	}
	_spec.Node.Schema = thu.schemaConfig.TemplateHistory
	ctx = internal.NewSchemaConfigContext(ctx, thu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, thu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{templatehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	thu.mutation.done = true
	return n, nil
}

// TemplateHistoryUpdateOne is the builder for updating a single TemplateHistory entity.
type TemplateHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TemplateHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (thuo *TemplateHistoryUpdateOne) SetUpdatedAt(t time.Time) *TemplateHistoryUpdateOne {
	thuo.mutation.SetUpdatedAt(t)
	return thuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (thuo *TemplateHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *TemplateHistoryUpdateOne {
	if t != nil {
		thuo.SetUpdatedAt(*t)
	}
	return thuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (thuo *TemplateHistoryUpdateOne) ClearUpdatedAt() *TemplateHistoryUpdateOne {
	thuo.mutation.ClearUpdatedAt()
	return thuo
}

// SetUpdatedBy sets the "updated_by" field.
func (thuo *TemplateHistoryUpdateOne) SetUpdatedBy(s string) *TemplateHistoryUpdateOne {
	thuo.mutation.SetUpdatedBy(s)
	return thuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (thuo *TemplateHistoryUpdateOne) SetNillableUpdatedBy(s *string) *TemplateHistoryUpdateOne {
	if s != nil {
		thuo.SetUpdatedBy(*s)
	}
	return thuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (thuo *TemplateHistoryUpdateOne) ClearUpdatedBy() *TemplateHistoryUpdateOne {
	thuo.mutation.ClearUpdatedBy()
	return thuo
}

// SetDeletedAt sets the "deleted_at" field.
func (thuo *TemplateHistoryUpdateOne) SetDeletedAt(t time.Time) *TemplateHistoryUpdateOne {
	thuo.mutation.SetDeletedAt(t)
	return thuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (thuo *TemplateHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *TemplateHistoryUpdateOne {
	if t != nil {
		thuo.SetDeletedAt(*t)
	}
	return thuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (thuo *TemplateHistoryUpdateOne) ClearDeletedAt() *TemplateHistoryUpdateOne {
	thuo.mutation.ClearDeletedAt()
	return thuo
}

// SetDeletedBy sets the "deleted_by" field.
func (thuo *TemplateHistoryUpdateOne) SetDeletedBy(s string) *TemplateHistoryUpdateOne {
	thuo.mutation.SetDeletedBy(s)
	return thuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (thuo *TemplateHistoryUpdateOne) SetNillableDeletedBy(s *string) *TemplateHistoryUpdateOne {
	if s != nil {
		thuo.SetDeletedBy(*s)
	}
	return thuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (thuo *TemplateHistoryUpdateOne) ClearDeletedBy() *TemplateHistoryUpdateOne {
	thuo.mutation.ClearDeletedBy()
	return thuo
}

// SetOwnerID sets the "owner_id" field.
func (thuo *TemplateHistoryUpdateOne) SetOwnerID(s string) *TemplateHistoryUpdateOne {
	thuo.mutation.SetOwnerID(s)
	return thuo
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (thuo *TemplateHistoryUpdateOne) SetNillableOwnerID(s *string) *TemplateHistoryUpdateOne {
	if s != nil {
		thuo.SetOwnerID(*s)
	}
	return thuo
}

// SetName sets the "name" field.
func (thuo *TemplateHistoryUpdateOne) SetName(s string) *TemplateHistoryUpdateOne {
	thuo.mutation.SetName(s)
	return thuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (thuo *TemplateHistoryUpdateOne) SetNillableName(s *string) *TemplateHistoryUpdateOne {
	if s != nil {
		thuo.SetName(*s)
	}
	return thuo
}

// SetTemplateType sets the "template_type" field.
func (thuo *TemplateHistoryUpdateOne) SetTemplateType(et enums.DocumentType) *TemplateHistoryUpdateOne {
	thuo.mutation.SetTemplateType(et)
	return thuo
}

// SetNillableTemplateType sets the "template_type" field if the given value is not nil.
func (thuo *TemplateHistoryUpdateOne) SetNillableTemplateType(et *enums.DocumentType) *TemplateHistoryUpdateOne {
	if et != nil {
		thuo.SetTemplateType(*et)
	}
	return thuo
}

// SetDescription sets the "description" field.
func (thuo *TemplateHistoryUpdateOne) SetDescription(s string) *TemplateHistoryUpdateOne {
	thuo.mutation.SetDescription(s)
	return thuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (thuo *TemplateHistoryUpdateOne) SetNillableDescription(s *string) *TemplateHistoryUpdateOne {
	if s != nil {
		thuo.SetDescription(*s)
	}
	return thuo
}

// ClearDescription clears the value of the "description" field.
func (thuo *TemplateHistoryUpdateOne) ClearDescription() *TemplateHistoryUpdateOne {
	thuo.mutation.ClearDescription()
	return thuo
}

// SetJsonconfig sets the "jsonconfig" field.
func (thuo *TemplateHistoryUpdateOne) SetJsonconfig(co customtypes.JSONObject) *TemplateHistoryUpdateOne {
	thuo.mutation.SetJsonconfig(co)
	return thuo
}

// SetUischema sets the "uischema" field.
func (thuo *TemplateHistoryUpdateOne) SetUischema(co customtypes.JSONObject) *TemplateHistoryUpdateOne {
	thuo.mutation.SetUischema(co)
	return thuo
}

// ClearUischema clears the value of the "uischema" field.
func (thuo *TemplateHistoryUpdateOne) ClearUischema() *TemplateHistoryUpdateOne {
	thuo.mutation.ClearUischema()
	return thuo
}

// Mutation returns the TemplateHistoryMutation object of the builder.
func (thuo *TemplateHistoryUpdateOne) Mutation() *TemplateHistoryMutation {
	return thuo.mutation
}

// Where appends a list predicates to the TemplateHistoryUpdate builder.
func (thuo *TemplateHistoryUpdateOne) Where(ps ...predicate.TemplateHistory) *TemplateHistoryUpdateOne {
	thuo.mutation.Where(ps...)
	return thuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (thuo *TemplateHistoryUpdateOne) Select(field string, fields ...string) *TemplateHistoryUpdateOne {
	thuo.fields = append([]string{field}, fields...)
	return thuo
}

// Save executes the query and returns the updated TemplateHistory entity.
func (thuo *TemplateHistoryUpdateOne) Save(ctx context.Context) (*TemplateHistory, error) {
	return withHooks(ctx, thuo.sqlSave, thuo.mutation, thuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (thuo *TemplateHistoryUpdateOne) SaveX(ctx context.Context) *TemplateHistory {
	node, err := thuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (thuo *TemplateHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := thuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (thuo *TemplateHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := thuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (thuo *TemplateHistoryUpdateOne) check() error {
	if v, ok := thuo.mutation.TemplateType(); ok {
		if err := templatehistory.TemplateTypeValidator(v); err != nil {
			return &ValidationError{Name: "template_type", err: fmt.Errorf(`generated: validator failed for field "TemplateHistory.template_type": %w`, err)}
		}
	}
	return nil
}

func (thuo *TemplateHistoryUpdateOne) sqlSave(ctx context.Context) (_node *TemplateHistory, err error) {
	if err := thuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(templatehistory.Table, templatehistory.Columns, sqlgraph.NewFieldSpec(templatehistory.FieldID, field.TypeString))
	id, ok := thuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "TemplateHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := thuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, templatehistory.FieldID)
		for _, f := range fields {
			if !templatehistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != templatehistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := thuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if thuo.mutation.RefCleared() {
		_spec.ClearField(templatehistory.FieldRef, field.TypeString)
	}
	if thuo.mutation.CreatedAtCleared() {
		_spec.ClearField(templatehistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := thuo.mutation.UpdatedAt(); ok {
		_spec.SetField(templatehistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if thuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(templatehistory.FieldUpdatedAt, field.TypeTime)
	}
	if thuo.mutation.CreatedByCleared() {
		_spec.ClearField(templatehistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := thuo.mutation.UpdatedBy(); ok {
		_spec.SetField(templatehistory.FieldUpdatedBy, field.TypeString, value)
	}
	if thuo.mutation.UpdatedByCleared() {
		_spec.ClearField(templatehistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := thuo.mutation.DeletedAt(); ok {
		_spec.SetField(templatehistory.FieldDeletedAt, field.TypeTime, value)
	}
	if thuo.mutation.DeletedAtCleared() {
		_spec.ClearField(templatehistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := thuo.mutation.DeletedBy(); ok {
		_spec.SetField(templatehistory.FieldDeletedBy, field.TypeString, value)
	}
	if thuo.mutation.DeletedByCleared() {
		_spec.ClearField(templatehistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := thuo.mutation.OwnerID(); ok {
		_spec.SetField(templatehistory.FieldOwnerID, field.TypeString, value)
	}
	if value, ok := thuo.mutation.Name(); ok {
		_spec.SetField(templatehistory.FieldName, field.TypeString, value)
	}
	if value, ok := thuo.mutation.TemplateType(); ok {
		_spec.SetField(templatehistory.FieldTemplateType, field.TypeEnum, value)
	}
	if value, ok := thuo.mutation.Description(); ok {
		_spec.SetField(templatehistory.FieldDescription, field.TypeString, value)
	}
	if thuo.mutation.DescriptionCleared() {
		_spec.ClearField(templatehistory.FieldDescription, field.TypeString)
	}
	if value, ok := thuo.mutation.Jsonconfig(); ok {
		_spec.SetField(templatehistory.FieldJsonconfig, field.TypeJSON, value)
	}
	if value, ok := thuo.mutation.Uischema(); ok {
		_spec.SetField(templatehistory.FieldUischema, field.TypeJSON, value)
	}
	if thuo.mutation.UischemaCleared() {
		_spec.ClearField(templatehistory.FieldUischema, field.TypeJSON)
	}
	_spec.Node.Schema = thuo.schemaConfig.TemplateHistory
	ctx = internal.NewSchemaConfigContext(ctx, thuo.schemaConfig)
	_node = &TemplateHistory{config: thuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, thuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{templatehistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	thuo.mutation.done = true
	return _node, nil
}
