// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"formulago/data/ent/logs"
	"formulago/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// LogsDelete is the builder for deleting a Logs entity.
type LogsDelete struct {
	config
	hooks    []Hook
	mutation *LogsMutation
}

// Where appends a list predicates to the LogsDelete builder.
func (ld *LogsDelete) Where(ps ...predicate.Logs) *LogsDelete {
	ld.mutation.Where(ps...)
	return ld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ld *LogsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, LogsMutation](ctx, ld.sqlExec, ld.mutation, ld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ld *LogsDelete) ExecX(ctx context.Context) int {
	n, err := ld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ld *LogsDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: logs.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: logs.FieldID,
			},
		},
	}
	if ps := ld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ld.mutation.done = true
	return affected, err
}

// LogsDeleteOne is the builder for deleting a single Logs entity.
type LogsDeleteOne struct {
	ld *LogsDelete
}

// Exec executes the deletion query.
func (ldo *LogsDeleteOne) Exec(ctx context.Context) error {
	n, err := ldo.ld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{logs.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ldo *LogsDeleteOne) ExecX(ctx context.Context) {
	ldo.ld.ExecX(ctx)
}