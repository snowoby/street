// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"street/ent/account"
	"street/ent/episode"
	"street/ent/file"
	"street/ent/predicate"
	"street/ent/profile"
	"street/ent/schema"
	"street/ent/series"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ProfileUpdate is the builder for updating Profile entities.
type ProfileUpdate struct {
	config
	hooks    []Hook
	mutation *ProfileMutation
}

// Where appends a list predicates to the ProfileUpdate builder.
func (pu *ProfileUpdate) Where(ps ...predicate.Profile) *ProfileUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetSid sets the "sid" field.
func (pu *ProfileUpdate) SetSid(s schema.ID) *ProfileUpdate {
	pu.mutation.SetSid(s)
	return pu
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableSid(s *schema.ID) *ProfileUpdate {
	if s != nil {
		pu.SetSid(*s)
	}
	return pu
}

// SetTitle sets the "title" field.
func (pu *ProfileUpdate) SetTitle(s string) *ProfileUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetCall sets the "call" field.
func (pu *ProfileUpdate) SetCall(s string) *ProfileUpdate {
	pu.mutation.SetCall(s)
	return pu
}

// SetCategory sets the "category" field.
func (pu *ProfileUpdate) SetCategory(s string) *ProfileUpdate {
	pu.mutation.SetCategory(s)
	return pu
}

// SetAvatar sets the "avatar" field.
func (pu *ProfileUpdate) SetAvatar(s string) *ProfileUpdate {
	pu.mutation.SetAvatar(s)
	return pu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (pu *ProfileUpdate) SetNillableAvatar(s *string) *ProfileUpdate {
	if s != nil {
		pu.SetAvatar(*s)
	}
	return pu
}

// ClearAvatar clears the value of the "avatar" field.
func (pu *ProfileUpdate) ClearAvatar() *ProfileUpdate {
	pu.mutation.ClearAvatar()
	return pu
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (pu *ProfileUpdate) SetAccountID(id uuid.UUID) *ProfileUpdate {
	pu.mutation.SetAccountID(id)
	return pu
}

// SetAccount sets the "account" edge to the Account entity.
func (pu *ProfileUpdate) SetAccount(a *Account) *ProfileUpdate {
	return pu.SetAccountID(a.ID)
}

// AddEpisodeIDs adds the "episode" edge to the Episode entity by IDs.
func (pu *ProfileUpdate) AddEpisodeIDs(ids ...uuid.UUID) *ProfileUpdate {
	pu.mutation.AddEpisodeIDs(ids...)
	return pu
}

// AddEpisode adds the "episode" edges to the Episode entity.
func (pu *ProfileUpdate) AddEpisode(e ...*Episode) *ProfileUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pu.AddEpisodeIDs(ids...)
}

// AddSeriesIDs adds the "series" edge to the Series entity by IDs.
func (pu *ProfileUpdate) AddSeriesIDs(ids ...uuid.UUID) *ProfileUpdate {
	pu.mutation.AddSeriesIDs(ids...)
	return pu
}

// AddSeries adds the "series" edges to the Series entity.
func (pu *ProfileUpdate) AddSeries(s ...*Series) *ProfileUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddSeriesIDs(ids...)
}

// AddFileIDs adds the "file" edge to the File entity by IDs.
func (pu *ProfileUpdate) AddFileIDs(ids ...uuid.UUID) *ProfileUpdate {
	pu.mutation.AddFileIDs(ids...)
	return pu
}

// AddFile adds the "file" edges to the File entity.
func (pu *ProfileUpdate) AddFile(f ...*File) *ProfileUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pu.AddFileIDs(ids...)
}

// Mutation returns the ProfileMutation object of the builder.
func (pu *ProfileUpdate) Mutation() *ProfileMutation {
	return pu.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (pu *ProfileUpdate) ClearAccount() *ProfileUpdate {
	pu.mutation.ClearAccount()
	return pu
}

// ClearEpisode clears all "episode" edges to the Episode entity.
func (pu *ProfileUpdate) ClearEpisode() *ProfileUpdate {
	pu.mutation.ClearEpisode()
	return pu
}

// RemoveEpisodeIDs removes the "episode" edge to Episode entities by IDs.
func (pu *ProfileUpdate) RemoveEpisodeIDs(ids ...uuid.UUID) *ProfileUpdate {
	pu.mutation.RemoveEpisodeIDs(ids...)
	return pu
}

// RemoveEpisode removes "episode" edges to Episode entities.
func (pu *ProfileUpdate) RemoveEpisode(e ...*Episode) *ProfileUpdate {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pu.RemoveEpisodeIDs(ids...)
}

// ClearSeries clears all "series" edges to the Series entity.
func (pu *ProfileUpdate) ClearSeries() *ProfileUpdate {
	pu.mutation.ClearSeries()
	return pu
}

// RemoveSeriesIDs removes the "series" edge to Series entities by IDs.
func (pu *ProfileUpdate) RemoveSeriesIDs(ids ...uuid.UUID) *ProfileUpdate {
	pu.mutation.RemoveSeriesIDs(ids...)
	return pu
}

// RemoveSeries removes "series" edges to Series entities.
func (pu *ProfileUpdate) RemoveSeries(s ...*Series) *ProfileUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveSeriesIDs(ids...)
}

// ClearFile clears all "file" edges to the File entity.
func (pu *ProfileUpdate) ClearFile() *ProfileUpdate {
	pu.mutation.ClearFile()
	return pu
}

// RemoveFileIDs removes the "file" edge to File entities by IDs.
func (pu *ProfileUpdate) RemoveFileIDs(ids ...uuid.UUID) *ProfileUpdate {
	pu.mutation.RemoveFileIDs(ids...)
	return pu
}

// RemoveFile removes "file" edges to File entities.
func (pu *ProfileUpdate) RemoveFile(f ...*File) *ProfileUpdate {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pu.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProfileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfileUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfileUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfileUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProfileUpdate) defaults() {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		v := profile.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProfileUpdate) check() error {
	if v, ok := pu.mutation.Title(); ok {
		if err := profile.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Call(); ok {
		if err := profile.CallValidator(v); err != nil {
			return &ValidationError{Name: "call", err: fmt.Errorf("ent: validator failed for field \"call\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Category(); ok {
		if err := profile.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf("ent: validator failed for field \"category\": %w", err)}
		}
	}
	if v, ok := pu.mutation.Avatar(); ok {
		if err := profile.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf("ent: validator failed for field \"avatar\": %w", err)}
		}
	}
	if _, ok := pu.mutation.AccountID(); pu.mutation.AccountCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"account\"")
	}
	return nil
}

func (pu *ProfileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profile.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Sid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: profile.FieldSid,
		})
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdateTime,
		})
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldTitle,
		})
	}
	if value, ok := pu.mutation.Call(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldCall,
		})
	}
	if value, ok := pu.mutation.Category(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldCategory,
		})
	}
	if value, ok := pu.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldAvatar,
		})
	}
	if pu.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldAvatar,
		})
	}
	if pu.mutation.AccountCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.AccountIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.EpisodeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedEpisodeIDs(); len(nodes) > 0 && !pu.mutation.EpisodeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.EpisodeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SeriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedSeriesIDs(); len(nodes) > 0 && !pu.mutation.SeriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SeriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.FileTable,
			Columns: []string{profile.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: file.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedFileIDs(); len(nodes) > 0 && !pu.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.FileTable,
			Columns: []string{profile.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.FileTable,
			Columns: []string{profile.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProfileUpdateOne is the builder for updating a single Profile entity.
type ProfileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProfileMutation
}

// SetSid sets the "sid" field.
func (puo *ProfileUpdateOne) SetSid(s schema.ID) *ProfileUpdateOne {
	puo.mutation.SetSid(s)
	return puo
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableSid(s *schema.ID) *ProfileUpdateOne {
	if s != nil {
		puo.SetSid(*s)
	}
	return puo
}

// SetTitle sets the "title" field.
func (puo *ProfileUpdateOne) SetTitle(s string) *ProfileUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetCall sets the "call" field.
func (puo *ProfileUpdateOne) SetCall(s string) *ProfileUpdateOne {
	puo.mutation.SetCall(s)
	return puo
}

// SetCategory sets the "category" field.
func (puo *ProfileUpdateOne) SetCategory(s string) *ProfileUpdateOne {
	puo.mutation.SetCategory(s)
	return puo
}

// SetAvatar sets the "avatar" field.
func (puo *ProfileUpdateOne) SetAvatar(s string) *ProfileUpdateOne {
	puo.mutation.SetAvatar(s)
	return puo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (puo *ProfileUpdateOne) SetNillableAvatar(s *string) *ProfileUpdateOne {
	if s != nil {
		puo.SetAvatar(*s)
	}
	return puo
}

// ClearAvatar clears the value of the "avatar" field.
func (puo *ProfileUpdateOne) ClearAvatar() *ProfileUpdateOne {
	puo.mutation.ClearAvatar()
	return puo
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (puo *ProfileUpdateOne) SetAccountID(id uuid.UUID) *ProfileUpdateOne {
	puo.mutation.SetAccountID(id)
	return puo
}

// SetAccount sets the "account" edge to the Account entity.
func (puo *ProfileUpdateOne) SetAccount(a *Account) *ProfileUpdateOne {
	return puo.SetAccountID(a.ID)
}

// AddEpisodeIDs adds the "episode" edge to the Episode entity by IDs.
func (puo *ProfileUpdateOne) AddEpisodeIDs(ids ...uuid.UUID) *ProfileUpdateOne {
	puo.mutation.AddEpisodeIDs(ids...)
	return puo
}

// AddEpisode adds the "episode" edges to the Episode entity.
func (puo *ProfileUpdateOne) AddEpisode(e ...*Episode) *ProfileUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return puo.AddEpisodeIDs(ids...)
}

// AddSeriesIDs adds the "series" edge to the Series entity by IDs.
func (puo *ProfileUpdateOne) AddSeriesIDs(ids ...uuid.UUID) *ProfileUpdateOne {
	puo.mutation.AddSeriesIDs(ids...)
	return puo
}

// AddSeries adds the "series" edges to the Series entity.
func (puo *ProfileUpdateOne) AddSeries(s ...*Series) *ProfileUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddSeriesIDs(ids...)
}

// AddFileIDs adds the "file" edge to the File entity by IDs.
func (puo *ProfileUpdateOne) AddFileIDs(ids ...uuid.UUID) *ProfileUpdateOne {
	puo.mutation.AddFileIDs(ids...)
	return puo
}

// AddFile adds the "file" edges to the File entity.
func (puo *ProfileUpdateOne) AddFile(f ...*File) *ProfileUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return puo.AddFileIDs(ids...)
}

// Mutation returns the ProfileMutation object of the builder.
func (puo *ProfileUpdateOne) Mutation() *ProfileMutation {
	return puo.mutation
}

// ClearAccount clears the "account" edge to the Account entity.
func (puo *ProfileUpdateOne) ClearAccount() *ProfileUpdateOne {
	puo.mutation.ClearAccount()
	return puo
}

// ClearEpisode clears all "episode" edges to the Episode entity.
func (puo *ProfileUpdateOne) ClearEpisode() *ProfileUpdateOne {
	puo.mutation.ClearEpisode()
	return puo
}

// RemoveEpisodeIDs removes the "episode" edge to Episode entities by IDs.
func (puo *ProfileUpdateOne) RemoveEpisodeIDs(ids ...uuid.UUID) *ProfileUpdateOne {
	puo.mutation.RemoveEpisodeIDs(ids...)
	return puo
}

// RemoveEpisode removes "episode" edges to Episode entities.
func (puo *ProfileUpdateOne) RemoveEpisode(e ...*Episode) *ProfileUpdateOne {
	ids := make([]uuid.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return puo.RemoveEpisodeIDs(ids...)
}

// ClearSeries clears all "series" edges to the Series entity.
func (puo *ProfileUpdateOne) ClearSeries() *ProfileUpdateOne {
	puo.mutation.ClearSeries()
	return puo
}

// RemoveSeriesIDs removes the "series" edge to Series entities by IDs.
func (puo *ProfileUpdateOne) RemoveSeriesIDs(ids ...uuid.UUID) *ProfileUpdateOne {
	puo.mutation.RemoveSeriesIDs(ids...)
	return puo
}

// RemoveSeries removes "series" edges to Series entities.
func (puo *ProfileUpdateOne) RemoveSeries(s ...*Series) *ProfileUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveSeriesIDs(ids...)
}

// ClearFile clears all "file" edges to the File entity.
func (puo *ProfileUpdateOne) ClearFile() *ProfileUpdateOne {
	puo.mutation.ClearFile()
	return puo
}

// RemoveFileIDs removes the "file" edge to File entities by IDs.
func (puo *ProfileUpdateOne) RemoveFileIDs(ids ...uuid.UUID) *ProfileUpdateOne {
	puo.mutation.RemoveFileIDs(ids...)
	return puo
}

// RemoveFile removes "file" edges to File entities.
func (puo *ProfileUpdateOne) RemoveFile(f ...*File) *ProfileUpdateOne {
	ids := make([]uuid.UUID, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return puo.RemoveFileIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProfileUpdateOne) Select(field string, fields ...string) *ProfileUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Profile entity.
func (puo *ProfileUpdateOne) Save(ctx context.Context) (*Profile, error) {
	var (
		err  error
		node *Profile
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProfileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfileUpdateOne) SaveX(ctx context.Context) *Profile {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProfileUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfileUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProfileUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		v := profile.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProfileUpdateOne) check() error {
	if v, ok := puo.mutation.Title(); ok {
		if err := profile.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf("ent: validator failed for field \"title\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Call(); ok {
		if err := profile.CallValidator(v); err != nil {
			return &ValidationError{Name: "call", err: fmt.Errorf("ent: validator failed for field \"call\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Category(); ok {
		if err := profile.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf("ent: validator failed for field \"category\": %w", err)}
		}
	}
	if v, ok := puo.mutation.Avatar(); ok {
		if err := profile.AvatarValidator(v); err != nil {
			return &ValidationError{Name: "avatar", err: fmt.Errorf("ent: validator failed for field \"avatar\": %w", err)}
		}
	}
	if _, ok := puo.mutation.AccountID(); puo.mutation.AccountCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"account\"")
	}
	return nil
}

func (puo *ProfileUpdateOne) sqlSave(ctx context.Context) (_node *Profile, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   profile.Table,
			Columns: profile.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: profile.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Profile.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, profile.FieldID)
		for _, f := range fields {
			if !profile.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != profile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Sid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: profile.FieldSid,
		})
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: profile.FieldUpdateTime,
		})
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldTitle,
		})
	}
	if value, ok := puo.mutation.Call(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldCall,
		})
	}
	if value, ok := puo.mutation.Category(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldCategory,
		})
	}
	if value, ok := puo.mutation.Avatar(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: profile.FieldAvatar,
		})
	}
	if puo.mutation.AvatarCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: profile.FieldAvatar,
		})
	}
	if puo.mutation.AccountCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.AccountIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.EpisodeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedEpisodeIDs(); len(nodes) > 0 && !puo.mutation.EpisodeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.EpisodeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SeriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedSeriesIDs(); len(nodes) > 0 && !puo.mutation.SeriesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SeriesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.FileTable,
			Columns: []string{profile.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: file.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedFileIDs(); len(nodes) > 0 && !puo.mutation.FileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.FileTable,
			Columns: []string{profile.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.FileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   profile.FileTable,
			Columns: []string{profile.FileColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Profile{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{profile.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
