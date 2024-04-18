// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/datumforge/datum/internal/ent/generated/predicate"

	"github.com/datumforge/datum/internal/ent/generated/grouphistory"
	"github.com/datumforge/datum/internal/ent/generated/internal"
)

// GroupHistoryDelete is the builder for deleting a GroupHistory entity.
type GroupHistoryDelete struct {
	config
	hooks    []Hook
	mutation *GroupHistoryMutation
}

// Where appends a list predicates to the GroupHistoryDelete builder.
func (ghd *GroupHistoryDelete) Where(ps ...predicate.GroupHistory) *GroupHistoryDelete {
	ghd.mutation.Where(ps...)
	return ghd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ghd *GroupHistoryDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ghd.sqlExec, ghd.mutation, ghd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ghd *GroupHistoryDelete) ExecX(ctx context.Context) int {
	n, err := ghd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ghd *GroupHistoryDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(grouphistory.Table, sqlgraph.NewFieldSpec(grouphistory.FieldID, field.TypeString))
	_spec.Node.Schema = ghd.schemaConfig.GroupHistory
	ctx = internal.NewSchemaConfigContext(ctx, ghd.schemaConfig)
	if ps := ghd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ghd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ghd.mutation.done = true
	return affected, err
}

// GroupHistoryDeleteOne is the builder for deleting a single GroupHistory entity.
type GroupHistoryDeleteOne struct {
	ghd *GroupHistoryDelete
}

// Where appends a list predicates to the GroupHistoryDelete builder.
func (ghdo *GroupHistoryDeleteOne) Where(ps ...predicate.GroupHistory) *GroupHistoryDeleteOne {
	ghdo.ghd.mutation.Where(ps...)
	return ghdo
}

// Exec executes the deletion query.
func (ghdo *GroupHistoryDeleteOne) Exec(ctx context.Context) error {
	n, err := ghdo.ghd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{grouphistory.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ghdo *GroupHistoryDeleteOne) ExecX(ctx context.Context) {
	if err := ghdo.Exec(ctx); err != nil {
		panic(err)
	}
}
