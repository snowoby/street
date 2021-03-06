// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
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

// EpisodeCreate is the builder for creating a Episode entity.
type EpisodeCreate struct {
	config
	mutation *EpisodeMutation
	hooks    []Hook
}

// SetSid sets the "sid" field.
func (ec *EpisodeCreate) SetSid(s schema.ID) *EpisodeCreate {
	ec.mutation.SetSid(s)
	return ec
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableSid(s *schema.ID) *EpisodeCreate {
	if s != nil {
		ec.SetSid(*s)
	}
	return ec
}

// SetCreateTime sets the "create_time" field.
func (ec *EpisodeCreate) SetCreateTime(t time.Time) *EpisodeCreate {
	ec.mutation.SetCreateTime(t)
	return ec
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableCreateTime(t *time.Time) *EpisodeCreate {
	if t != nil {
		ec.SetCreateTime(*t)
	}
	return ec
}

// SetUpdateTime sets the "update_time" field.
func (ec *EpisodeCreate) SetUpdateTime(t time.Time) *EpisodeCreate {
	ec.mutation.SetUpdateTime(t)
	return ec
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableUpdateTime(t *time.Time) *EpisodeCreate {
	if t != nil {
		ec.SetUpdateTime(*t)
	}
	return ec
}

// SetCover sets the "cover" field.
func (ec *EpisodeCreate) SetCover(s string) *EpisodeCreate {
	ec.mutation.SetCover(s)
	return ec
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableCover(s *string) *EpisodeCreate {
	if s != nil {
		ec.SetCover(*s)
	}
	return ec
}

// SetTitle sets the "title" field.
func (ec *EpisodeCreate) SetTitle(s string) *EpisodeCreate {
	ec.mutation.SetTitle(s)
	return ec
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableTitle(s *string) *EpisodeCreate {
	if s != nil {
		ec.SetTitle(*s)
	}
	return ec
}

// SetContent sets the "content" field.
func (ec *EpisodeCreate) SetContent(s string) *EpisodeCreate {
	ec.mutation.SetContent(s)
	return ec
}

// SetFiles sets the "files" field.
func (ec *EpisodeCreate) SetFiles(s schema.Medias) *EpisodeCreate {
	ec.mutation.SetFiles(s)
	return ec
}

// SetID sets the "id" field.
func (ec *EpisodeCreate) SetID(u uuid.UUID) *EpisodeCreate {
	ec.mutation.SetID(u)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableID(u *uuid.UUID) *EpisodeCreate {
	if u != nil {
		ec.SetID(*u)
	}
	return ec
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (ec *EpisodeCreate) SetProfileID(id uuid.UUID) *EpisodeCreate {
	ec.mutation.SetProfileID(id)
	return ec
}

// SetProfile sets the "profile" edge to the Profile entity.
func (ec *EpisodeCreate) SetProfile(p *Profile) *EpisodeCreate {
	return ec.SetProfileID(p.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (ec *EpisodeCreate) AddCommentIDs(ids ...uuid.UUID) *EpisodeCreate {
	ec.mutation.AddCommentIDs(ids...)
	return ec
}

// AddComments adds the "comments" edges to the Comment entity.
func (ec *EpisodeCreate) AddComments(c ...*Comment) *EpisodeCreate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ec.AddCommentIDs(ids...)
}

// SetSeriesID sets the "series" edge to the Series entity by ID.
func (ec *EpisodeCreate) SetSeriesID(id uuid.UUID) *EpisodeCreate {
	ec.mutation.SetSeriesID(id)
	return ec
}

// SetNillableSeriesID sets the "series" edge to the Series entity by ID if the given value is not nil.
func (ec *EpisodeCreate) SetNillableSeriesID(id *uuid.UUID) *EpisodeCreate {
	if id != nil {
		ec = ec.SetSeriesID(*id)
	}
	return ec
}

// SetSeries sets the "series" edge to the Series entity.
func (ec *EpisodeCreate) SetSeries(s *Series) *EpisodeCreate {
	return ec.SetSeriesID(s.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (ec *EpisodeCreate) Mutation() *EpisodeMutation {
	return ec.mutation
}

// Save creates the Episode in the database.
func (ec *EpisodeCreate) Save(ctx context.Context) (*Episode, error) {
	var (
		err  error
		node *Episode
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EpisodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EpisodeCreate) SaveX(ctx context.Context) *Episode {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EpisodeCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EpisodeCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EpisodeCreate) defaults() {
	if _, ok := ec.mutation.Sid(); !ok {
		v := episode.DefaultSid()
		ec.mutation.SetSid(v)
	}
	if _, ok := ec.mutation.CreateTime(); !ok {
		v := episode.DefaultCreateTime()
		ec.mutation.SetCreateTime(v)
	}
	if _, ok := ec.mutation.UpdateTime(); !ok {
		v := episode.DefaultUpdateTime()
		ec.mutation.SetUpdateTime(v)
	}
	if _, ok := ec.mutation.Files(); !ok {
		v := episode.DefaultFiles()
		ec.mutation.SetFiles(v)
	}
	if _, ok := ec.mutation.ID(); !ok {
		v := episode.DefaultID()
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EpisodeCreate) check() error {
	if _, ok := ec.mutation.Sid(); !ok {
		return &ValidationError{Name: "sid", err: errors.New(`ent: missing required field "Episode.sid"`)}
	}
	if _, ok := ec.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Episode.create_time"`)}
	}
	if _, ok := ec.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Episode.update_time"`)}
	}
	if v, ok := ec.mutation.Cover(); ok {
		if err := episode.CoverValidator(v); err != nil {
			return &ValidationError{Name: "cover", err: fmt.Errorf(`ent: validator failed for field "Episode.cover": %w`, err)}
		}
	}
	if v, ok := ec.mutation.Title(); ok {
		if err := episode.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Episode.title": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Episode.content"`)}
	}
	if v, ok := ec.mutation.Content(); ok {
		if err := episode.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Episode.content": %w`, err)}
		}
	}
	if _, ok := ec.mutation.Files(); !ok {
		return &ValidationError{Name: "files", err: errors.New(`ent: missing required field "Episode.files"`)}
	}
	if _, ok := ec.mutation.ProfileID(); !ok {
		return &ValidationError{Name: "profile", err: errors.New(`ent: missing required edge "Episode.profile"`)}
	}
	return nil
}

func (ec *EpisodeCreate) sqlSave(ctx context.Context) (*Episode, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
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

func (ec *EpisodeCreate) createSpec() (*Episode, *sqlgraph.CreateSpec) {
	var (
		_node = &Episode{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: episode.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: episode.FieldID,
			},
		}
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ec.mutation.Sid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: episode.FieldSid,
		})
		_node.Sid = value
	}
	if value, ok := ec.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: episode.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ec.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: episode.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ec.mutation.Cover(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldCover,
		})
		_node.Cover = &value
	}
	if value, ok := ec.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldTitle,
		})
		_node.Title = &value
	}
	if value, ok := ec.mutation.Content(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldContent,
		})
		_node.Content = value
	}
	if value, ok := ec.mutation.Files(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBytes,
			Value:  value,
			Column: episode.FieldFiles,
		})
		_node.Files = value
	}
	if nodes := ec.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.ProfileTable,
			Columns: []string{episode.ProfileColumn},
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
		_node.profile_episode = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   episode.CommentsTable,
			Columns: []string{episode.CommentsColumn},
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
	if nodes := ec.mutation.SeriesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.SeriesTable,
			Columns: []string{episode.SeriesColumn},
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
		_node.series_episodes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EpisodeCreateBulk is the builder for creating many Episode entities in bulk.
type EpisodeCreateBulk struct {
	config
	builders []*EpisodeCreate
}

// Save creates the Episode entities in the database.
func (ecb *EpisodeCreateBulk) Save(ctx context.Context) ([]*Episode, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Episode, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EpisodeMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EpisodeCreateBulk) SaveX(ctx context.Context) []*Episode {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EpisodeCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EpisodeCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
