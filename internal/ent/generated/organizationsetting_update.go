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
	"github.com/datumforge/datum/internal/ent/generated/organization"
	"github.com/datumforge/datum/internal/ent/generated/organizationsetting"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// OrganizationSettingUpdate is the builder for updating OrganizationSetting entities.
type OrganizationSettingUpdate struct {
	config
	hooks    []Hook
	mutation *OrganizationSettingMutation
}

// Where appends a list predicates to the OrganizationSettingUpdate builder.
func (osu *OrganizationSettingUpdate) Where(ps ...predicate.OrganizationSetting) *OrganizationSettingUpdate {
	osu.mutation.Where(ps...)
	return osu
}

// SetUpdatedAt sets the "updated_at" field.
func (osu *OrganizationSettingUpdate) SetUpdatedAt(t time.Time) *OrganizationSettingUpdate {
	osu.mutation.SetUpdatedAt(t)
	return osu
}

// SetUpdatedBy sets the "updated_by" field.
func (osu *OrganizationSettingUpdate) SetUpdatedBy(s string) *OrganizationSettingUpdate {
	osu.mutation.SetUpdatedBy(s)
	return osu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (osu *OrganizationSettingUpdate) SetNillableUpdatedBy(s *string) *OrganizationSettingUpdate {
	if s != nil {
		osu.SetUpdatedBy(*s)
	}
	return osu
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (osu *OrganizationSettingUpdate) ClearUpdatedBy() *OrganizationSettingUpdate {
	osu.mutation.ClearUpdatedBy()
	return osu
}

// SetDomains sets the "domains" field.
func (osu *OrganizationSettingUpdate) SetDomains(s []string) *OrganizationSettingUpdate {
	osu.mutation.SetDomains(s)
	return osu
}

// AppendDomains appends s to the "domains" field.
func (osu *OrganizationSettingUpdate) AppendDomains(s []string) *OrganizationSettingUpdate {
	osu.mutation.AppendDomains(s)
	return osu
}

// SetSSOCert sets the "sso_cert" field.
func (osu *OrganizationSettingUpdate) SetSSOCert(s string) *OrganizationSettingUpdate {
	osu.mutation.SetSSOCert(s)
	return osu
}

// SetNillableSSOCert sets the "sso_cert" field if the given value is not nil.
func (osu *OrganizationSettingUpdate) SetNillableSSOCert(s *string) *OrganizationSettingUpdate {
	if s != nil {
		osu.SetSSOCert(*s)
	}
	return osu
}

// SetSSOEntrypoint sets the "sso_entrypoint" field.
func (osu *OrganizationSettingUpdate) SetSSOEntrypoint(s string) *OrganizationSettingUpdate {
	osu.mutation.SetSSOEntrypoint(s)
	return osu
}

// SetNillableSSOEntrypoint sets the "sso_entrypoint" field if the given value is not nil.
func (osu *OrganizationSettingUpdate) SetNillableSSOEntrypoint(s *string) *OrganizationSettingUpdate {
	if s != nil {
		osu.SetSSOEntrypoint(*s)
	}
	return osu
}

// SetSSOIssuer sets the "sso_issuer" field.
func (osu *OrganizationSettingUpdate) SetSSOIssuer(s string) *OrganizationSettingUpdate {
	osu.mutation.SetSSOIssuer(s)
	return osu
}

// SetNillableSSOIssuer sets the "sso_issuer" field if the given value is not nil.
func (osu *OrganizationSettingUpdate) SetNillableSSOIssuer(s *string) *OrganizationSettingUpdate {
	if s != nil {
		osu.SetSSOIssuer(*s)
	}
	return osu
}

// SetBillingContact sets the "billing_contact" field.
func (osu *OrganizationSettingUpdate) SetBillingContact(s string) *OrganizationSettingUpdate {
	osu.mutation.SetBillingContact(s)
	return osu
}

// SetNillableBillingContact sets the "billing_contact" field if the given value is not nil.
func (osu *OrganizationSettingUpdate) SetNillableBillingContact(s *string) *OrganizationSettingUpdate {
	if s != nil {
		osu.SetBillingContact(*s)
	}
	return osu
}

// SetBillingEmail sets the "billing_email" field.
func (osu *OrganizationSettingUpdate) SetBillingEmail(s string) *OrganizationSettingUpdate {
	osu.mutation.SetBillingEmail(s)
	return osu
}

// SetNillableBillingEmail sets the "billing_email" field if the given value is not nil.
func (osu *OrganizationSettingUpdate) SetNillableBillingEmail(s *string) *OrganizationSettingUpdate {
	if s != nil {
		osu.SetBillingEmail(*s)
	}
	return osu
}

// SetBillingPhone sets the "billing_phone" field.
func (osu *OrganizationSettingUpdate) SetBillingPhone(s string) *OrganizationSettingUpdate {
	osu.mutation.SetBillingPhone(s)
	return osu
}

// SetNillableBillingPhone sets the "billing_phone" field if the given value is not nil.
func (osu *OrganizationSettingUpdate) SetNillableBillingPhone(s *string) *OrganizationSettingUpdate {
	if s != nil {
		osu.SetBillingPhone(*s)
	}
	return osu
}

// SetBillingAddress sets the "billing_address" field.
func (osu *OrganizationSettingUpdate) SetBillingAddress(s string) *OrganizationSettingUpdate {
	osu.mutation.SetBillingAddress(s)
	return osu
}

// SetNillableBillingAddress sets the "billing_address" field if the given value is not nil.
func (osu *OrganizationSettingUpdate) SetNillableBillingAddress(s *string) *OrganizationSettingUpdate {
	if s != nil {
		osu.SetBillingAddress(*s)
	}
	return osu
}

// SetTaxIdentifier sets the "tax_identifier" field.
func (osu *OrganizationSettingUpdate) SetTaxIdentifier(s string) *OrganizationSettingUpdate {
	osu.mutation.SetTaxIdentifier(s)
	return osu
}

// SetNillableTaxIdentifier sets the "tax_identifier" field if the given value is not nil.
func (osu *OrganizationSettingUpdate) SetNillableTaxIdentifier(s *string) *OrganizationSettingUpdate {
	if s != nil {
		osu.SetTaxIdentifier(*s)
	}
	return osu
}

// SetTags sets the "tags" field.
func (osu *OrganizationSettingUpdate) SetTags(s []string) *OrganizationSettingUpdate {
	osu.mutation.SetTags(s)
	return osu
}

// AppendTags appends s to the "tags" field.
func (osu *OrganizationSettingUpdate) AppendTags(s []string) *OrganizationSettingUpdate {
	osu.mutation.AppendTags(s)
	return osu
}

// ClearTags clears the value of the "tags" field.
func (osu *OrganizationSettingUpdate) ClearTags() *OrganizationSettingUpdate {
	osu.mutation.ClearTags()
	return osu
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (osu *OrganizationSettingUpdate) SetOrganizationID(id string) *OrganizationSettingUpdate {
	osu.mutation.SetOrganizationID(id)
	return osu
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (osu *OrganizationSettingUpdate) SetNillableOrganizationID(id *string) *OrganizationSettingUpdate {
	if id != nil {
		osu = osu.SetOrganizationID(*id)
	}
	return osu
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (osu *OrganizationSettingUpdate) SetOrganization(o *Organization) *OrganizationSettingUpdate {
	return osu.SetOrganizationID(o.ID)
}

// Mutation returns the OrganizationSettingMutation object of the builder.
func (osu *OrganizationSettingUpdate) Mutation() *OrganizationSettingMutation {
	return osu.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (osu *OrganizationSettingUpdate) ClearOrganization() *OrganizationSettingUpdate {
	osu.mutation.ClearOrganization()
	return osu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (osu *OrganizationSettingUpdate) Save(ctx context.Context) (int, error) {
	if err := osu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, osu.sqlSave, osu.mutation, osu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (osu *OrganizationSettingUpdate) SaveX(ctx context.Context) int {
	affected, err := osu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (osu *OrganizationSettingUpdate) Exec(ctx context.Context) error {
	_, err := osu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osu *OrganizationSettingUpdate) ExecX(ctx context.Context) {
	if err := osu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osu *OrganizationSettingUpdate) defaults() error {
	if _, ok := osu.mutation.UpdatedAt(); !ok {
		if organizationsetting.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized organizationsetting.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := organizationsetting.UpdateDefaultUpdatedAt()
		osu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (osu *OrganizationSettingUpdate) check() error {
	if v, ok := osu.mutation.BillingContact(); ok {
		if err := organizationsetting.BillingContactValidator(v); err != nil {
			return &ValidationError{Name: "billing_contact", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.billing_contact": %w`, err)}
		}
	}
	if v, ok := osu.mutation.BillingEmail(); ok {
		if err := organizationsetting.BillingEmailValidator(v); err != nil {
			return &ValidationError{Name: "billing_email", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.billing_email": %w`, err)}
		}
	}
	if v, ok := osu.mutation.BillingPhone(); ok {
		if err := organizationsetting.BillingPhoneValidator(v); err != nil {
			return &ValidationError{Name: "billing_phone", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.billing_phone": %w`, err)}
		}
	}
	if v, ok := osu.mutation.BillingAddress(); ok {
		if err := organizationsetting.BillingAddressValidator(v); err != nil {
			return &ValidationError{Name: "billing_address", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.billing_address": %w`, err)}
		}
	}
	return nil
}

func (osu *OrganizationSettingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := osu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(organizationsetting.Table, organizationsetting.Columns, sqlgraph.NewFieldSpec(organizationsetting.FieldID, field.TypeString))
	if ps := osu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osu.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationsetting.FieldUpdatedAt, field.TypeTime, value)
	}
	if osu.mutation.CreatedByCleared() {
		_spec.ClearField(organizationsetting.FieldCreatedBy, field.TypeString)
	}
	if value, ok := osu.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationsetting.FieldUpdatedBy, field.TypeString, value)
	}
	if osu.mutation.UpdatedByCleared() {
		_spec.ClearField(organizationsetting.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := osu.mutation.Domains(); ok {
		_spec.SetField(organizationsetting.FieldDomains, field.TypeJSON, value)
	}
	if value, ok := osu.mutation.AppendedDomains(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationsetting.FieldDomains, value)
		})
	}
	if value, ok := osu.mutation.SSOCert(); ok {
		_spec.SetField(organizationsetting.FieldSSOCert, field.TypeString, value)
	}
	if value, ok := osu.mutation.SSOEntrypoint(); ok {
		_spec.SetField(organizationsetting.FieldSSOEntrypoint, field.TypeString, value)
	}
	if value, ok := osu.mutation.SSOIssuer(); ok {
		_spec.SetField(organizationsetting.FieldSSOIssuer, field.TypeString, value)
	}
	if value, ok := osu.mutation.BillingContact(); ok {
		_spec.SetField(organizationsetting.FieldBillingContact, field.TypeString, value)
	}
	if value, ok := osu.mutation.BillingEmail(); ok {
		_spec.SetField(organizationsetting.FieldBillingEmail, field.TypeString, value)
	}
	if value, ok := osu.mutation.BillingPhone(); ok {
		_spec.SetField(organizationsetting.FieldBillingPhone, field.TypeString, value)
	}
	if value, ok := osu.mutation.BillingAddress(); ok {
		_spec.SetField(organizationsetting.FieldBillingAddress, field.TypeString, value)
	}
	if value, ok := osu.mutation.TaxIdentifier(); ok {
		_spec.SetField(organizationsetting.FieldTaxIdentifier, field.TypeString, value)
	}
	if value, ok := osu.mutation.Tags(); ok {
		_spec.SetField(organizationsetting.FieldTags, field.TypeJSON, value)
	}
	if value, ok := osu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationsetting.FieldTags, value)
		})
	}
	if osu.mutation.TagsCleared() {
		_spec.ClearField(organizationsetting.FieldTags, field.TypeJSON)
	}
	if osu.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   organizationsetting.OrganizationTable,
			Columns: []string{organizationsetting.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = osu.schemaConfig.OrganizationSetting
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osu.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   organizationsetting.OrganizationTable,
			Columns: []string{organizationsetting.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = osu.schemaConfig.OrganizationSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = osu.schemaConfig.OrganizationSetting
	ctx = internal.NewSchemaConfigContext(ctx, osu.schemaConfig)
	if n, err = sqlgraph.UpdateNodes(ctx, osu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	osu.mutation.done = true
	return n, nil
}

// OrganizationSettingUpdateOne is the builder for updating a single OrganizationSetting entity.
type OrganizationSettingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OrganizationSettingMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (osuo *OrganizationSettingUpdateOne) SetUpdatedAt(t time.Time) *OrganizationSettingUpdateOne {
	osuo.mutation.SetUpdatedAt(t)
	return osuo
}

// SetUpdatedBy sets the "updated_by" field.
func (osuo *OrganizationSettingUpdateOne) SetUpdatedBy(s string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetUpdatedBy(s)
	return osuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (osuo *OrganizationSettingUpdateOne) SetNillableUpdatedBy(s *string) *OrganizationSettingUpdateOne {
	if s != nil {
		osuo.SetUpdatedBy(*s)
	}
	return osuo
}

// ClearUpdatedBy clears the value of the "updated_by" field.
func (osuo *OrganizationSettingUpdateOne) ClearUpdatedBy() *OrganizationSettingUpdateOne {
	osuo.mutation.ClearUpdatedBy()
	return osuo
}

// SetDomains sets the "domains" field.
func (osuo *OrganizationSettingUpdateOne) SetDomains(s []string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetDomains(s)
	return osuo
}

// AppendDomains appends s to the "domains" field.
func (osuo *OrganizationSettingUpdateOne) AppendDomains(s []string) *OrganizationSettingUpdateOne {
	osuo.mutation.AppendDomains(s)
	return osuo
}

// SetSSOCert sets the "sso_cert" field.
func (osuo *OrganizationSettingUpdateOne) SetSSOCert(s string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetSSOCert(s)
	return osuo
}

// SetNillableSSOCert sets the "sso_cert" field if the given value is not nil.
func (osuo *OrganizationSettingUpdateOne) SetNillableSSOCert(s *string) *OrganizationSettingUpdateOne {
	if s != nil {
		osuo.SetSSOCert(*s)
	}
	return osuo
}

// SetSSOEntrypoint sets the "sso_entrypoint" field.
func (osuo *OrganizationSettingUpdateOne) SetSSOEntrypoint(s string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetSSOEntrypoint(s)
	return osuo
}

// SetNillableSSOEntrypoint sets the "sso_entrypoint" field if the given value is not nil.
func (osuo *OrganizationSettingUpdateOne) SetNillableSSOEntrypoint(s *string) *OrganizationSettingUpdateOne {
	if s != nil {
		osuo.SetSSOEntrypoint(*s)
	}
	return osuo
}

// SetSSOIssuer sets the "sso_issuer" field.
func (osuo *OrganizationSettingUpdateOne) SetSSOIssuer(s string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetSSOIssuer(s)
	return osuo
}

// SetNillableSSOIssuer sets the "sso_issuer" field if the given value is not nil.
func (osuo *OrganizationSettingUpdateOne) SetNillableSSOIssuer(s *string) *OrganizationSettingUpdateOne {
	if s != nil {
		osuo.SetSSOIssuer(*s)
	}
	return osuo
}

// SetBillingContact sets the "billing_contact" field.
func (osuo *OrganizationSettingUpdateOne) SetBillingContact(s string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetBillingContact(s)
	return osuo
}

// SetNillableBillingContact sets the "billing_contact" field if the given value is not nil.
func (osuo *OrganizationSettingUpdateOne) SetNillableBillingContact(s *string) *OrganizationSettingUpdateOne {
	if s != nil {
		osuo.SetBillingContact(*s)
	}
	return osuo
}

// SetBillingEmail sets the "billing_email" field.
func (osuo *OrganizationSettingUpdateOne) SetBillingEmail(s string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetBillingEmail(s)
	return osuo
}

// SetNillableBillingEmail sets the "billing_email" field if the given value is not nil.
func (osuo *OrganizationSettingUpdateOne) SetNillableBillingEmail(s *string) *OrganizationSettingUpdateOne {
	if s != nil {
		osuo.SetBillingEmail(*s)
	}
	return osuo
}

// SetBillingPhone sets the "billing_phone" field.
func (osuo *OrganizationSettingUpdateOne) SetBillingPhone(s string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetBillingPhone(s)
	return osuo
}

// SetNillableBillingPhone sets the "billing_phone" field if the given value is not nil.
func (osuo *OrganizationSettingUpdateOne) SetNillableBillingPhone(s *string) *OrganizationSettingUpdateOne {
	if s != nil {
		osuo.SetBillingPhone(*s)
	}
	return osuo
}

// SetBillingAddress sets the "billing_address" field.
func (osuo *OrganizationSettingUpdateOne) SetBillingAddress(s string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetBillingAddress(s)
	return osuo
}

// SetNillableBillingAddress sets the "billing_address" field if the given value is not nil.
func (osuo *OrganizationSettingUpdateOne) SetNillableBillingAddress(s *string) *OrganizationSettingUpdateOne {
	if s != nil {
		osuo.SetBillingAddress(*s)
	}
	return osuo
}

// SetTaxIdentifier sets the "tax_identifier" field.
func (osuo *OrganizationSettingUpdateOne) SetTaxIdentifier(s string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetTaxIdentifier(s)
	return osuo
}

// SetNillableTaxIdentifier sets the "tax_identifier" field if the given value is not nil.
func (osuo *OrganizationSettingUpdateOne) SetNillableTaxIdentifier(s *string) *OrganizationSettingUpdateOne {
	if s != nil {
		osuo.SetTaxIdentifier(*s)
	}
	return osuo
}

// SetTags sets the "tags" field.
func (osuo *OrganizationSettingUpdateOne) SetTags(s []string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetTags(s)
	return osuo
}

// AppendTags appends s to the "tags" field.
func (osuo *OrganizationSettingUpdateOne) AppendTags(s []string) *OrganizationSettingUpdateOne {
	osuo.mutation.AppendTags(s)
	return osuo
}

// ClearTags clears the value of the "tags" field.
func (osuo *OrganizationSettingUpdateOne) ClearTags() *OrganizationSettingUpdateOne {
	osuo.mutation.ClearTags()
	return osuo
}

// SetOrganizationID sets the "organization" edge to the Organization entity by ID.
func (osuo *OrganizationSettingUpdateOne) SetOrganizationID(id string) *OrganizationSettingUpdateOne {
	osuo.mutation.SetOrganizationID(id)
	return osuo
}

// SetNillableOrganizationID sets the "organization" edge to the Organization entity by ID if the given value is not nil.
func (osuo *OrganizationSettingUpdateOne) SetNillableOrganizationID(id *string) *OrganizationSettingUpdateOne {
	if id != nil {
		osuo = osuo.SetOrganizationID(*id)
	}
	return osuo
}

// SetOrganization sets the "organization" edge to the Organization entity.
func (osuo *OrganizationSettingUpdateOne) SetOrganization(o *Organization) *OrganizationSettingUpdateOne {
	return osuo.SetOrganizationID(o.ID)
}

// Mutation returns the OrganizationSettingMutation object of the builder.
func (osuo *OrganizationSettingUpdateOne) Mutation() *OrganizationSettingMutation {
	return osuo.mutation
}

// ClearOrganization clears the "organization" edge to the Organization entity.
func (osuo *OrganizationSettingUpdateOne) ClearOrganization() *OrganizationSettingUpdateOne {
	osuo.mutation.ClearOrganization()
	return osuo
}

// Where appends a list predicates to the OrganizationSettingUpdate builder.
func (osuo *OrganizationSettingUpdateOne) Where(ps ...predicate.OrganizationSetting) *OrganizationSettingUpdateOne {
	osuo.mutation.Where(ps...)
	return osuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (osuo *OrganizationSettingUpdateOne) Select(field string, fields ...string) *OrganizationSettingUpdateOne {
	osuo.fields = append([]string{field}, fields...)
	return osuo
}

// Save executes the query and returns the updated OrganizationSetting entity.
func (osuo *OrganizationSettingUpdateOne) Save(ctx context.Context) (*OrganizationSetting, error) {
	if err := osuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, osuo.sqlSave, osuo.mutation, osuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (osuo *OrganizationSettingUpdateOne) SaveX(ctx context.Context) *OrganizationSetting {
	node, err := osuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (osuo *OrganizationSettingUpdateOne) Exec(ctx context.Context) error {
	_, err := osuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (osuo *OrganizationSettingUpdateOne) ExecX(ctx context.Context) {
	if err := osuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (osuo *OrganizationSettingUpdateOne) defaults() error {
	if _, ok := osuo.mutation.UpdatedAt(); !ok {
		if organizationsetting.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized organizationsetting.UpdateDefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := organizationsetting.UpdateDefaultUpdatedAt()
		osuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (osuo *OrganizationSettingUpdateOne) check() error {
	if v, ok := osuo.mutation.BillingContact(); ok {
		if err := organizationsetting.BillingContactValidator(v); err != nil {
			return &ValidationError{Name: "billing_contact", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.billing_contact": %w`, err)}
		}
	}
	if v, ok := osuo.mutation.BillingEmail(); ok {
		if err := organizationsetting.BillingEmailValidator(v); err != nil {
			return &ValidationError{Name: "billing_email", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.billing_email": %w`, err)}
		}
	}
	if v, ok := osuo.mutation.BillingPhone(); ok {
		if err := organizationsetting.BillingPhoneValidator(v); err != nil {
			return &ValidationError{Name: "billing_phone", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.billing_phone": %w`, err)}
		}
	}
	if v, ok := osuo.mutation.BillingAddress(); ok {
		if err := organizationsetting.BillingAddressValidator(v); err != nil {
			return &ValidationError{Name: "billing_address", err: fmt.Errorf(`generated: validator failed for field "OrganizationSetting.billing_address": %w`, err)}
		}
	}
	return nil
}

func (osuo *OrganizationSettingUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationSetting, err error) {
	if err := osuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(organizationsetting.Table, organizationsetting.Columns, sqlgraph.NewFieldSpec(organizationsetting.FieldID, field.TypeString))
	id, ok := osuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "OrganizationSetting.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := osuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationsetting.FieldID)
		for _, f := range fields {
			if !organizationsetting.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != organizationsetting.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := osuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := osuo.mutation.UpdatedAt(); ok {
		_spec.SetField(organizationsetting.FieldUpdatedAt, field.TypeTime, value)
	}
	if osuo.mutation.CreatedByCleared() {
		_spec.ClearField(organizationsetting.FieldCreatedBy, field.TypeString)
	}
	if value, ok := osuo.mutation.UpdatedBy(); ok {
		_spec.SetField(organizationsetting.FieldUpdatedBy, field.TypeString, value)
	}
	if osuo.mutation.UpdatedByCleared() {
		_spec.ClearField(organizationsetting.FieldUpdatedBy, field.TypeString)
	}
	if value, ok := osuo.mutation.Domains(); ok {
		_spec.SetField(organizationsetting.FieldDomains, field.TypeJSON, value)
	}
	if value, ok := osuo.mutation.AppendedDomains(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationsetting.FieldDomains, value)
		})
	}
	if value, ok := osuo.mutation.SSOCert(); ok {
		_spec.SetField(organizationsetting.FieldSSOCert, field.TypeString, value)
	}
	if value, ok := osuo.mutation.SSOEntrypoint(); ok {
		_spec.SetField(organizationsetting.FieldSSOEntrypoint, field.TypeString, value)
	}
	if value, ok := osuo.mutation.SSOIssuer(); ok {
		_spec.SetField(organizationsetting.FieldSSOIssuer, field.TypeString, value)
	}
	if value, ok := osuo.mutation.BillingContact(); ok {
		_spec.SetField(organizationsetting.FieldBillingContact, field.TypeString, value)
	}
	if value, ok := osuo.mutation.BillingEmail(); ok {
		_spec.SetField(organizationsetting.FieldBillingEmail, field.TypeString, value)
	}
	if value, ok := osuo.mutation.BillingPhone(); ok {
		_spec.SetField(organizationsetting.FieldBillingPhone, field.TypeString, value)
	}
	if value, ok := osuo.mutation.BillingAddress(); ok {
		_spec.SetField(organizationsetting.FieldBillingAddress, field.TypeString, value)
	}
	if value, ok := osuo.mutation.TaxIdentifier(); ok {
		_spec.SetField(organizationsetting.FieldTaxIdentifier, field.TypeString, value)
	}
	if value, ok := osuo.mutation.Tags(); ok {
		_spec.SetField(organizationsetting.FieldTags, field.TypeJSON, value)
	}
	if value, ok := osuo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, organizationsetting.FieldTags, value)
		})
	}
	if osuo.mutation.TagsCleared() {
		_spec.ClearField(organizationsetting.FieldTags, field.TypeJSON)
	}
	if osuo.mutation.OrganizationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   organizationsetting.OrganizationTable,
			Columns: []string{organizationsetting.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = osuo.schemaConfig.OrganizationSetting
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := osuo.mutation.OrganizationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   organizationsetting.OrganizationTable,
			Columns: []string{organizationsetting.OrganizationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeString),
			},
		}
		edge.Schema = osuo.schemaConfig.OrganizationSetting
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = osuo.schemaConfig.OrganizationSetting
	ctx = internal.NewSchemaConfigContext(ctx, osuo.schemaConfig)
	_node = &OrganizationSetting{config: osuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, osuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationsetting.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	osuo.mutation.done = true
	return _node, nil
}