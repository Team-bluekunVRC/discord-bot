// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Team-bluekunVRC/discord-bot/ent/discorduser"
	"github.com/Team-bluekunVRC/discord-bot/ent/user"
	"github.com/Team-bluekunVRC/discord-bot/ent/vrcuser"
	"github.com/google/uuid"
)

// VRCUserCreate is the builder for creating a VRCUser entity.
type VRCUserCreate struct {
	config
	mutation *VRCUserMutation
	hooks    []Hook
}

// SetID sets the "id" field.
func (vuc *VRCUserCreate) SetID(s string) *VRCUserCreate {
	vuc.mutation.SetID(s)
	return vuc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vuc *VRCUserCreate) SetUserID(id uuid.UUID) *VRCUserCreate {
	vuc.mutation.SetUserID(id)
	return vuc
}

// SetUser sets the "user" edge to the User entity.
func (vuc *VRCUserCreate) SetUser(u *User) *VRCUserCreate {
	return vuc.SetUserID(u.ID)
}

// SetDiscordID sets the "discord" edge to the DiscordUser entity by ID.
func (vuc *VRCUserCreate) SetDiscordID(id string) *VRCUserCreate {
	vuc.mutation.SetDiscordID(id)
	return vuc
}

// SetNillableDiscordID sets the "discord" edge to the DiscordUser entity by ID if the given value is not nil.
func (vuc *VRCUserCreate) SetNillableDiscordID(id *string) *VRCUserCreate {
	if id != nil {
		vuc = vuc.SetDiscordID(*id)
	}
	return vuc
}

// SetDiscord sets the "discord" edge to the DiscordUser entity.
func (vuc *VRCUserCreate) SetDiscord(d *DiscordUser) *VRCUserCreate {
	return vuc.SetDiscordID(d.ID)
}

// Mutation returns the VRCUserMutation object of the builder.
func (vuc *VRCUserCreate) Mutation() *VRCUserMutation {
	return vuc.mutation
}

// Save creates the VRCUser in the database.
func (vuc *VRCUserCreate) Save(ctx context.Context) (*VRCUser, error) {
	var (
		err  error
		node *VRCUser
	)
	if len(vuc.hooks) == 0 {
		if err = vuc.check(); err != nil {
			return nil, err
		}
		node, err = vuc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VRCUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuc.check(); err != nil {
				return nil, err
			}
			vuc.mutation = mutation
			if node, err = vuc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(vuc.hooks) - 1; i >= 0; i-- {
			if vuc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vuc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (vuc *VRCUserCreate) SaveX(ctx context.Context) *VRCUser {
	v, err := vuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vuc *VRCUserCreate) Exec(ctx context.Context) error {
	_, err := vuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuc *VRCUserCreate) ExecX(ctx context.Context) {
	if err := vuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuc *VRCUserCreate) check() error {
	if v, ok := vuc.mutation.ID(); ok {
		if err := vrcuser.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "id": %w`, err)}
		}
	}
	if _, ok := vuc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New("ent: missing required edge \"user\"")}
	}
	return nil
}

func (vuc *VRCUserCreate) sqlSave(ctx context.Context) (*VRCUser, error) {
	_node, _spec := vuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, vuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		_node.ID = _spec.ID.Value.(string)
	}
	return _node, nil
}

func (vuc *VRCUserCreate) createSpec() (*VRCUser, *sqlgraph.CreateSpec) {
	var (
		_node = &VRCUser{config: vuc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: vrcuser.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: vrcuser.FieldID,
			},
		}
	)
	if id, ok := vuc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if nodes := vuc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   vrcuser.UserTable,
			Columns: []string{vrcuser.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.vrc_user_user = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := vuc.mutation.DiscordIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   vrcuser.DiscordTable,
			Columns: []string{vrcuser.DiscordColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: discorduser.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.discord_user_vrc = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// VRCUserCreateBulk is the builder for creating many VRCUser entities in bulk.
type VRCUserCreateBulk struct {
	config
	builders []*VRCUserCreate
}

// Save creates the VRCUser entities in the database.
func (vucb *VRCUserCreateBulk) Save(ctx context.Context) ([]*VRCUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(vucb.builders))
	nodes := make([]*VRCUser, len(vucb.builders))
	mutators := make([]Mutator, len(vucb.builders))
	for i := range vucb.builders {
		func(i int, root context.Context) {
			builder := vucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*VRCUserMutation)
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
					_, err = mutators[i+1].Mutate(root, vucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, vucb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, vucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (vucb *VRCUserCreateBulk) SaveX(ctx context.Context) []*VRCUser {
	v, err := vucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (vucb *VRCUserCreateBulk) Exec(ctx context.Context) error {
	_, err := vucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vucb *VRCUserCreateBulk) ExecX(ctx context.Context) {
	if err := vucb.Exec(ctx); err != nil {
		panic(err)
	}
}
