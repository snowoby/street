// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"street/ent/episode"
	"street/ent/profile"
	"street/ent/schema"
	"street/ent/series"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SeriesCreate is the builder for creating a Series entity.
type SeriesCreate struct {
	config
	mutation *SeriesMutation
	hooks    []Hook
}

// SetSid sets the "sid" field.
func (sc *SeriesCreate) SetSid(s schema.ID) *SeriesCreate {
	sc.mutation.SetSid(s)
	return sc
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableSid(s *schema.ID) *SeriesCreate {
	if s != nil {
		sc.SetSid(*s)
	}
	return sc
}

// SetCreateTime sets the "create_time" field.
func (sc *SeriesCreate) SetCreateTime(t time.Time) *SeriesCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableCreateTime(t *time.Time) *SeriesCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *SeriesCreate) SetUpdateTime(t time.Time) *SeriesCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableUpdateTime(t *time.Time) *SeriesCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetTitle sets the "title" field.
func (sc *SeriesCreate) SetTitle(s string) *SeriesCreate {
	sc.mutation.SetTitle(s)
	return sc
}

// SetType sets the "type" field.
func (sc *SeriesCreate) SetType(s string) *SeriesCreate {
	sc.mutation.SetType(s)
	return sc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableType(s *string) *SeriesCreate {
	if s != nil {
		sc.SetType(*s)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SeriesCreate) SetID(u uuid.UUID) *SeriesCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SeriesCreate) SetNillableID(u *uuid.UUID) *SeriesCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// AddEpisodeIDs adds the "episodes" edge to the Episode entity by IDs.
func (sc *SeriesCreate) AddEpisodeIDs(ids ...uuid.UUID) *SeriesCreate {
	sc.mutation.AddEpisodeIDs(ids...)
	return sc
}

// AddEpisodes adds the "episodes" edges to the Episode entity.
func (sc *SeriesCreate) AddEpisodes(e ...*Episode) *SeriesCreate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return sc.AddEpisodeIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Profile entity by ID.
func (sc *SeriesCreate) SetOwnerID(id uuid.UUID) *SeriesCreate {
	sc.mutation.SetOwnerID(id)
	return sc
}

// SetOwner sets the "owner" edge to the Profile entity.
func (sc *SeriesCreate) SetOwner(p *Profile) *SeriesCreate {
	return sc.SetOwnerID(p.ID)
}

// AddParticipantIDs adds the "participant" edge to the Profile entity by IDs.
func (sc *SeriesCreate) AddParticipantIDs(ids ...uuid.UUID) *SeriesCreate {
	sc.mutation.AddParticipantIDs(ids...)
	return sc
}

// AddParticipant adds the "participant" edges to the Profile entity.
func (sc *SeriesCreate) AddParticipant(p ...*Profile) *SeriesCreate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return sc.AddParticipantIDs(ids...)
}

// Mutation returns the SeriesMutation object of the builder.
func (sc *SeriesCreate) Mutation() *SeriesMutation {
	return sc.mutation
}

// Save creates the Series in the database.
func (sc *SeriesCreate) Save(ctx context.Context) (*Series, error) {
	var (
		err  error
		node *Series
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SeriesCreate) SaveX(ctx context.Context) *Series {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SeriesCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SeriesCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SeriesCreate) defaults() {
	if _, ok := sc.mutation.Sid(); !ok {
		v := series.DefaultSid()
		sc.mutation.SetSid(v)
	}
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := series.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := series.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := series.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SeriesCreate) check() error {
	if _, ok := sc.mutation.Sid(); !ok {
		return &ValidationError{Name: "sid", err: errors.New(`ent: missing required field "Series.sid"`)}
	}
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Series.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Series.update_time"`)}
	}
	if _, ok := sc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Series.title"`)}
	}
	if v, ok := sc.mutation.Title(); ok {
		if err := series.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Series.title": %w`, err)}
		}
	}
	if v, ok := sc.mutation.GetType(); ok {
		if err := series.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Series.type": %w`, err)}
		}
	}
	if _, ok := sc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Series.owner"`)}
	}
	return nil
}

func (sc *SeriesCreate) sqlSave(ctx context.Context) (*Series, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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

func (sc *SeriesCreate) createSpec() (*Series, *sqlgraph.CreateSpec) {
	var (
		_node = &Series{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: series.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: series.FieldID,
			},
		}
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Sid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: series.FieldSid,
		})
		_node.Sid = value
	}
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: series.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: series.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: series.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := sc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: series.FieldType,
		})
		_node.Type = &value
	}
	if nodes := sc.mutation.EpisodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   series.EpisodesTable,
			Columns: []string{series.EpisodesColumn},
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
	if nodes := sc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   series.OwnerTable,
			Columns: []string{series.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: profile.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.profile_series = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.ParticipantIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   series.ParticipantTable,
			Columns: series.ParticipantPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: profile.FieldID,
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

// SeriesCreateBulk is the builder for creating many Series entities in bulk.
type SeriesCreateBulk struct {
	config
	builders []*SeriesCreate
}

// Save creates the Series entities in the database.
func (scb *SeriesCreateBulk) Save(ctx context.Context) ([]*Series, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Series, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SeriesMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SeriesCreateBulk) SaveX(ctx context.Context) []*Series {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SeriesCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SeriesCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
