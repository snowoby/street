// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"street/ent/episode"
	"street/ent/predicate"
	"street/ent/profile"
	"street/ent/schema"
	"street/ent/series"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// EpisodeUpdate is the builder for updating Episode entities.
type EpisodeUpdate struct {
	config
	hooks    []Hook
	mutation *EpisodeMutation
}

// Where appends a list predicates to the EpisodeUpdate builder.
func (eu *EpisodeUpdate) Where(ps ...predicate.Episode) *EpisodeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetSid sets the "sid" field.
func (eu *EpisodeUpdate) SetSid(s schema.ID) *EpisodeUpdate {
	eu.mutation.SetSid(s)
	return eu
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableSid(s *schema.ID) *EpisodeUpdate {
	if s != nil {
		eu.SetSid(*s)
	}
	return eu
}

// SetTitle sets the "title" field.
func (eu *EpisodeUpdate) SetTitle(s string) *EpisodeUpdate {
	eu.mutation.SetTitle(s)
	return eu
}

// SetContent sets the "content" field.
func (eu *EpisodeUpdate) SetContent(s string) *EpisodeUpdate {
	eu.mutation.SetContent(s)
	return eu
}

// SetCover sets the "cover" field.
func (eu *EpisodeUpdate) SetCover(s string) *EpisodeUpdate {
	eu.mutation.SetCover(s)
	return eu
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableCover(s *string) *EpisodeUpdate {
	if s != nil {
		eu.SetCover(*s)
	}
	return eu
}

// ClearCover clears the value of the "cover" field.
func (eu *EpisodeUpdate) ClearCover() *EpisodeUpdate {
	eu.mutation.ClearCover()
	return eu
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (eu *EpisodeUpdate) SetProfileID(id uuid.UUID) *EpisodeUpdate {
	eu.mutation.SetProfileID(id)
	return eu
}

// SetProfile sets the "profile" edge to the Profile entity.
func (eu *EpisodeUpdate) SetProfile(p *Profile) *EpisodeUpdate {
	return eu.SetProfileID(p.ID)
}

// SetSeriesID sets the "series" edge to the Series entity by ID.
func (eu *EpisodeUpdate) SetSeriesID(id uuid.UUID) *EpisodeUpdate {
	eu.mutation.SetSeriesID(id)
	return eu
}

// SetNillableSeriesID sets the "series" edge to the Series entity by ID if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableSeriesID(id *uuid.UUID) *EpisodeUpdate {
	if id != nil {
		eu = eu.SetSeriesID(*id)
	}
	return eu
}

// SetSeries sets the "series" edge to the Series entity.
func (eu *EpisodeUpdate) SetSeries(s *Series) *EpisodeUpdate {
	return eu.SetSeriesID(s.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (eu *EpisodeUpdate) Mutation() *EpisodeMutation {
	return eu.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (eu *EpisodeUpdate) ClearProfile() *EpisodeUpdate {
	eu.mutation.ClearProfile()
	return eu
}

// ClearSeries clears the "series" edge to the Series entity.
func (eu *EpisodeUpdate) ClearSeries() *EpisodeUpdate {
	eu.mutation.ClearSeries()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EpisodeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	eu.defaults()
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EpisodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EpisodeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EpisodeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EpisodeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EpisodeUpdate) defaults() {
	if _, ok := eu.mutation.UpdateTime(); !ok {
		v := episode.UpdateDefaultUpdateTime()
		eu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EpisodeUpdate) check() error {
	if v, ok := eu.mutation.Title(); ok {
		if err := episode.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := eu.mutation.Content(); ok {
		if err := episode.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf("ent: validator failed for field \"content\": %w", err)}
		}
	}
	if v, ok := eu.mutation.Cover(); ok {
		if err := episode.CoverValidator(v); err != nil {
			return &ValidationError{Name: "cover", err: fmt.Errorf("ent: validator failed for field \"cover\": %w", err)}
		}
	}
	if _, ok := eu.mutation.ProfileID(); eu.mutation.ProfileCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"profile\"")
	}
	return nil
}

func (eu *EpisodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   episode.Table,
			Columns: episode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: episode.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Sid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: episode.FieldSid,
		})
	}
	if value, ok := eu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: episode.FieldUpdateTime,
		})
	}
	if value, ok := eu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldTitle,
		})
	}
	if value, ok := eu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldContent,
		})
	}
	if value, ok := eu.mutation.Cover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldCover,
		})
	}
	if eu.mutation.CoverCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: episode.FieldCover,
		})
	}
	if eu.mutation.ProfileCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ProfileIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.SeriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.SeriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{episode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EpisodeUpdateOne is the builder for updating a single Episode entity.
type EpisodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EpisodeMutation
}

// SetSid sets the "sid" field.
func (euo *EpisodeUpdateOne) SetSid(s schema.ID) *EpisodeUpdateOne {
	euo.mutation.SetSid(s)
	return euo
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableSid(s *schema.ID) *EpisodeUpdateOne {
	if s != nil {
		euo.SetSid(*s)
	}
	return euo
}

// SetTitle sets the "title" field.
func (euo *EpisodeUpdateOne) SetTitle(s string) *EpisodeUpdateOne {
	euo.mutation.SetTitle(s)
	return euo
}

// SetContent sets the "content" field.
func (euo *EpisodeUpdateOne) SetContent(s string) *EpisodeUpdateOne {
	euo.mutation.SetContent(s)
	return euo
}

// SetCover sets the "cover" field.
func (euo *EpisodeUpdateOne) SetCover(s string) *EpisodeUpdateOne {
	euo.mutation.SetCover(s)
	return euo
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableCover(s *string) *EpisodeUpdateOne {
	if s != nil {
		euo.SetCover(*s)
	}
	return euo
}

// ClearCover clears the value of the "cover" field.
func (euo *EpisodeUpdateOne) ClearCover() *EpisodeUpdateOne {
	euo.mutation.ClearCover()
	return euo
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (euo *EpisodeUpdateOne) SetProfileID(id uuid.UUID) *EpisodeUpdateOne {
	euo.mutation.SetProfileID(id)
	return euo
}

// SetProfile sets the "profile" edge to the Profile entity.
func (euo *EpisodeUpdateOne) SetProfile(p *Profile) *EpisodeUpdateOne {
	return euo.SetProfileID(p.ID)
}

// SetSeriesID sets the "series" edge to the Series entity by ID.
func (euo *EpisodeUpdateOne) SetSeriesID(id uuid.UUID) *EpisodeUpdateOne {
	euo.mutation.SetSeriesID(id)
	return euo
}

// SetNillableSeriesID sets the "series" edge to the Series entity by ID if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableSeriesID(id *uuid.UUID) *EpisodeUpdateOne {
	if id != nil {
		euo = euo.SetSeriesID(*id)
	}
	return euo
}

// SetSeries sets the "series" edge to the Series entity.
func (euo *EpisodeUpdateOne) SetSeries(s *Series) *EpisodeUpdateOne {
	return euo.SetSeriesID(s.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (euo *EpisodeUpdateOne) Mutation() *EpisodeMutation {
	return euo.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (euo *EpisodeUpdateOne) ClearProfile() *EpisodeUpdateOne {
	euo.mutation.ClearProfile()
	return euo
}

// ClearSeries clears the "series" edge to the Series entity.
func (euo *EpisodeUpdateOne) ClearSeries() *EpisodeUpdateOne {
	euo.mutation.ClearSeries()
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EpisodeUpdateOne) Select(field string, fields ...string) *EpisodeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Episode entity.
func (euo *EpisodeUpdateOne) Save(ctx context.Context) (*Episode, error) {
	var (
		err  error
		node *Episode
	)
	euo.defaults()
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EpisodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EpisodeUpdateOne) SaveX(ctx context.Context) *Episode {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EpisodeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EpisodeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EpisodeUpdateOne) defaults() {
	if _, ok := euo.mutation.UpdateTime(); !ok {
		v := episode.UpdateDefaultUpdateTime()
		euo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EpisodeUpdateOne) check() error {
	if v, ok := euo.mutation.Title(); ok {
		if err := episode.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := euo.mutation.Content(); ok {
		if err := episode.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf("ent: validator failed for field \"content\": %w", err)}
		}
	}
	if v, ok := euo.mutation.Cover(); ok {
		if err := episode.CoverValidator(v); err != nil {
			return &ValidationError{Name: "cover", err: fmt.Errorf("ent: validator failed for field \"cover\": %w", err)}
		}
	}
	if _, ok := euo.mutation.ProfileID(); euo.mutation.ProfileCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"profile\"")
	}
	return nil
}

func (euo *EpisodeUpdateOne) sqlSave(ctx context.Context) (_node *Episode, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   episode.Table,
			Columns: episode.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: episode.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Episode.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, episode.FieldID)
		for _, f := range fields {
			if !episode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != episode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Sid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: episode.FieldSid,
		})
	}
	if value, ok := euo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: episode.FieldUpdateTime,
		})
	}
	if value, ok := euo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldTitle,
		})
	}
	if value, ok := euo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldContent,
		})
	}
	if value, ok := euo.mutation.Cover(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: episode.FieldCover,
		})
	}
	if euo.mutation.CoverCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: episode.FieldCover,
		})
	}
	if euo.mutation.ProfileCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ProfileIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.SeriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.SeriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Episode{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{episode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
