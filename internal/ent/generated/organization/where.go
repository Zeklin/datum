// Code generated by ent, DO NOT EDIT.

package organization

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDeletedBy, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// DisplayName applies equality check predicate on the "display_name" field. It's identical to DisplayNameEQ.
func DisplayName(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDisplayName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDescription, v))
}

// ParentOrganizationID applies equality check predicate on the "parent_organization_id" field. It's identical to ParentOrganizationIDEQ.
func ParentOrganizationID(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldParentOrganizationID, v))
}

// PersonalOrg applies equality check predicate on the "personal_org" field. It's identical to PersonalOrgEQ.
func PersonalOrg(v bool) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldPersonalOrg, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v time.Time) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedAtIsNil applies the IsNil predicate on the "deleted_at" field.
func DeletedAtIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldDeletedAt))
}

// DeletedAtNotNil applies the NotNil predicate on the "deleted_at" field.
func DeletedAtNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldDeletedAt))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDeletedBy, v))
}

// DeletedByContains applies the Contains predicate on the "deleted_by" field.
func DeletedByContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldDeletedBy, v))
}

// DeletedByHasPrefix applies the HasPrefix predicate on the "deleted_by" field.
func DeletedByHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldDeletedBy, v))
}

// DeletedByHasSuffix applies the HasSuffix predicate on the "deleted_by" field.
func DeletedByHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldDeletedBy, v))
}

// DeletedByIsNil applies the IsNil predicate on the "deleted_by" field.
func DeletedByIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldDeletedBy))
}

// DeletedByNotNil applies the NotNil predicate on the "deleted_by" field.
func DeletedByNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldDeletedBy))
}

// DeletedByEqualFold applies the EqualFold predicate on the "deleted_by" field.
func DeletedByEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldDeletedBy, v))
}

// DeletedByContainsFold applies the ContainsFold predicate on the "deleted_by" field.
func DeletedByContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldDeletedBy, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldName, v))
}

// DisplayNameEQ applies the EQ predicate on the "display_name" field.
func DisplayNameEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDisplayName, v))
}

// DisplayNameNEQ applies the NEQ predicate on the "display_name" field.
func DisplayNameNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDisplayName, v))
}

// DisplayNameIn applies the In predicate on the "display_name" field.
func DisplayNameIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDisplayName, vs...))
}

// DisplayNameNotIn applies the NotIn predicate on the "display_name" field.
func DisplayNameNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDisplayName, vs...))
}

// DisplayNameGT applies the GT predicate on the "display_name" field.
func DisplayNameGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDisplayName, v))
}

// DisplayNameGTE applies the GTE predicate on the "display_name" field.
func DisplayNameGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDisplayName, v))
}

// DisplayNameLT applies the LT predicate on the "display_name" field.
func DisplayNameLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDisplayName, v))
}

// DisplayNameLTE applies the LTE predicate on the "display_name" field.
func DisplayNameLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDisplayName, v))
}

// DisplayNameContains applies the Contains predicate on the "display_name" field.
func DisplayNameContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldDisplayName, v))
}

// DisplayNameHasPrefix applies the HasPrefix predicate on the "display_name" field.
func DisplayNameHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldDisplayName, v))
}

// DisplayNameHasSuffix applies the HasSuffix predicate on the "display_name" field.
func DisplayNameHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldDisplayName, v))
}

// DisplayNameEqualFold applies the EqualFold predicate on the "display_name" field.
func DisplayNameEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldDisplayName, v))
}

// DisplayNameContainsFold applies the ContainsFold predicate on the "display_name" field.
func DisplayNameContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldDisplayName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldDescription, v))
}

// ParentOrganizationIDEQ applies the EQ predicate on the "parent_organization_id" field.
func ParentOrganizationIDEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldParentOrganizationID, v))
}

// ParentOrganizationIDNEQ applies the NEQ predicate on the "parent_organization_id" field.
func ParentOrganizationIDNEQ(v string) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldParentOrganizationID, v))
}

// ParentOrganizationIDIn applies the In predicate on the "parent_organization_id" field.
func ParentOrganizationIDIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldIn(FieldParentOrganizationID, vs...))
}

// ParentOrganizationIDNotIn applies the NotIn predicate on the "parent_organization_id" field.
func ParentOrganizationIDNotIn(vs ...string) predicate.Organization {
	return predicate.Organization(sql.FieldNotIn(FieldParentOrganizationID, vs...))
}

// ParentOrganizationIDGT applies the GT predicate on the "parent_organization_id" field.
func ParentOrganizationIDGT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGT(FieldParentOrganizationID, v))
}

// ParentOrganizationIDGTE applies the GTE predicate on the "parent_organization_id" field.
func ParentOrganizationIDGTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldGTE(FieldParentOrganizationID, v))
}

// ParentOrganizationIDLT applies the LT predicate on the "parent_organization_id" field.
func ParentOrganizationIDLT(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLT(FieldParentOrganizationID, v))
}

// ParentOrganizationIDLTE applies the LTE predicate on the "parent_organization_id" field.
func ParentOrganizationIDLTE(v string) predicate.Organization {
	return predicate.Organization(sql.FieldLTE(FieldParentOrganizationID, v))
}

// ParentOrganizationIDContains applies the Contains predicate on the "parent_organization_id" field.
func ParentOrganizationIDContains(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContains(FieldParentOrganizationID, v))
}

// ParentOrganizationIDHasPrefix applies the HasPrefix predicate on the "parent_organization_id" field.
func ParentOrganizationIDHasPrefix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasPrefix(FieldParentOrganizationID, v))
}

// ParentOrganizationIDHasSuffix applies the HasSuffix predicate on the "parent_organization_id" field.
func ParentOrganizationIDHasSuffix(v string) predicate.Organization {
	return predicate.Organization(sql.FieldHasSuffix(FieldParentOrganizationID, v))
}

// ParentOrganizationIDIsNil applies the IsNil predicate on the "parent_organization_id" field.
func ParentOrganizationIDIsNil() predicate.Organization {
	return predicate.Organization(sql.FieldIsNull(FieldParentOrganizationID))
}

// ParentOrganizationIDNotNil applies the NotNil predicate on the "parent_organization_id" field.
func ParentOrganizationIDNotNil() predicate.Organization {
	return predicate.Organization(sql.FieldNotNull(FieldParentOrganizationID))
}

// ParentOrganizationIDEqualFold applies the EqualFold predicate on the "parent_organization_id" field.
func ParentOrganizationIDEqualFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldEqualFold(FieldParentOrganizationID, v))
}

// ParentOrganizationIDContainsFold applies the ContainsFold predicate on the "parent_organization_id" field.
func ParentOrganizationIDContainsFold(v string) predicate.Organization {
	return predicate.Organization(sql.FieldContainsFold(FieldParentOrganizationID, v))
}

// PersonalOrgEQ applies the EQ predicate on the "personal_org" field.
func PersonalOrgEQ(v bool) predicate.Organization {
	return predicate.Organization(sql.FieldEQ(FieldPersonalOrg, v))
}

// PersonalOrgNEQ applies the NEQ predicate on the "personal_org" field.
func PersonalOrgNEQ(v bool) predicate.Organization {
	return predicate.Organization(sql.FieldNEQ(FieldPersonalOrg, v))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Organization
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newParentStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Organization
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Organization
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Organization) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newChildrenStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Organization
		step.Edge.Schema = schemaConfig.Organization
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UserOrganizations
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newUsersStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UserOrganizations
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, GroupsTable, GroupsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Group
		step.Edge.Schema = schemaConfig.Group
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newGroupsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Group
		step.Edge.Schema = schemaConfig.Group
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasIntegrations applies the HasEdge predicate on the "integrations" edge.
func HasIntegrations() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, IntegrationsTable, IntegrationsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Integration
		step.Edge.Schema = schemaConfig.Integration
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasIntegrationsWith applies the HasEdge predicate on the "integrations" edge with a given conditions (other predicates).
func HasIntegrationsWith(preds ...predicate.Integration) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newIntegrationsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Integration
		step.Edge.Schema = schemaConfig.Integration
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSetting applies the HasEdge predicate on the "setting" edge.
func HasSetting() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SettingTable, SettingColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.OrganizationSetting
		step.Edge.Schema = schemaConfig.OrganizationSetting
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSettingWith applies the HasEdge predicate on the "setting" edge with a given conditions (other predicates).
func HasSettingWith(preds ...predicate.OrganizationSetting) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newSettingStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.OrganizationSetting
		step.Edge.Schema = schemaConfig.OrganizationSetting
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasEntitlements applies the HasEdge predicate on the "entitlements" edge.
func HasEntitlements() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, EntitlementsTable, EntitlementsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Entitlement
		step.Edge.Schema = schemaConfig.Entitlement
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasEntitlementsWith applies the HasEdge predicate on the "entitlements" edge with a given conditions (other predicates).
func HasEntitlementsWith(preds ...predicate.Entitlement) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newEntitlementsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Entitlement
		step.Edge.Schema = schemaConfig.Entitlement
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOauthprovider applies the HasEdge predicate on the "oauthprovider" edge.
func HasOauthprovider() predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, OauthproviderTable, OauthproviderColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.OauthProvider
		step.Edge.Schema = schemaConfig.OauthProvider
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOauthproviderWith applies the HasEdge predicate on the "oauthprovider" edge with a given conditions (other predicates).
func HasOauthproviderWith(preds ...predicate.OauthProvider) predicate.Organization {
	return predicate.Organization(func(s *sql.Selector) {
		step := newOauthproviderStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.OauthProvider
		step.Edge.Schema = schemaConfig.OauthProvider
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Organization) predicate.Organization {
	return predicate.Organization(sql.NotPredicates(p))
}
