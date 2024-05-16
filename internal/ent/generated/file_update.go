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
	"github.com/datumforge/datum/internal/ent/generated/file"
	"github.com/datumforge/datum/internal/ent/generated/group"
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/predicate"
	"github.com/datumforge/datum/internal/ent/generated/user"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FileUpdate) SetUpdatedAt(t time.Time) *FileUpdate {
	fu.mutation.SetUpdatedAt(t)
	return fu
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fu *FileUpdate) ClearUpdatedAt() *FileUpdate {
	fu.mutation.ClearUpdatedAt()
	return fu
}

// SetUpdatedBy sets the "updated_by" field.
func (fu *FileUpdate) SetUpdatedBy(s string) *FileUpdate {
	fu.mutation.SetUpdatedBy(s)
	return fu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fu *FileUpdate) SetNillableUpdatedBy(s *string) *FileUpdate {
	if s != nil {
		fu.SetUpdatedBy(*s)
	}
	return fu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (fu *FileUpdate) ClearUpdatedBy() *FileUpdate {
	fu.mutation.ClearUpdatedBy()
	return fu
}

// SetDeletedAt sets the "deleted_at" field.
func (fu *FileUpdate) SetDeletedAt(t time.Time) *FileUpdate {
	fu.mutation.SetDeletedAt(t)
	return fu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fu *FileUpdate) SetNillableDeletedAt(t *time.Time) *FileUpdate {
	if t != nil {
		fu.SetDeletedAt(*t)
	}
	return fu
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fu *FileUpdate) ClearDeletedAt() *FileUpdate {
	fu.mutation.ClearDeletedAt()
	return fu
}

// SetDeletedBy sets the "deleted_by" field.
func (fu *FileUpdate) SetDeletedBy(s string) *FileUpdate {
	fu.mutation.SetDeletedBy(s)
	return fu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (fu *FileUpdate) SetNillableDeletedBy(s *string) *FileUpdate {
	if s != nil {
		fu.SetDeletedBy(*s)
	}
	return fu
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (fu *FileUpdate) ClearDeletedBy() *FileUpdate {
	fu.mutation.ClearDeletedBy()
	return fu
}

// SetTags sets the "tags" field.
func (fu *FileUpdate) SetTags(s []string) *FileUpdate {
	fu.mutation.SetTags(s)
	return fu
}

// AppendTags appends s to the "tags" field.
func (fu *FileUpdate) AppendTags(s []string) *FileUpdate {
	fu.mutation.AppendTags(s)
	return fu
}

// ClearTags clears the value of the "tags" field.
func (fu *FileUpdate) ClearTags() *FileUpdate {
	fu.mutation.ClearTags()
	return fu
}

// SetFileName sets the "file_name" field.
func (fu *FileUpdate) SetFileName(s string) *FileUpdate {
	fu.mutation.SetFileName(s)
	return fu
}

// SetNillableFileName sets the "file_name" field if the given value is not nil.
func (fu *FileUpdate) SetNillableFileName(s *string) *FileUpdate {
	if s != nil {
		fu.SetFileName(*s)
	}
	return fu
}

// SetFileExtension sets the "file_extension" field.
func (fu *FileUpdate) SetFileExtension(s string) *FileUpdate {
	fu.mutation.SetFileExtension(s)
	return fu
}

// SetNillableFileExtension sets the "file_extension" field if the given value is not nil.
func (fu *FileUpdate) SetNillableFileExtension(s *string) *FileUpdate {
	if s != nil {
		fu.SetFileExtension(*s)
	}
	return fu
}

// SetFileSize sets the "file_size" field.
func (fu *FileUpdate) SetFileSize(i int) *FileUpdate {
	fu.mutation.ResetFileSize()
	fu.mutation.SetFileSize(i)
	return fu
}

// SetNillableFileSize sets the "file_size" field if the given value is not nil.
func (fu *FileUpdate) SetNillableFileSize(i *int) *FileUpdate {
	if i != nil {
		fu.SetFileSize(*i)
	}
	return fu
}

// AddFileSize adds i to the "file_size" field.
func (fu *FileUpdate) AddFileSize(i int) *FileUpdate {
	fu.mutation.AddFileSize(i)
	return fu
}

// ClearFileSize clears the value of the "file_size" field.
func (fu *FileUpdate) ClearFileSize() *FileUpdate {
	fu.mutation.ClearFileSize()
	return fu
}

// SetContentType sets the "content_type" field.
func (fu *FileUpdate) SetContentType(s string) *FileUpdate {
	fu.mutation.SetContentType(s)
	return fu
}

// SetNillableContentType sets the "content_type" field if the given value is not nil.
func (fu *FileUpdate) SetNillableContentType(s *string) *FileUpdate {
	if s != nil {
		fu.SetContentType(*s)
	}
	return fu
}

// SetStoreKey sets the "store_key" field.
func (fu *FileUpdate) SetStoreKey(s string) *FileUpdate {
	fu.mutation.SetStoreKey(s)
	return fu
}

// SetNillableStoreKey sets the "store_key" field if the given value is not nil.
func (fu *FileUpdate) SetNillableStoreKey(s *string) *FileUpdate {
	if s != nil {
		fu.SetStoreKey(*s)
	}
	return fu
}

// SetCategory sets the "category" field.
func (fu *FileUpdate) SetCategory(s string) *FileUpdate {
	fu.mutation.SetCategory(s)
	return fu
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (fu *FileUpdate) SetNillableCategory(s *string) *FileUpdate {
	if s != nil {
		fu.SetCategory(*s)
	}
	return fu
}

// ClearCategory clears the value of the "category" field.
func (fu *FileUpdate) ClearCategory() *FileUpdate {
	fu.mutation.ClearCategory()
	return fu
}

// SetAnnotation sets the "annotation" field.
func (fu *FileUpdate) SetAnnotation(s string) *FileUpdate {
	fu.mutation.SetAnnotation(s)
	return fu
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (fu *FileUpdate) SetNillableAnnotation(s *string) *FileUpdate {
	if s != nil {
		fu.SetAnnotation(*s)
	}
	return fu
}

// ClearAnnotation clears the value of the "annotation" field.
func (fu *FileUpdate) ClearAnnotation() *FileUpdate {
	fu.mutation.ClearAnnotation()
	return fu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (fu *FileUpdate) SetUserID(id string) *FileUpdate {
	fu.mutation.SetUserID(id)
	return fu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (fu *FileUpdate) SetNillableUserID(id *string) *FileUpdate {
	if id != nil {
		fu = fu.SetUserID(*id)
	}
	return fu
}

// SetUser sets the "user" edge to the User entity.
func (fu *FileUpdate) SetUser(u *User) *FileUpdate {
	return fu.SetUserID(u.ID)
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (fu *FileUpdate) AddOrganizationIDs(ids ...string) *FileUpdate {
	fu.mutation.AddOrganizationIDs(ids...)
	return fu
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (fu *FileUpdate) AddOrganization(o ...*Organization) *FileUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return fu.AddOrganizationIDs(ids...)
}

// AddGroupIDs adds the "group" edge to the Group entity by IDs.
func (fu *FileUpdate) AddGroupIDs(ids ...string) *FileUpdate {
	fu.mutation.AddGroupIDs(ids...)
	return fu
}

// AddGroup adds the "group" edges to the Group entity.
func (fu *FileUpdate) AddGroup(g ...*Group) *FileUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return fu.AddGroupIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fu *FileUpdate) ClearUser() *FileUpdate {
	fu.mutation.ClearUser()
	return fu
}

// ClearOrganization clears all "organization" edges to the Organization entity.
func (fu *FileUpdate) ClearOrganization() *FileUpdate {
	fu.mutation.ClearOrganization()
	return fu
}

// RemoveOrganizationIDs removes the "organization" edge to Organization entities by IDs.
func (fu *FileUpdate) RemoveOrganizationIDs(ids ...string) *FileUpdate {
	fu.mutation.RemoveOrganizationIDs(ids...)
	return fu
}

// RemoveOrganization removes "organization" edges to Organization entities.
func (fu *FileUpdate) RemoveOrganization(o ...*Organization) *FileUpdate {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return fu.RemoveOrganizationIDs(ids...)
}

// ClearGroup clears all "group" edges to the Group entity.
func (fu *FileUpdate) ClearGroup() *FileUpdate {
	fu.mutation.ClearGroup()
	return fu
}

// RemoveGroupIDs removes the "group" edge to Group entities by IDs.
func (fu *FileUpdate) RemoveGroupIDs(ids ...string) *FileUpdate {
	fu.mutation.RemoveGroupIDs(ids...)
	return fu
}

// RemoveGroup removes "group" edges to Group entities.
func (fu *FileUpdate) RemoveGroup(g ...*Group) *FileUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return fu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	if err := fu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FileUpdate) defaults() error {
	if _, ok := fu.mutation.UpdatedAt(); !ok && !fu.mutation.UpdatedAtCleared() {
		if file.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized file.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := file.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (fu *FileUpdate) check() error {
	if v, ok := fu.mutation.FileSize(); ok {
		if err := file.FileSizeValidator(v); err != nil {
			return &ValidationError{Name: "file_size", err: fmt.Errorf(`generated: validator failed for field "File.file_size": %w`, err)}
		}
	}
	return nil
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeString))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fu.mutation.CreatedAtCleared() {
		_spec.ClearField(file.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
	}
	if fu.mutation.UpdatedAtCleared() {
		_spec.ClearField(file.FieldUpdatedAt, field.TypeTime)
	}
	if fu.mutation.CreatedByCleared() {
		_spec.ClearField(file.FieldCreatedBy, field.TypeString)
	}
	if value, ok := fu.mutation.UpdatedBy(); ok {
		_spec.SetField(file.FieldUpdatedBy, field.TypeString, value)
	}
	if fu.mutation.UpdatedByCleared() {
		_spec.ClearField(file.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := fu.mutation.DeletedAt(); ok {
		_spec.SetField(file.FieldDeletedAt, field.TypeTime, value)
	}
	if fu.mutation.DeletedAtCleared() {
		_spec.ClearField(file.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := fu.mutation.DeletedBy(); ok {
		_spec.SetField(file.FieldDeletedBy, field.TypeString, value)
	}
	if fu.mutation.DeletedByCleared() {
		_spec.ClearField(file.FieldDeletedBy, field.TypeString)
	}
	if value, ok := fu.mutation.Tags(); ok {
		_spec.SetField(file.FieldTags, field.TypeJSON, value)
	}
	if value, ok := fu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, file.FieldTags, value)
		})
	}
	if fu.mutation.TagsCleared() {
		_spec.ClearField(file.FieldTags, field.TypeJSON)
	}
	if value, ok := fu.mutation.FileName(); ok {
		_spec.SetField(file.FieldFileName, field.TypeString, value)
	}
	if value, ok := fu.mutation.FileExtension(); ok {
		_spec.SetField(file.FieldFileExtension, field.TypeString, value)
	}
	if value, ok := fu.mutation.FileSize(); ok {
		_spec.SetField(file.FieldFileSize, field.TypeInt, value)
	}
	if value, ok := fu.mutation.AddedFileSize(); ok {
		_spec.AddField(file.FieldFileSize, field.TypeInt, value)
	}
	if fu.mutation.FileSizeCleared() {
		_spec.ClearField(file.FieldFileSize, field.TypeInt)
	}
	if value, ok := fu.mutation.ContentType(); ok {
		_spec.SetField(file.FieldContentType, field.TypeString, value)
	}
	if value, ok := fu.mutation.StoreKey(); ok {
		_spec.SetField(file.FieldStoreKey, field.TypeString, value)
	}
	if value, ok := fu.mutation.Category(); ok {
		_spec.SetField(file.FieldCategory, field.TypeString, value)
	}
	if fu.mutation.CategoryCleared() {
		_spec.ClearField(file.FieldCategory, field.TypeString)
	}
	if value, ok := fu.mutation.Annotation(); ok {
		_spec.SetField(file.FieldAnnotation, field.TypeString, value)
	}
	if fu.mutation.AnnotationCleared() {
		_spec.ClearField(file.FieldAnnotation, field.TypeString)
	}
	if fu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.UserTable,
			Columns: []string{file.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = fu.schemaConfig.File
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.UserTable,
			Columns: []string{file.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = fu.schemaConfig.File
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.OrganizationTable,
			Columns: file.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = fu.schemaConfig.OrganizationFiles
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedOrganizationIDs(); len(nodes) > 0 && !fu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.OrganizationTable,
			Columns: file.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = fu.schemaConfig.OrganizationFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.OrganizationTable,
			Columns: file.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = fu.schemaConfig.OrganizationFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.GroupTable,
			Columns: file.GroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = fu.schemaConfig.GroupFiles
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedGroupIDs(); len(nodes) > 0 && !fu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.GroupTable,
			Columns: file.GroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = fu.schemaConfig.GroupFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.GroupTable,
			Columns: file.GroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = fu.schemaConfig.GroupFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = fu.schemaConfig.File
	ctx = internal.NewSchemaConfigContext(ctx, fu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FileUpdateOne) SetUpdatedAt(t time.Time) *FileUpdateOne {
	fuo.mutation.SetUpdatedAt(t)
	return fuo
}

// ClearUpdatedAt clears the value of the "updated_at" field.
func (fuo *FileUpdateOne) ClearUpdatedAt() *FileUpdateOne {
	fuo.mutation.ClearUpdatedAt()
	return fuo
}

// SetUpdatedBy sets the "updated_by" field.
func (fuo *FileUpdateOne) SetUpdatedBy(s string) *FileUpdateOne {
	fuo.mutation.SetUpdatedBy(s)
	return fuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableUpdatedBy(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetUpdatedBy(*s)
	}
	return fuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (fuo *FileUpdateOne) ClearUpdatedBy() *FileUpdateOne {
	fuo.mutation.ClearUpdatedBy()
	return fuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fuo *FileUpdateOne) SetDeletedAt(t time.Time) *FileUpdateOne {
	fuo.mutation.SetDeletedAt(t)
	return fuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableDeletedAt(t *time.Time) *FileUpdateOne {
	if t != nil {
		fuo.SetDeletedAt(*t)
	}
	return fuo
}

// ClearDeletedAt clears the value of the "deleted_at" field.
func (fuo *FileUpdateOne) ClearDeletedAt() *FileUpdateOne {
	fuo.mutation.ClearDeletedAt()
	return fuo
}

// SetDeletedBy sets the "deleted_by" field.
func (fuo *FileUpdateOne) SetDeletedBy(s string) *FileUpdateOne {
	fuo.mutation.SetDeletedBy(s)
	return fuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableDeletedBy(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetDeletedBy(*s)
	}
	return fuo
}

// ClearDeletedBy clears the value of the "deleted_by" field.
func (fuo *FileUpdateOne) ClearDeletedBy() *FileUpdateOne {
	fuo.mutation.ClearDeletedBy()
	return fuo
}

// SetTags sets the "tags" field.
func (fuo *FileUpdateOne) SetTags(s []string) *FileUpdateOne {
	fuo.mutation.SetTags(s)
	return fuo
}

// AppendTags appends s to the "tags" field.
func (fuo *FileUpdateOne) AppendTags(s []string) *FileUpdateOne {
	fuo.mutation.AppendTags(s)
	return fuo
}

// ClearTags clears the value of the "tags" field.
func (fuo *FileUpdateOne) ClearTags() *FileUpdateOne {
	fuo.mutation.ClearTags()
	return fuo
}

// SetFileName sets the "file_name" field.
func (fuo *FileUpdateOne) SetFileName(s string) *FileUpdateOne {
	fuo.mutation.SetFileName(s)
	return fuo
}

// SetNillableFileName sets the "file_name" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableFileName(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetFileName(*s)
	}
	return fuo
}

// SetFileExtension sets the "file_extension" field.
func (fuo *FileUpdateOne) SetFileExtension(s string) *FileUpdateOne {
	fuo.mutation.SetFileExtension(s)
	return fuo
}

// SetNillableFileExtension sets the "file_extension" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableFileExtension(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetFileExtension(*s)
	}
	return fuo
}

// SetFileSize sets the "file_size" field.
func (fuo *FileUpdateOne) SetFileSize(i int) *FileUpdateOne {
	fuo.mutation.ResetFileSize()
	fuo.mutation.SetFileSize(i)
	return fuo
}

// SetNillableFileSize sets the "file_size" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableFileSize(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetFileSize(*i)
	}
	return fuo
}

// AddFileSize adds i to the "file_size" field.
func (fuo *FileUpdateOne) AddFileSize(i int) *FileUpdateOne {
	fuo.mutation.AddFileSize(i)
	return fuo
}

// ClearFileSize clears the value of the "file_size" field.
func (fuo *FileUpdateOne) ClearFileSize() *FileUpdateOne {
	fuo.mutation.ClearFileSize()
	return fuo
}

// SetContentType sets the "content_type" field.
func (fuo *FileUpdateOne) SetContentType(s string) *FileUpdateOne {
	fuo.mutation.SetContentType(s)
	return fuo
}

// SetNillableContentType sets the "content_type" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableContentType(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetContentType(*s)
	}
	return fuo
}

// SetStoreKey sets the "store_key" field.
func (fuo *FileUpdateOne) SetStoreKey(s string) *FileUpdateOne {
	fuo.mutation.SetStoreKey(s)
	return fuo
}

// SetNillableStoreKey sets the "store_key" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableStoreKey(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetStoreKey(*s)
	}
	return fuo
}

// SetCategory sets the "category" field.
func (fuo *FileUpdateOne) SetCategory(s string) *FileUpdateOne {
	fuo.mutation.SetCategory(s)
	return fuo
}

// SetNillableCategory sets the "category" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableCategory(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetCategory(*s)
	}
	return fuo
}

// ClearCategory clears the value of the "category" field.
func (fuo *FileUpdateOne) ClearCategory() *FileUpdateOne {
	fuo.mutation.ClearCategory()
	return fuo
}

// SetAnnotation sets the "annotation" field.
func (fuo *FileUpdateOne) SetAnnotation(s string) *FileUpdateOne {
	fuo.mutation.SetAnnotation(s)
	return fuo
}

// SetNillableAnnotation sets the "annotation" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableAnnotation(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetAnnotation(*s)
	}
	return fuo
}

// ClearAnnotation clears the value of the "annotation" field.
func (fuo *FileUpdateOne) ClearAnnotation() *FileUpdateOne {
	fuo.mutation.ClearAnnotation()
	return fuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (fuo *FileUpdateOne) SetUserID(id string) *FileUpdateOne {
	fuo.mutation.SetUserID(id)
	return fuo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableUserID(id *string) *FileUpdateOne {
	if id != nil {
		fuo = fuo.SetUserID(*id)
	}
	return fuo
}

// SetUser sets the "user" edge to the User entity.
func (fuo *FileUpdateOne) SetUser(u *User) *FileUpdateOne {
	return fuo.SetUserID(u.ID)
}

// AddOrganizationIDs adds the "organization" edge to the Organization entity by IDs.
func (fuo *FileUpdateOne) AddOrganizationIDs(ids ...string) *FileUpdateOne {
	fuo.mutation.AddOrganizationIDs(ids...)
	return fuo
}

// AddOrganization adds the "organization" edges to the Organization entity.
func (fuo *FileUpdateOne) AddOrganization(o ...*Organization) *FileUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return fuo.AddOrganizationIDs(ids...)
}

// AddGroupIDs adds the "group" edge to the Group entity by IDs.
func (fuo *FileUpdateOne) AddGroupIDs(ids ...string) *FileUpdateOne {
	fuo.mutation.AddGroupIDs(ids...)
	return fuo
}

// AddGroup adds the "group" edges to the Group entity.
func (fuo *FileUpdateOne) AddGroup(g ...*Group) *FileUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return fuo.AddGroupIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (fuo *FileUpdateOne) ClearUser() *FileUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// ClearOrganization clears all "organization" edges to the Organization entity.
func (fuo *FileUpdateOne) ClearOrganization() *FileUpdateOne {
	fuo.mutation.ClearOrganization()
	return fuo
}

// RemoveOrganizationIDs removes the "organization" edge to Organization entities by IDs.
func (fuo *FileUpdateOne) RemoveOrganizationIDs(ids ...string) *FileUpdateOne {
	fuo.mutation.RemoveOrganizationIDs(ids...)
	return fuo
}

// RemoveOrganization removes "organization" edges to Organization entities.
func (fuo *FileUpdateOne) RemoveOrganization(o ...*Organization) *FileUpdateOne {
	ids := make([]string, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return fuo.RemoveOrganizationIDs(ids...)
}

// ClearGroup clears all "group" edges to the Group entity.
func (fuo *FileUpdateOne) ClearGroup() *FileUpdateOne {
	fuo.mutation.ClearGroup()
	return fuo
}

// RemoveGroupIDs removes the "group" edge to Group entities by IDs.
func (fuo *FileUpdateOne) RemoveGroupIDs(ids ...string) *FileUpdateOne {
	fuo.mutation.RemoveGroupIDs(ids...)
	return fuo
}

// RemoveGroup removes "group" edges to Group entities.
func (fuo *FileUpdateOne) RemoveGroup(g ...*Group) *FileUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return fuo.RemoveGroupIDs(ids...)
}

// Where appends a list predicates to the FileUpdate builder.
func (fuo *FileUpdateOne) Where(ps ...predicate.File) *FileUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	if err := fuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FileUpdateOne) defaults() error {
	if _, ok := fuo.mutation.UpdatedAt(); !ok && !fuo.mutation.UpdatedAtCleared() {
		if file.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized file.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := file.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FileUpdateOne) check() error {
	if v, ok := fuo.mutation.FileSize(); ok {
		if err := file.FileSizeValidator(v); err != nil {
			return &ValidationError{Name: "file_size", err: fmt.Errorf(`generated: validator failed for field "File.file_size": %w`, err)}
		}
	}
	return nil
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (_node *File, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeString))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "File.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for _, f := range fields {
			if !file.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if fuo.mutation.CreatedAtCleared() {
		_spec.ClearField(file.FieldCreatedAt, field.TypeTime)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(file.FieldUpdatedAt, field.TypeTime, value)
	}
	if fuo.mutation.UpdatedAtCleared() {
		_spec.ClearField(file.FieldUpdatedAt, field.TypeTime)
	}
	if fuo.mutation.CreatedByCleared() {
		_spec.ClearField(file.FieldCreatedBy, field.TypeString)
	}
	if value, ok := fuo.mutation.UpdatedBy(); ok {
		_spec.SetField(file.FieldUpdatedBy, field.TypeString, value)
	}
	if fuo.mutation.UpdatedByCleared() {
		_spec.ClearField(file.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := fuo.mutation.DeletedAt(); ok {
		_spec.SetField(file.FieldDeletedAt, field.TypeTime, value)
	}
	if fuo.mutation.DeletedAtCleared() {
		_spec.ClearField(file.FieldDeletedAt, field.TypeTime)
	}
	if value, ok := fuo.mutation.DeletedBy(); ok {
		_spec.SetField(file.FieldDeletedBy, field.TypeString, value)
	}
	if fuo.mutation.DeletedByCleared() {
		_spec.ClearField(file.FieldDeletedBy, field.TypeString)
	}
	if value, ok := fuo.mutation.Tags(); ok {
		_spec.SetField(file.FieldTags, field.TypeJSON, value)
	}
	if value, ok := fuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, file.FieldTags, value)
		})
	}
	if fuo.mutation.TagsCleared() {
		_spec.ClearField(file.FieldTags, field.TypeJSON)
	}
	if value, ok := fuo.mutation.FileName(); ok {
		_spec.SetField(file.FieldFileName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.FileExtension(); ok {
		_spec.SetField(file.FieldFileExtension, field.TypeString, value)
	}
	if value, ok := fuo.mutation.FileSize(); ok {
		_spec.SetField(file.FieldFileSize, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.AddedFileSize(); ok {
		_spec.AddField(file.FieldFileSize, field.TypeInt, value)
	}
	if fuo.mutation.FileSizeCleared() {
		_spec.ClearField(file.FieldFileSize, field.TypeInt)
	}
	if value, ok := fuo.mutation.ContentType(); ok {
		_spec.SetField(file.FieldContentType, field.TypeString, value)
	}
	if value, ok := fuo.mutation.StoreKey(); ok {
		_spec.SetField(file.FieldStoreKey, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Category(); ok {
		_spec.SetField(file.FieldCategory, field.TypeString, value)
	}
	if fuo.mutation.CategoryCleared() {
		_spec.ClearField(file.FieldCategory, field.TypeString)
	}
	if value, ok := fuo.mutation.Annotation(); ok {
		_spec.SetField(file.FieldAnnotation, field.TypeString, value)
	}
	if fuo.mutation.AnnotationCleared() {
		_spec.ClearField(file.FieldAnnotation, field.TypeString)
	}
	if fuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.UserTable,
			Columns: []string{file.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = fuo.schemaConfig.File
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.UserTable,
			Columns: []string{file.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		edge.Schema = fuo.schemaConfig.File
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.OrganizationTable,
			Columns: file.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = fuo.schemaConfig.OrganizationFiles
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedOrganizationIDs(); len(nodes) > 0 && !fuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.OrganizationTable,
			Columns: file.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = fuo.schemaConfig.OrganizationFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.OrganizationTable,
			Columns: file.OrganizationPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = fuo.schemaConfig.OrganizationFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.GroupTable,
			Columns: file.GroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = fuo.schemaConfig.GroupFiles
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedGroupIDs(); len(nodes) > 0 && !fuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.GroupTable,
			Columns: file.GroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = fuo.schemaConfig.GroupFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   file.GroupTable,
			Columns: file.GroupPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeString),
			},
		}
		edge.Schema = fuo.schemaConfig.GroupFiles
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = fuo.schemaConfig.File
	ctx = internal.NewSchemaConfigContext(ctx, fuo.schemaConfig)
	_node = &File{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
