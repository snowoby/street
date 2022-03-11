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
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SeriesUpdate is the builder for updating Series entities.
type SeriesUpdate struct {
	config
	hooks    []Hook
	mutation *SeriesMutation
}

// Where appends a list predicates to the SeriesUpdate builder.
func (su *SeriesUpdate) Where(ps ...predicate.Series) *SeriesUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetSid sets the "sid" field.
func (su *SeriesUpdate) SetSid(s schema.ID) *SeriesUpdate {
	su.mutation.SetSid(s)
	return su
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (su *SeriesUpdate) SetNillableSid(s *schema.ID) *SeriesUpdate {
	if s != nil {
		su.SetSid(*s)
	}
	return su
}

// SetUpdateTime sets the "update_time" field.
func (su *SeriesUpdate) SetUpdateTime(t time.Time) *SeriesUpdate {
	su.mutation.SetUpdateTime(t)
	return su
}

// SetTitle sets the "title" field.
func (su *SeriesUpdate) SetTitle(s string) *SeriesUpdate {
	su.mutation.SetTitle(s)
	return su
}

// SetType sets the "type" field.
func (su *SeriesUpdate) SetType(s string) *SeriesUpdate {
	su.mutation.SetType(s)
	return su
}

// SetNillableType sets the "type" field if the given value is not nil.
func (su *SeriesUpdate) SetNillableType(s *string) *SeriesUpdate {
	if s != nil {
		su.SetType(*s)
	}
	return su
}

// ClearType clears the value of the "type" field.
func (su *SeriesUpdate) ClearType() *SeriesUpdate {
	su.mutation.ClearType()
	return su
}

// AddEpisodeIDs adds the "episodes" edge to the Episode entity by IDs.
func (su *SeriesUpdate) AddEpisodeIDs(ids ...uuid.UUID) *SeriesUpdate {
	su.mutation.AddEpisodeIDs(ids...)
	return su
}

// AddEpisodes adds the "episodes" edges to the Episode entity.
func (su *SeriesUpdate) AddEpisodes(e ...*Episode) *SeriesUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.AddEpisodeIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Profile entity by ID.
func (su *SeriesUpdate) SetOwnerID(id uuid.UUID) *SeriesUpdate {
	su.mutation.SetOwnerID(id)
	return su
}

// SetOwner sets the "owner" edge to the Profile entity.
func (su *SeriesUpdate) SetOwner(p *Profile) *SeriesUpdate {
	return su.SetOwnerID(p.ID)
}

// AddParticipantIDs adds the "participant" edge to the Profile entity by IDs.
func (su *SeriesUpdate) AddParticipantIDs(ids ...uuid.UUID) *SeriesUpdate {
	su.mutation.AddParticipantIDs(ids...)
	return su
}

// AddParticipant adds the "participant" edges to the Profile entity.
func (su *SeriesUpdate) AddParticipant(p ...*Profile) *SeriesUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.AddParticipantIDs(ids...)
}

// Mutation returns the SeriesMutation object of the builder.
func (su *SeriesUpdate) Mutation() *SeriesMutation {
	return su.mutation
}

// ClearEpisodes clears all "episodes" edges to the Episode entity.
func (su *SeriesUpdate) ClearEpisodes() *SeriesUpdate {
	su.mutation.ClearEpisodes()
	return su
}

// RemoveEpisodeIDs removes the "episodes" edge to Episode entities by IDs.
func (su *SeriesUpdate) RemoveEpisodeIDs(ids ...uuid.UUID) *SeriesUpdate {
	su.mutation.RemoveEpisodeIDs(ids...)
	return su
}

// RemoveEpisodes removes "episodes" edges to Episode entities.
func (su *SeriesUpdate) RemoveEpisodes(e ...*Episode) *SeriesUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return su.RemoveEpisodeIDs(ids...)
}

// ClearOwner clears the "owner" edge to the Profile entity.
func (su *SeriesUpdate) ClearOwner() *SeriesUpdate {
	su.mutation.ClearOwner()
	return su
}

// ClearParticipant clears all "participant" edges to the Profile entity.
func (su *SeriesUpdate) ClearParticipant() *SeriesUpdate {
	su.mutation.ClearParticipant()
	return su
}

// RemoveParticipantIDs removes the "participant" edge to Profile entities by IDs.
func (su *SeriesUpdate) RemoveParticipantIDs(ids ...uuid.UUID) *SeriesUpdate {
	su.mutation.RemoveParticipantIDs(ids...)
	return su
}

// RemoveParticipant removes "participant" edges to Profile entities.
func (su *SeriesUpdate) RemoveParticipant(p ...*Profile) *SeriesUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return su.RemoveParticipantIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SeriesUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	su.defaults()
	if len(su.hooks) == 0 {
		if err = su.check(); err != nil {
			return 0, err
		}
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = su.check(); err != nil {
				return 0, err
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SeriesUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SeriesUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SeriesUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (su *SeriesUpdate) defaults() {
	if _, ok := su.mutation.UpdateTime(); !ok {
		v := series.UpdateDefaultUpdateTime()
		su.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (su *SeriesUpdate) check() error {
	if v, ok := su.mutation.Title(); ok {
		if err := series.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Series.title": %w`, err)}
		}
	}
	if v, ok := su.mutation.GetType(); ok {
		if err := series.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Series.type": %w`, err)}
		}
	}
	if _, ok := su.mutation.OwnerID(); su.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Series.owner"`)
	}
	return nil
}

func (su *SeriesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   series.Table,
			Columns: series.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: series.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Sid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: series.FieldSid,
		})
	}
	if value, ok := su.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: series.FieldUpdateTime,
		})
	}
	if value, ok := su.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: series.FieldTitle,
		})
	}
	if value, ok := su.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: series.FieldType,
		})
	}
	if su.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: series.FieldType,
		})
	}
	if su.mutation.EpisodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedEpisodesIDs(); len(nodes) > 0 && !su.mutation.EpisodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.EpisodesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.ParticipantCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedParticipantIDs(); len(nodes) > 0 && !su.mutation.ParticipantCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ParticipantIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{series.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// SeriesUpdateOne is the builder for updating a single Series entity.
type SeriesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SeriesMutation
}

// SetSid sets the "sid" field.
func (suo *SeriesUpdateOne) SetSid(s schema.ID) *SeriesUpdateOne {
	suo.mutation.SetSid(s)
	return suo
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillableSid(s *schema.ID) *SeriesUpdateOne {
	if s != nil {
		suo.SetSid(*s)
	}
	return suo
}

// SetUpdateTime sets the "update_time" field.
func (suo *SeriesUpdateOne) SetUpdateTime(t time.Time) *SeriesUpdateOne {
	suo.mutation.SetUpdateTime(t)
	return suo
}

// SetTitle sets the "title" field.
func (suo *SeriesUpdateOne) SetTitle(s string) *SeriesUpdateOne {
	suo.mutation.SetTitle(s)
	return suo
}

// SetType sets the "type" field.
func (suo *SeriesUpdateOne) SetType(s string) *SeriesUpdateOne {
	suo.mutation.SetType(s)
	return suo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (suo *SeriesUpdateOne) SetNillableType(s *string) *SeriesUpdateOne {
	if s != nil {
		suo.SetType(*s)
	}
	return suo
}

// ClearType clears the value of the "type" field.
func (suo *SeriesUpdateOne) ClearType() *SeriesUpdateOne {
	suo.mutation.ClearType()
	return suo
}

// AddEpisodeIDs adds the "episodes" edge to the Episode entity by IDs.
func (suo *SeriesUpdateOne) AddEpisodeIDs(ids ...uuid.UUID) *SeriesUpdateOne {
	suo.mutation.AddEpisodeIDs(ids...)
	return suo
}

// AddEpisodes adds the "episodes" edges to the Episode entity.
func (suo *SeriesUpdateOne) AddEpisodes(e ...*Episode) *SeriesUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.AddEpisodeIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the Profile entity by ID.
func (suo *SeriesUpdateOne) SetOwnerID(id uuid.UUID) *SeriesUpdateOne {
	suo.mutation.SetOwnerID(id)
	return suo
}

// SetOwner sets the "owner" edge to the Profile entity.
func (suo *SeriesUpdateOne) SetOwner(p *Profile) *SeriesUpdateOne {
	return suo.SetOwnerID(p.ID)
}

// AddParticipantIDs adds the "participant" edge to the Profile entity by IDs.
func (suo *SeriesUpdateOne) AddParticipantIDs(ids ...uuid.UUID) *SeriesUpdateOne {
	suo.mutation.AddParticipantIDs(ids...)
	return suo
}

// AddParticipant adds the "participant" edges to the Profile entity.
func (suo *SeriesUpdateOne) AddParticipant(p ...*Profile) *SeriesUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.AddParticipantIDs(ids...)
}

// Mutation returns the SeriesMutation object of the builder.
func (suo *SeriesUpdateOne) Mutation() *SeriesMutation {
	return suo.mutation
}

// ClearEpisodes clears all "episodes" edges to the Episode entity.
func (suo *SeriesUpdateOne) ClearEpisodes() *SeriesUpdateOne {
	suo.mutation.ClearEpisodes()
	return suo
}

// RemoveEpisodeIDs removes the "episodes" edge to Episode entities by IDs.
func (suo *SeriesUpdateOne) RemoveEpisodeIDs(ids ...uuid.UUID) *SeriesUpdateOne {
	suo.mutation.RemoveEpisodeIDs(ids...)
	return suo
}

// RemoveEpisodes removes "episodes" edges to Episode entities.
func (suo *SeriesUpdateOne) RemoveEpisodes(e ...*Episode) *SeriesUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return suo.RemoveEpisodeIDs(ids...)
}

// ClearOwner clears the "owner" edge to the Profile entity.
func (suo *SeriesUpdateOne) ClearOwner() *SeriesUpdateOne {
	suo.mutation.ClearOwner()
	return suo
}

// ClearParticipant clears all "participant" edges to the Profile entity.
func (suo *SeriesUpdateOne) ClearParticipant() *SeriesUpdateOne {
	suo.mutation.ClearParticipant()
	return suo
}

// RemoveParticipantIDs removes the "participant" edge to Profile entities by IDs.
func (suo *SeriesUpdateOne) RemoveParticipantIDs(ids ...uuid.UUID) *SeriesUpdateOne {
	suo.mutation.RemoveParticipantIDs(ids...)
	return suo
}

// RemoveParticipant removes "participant" edges to Profile entities.
func (suo *SeriesUpdateOne) RemoveParticipant(p ...*Profile) *SeriesUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return suo.RemoveParticipantIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SeriesUpdateOne) Select(field string, fields ...string) *SeriesUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Series entity.
func (suo *SeriesUpdateOne) Save(ctx context.Context) (*Series, error) {
	var (
		err  error
		node *Series
	)
	suo.defaults()
	if len(suo.hooks) == 0 {
		if err = suo.check(); err != nil {
			return nil, err
		}
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SeriesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = suo.check(); err != nil {
				return nil, err
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *SeriesUpdateOne) SaveX(ctx context.Context) *Series {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SeriesUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SeriesUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (suo *SeriesUpdateOne) defaults() {
	if _, ok := suo.mutation.UpdateTime(); !ok {
		v := series.UpdateDefaultUpdateTime()
		suo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (suo *SeriesUpdateOne) check() error {
	if v, ok := suo.mutation.Title(); ok {
		if err := series.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Series.title": %w`, err)}
		}
	}
	if v, ok := suo.mutation.GetType(); ok {
		if err := series.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Series.type": %w`, err)}
		}
	}
	if _, ok := suo.mutation.OwnerID(); suo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Series.owner"`)
	}
	return nil
}

func (suo *SeriesUpdateOne) sqlSave(ctx context.Context) (_node *Series, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   series.Table,
			Columns: series.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: series.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Series.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, series.FieldID)
		for _, f := range fields {
			if !series.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != series.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.Sid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: series.FieldSid,
		})
	}
	if value, ok := suo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: series.FieldUpdateTime,
		})
	}
	if value, ok := suo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: series.FieldTitle,
		})
	}
	if value, ok := suo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: series.FieldType,
		})
	}
	if suo.mutation.TypeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: series.FieldType,
		})
	}
	if suo.mutation.EpisodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedEpisodesIDs(); len(nodes) > 0 && !suo.mutation.EpisodesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.EpisodesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.ParticipantCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedParticipantIDs(); len(nodes) > 0 && !suo.mutation.ParticipantCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ParticipantIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Series{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{series.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
