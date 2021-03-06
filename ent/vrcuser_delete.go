// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Team-bluekunVRC/discord-bot/ent/predicate"
	"github.com/Team-bluekunVRC/discord-bot/ent/vrcuser"
)

// VRCUserDelete is the builder for deleting a VRCUser entity.
type VRCUserDelete struct {
	config
	hooks    []Hook
	mutation *VRCUserMutation
}

// Where appends a list predicates to the VRCUserDelete builder.
func (vud *VRCUserDelete) Where(ps ...predicate.VRCUser) *VRCUserDelete {
	vud.mutation.Where(ps...)
	return vud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (vud *VRCUserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vud.hooks) == 0 {
		affected, err = vud.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VRCUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			vud.mutation = mutation
			affected, err = vud.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vud.hooks) - 1; i >= 0; i-- {
			if vud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (vud *VRCUserDelete) ExecX(ctx context.Context) int {
	n, err := vud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (vud *VRCUserDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: vrcuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: vrcuser.FieldID,
			},
		},
	}
	if ps := vud.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, vud.driver, _spec)
}

// VRCUserDeleteOne is the builder for deleting a single VRCUser entity.
type VRCUserDeleteOne struct {
	vud *VRCUserDelete
}

// Exec executes the deletion query.
func (vudo *VRCUserDeleteOne) Exec(ctx context.Context) error {
	n, err := vudo.vud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{vrcuser.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (vudo *VRCUserDeleteOne) ExecX(ctx context.Context) {
	vudo.vud.ExecX(ctx)
}
