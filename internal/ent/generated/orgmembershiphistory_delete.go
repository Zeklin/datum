// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/internal"
	"github.com/datumforge/datum/internal/ent/generated/orgmembershiphistory"
)

// OrgMembershipHistoryDelete is the builder for deleting a OrgMembershipHistory entity.
type OrgMembershipHistoryDelete struct {
	config
	hooks    []Hook
	mutation *OrgMembershipHistoryMutation
}

// Where appends a list predicates to the OrgMembershipHistoryDelete builder.
func (omhd *OrgMembershipHistoryDelete) Where(ps ...predicate.OrgMembershipHistory) *OrgMembershipHistoryDelete {
	omhd.mutation.Where(ps...)
	return omhd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (omhd *OrgMembershipHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, omhd.sqlExec, omhd.mutation, omhd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (omhd *OrgMembershipHistoryDelete) ExecX(ctx context.Context) int {
	n, err := omhd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (omhd *OrgMembershipHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(orgmembershiphistory.Table, sqlgraph.NewFieldSpec(orgmembershiphistory.FieldID, field.TypeString))
	_spec.Node.Schema = omhd.schemaConfig.OrgMembershipHistory
	ctx = internal.NewSchemaConfigContext(ctx, omhd.schemaConfig)
	if ps := omhd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, omhd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	omhd.mutation.done = true
	return affected, err
}

// OrgMembershipHistoryDeleteOne is the builder for deleting a single OrgMembershipHistory entity.
type OrgMembershipHistoryDeleteOne struct {
	omhd *OrgMembershipHistoryDelete
}

// Where appends a list predicates to the OrgMembershipHistoryDelete builder.
func (omhdo *OrgMembershipHistoryDeleteOne) Where(ps ...predicate.OrgMembershipHistory) *OrgMembershipHistoryDeleteOne {
	omhdo.omhd.mutation.Where(ps...)
	return omhdo
}

// Exec executes the deletion query.
func (omhdo *OrgMembershipHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := omhdo.omhd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{orgmembershiphistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (omhdo *OrgMembershipHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := omhdo.Exec(ctx); err != nil {
		panic(err)
	}
}
