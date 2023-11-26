// Code generated by ent, DO NOT EDIT.

package usersetting

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContainsFold(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldUpdatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldUpdatedBy, v))
}

// Locked applies equality check predicate on the "locked" field. It's identical to LockedEQ.
func Locked(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldLocked, v))
}

// SilencedAt applies equality check predicate on the "silenced_at" field. It's identical to SilencedAtEQ.
func SilencedAt(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldSilencedAt, v))
}

// SuspendedAt applies equality check predicate on the "suspended_at" field. It's identical to SuspendedAtEQ.
func SuspendedAt(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldSuspendedAt, v))
}

// RecoveryCode applies equality check predicate on the "recovery_code" field. It's identical to RecoveryCodeEQ.
func RecoveryCode(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldRecoveryCode, v))
}

// EmailConfirmed applies equality check predicate on the "email_confirmed" field. It's identical to EmailConfirmedEQ.
func EmailConfirmed(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldEmailConfirmed, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldUpdatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldCreatedBy, v))
}

// CreatedByContains applies the Contains predicate on the "created_by" field.
func CreatedByContains(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContains(FieldCreatedBy, v))
}

// CreatedByHasPrefix applies the HasPrefix predicate on the "created_by" field.
func CreatedByHasPrefix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasPrefix(FieldCreatedBy, v))
}

// CreatedByHasSuffix applies the HasSuffix predicate on the "created_by" field.
func CreatedByHasSuffix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasSuffix(FieldCreatedBy, v))
}

// CreatedByIsNil applies the IsNil predicate on the "created_by" field.
func CreatedByIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldCreatedBy))
}

// CreatedByNotNil applies the NotNil predicate on the "created_by" field.
func CreatedByNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldCreatedBy))
}

// CreatedByEqualFold applies the EqualFold predicate on the "created_by" field.
func CreatedByEqualFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEqualFold(FieldCreatedBy, v))
}

// CreatedByContainsFold applies the ContainsFold predicate on the "created_by" field.
func CreatedByContainsFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContainsFold(FieldCreatedBy, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldUpdatedBy, v))
}

// UpdatedByContains applies the Contains predicate on the "updated_by" field.
func UpdatedByContains(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContains(FieldUpdatedBy, v))
}

// UpdatedByHasPrefix applies the HasPrefix predicate on the "updated_by" field.
func UpdatedByHasPrefix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasPrefix(FieldUpdatedBy, v))
}

// UpdatedByHasSuffix applies the HasSuffix predicate on the "updated_by" field.
func UpdatedByHasSuffix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasSuffix(FieldUpdatedBy, v))
}

// UpdatedByIsNil applies the IsNil predicate on the "updated_by" field.
func UpdatedByIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldUpdatedBy))
}

// UpdatedByNotNil applies the NotNil predicate on the "updated_by" field.
func UpdatedByNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldUpdatedBy))
}

// UpdatedByEqualFold applies the EqualFold predicate on the "updated_by" field.
func UpdatedByEqualFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEqualFold(FieldUpdatedBy, v))
}

// UpdatedByContainsFold applies the ContainsFold predicate on the "updated_by" field.
func UpdatedByContainsFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContainsFold(FieldUpdatedBy, v))
}

// LockedEQ applies the EQ predicate on the "locked" field.
func LockedEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldLocked, v))
}

// LockedNEQ applies the NEQ predicate on the "locked" field.
func LockedNEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldLocked, v))
}

// SilencedAtEQ applies the EQ predicate on the "silenced_at" field.
func SilencedAtEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldSilencedAt, v))
}

// SilencedAtNEQ applies the NEQ predicate on the "silenced_at" field.
func SilencedAtNEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldSilencedAt, v))
}

// SilencedAtIn applies the In predicate on the "silenced_at" field.
func SilencedAtIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldSilencedAt, vs...))
}

// SilencedAtNotIn applies the NotIn predicate on the "silenced_at" field.
func SilencedAtNotIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldSilencedAt, vs...))
}

// SilencedAtGT applies the GT predicate on the "silenced_at" field.
func SilencedAtGT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldSilencedAt, v))
}

// SilencedAtGTE applies the GTE predicate on the "silenced_at" field.
func SilencedAtGTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldSilencedAt, v))
}

// SilencedAtLT applies the LT predicate on the "silenced_at" field.
func SilencedAtLT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldSilencedAt, v))
}

// SilencedAtLTE applies the LTE predicate on the "silenced_at" field.
func SilencedAtLTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldSilencedAt, v))
}

// SilencedAtIsNil applies the IsNil predicate on the "silenced_at" field.
func SilencedAtIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldSilencedAt))
}

// SilencedAtNotNil applies the NotNil predicate on the "silenced_at" field.
func SilencedAtNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldSilencedAt))
}

// SuspendedAtEQ applies the EQ predicate on the "suspended_at" field.
func SuspendedAtEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldSuspendedAt, v))
}

// SuspendedAtNEQ applies the NEQ predicate on the "suspended_at" field.
func SuspendedAtNEQ(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldSuspendedAt, v))
}

// SuspendedAtIn applies the In predicate on the "suspended_at" field.
func SuspendedAtIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldSuspendedAt, vs...))
}

// SuspendedAtNotIn applies the NotIn predicate on the "suspended_at" field.
func SuspendedAtNotIn(vs ...time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldSuspendedAt, vs...))
}

// SuspendedAtGT applies the GT predicate on the "suspended_at" field.
func SuspendedAtGT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldSuspendedAt, v))
}

// SuspendedAtGTE applies the GTE predicate on the "suspended_at" field.
func SuspendedAtGTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldSuspendedAt, v))
}

// SuspendedAtLT applies the LT predicate on the "suspended_at" field.
func SuspendedAtLT(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldSuspendedAt, v))
}

// SuspendedAtLTE applies the LTE predicate on the "suspended_at" field.
func SuspendedAtLTE(v time.Time) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldSuspendedAt, v))
}

// SuspendedAtIsNil applies the IsNil predicate on the "suspended_at" field.
func SuspendedAtIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldSuspendedAt))
}

// SuspendedAtNotNil applies the NotNil predicate on the "suspended_at" field.
func SuspendedAtNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldSuspendedAt))
}

// RecoveryCodeEQ applies the EQ predicate on the "recovery_code" field.
func RecoveryCodeEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldRecoveryCode, v))
}

// RecoveryCodeNEQ applies the NEQ predicate on the "recovery_code" field.
func RecoveryCodeNEQ(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldRecoveryCode, v))
}

// RecoveryCodeIn applies the In predicate on the "recovery_code" field.
func RecoveryCodeIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldRecoveryCode, vs...))
}

// RecoveryCodeNotIn applies the NotIn predicate on the "recovery_code" field.
func RecoveryCodeNotIn(vs ...string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldRecoveryCode, vs...))
}

// RecoveryCodeGT applies the GT predicate on the "recovery_code" field.
func RecoveryCodeGT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGT(FieldRecoveryCode, v))
}

// RecoveryCodeGTE applies the GTE predicate on the "recovery_code" field.
func RecoveryCodeGTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldGTE(FieldRecoveryCode, v))
}

// RecoveryCodeLT applies the LT predicate on the "recovery_code" field.
func RecoveryCodeLT(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLT(FieldRecoveryCode, v))
}

// RecoveryCodeLTE applies the LTE predicate on the "recovery_code" field.
func RecoveryCodeLTE(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldLTE(FieldRecoveryCode, v))
}

// RecoveryCodeContains applies the Contains predicate on the "recovery_code" field.
func RecoveryCodeContains(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContains(FieldRecoveryCode, v))
}

// RecoveryCodeHasPrefix applies the HasPrefix predicate on the "recovery_code" field.
func RecoveryCodeHasPrefix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasPrefix(FieldRecoveryCode, v))
}

// RecoveryCodeHasSuffix applies the HasSuffix predicate on the "recovery_code" field.
func RecoveryCodeHasSuffix(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldHasSuffix(FieldRecoveryCode, v))
}

// RecoveryCodeIsNil applies the IsNil predicate on the "recovery_code" field.
func RecoveryCodeIsNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIsNull(FieldRecoveryCode))
}

// RecoveryCodeNotNil applies the NotNil predicate on the "recovery_code" field.
func RecoveryCodeNotNil() predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotNull(FieldRecoveryCode))
}

// RecoveryCodeEqualFold applies the EqualFold predicate on the "recovery_code" field.
func RecoveryCodeEqualFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEqualFold(FieldRecoveryCode, v))
}

// RecoveryCodeContainsFold applies the ContainsFold predicate on the "recovery_code" field.
func RecoveryCodeContainsFold(v string) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldContainsFold(FieldRecoveryCode, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldStatus, vs...))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNotIn(FieldRole, vs...))
}

// EmailConfirmedEQ applies the EQ predicate on the "email_confirmed" field.
func EmailConfirmedEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldEQ(FieldEmailConfirmed, v))
}

// EmailConfirmedNEQ applies the NEQ predicate on the "email_confirmed" field.
func EmailConfirmedNEQ(v bool) predicate.UserSetting {
	return predicate.UserSetting(sql.FieldNEQ(FieldEmailConfirmed, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserSetting {
	return predicate.UserSetting(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, UserTable, UserColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UserSetting
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserSetting {
	return predicate.UserSetting(func(s *sql.Selector) {
		step := newUserStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.User
		step.Edge.Schema = schemaConfig.UserSetting
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserSetting) predicate.UserSetting {
	return predicate.UserSetting(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserSetting) predicate.UserSetting {
	return predicate.UserSetting(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserSetting) predicate.UserSetting {
	return predicate.UserSetting(sql.NotPredicates(p))
}