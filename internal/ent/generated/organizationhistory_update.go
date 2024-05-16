// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/organizationhistory"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OrganizationHistoryUpdate is the builder for updating OrganizationHistory entities.
type OrganizationHistoryUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationHistoryMutation
}

// Where appends a list predicates to the OrganizationHistoryUpdate builder.
func (ohu *OrganizationHistoryUpdate) Where(ps ...predicate.OrganizationHistory) *OrganizationHistoryUpdate {
	ohu.mutation.Where(ps...)
	return ohu
}

// SetUpdatedAt sets the "updated_at" field.
func (ohu *OrganizationHistoryUpdate) SetUpdatedAt(t time.Time) *OrganizationHistoryUpdate {
	ohu.mutation.SetUpdatedAt(t)
	return ohu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableUpdatedAt(t *time.Time) *OrganizationHistoryUpdate {
	if t != nil {
		ohu.SetUpdatedAt(*t)
	}
	return ohu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ohu *OrganizationHistoryUpdate) ClearUpdatedAt() *OrganizationHistoryUpdate {
	ohu.mutation.ClearUpdatedAt()
	return ohu
}

// SetUpdatedBy sets the "updated_by" field.
func (ohu *OrganizationHistoryUpdate) SetUpdatedBy(s string) *OrganizationHistoryUpdate {
	ohu.mutation.SetUpdatedBy(s)
	return ohu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableUpdatedBy(s *string) *OrganizationHistoryUpdate {
	if s != nil {
		ohu.SetUpdatedBy(*s)
	}
	return ohu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ohu *OrganizationHistoryUpdate) ClearUpdatedBy() *OrganizationHistoryUpdate {
	ohu.mutation.ClearUpdatedBy()
	return ohu
}

// SetTags sets the "tags" field.
func (ohu *OrganizationHistoryUpdate) SetTags(s []string) *OrganizationHistoryUpdate {
	ohu.mutation.SetTags(s)
	return ohu
}

// AppendTags appends s to the "tags" field.
func (ohu *OrganizationHistoryUpdate) AppendTags(s []string) *OrganizationHistoryUpdate {
	ohu.mutation.AppendTags(s)
	return ohu
}

// ClearTags clears the value of the "tags" field.
func (ohu *OrganizationHistoryUpdate) ClearTags() *OrganizationHistoryUpdate {
	ohu.mutation.ClearTags()
	return ohu
}

// SetDeletedAt sets the "deleted_at" field.
func (ohu *OrganizationHistoryUpdate) SetDeletedAt(t time.Time) *OrganizationHistoryUpdate {
	ohu.mutation.SetDeletedAt(t)
	return ohu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableDeletedAt(t *time.Time) *OrganizationHistoryUpdate {
	if t != nil {
		ohu.SetDeletedAt(*t)
	}
	return ohu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ohu *OrganizationHistoryUpdate) ClearDeletedAt() *OrganizationHistoryUpdate {
	ohu.mutation.ClearDeletedAt()
	return ohu
}

// SetDeletedBy sets the "deleted_by" field.
func (ohu *OrganizationHistoryUpdate) SetDeletedBy(s string) *OrganizationHistoryUpdate {
	ohu.mutation.SetDeletedBy(s)
	return ohu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableDeletedBy(s *string) *OrganizationHistoryUpdate {
	if s != nil {
		ohu.SetDeletedBy(*s)
	}
	return ohu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ohu *OrganizationHistoryUpdate) ClearDeletedBy() *OrganizationHistoryUpdate {
	ohu.mutation.ClearDeletedBy()
	return ohu
}

// SetName sets the "name" field.
func (ohu *OrganizationHistoryUpdate) SetName(s string) *OrganizationHistoryUpdate {
	ohu.mutation.SetName(s)
	return ohu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableName(s *string) *OrganizationHistoryUpdate {
	if s != nil {
		ohu.SetName(*s)
	}
	return ohu
}

// SetDisplayName sets the "display_name" field.
func (ohu *OrganizationHistoryUpdate) SetDisplayName(s string) *OrganizationHistoryUpdate {
	ohu.mutation.SetDisplayName(s)
	return ohu
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableDisplayName(s *string) *OrganizationHistoryUpdate {
	if s != nil {
		ohu.SetDisplayName(*s)
	}
	return ohu
}

// SetDescription sets the "description" field.
func (ohu *OrganizationHistoryUpdate) SetDescription(s string) *OrganizationHistoryUpdate {
	ohu.mutation.SetDescription(s)
	return ohu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableDescription(s *string) *OrganizationHistoryUpdate {
	if s != nil {
		ohu.SetDescription(*s)
	}
	return ohu
}

// ClearDescription clears the value of the "description" field.
func (ohu *OrganizationHistoryUpdate) ClearDescription() *OrganizationHistoryUpdate {
	ohu.mutation.ClearDescription()
	return ohu
}

// SetAvatarRemoteURL sets the "avatar_remote_url" field.
func (ohu *OrganizationHistoryUpdate) SetAvatarRemoteURL(s string) *OrganizationHistoryUpdate {
	ohu.mutation.SetAvatarRemoteURL(s)
	return ohu
}

// SetNillableAvatarRemoteURL sets the "avatar_remote_url" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableAvatarRemoteURL(s *string) *OrganizationHistoryUpdate {
	if s != nil {
		ohu.SetAvatarRemoteURL(*s)
	}
	return ohu
}

// ClearAvatarRemoteURL clears the value of the "avatar_remote_url" field.
func (ohu *OrganizationHistoryUpdate) ClearAvatarRemoteURL() *OrganizationHistoryUpdate {
	ohu.mutation.ClearAvatarRemoteURL()
	return ohu
}

// SetDedicatedDb sets the "dedicated_db" field.
func (ohu *OrganizationHistoryUpdate) SetDedicatedDb(b bool) *OrganizationHistoryUpdate {
	ohu.mutation.SetDedicatedDb(b)
	return ohu
}

// SetNillableDedicatedDb sets the "dedicated_db" field if the given value is not nil.
func (ohu *OrganizationHistoryUpdate) SetNillableDedicatedDb(b *bool) *OrganizationHistoryUpdate {
	if b != nil {
		ohu.SetDedicatedDb(*b)
	}
	return ohu
}

// Mutation returns the OrganizationHistoryMutation object of the builder.
func (ohu *OrganizationHistoryUpdate) Mutation() *OrganizationHistoryMutation {
	return ohu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ohu *OrganizationHistoryUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ohu.sqlSave, ohu.mutation, ohu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ohu *OrganizationHistoryUpdate) SaveX(ctx context.Context) int {
	affected, err := ohu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ohu *OrganizationHistoryUpdate) Exec(ctx context.Context) error {
	_, err := ohu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohu *OrganizationHistoryUpdate) ExecX(ctx context.Context) {
	if err := ohu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ohu *OrganizationHistoryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(organizationhistory.Table, organizationhistory.Columns, sqlgraph.NewFieldSpec(organizationhistory.FieldID, field.TypeString))
	if ps := ohu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ohu.mutation.RefCleared() {
		_spec.ClearField(organizationhistory.FieldRef, field.TypeString)
	}
	if ohu.mutation.CreatedAtCleared() {
		_spec.ClearField(organizationhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ohu.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ohu.mutation.UpdatedAtCleared() {
		_spec.ClearField(organizationhistory.FieldUpdatedAt, field.TypeTime)
	}
	if ohu.mutation.CreatedByCleared() {
		_spec.ClearField(organizationhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ohu.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ohu.mutation.UpdatedByCleared() {
		_spec.ClearField(organizationhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ohu.mutation.Tags(); ok {
		_spec.SetField(organizationhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := ohu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationhistory.FieldTags, value)
		})
	}
	if ohu.mutation.TagsCleared() {
		_spec.ClearField(organizationhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := ohu.mutation.DeletedAt(); ok {
		_spec.SetField(organizationhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ohu.mutation.DeletedAtCleared() {
		_spec.ClearField(organizationhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ohu.mutation.DeletedBy(); ok {
		_spec.SetField(organizationhistory.FieldDeletedBy, field.TypeString, value)
	}
	if ohu.mutation.DeletedByCleared() {
		_spec.ClearField(organizationhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ohu.mutation.Name(); ok {
		_spec.SetField(organizationhistory.FieldName, field.TypeString, value)
	}
	if value, ok := ohu.mutation.DisplayName(); ok {
		_spec.SetField(organizationhistory.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := ohu.mutation.Description(); ok {
		_spec.SetField(organizationhistory.FieldDescription, field.TypeString, value)
	}
	if ohu.mutation.DescriptionCleared() {
		_spec.ClearField(organizationhistory.FieldDescription, field.TypeString)
	}
	if ohu.mutation.ParentOrganizationIDCleared() {
		_spec.ClearField(organizationhistory.FieldParentOrganizationID, field.TypeString)
	}
	if ohu.mutation.PersonalOrgCleared() {
		_spec.ClearField(organizationhistory.FieldPersonalOrg, field.TypeBool)
	}
	if value, ok := ohu.mutation.AvatarRemoteURL(); ok {
		_spec.SetField(organizationhistory.FieldAvatarRemoteURL, field.TypeString, value)
	}
	if ohu.mutation.AvatarRemoteURLCleared() {
		_spec.ClearField(organizationhistory.FieldAvatarRemoteURL, field.TypeString)
	}
	if value, ok := ohu.mutation.DedicatedDb(); ok {
		_spec.SetField(organizationhistory.FieldDedicatedDb, field.TypeBool, value)
	}
	_spec.Node.Schema = ohu.schemaConfig.OrganizationHistory
	ctx = internal.NewSchemaConfigContext(ctx, ohu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, ohu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ohu.mutation.done = true
	return n, nil
}

// OrganizationHistoryUpdateOne is the builder for updating a single OrganizationHistory entity.
type OrganizationHistoryUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationHistoryMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ohuo *OrganizationHistoryUpdateOne) SetUpdatedAt(t time.Time) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetUpdatedAt(t)
	return ohuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableUpdatedAt(t *time.Time) *OrganizationHistoryUpdateOne {
	if t != nil {
		ohuo.SetUpdatedAt(*t)
	}
	return ohuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (ohuo *OrganizationHistoryUpdateOne) ClearUpdatedAt() *OrganizationHistoryUpdateOne {
	ohuo.mutation.ClearUpdatedAt()
	return ohuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ohuo *OrganizationHistoryUpdateOne) SetUpdatedBy(s string) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetUpdatedBy(s)
	return ohuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableUpdatedBy(s *string) *OrganizationHistoryUpdateOne {
	if s != nil {
		ohuo.SetUpdatedBy(*s)
	}
	return ohuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (ohuo *OrganizationHistoryUpdateOne) ClearUpdatedBy() *OrganizationHistoryUpdateOne {
	ohuo.mutation.ClearUpdatedBy()
	return ohuo
}

// SetTags sets the "tags" field.
func (ohuo *OrganizationHistoryUpdateOne) SetTags(s []string) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetTags(s)
	return ohuo
}

// AppendTags appends s to the "tags" field.
func (ohuo *OrganizationHistoryUpdateOne) AppendTags(s []string) *OrganizationHistoryUpdateOne {
	ohuo.mutation.AppendTags(s)
	return ohuo
}

// ClearTags clears the value of the "tags" field.
func (ohuo *OrganizationHistoryUpdateOne) ClearTags() *OrganizationHistoryUpdateOne {
	ohuo.mutation.ClearTags()
	return ohuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ohuo *OrganizationHistoryUpdateOne) SetDeletedAt(t time.Time) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetDeletedAt(t)
	return ohuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableDeletedAt(t *time.Time) *OrganizationHistoryUpdateOne {
	if t != nil {
		ohuo.SetDeletedAt(*t)
	}
	return ohuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (ohuo *OrganizationHistoryUpdateOne) ClearDeletedAt() *OrganizationHistoryUpdateOne {
	ohuo.mutation.ClearDeletedAt()
	return ohuo
}

// SetDeletedBy sets the "deleted_by" field.
func (ohuo *OrganizationHistoryUpdateOne) SetDeletedBy(s string) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetDeletedBy(s)
	return ohuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableDeletedBy(s *string) *OrganizationHistoryUpdateOne {
	if s != nil {
		ohuo.SetDeletedBy(*s)
	}
	return ohuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (ohuo *OrganizationHistoryUpdateOne) ClearDeletedBy() *OrganizationHistoryUpdateOne {
	ohuo.mutation.ClearDeletedBy()
	return ohuo
}

// SetName sets the "name" field.
func (ohuo *OrganizationHistoryUpdateOne) SetName(s string) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetName(s)
	return ohuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableName(s *string) *OrganizationHistoryUpdateOne {
	if s != nil {
		ohuo.SetName(*s)
	}
	return ohuo
}

// SetDisplayName sets the "display_name" field.
func (ohuo *OrganizationHistoryUpdateOne) SetDisplayName(s string) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetDisplayName(s)
	return ohuo
}

// SetNillableDisplayName sets the "display_name" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableDisplayName(s *string) *OrganizationHistoryUpdateOne {
	if s != nil {
		ohuo.SetDisplayName(*s)
	}
	return ohuo
}

// SetDescription sets the "description" field.
func (ohuo *OrganizationHistoryUpdateOne) SetDescription(s string) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetDescription(s)
	return ohuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableDescription(s *string) *OrganizationHistoryUpdateOne {
	if s != nil {
		ohuo.SetDescription(*s)
	}
	return ohuo
}

// ClearDescription clears the value of the "description" field.
func (ohuo *OrganizationHistoryUpdateOne) ClearDescription() *OrganizationHistoryUpdateOne {
	ohuo.mutation.ClearDescription()
	return ohuo
}

// SetAvatarRemoteURL sets the "avatar_remote_url" field.
func (ohuo *OrganizationHistoryUpdateOne) SetAvatarRemoteURL(s string) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetAvatarRemoteURL(s)
	return ohuo
}

// SetNillableAvatarRemoteURL sets the "avatar_remote_url" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableAvatarRemoteURL(s *string) *OrganizationHistoryUpdateOne {
	if s != nil {
		ohuo.SetAvatarRemoteURL(*s)
	}
	return ohuo
}

// ClearAvatarRemoteURL clears the value of the "avatar_remote_url" field.
func (ohuo *OrganizationHistoryUpdateOne) ClearAvatarRemoteURL() *OrganizationHistoryUpdateOne {
	ohuo.mutation.ClearAvatarRemoteURL()
	return ohuo
}

// SetDedicatedDb sets the "dedicated_db" field.
func (ohuo *OrganizationHistoryUpdateOne) SetDedicatedDb(b bool) *OrganizationHistoryUpdateOne {
	ohuo.mutation.SetDedicatedDb(b)
	return ohuo
}

// SetNillableDedicatedDb sets the "dedicated_db" field if the given value is not nil.
func (ohuo *OrganizationHistoryUpdateOne) SetNillableDedicatedDb(b *bool) *OrganizationHistoryUpdateOne {
	if b != nil {
		ohuo.SetDedicatedDb(*b)
	}
	return ohuo
}

// Mutation returns the OrganizationHistoryMutation object of the builder.
func (ohuo *OrganizationHistoryUpdateOne) Mutation() *OrganizationHistoryMutation {
	return ohuo.mutation
}

// Where appends a list predicates to the OrganizationHistoryUpdate builder.
func (ohuo *OrganizationHistoryUpdateOne) Where(ps ...predicate.OrganizationHistory) *OrganizationHistoryUpdateOne {
	ohuo.mutation.Where(ps...)
	return ohuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ohuo *OrganizationHistoryUpdateOne) Select(field string, fields ...string) *OrganizationHistoryUpdateOne {
	ohuo.fields = append([]string{field}, fields...)
	return ohuo
}

// Save executes the query and returns the updated OrganizationHistory entity.
func (ohuo *OrganizationHistoryUpdateOne) Save(ctx context.Context) (*OrganizationHistory, error) {
	return withHooks(ctx, ohuo.sqlSave, ohuo.mutation, ohuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ohuo *OrganizationHistoryUpdateOne) SaveX(ctx context.Context) *OrganizationHistory {
	node, err := ohuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ohuo *OrganizationHistoryUpdateOne) Exec(ctx context.Context) error {
	_, err := ohuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ohuo *OrganizationHistoryUpdateOne) ExecX(ctx context.Context) {
	if err := ohuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ohuo *OrganizationHistoryUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationHistory, err error) {
	_spec := sqlgraph.NewUpdateSpec(organizationhistory.Table, organizationhistory.Columns, sqlgraph.NewFieldSpec(organizationhistory.FieldID, field.TypeString))
	id, ok := ohuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "OrganizationHistory.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ohuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationhistory.FieldID)
		for _, f := range fields {
			if !organizationhistory.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != organizationhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ohuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if ohuo.mutation.RefCleared() {
		_spec.ClearField(organizationhistory.FieldRef, field.TypeString)
	}
	if ohuo.mutation.CreatedAtCleared() {
		_spec.ClearField(organizationhistory.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := ohuo.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationhistory.FieldUpdatedAt, field.TypeTime, value)
	}
	if ohuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(organizationhistory.FieldUpdatedAt, field.TypeTime)
	}
	if ohuo.mutation.CreatedByCleared() {
		_spec.ClearField(organizationhistory.FieldCreatedBy, field.TypeString)
	}
	if value, ok := ohuo.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationhistory.FieldUpdatedBy, field.TypeString, value)
	}
	if ohuo.mutation.UpdatedByCleared() {
		_spec.ClearField(organizationhistory.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := ohuo.mutation.Tags(); ok {
		_spec.SetField(organizationhistory.FieldTags, field.TypeJSON, value)
	}
	if value, ok := ohuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationhistory.FieldTags, value)
		})
	}
	if ohuo.mutation.TagsCleared() {
		_spec.ClearField(organizationhistory.FieldTags, field.TypeJSON)
	}
	if value, ok := ohuo.mutation.DeletedAt(); ok {
		_spec.SetField(organizationhistory.FieldDeletedAt, field.TypeTime, value)
	}
	if ohuo.mutation.DeletedAtCleared() {
		_spec.ClearField(organizationhistory.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := ohuo.mutation.DeletedBy(); ok {
		_spec.SetField(organizationhistory.FieldDeletedBy, field.TypeString, value)
	}
	if ohuo.mutation.DeletedByCleared() {
		_spec.ClearField(organizationhistory.FieldDeletedBy, field.TypeString)
	}
	if value, ok := ohuo.mutation.Name(); ok {
		_spec.SetField(organizationhistory.FieldName, field.TypeString, value)
	}
	if value, ok := ohuo.mutation.DisplayName(); ok {
		_spec.SetField(organizationhistory.FieldDisplayName, field.TypeString, value)
	}
	if value, ok := ohuo.mutation.Description(); ok {
		_spec.SetField(organizationhistory.FieldDescription, field.TypeString, value)
	}
	if ohuo.mutation.DescriptionCleared() {
		_spec.ClearField(organizationhistory.FieldDescription, field.TypeString)
	}
	if ohuo.mutation.ParentOrganizationIDCleared() {
		_spec.ClearField(organizationhistory.FieldParentOrganizationID, field.TypeString)
	}
	if ohuo.mutation.PersonalOrgCleared() {
		_spec.ClearField(organizationhistory.FieldPersonalOrg, field.TypeBool)
	}
	if value, ok := ohuo.mutation.AvatarRemoteURL(); ok {
		_spec.SetField(organizationhistory.FieldAvatarRemoteURL, field.TypeString, value)
	}
	if ohuo.mutation.AvatarRemoteURLCleared() {
		_spec.ClearField(organizationhistory.FieldAvatarRemoteURL, field.TypeString)
	}
	if value, ok := ohuo.mutation.DedicatedDb(); ok {
		_spec.SetField(organizationhistory.FieldDedicatedDb, field.TypeBool, value)
	}
	_spec.Node.Schema = ohuo.schemaConfig.OrganizationHistory
	ctx = internal.NewSchemaConfigContext(ctx, ohuo.schemaConfig)
	_node = &OrganizationHistory{config: ohuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ohuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationhistory.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ohuo.mutation.done = true
	return _node, nil
}
