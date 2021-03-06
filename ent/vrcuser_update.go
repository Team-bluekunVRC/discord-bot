// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Team-bluekunVRC/discord-bot/ent/discorduser"
	"github.com/Team-bluekunVRC/discord-bot/ent/predicate"
	"github.com/Team-bluekunVRC/discord-bot/ent/user"
	"github.com/Team-bluekunVRC/discord-bot/ent/vrcuser"
	"github.com/google/uuid"
)

// VRCUserUpdate is the builder for updating VRCUser entities.
type VRCUserUpdate struct {
	config
	hooks    []Hook
	mutation *VRCUserMutation
}

// Where appends a list predicates to the VRCUserUpdate builder.
func (vuu *VRCUserUpdate) Where(ps ...predicate.VRCUser) *VRCUserUpdate {
	vuu.mutation.Where(ps...)
	return vuu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vuu *VRCUserUpdate) SetUserID(id uuid.UUID) *VRCUserUpdate {
	vuu.mutation.SetUserID(id)
	return vuu
}

// SetUser sets the "user" edge to the User entity.
func (vuu *VRCUserUpdate) SetUser(u *User) *VRCUserUpdate {
	return vuu.SetUserID(u.ID)
}

// SetDiscordID sets the "discord" edge to the DiscordUser entity by ID.
func (vuu *VRCUserUpdate) SetDiscordID(id string) *VRCUserUpdate {
	vuu.mutation.SetDiscordID(id)
	return vuu
}

// SetNillableDiscordID sets the "discord" edge to the DiscordUser entity by ID if the given value is not nil.
func (vuu *VRCUserUpdate) SetNillableDiscordID(id *string) *VRCUserUpdate {
	if id != nil {
		vuu = vuu.SetDiscordID(*id)
	}
	return vuu
}

// SetDiscord sets the "discord" edge to the DiscordUser entity.
func (vuu *VRCUserUpdate) SetDiscord(d *DiscordUser) *VRCUserUpdate {
	return vuu.SetDiscordID(d.ID)
}

// Mutation returns the VRCUserMutation object of the builder.
func (vuu *VRCUserUpdate) Mutation() *VRCUserMutation {
	return vuu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vuu *VRCUserUpdate) ClearUser() *VRCUserUpdate {
	vuu.mutation.ClearUser()
	return vuu
}

// ClearDiscord clears the "discord" edge to the DiscordUser entity.
func (vuu *VRCUserUpdate) ClearDiscord() *VRCUserUpdate {
	vuu.mutation.ClearDiscord()
	return vuu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vuu *VRCUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(vuu.hooks) == 0 {
		if err = vuu.check(); err != nil {
			return 0, err
		}
		affected, err = vuu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VRCUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuu.check(); err != nil {
				return 0, err
			}
			vuu.mutation = mutation
			affected, err = vuu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(vuu.hooks) - 1; i >= 0; i-- {
			if vuu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vuu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuu *VRCUserUpdate) SaveX(ctx context.Context) int {
	affected, err := vuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vuu *VRCUserUpdate) Exec(ctx context.Context) error {
	_, err := vuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuu *VRCUserUpdate) ExecX(ctx context.Context) {
	if err := vuu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuu *VRCUserUpdate) check() error {
	if _, ok := vuu.mutation.UserID(); vuu.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (vuu *VRCUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vrcuser.Table,
			Columns: vrcuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: vrcuser.FieldID,
			},
		},
	}
	if ps := vuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if vuu.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuu.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuu.mutation.DiscordCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuu.mutation.DiscordIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vrcuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// VRCUserUpdateOne is the builder for updating a single VRCUser entity.
type VRCUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VRCUserMutation
}

// SetUserID sets the "user" edge to the User entity by ID.
func (vuuo *VRCUserUpdateOne) SetUserID(id uuid.UUID) *VRCUserUpdateOne {
	vuuo.mutation.SetUserID(id)
	return vuuo
}

// SetUser sets the "user" edge to the User entity.
func (vuuo *VRCUserUpdateOne) SetUser(u *User) *VRCUserUpdateOne {
	return vuuo.SetUserID(u.ID)
}

// SetDiscordID sets the "discord" edge to the DiscordUser entity by ID.
func (vuuo *VRCUserUpdateOne) SetDiscordID(id string) *VRCUserUpdateOne {
	vuuo.mutation.SetDiscordID(id)
	return vuuo
}

// SetNillableDiscordID sets the "discord" edge to the DiscordUser entity by ID if the given value is not nil.
func (vuuo *VRCUserUpdateOne) SetNillableDiscordID(id *string) *VRCUserUpdateOne {
	if id != nil {
		vuuo = vuuo.SetDiscordID(*id)
	}
	return vuuo
}

// SetDiscord sets the "discord" edge to the DiscordUser entity.
func (vuuo *VRCUserUpdateOne) SetDiscord(d *DiscordUser) *VRCUserUpdateOne {
	return vuuo.SetDiscordID(d.ID)
}

// Mutation returns the VRCUserMutation object of the builder.
func (vuuo *VRCUserUpdateOne) Mutation() *VRCUserMutation {
	return vuuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (vuuo *VRCUserUpdateOne) ClearUser() *VRCUserUpdateOne {
	vuuo.mutation.ClearUser()
	return vuuo
}

// ClearDiscord clears the "discord" edge to the DiscordUser entity.
func (vuuo *VRCUserUpdateOne) ClearDiscord() *VRCUserUpdateOne {
	vuuo.mutation.ClearDiscord()
	return vuuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuuo *VRCUserUpdateOne) Select(field string, fields ...string) *VRCUserUpdateOne {
	vuuo.fields = append([]string{field}, fields...)
	return vuuo
}

// Save executes the query and returns the updated VRCUser entity.
func (vuuo *VRCUserUpdateOne) Save(ctx context.Context) (*VRCUser, error) {
	var (
		err  error
		node *VRCUser
	)
	if len(vuuo.hooks) == 0 {
		if err = vuuo.check(); err != nil {
			return nil, err
		}
		node, err = vuuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*VRCUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = vuuo.check(); err != nil {
				return nil, err
			}
			vuuo.mutation = mutation
			node, err = vuuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(vuuo.hooks) - 1; i >= 0; i-- {
			if vuuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = vuuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, vuuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (vuuo *VRCUserUpdateOne) SaveX(ctx context.Context) *VRCUser {
	node, err := vuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuuo *VRCUserUpdateOne) Exec(ctx context.Context) error {
	_, err := vuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuuo *VRCUserUpdateOne) ExecX(ctx context.Context) {
	if err := vuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (vuuo *VRCUserUpdateOne) check() error {
	if _, ok := vuuo.mutation.UserID(); vuuo.mutation.UserCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"user\"")
	}
	return nil
}

func (vuuo *VRCUserUpdateOne) sqlSave(ctx context.Context) (_node *VRCUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   vrcuser.Table,
			Columns: vrcuser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: vrcuser.FieldID,
			},
		},
	}
	id, ok := vuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing VRCUser.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := vuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, vrcuser.FieldID)
		for _, f := range fields {
			if !vrcuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != vrcuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if vuuo.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuuo.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if vuuo.mutation.DiscordCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := vuuo.mutation.DiscordIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &VRCUser{config: vuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{vrcuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
