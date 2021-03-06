// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Team-bluekunVRC/discord-bot/ent/audioclip"
	"github.com/google/uuid"
)

// AudioClipCreate is the builder for creating a AudioClip entity.
type AudioClipCreate struct {
	config
	mutation *AudioClipMutation
	hooks    []Hook
}

// SetStoragePath sets the "storage_path" field.
func (acc *AudioClipCreate) SetStoragePath(s string) *AudioClipCreate {
	acc.mutation.SetStoragePath(s)
	return acc
}

// SetType sets the "type" field.
func (acc *AudioClipCreate) SetType(a audioclip.Type) *AudioClipCreate {
	acc.mutation.SetType(a)
	return acc
}

// SetID sets the "id" field.
func (acc *AudioClipCreate) SetID(u uuid.UUID) *AudioClipCreate {
	acc.mutation.SetID(u)
	return acc
}

// Mutation returns the AudioClipMutation object of the builder.
func (acc *AudioClipCreate) Mutation() *AudioClipMutation {
	return acc.mutation
}

// Save creates the AudioClip in the database.
func (acc *AudioClipCreate) Save(ctx context.Context) (*AudioClip, error) {
	var (
		err  error
		node *AudioClip
	)
	acc.defaults()
	if len(acc.hooks) == 0 {
		if err = acc.check(); err != nil {
			return nil, err
		}
		node, err = acc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AudioClipMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = acc.check(); err != nil {
				return nil, err
			}
			acc.mutation = mutation
			if node, err = acc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(acc.hooks) - 1; i >= 0; i-- {
			if acc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = acc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, acc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (acc *AudioClipCreate) SaveX(ctx context.Context) *AudioClip {
	v, err := acc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acc *AudioClipCreate) Exec(ctx context.Context) error {
	_, err := acc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acc *AudioClipCreate) ExecX(ctx context.Context) {
	if err := acc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (acc *AudioClipCreate) defaults() {
	if _, ok := acc.mutation.ID(); !ok {
		v := audioclip.DefaultID()
		acc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (acc *AudioClipCreate) check() error {
	if _, ok := acc.mutation.StoragePath(); !ok {
		return &ValidationError{Name: "storage_path", err: errors.New(`ent: missing required field "storage_path"`)}
	}
	if v, ok := acc.mutation.StoragePath(); ok {
		if err := audioclip.StoragePathValidator(v); err != nil {
			return &ValidationError{Name: "storage_path", err: fmt.Errorf(`ent: validator failed for field "storage_path": %w`, err)}
		}
	}
	if _, ok := acc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := acc.mutation.GetType(); ok {
		if err := audioclip.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	return nil
}

func (acc *AudioClipCreate) sqlSave(ctx context.Context) (*AudioClip, error) {
	_node, _spec := acc.createSpec()
	if err := sqlgraph.CreateNode(ctx, acc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(uuid.UUID)
	}
	return _node, nil
}

func (acc *AudioClipCreate) createSpec() (*AudioClip, *sqlgraph.CreateSpec) {
	var (
		_node = &AudioClip{config: acc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: audioclip.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: audioclip.FieldID,
			},
		}
	)
	if id, ok := acc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := acc.mutation.StoragePath(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: audioclip.FieldStoragePath,
		})
		_node.StoragePath = value
	}
	if value, ok := acc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: audioclip.FieldType,
		})
		_node.Type = value
	}
	return _node, _spec
}

// AudioClipCreateBulk is the builder for creating many AudioClip entities in bulk.
type AudioClipCreateBulk struct {
	config
	builders []*AudioClipCreate
}

// Save creates the AudioClip entities in the database.
func (accb *AudioClipCreateBulk) Save(ctx context.Context) ([]*AudioClip, error) {
	specs := make([]*sqlgraph.CreateSpec, len(accb.builders))
	nodes := make([]*AudioClip, len(accb.builders))
	mutators := make([]Mutator, len(accb.builders))
	for i := range accb.builders {
		func(i int, root context.Context) {
			builder := accb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AudioClipMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, accb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, accb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, accb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (accb *AudioClipCreateBulk) SaveX(ctx context.Context) []*AudioClip {
	v, err := accb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (accb *AudioClipCreateBulk) Exec(ctx context.Context) error {
	_, err := accb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (accb *AudioClipCreateBulk) ExecX(ctx context.Context) {
	if err := accb.Exec(ctx); err != nil {
		panic(err)
	}
}
