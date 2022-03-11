// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"street/ent/account"
	"street/ent/comment"
	"street/ent/episode"
	"street/ent/profile"
	"street/ent/schema"
	"street/ent/series"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProfileCreate is the builder for creating a Profile entity.
type ProfileCreate struct {
	config
	mutation *ProfileMutation
	hooks    []Hook
}

// SetSid sets the "sid" field.
func (pc *ProfileCreate) SetSid(s schema.ID) *ProfileCreate {
	pc.mutation.SetSid(s)
	return pc
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableSid(s *schema.ID) *ProfileCreate {
	if s != nil {
		pc.SetSid(*s)
	}
	return pc
}

// SetCreateTime sets the "create_time" field.
func (pc *ProfileCreate) SetCreateTime(t time.Time) *ProfileCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableCreateTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *ProfileCreate) SetUpdateTime(t time.Time) *ProfileCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableUpdateTime(t *time.Time) *ProfileCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetTitle sets the "title" field.
func (pc *ProfileCreate) SetTitle(s string) *ProfileCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetCall sets the "call" field.
func (pc *ProfileCreate) SetCall(s string) *ProfileCreate {
	pc.mutation.SetCall(s)
	return pc
}

// SetCategory sets the "category" field.
func (pc *ProfileCreate) SetCategory(s string) *ProfileCreate {
	pc.mutation.SetCategory(s)
	return pc
}

// SetAvatar sets the "avatar" field.
func (pc *ProfileCreate) SetAvatar(s string) *ProfileCreate {
	pc.mutation.SetAvatar(s)
	return pc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableAvatar(s *string) *ProfileCreate {
	if s != nil {
		pc.SetAvatar(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProfileCreate) SetID(u uuid.UUID) *ProfileCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *ProfileCreate) SetNillableID(u *uuid.UUID) *ProfileCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (pc *ProfileCreate) SetAccountID(id uuid.UUID) *ProfileCreate {
	pc.mutation.SetAccountID(id)
	return pc
}

// SetAccount sets the "account" edge to the Account entity.
func (pc *ProfileCreate) SetAccount(a *Account) *ProfileCreate {
	return pc.SetAccountID(a.ID)
}

// AddEpisodeIDs adds the "episode" edge to the Episode entity by IDs.
func (pc *ProfileCreate) AddEpisodeIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddEpisodeIDs(ids...)
	return pc
}

// AddEpisode adds the "episode" edges to the Episode entity.
func (pc *ProfileCreate) AddEpisode(e ...*Episode) *ProfileCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pc.AddEpisodeIDs(ids...)
}

// AddCommenterIDs adds the "commenter" edge to the Comment entity by IDs.
func (pc *ProfileCreate) AddCommenterIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddCommenterIDs(ids...)
	return pc
}

// AddCommenter adds the "commenter" edges to the Comment entity.
func (pc *ProfileCreate) AddCommenter(c ...*Comment) *ProfileCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCommenterIDs(ids...)
}

// AddSeriesIDs adds the "series" edge to the Series entity by IDs.
func (pc *ProfileCreate) AddSeriesIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddSeriesIDs(ids...)
	return pc
}

// AddSeries adds the "series" edges to the Series entity.
func (pc *ProfileCreate) AddSeries(s ...*Series) *ProfileCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddSeriesIDs(ids...)
}

// AddJoinedSeriesIDs adds the "joined_series" edge to the Series entity by IDs.
func (pc *ProfileCreate) AddJoinedSeriesIDs(ids ...uuid.UUID) *ProfileCreate {
	pc.mutation.AddJoinedSeriesIDs(ids...)
	return pc
}

// AddJoinedSeries adds the "joined_series" edges to the Series entity.
func (pc *ProfileCreate) AddJoinedSeries(s ...*Series) *ProfileCreate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddJoinedSeriesIDs(ids...)
}

// Mutation returns the ProfileMutation object of the builder.
func (pc *ProfileCreate) Mutation() *ProfileMutation {
	return pc.mutation
}

// Save creates the Profile in the database.
func (pc *ProfileCreate) Save(ctx context.Context) (*Profile, error) {
	var (
		err  error
		node *Profile
	)
	pc.defaults()
	if len(pc.hooks) == 0 {
		if err = pc.check(); err != nil {
			return nil, err
		}
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pc.check(); err != nil {
				return nil, err
			}
			pc.mutation = mutation
			if node, err = pc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			if pc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfileCreate) SaveX(ctx context.Context) *Profile {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfileCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfileCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfileCreate) defaults() {
	if _, ok := pc.mutation.Sid(); !ok {
		v := profile.DefaultSid()
		pc.mutation.SetSid(v)
	}
	if _, ok := pc.mutation.CreateTime(); !ok {
		v := profile.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		v := profile.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := profile.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfileCreate) check() error {
	if _, ok := pc.mutation.Sid(); !ok {
		return &ValidationError{Name: "sid", err: errors.New(`ent: missing required field "Profile.sid"`)}
	}
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Profile.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Profile.update_time"`)}
	}
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Profile.title"`)}
	}
	if v, ok := pc.mutation.Title(); ok {
		if err := profile.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Profile.title": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Call(); !ok {
		return &ValidationError{Name: "call", err: errors.New(`ent: missing required field "Profile.call"`)}
	}
	if v, ok := pc.mutation.Call(); ok {
		if err := profile.CallValidator(v); err != nil {
			return &ValidationError{Name: "call", err: fmt.Errorf(`ent: validator failed for field "Profile.call": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`ent: missing required field "Profile.category"`)}
	}
	if v, ok := pc.mutation.Category(); ok {
		if err := profile.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`ent: validator failed for field "Profile.category": %w`, err)}
		}
	}
	if v, ok := pc.mutation.Avatar(); ok {
		if err := profile.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf(`ent: validator failed for field "Profile.avatar": %w`, err)}
		}
	}
	if _, ok := pc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required edge "Profile.account"`)}
	}
	return nil
}

func (pc *ProfileCreate) sqlSave(ctx context.Context) (*Profile, error) {
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (pc *ProfileCreate) createSpec() (*Profile, *sqlgraph.CreateSpec) {
	var (
		_node = &Profile{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: profile.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profile.FieldID,
			},
		}
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Sid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: profile.FieldSid,
		})
		_node.Sid = value
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := pc.mutation.Call(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldCall,
		})
		_node.Call = value
	}
	if value, ok := pc.mutation.Category(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldCategory,
		})
		_node.Category = value
	}
	if value, ok := pc.mutation.Avatar(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldAvatar,
		})
		_node.Avatar = &value
	}
	if nodes := pc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   profile.AccountTable,
			Columns: []string{profile.AccountColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.account_profile = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.EpisodeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.EpisodeTable,
			Columns: []string{profile.EpisodeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: episode.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.CommenterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.CommenterTable,
			Columns: []string{profile.CommenterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: comment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SeriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.SeriesTable,
			Columns: []string{profile.SeriesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: series.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.JoinedSeriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   profile.JoinedSeriesTable,
			Columns: profile.JoinedSeriesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: series.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProfileCreateBulk is the builder for creating many Profile entities in bulk.
type ProfileCreateBulk struct {
	config
	builders []*ProfileCreate
}

// Save creates the Profile entities in the database.
func (pcb *ProfileCreateBulk) Save(ctx context.Context) ([]*Profile, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Profile, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfileMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfileCreateBulk) SaveX(ctx context.Context) []*Profile {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfileCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfileCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
