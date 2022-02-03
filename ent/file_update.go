// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"street/ent/file"
	"street/ent/predicate"
	"street/ent/profile"
	"street/ent/schema"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks    []Hook
	mutation *FileMutation
}

// Where appends a list predicates to the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetSid sets the "sid" field.
func (fu *FileUpdate) SetSid(s schema.ID) *FileUpdate {
	fu.mutation.SetSid(s)
	return fu
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (fu *FileUpdate) SetNillableSid(s *schema.ID) *FileUpdate {
	if s != nil {
		fu.SetSid(*s)
	}
	return fu
}

// SetFilename sets the "filename" field.
func (fu *FileUpdate) SetFilename(s string) *FileUpdate {
	fu.mutation.SetFilename(s)
	return fu
}

// SetNillableFilename sets the "filename" field if the given value is not nil.
func (fu *FileUpdate) SetNillableFilename(s *string) *FileUpdate {
	if s != nil {
		fu.SetFilename(*s)
	}
	return fu
}

// SetPath sets the "path" field.
func (fu *FileUpdate) SetPath(s string) *FileUpdate {
	fu.mutation.SetPath(s)
	return fu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (fu *FileUpdate) SetNillablePath(s *string) *FileUpdate {
	if s != nil {
		fu.SetPath(*s)
	}
	return fu
}

// SetMime sets the "mime" field.
func (fu *FileUpdate) SetMime(s string) *FileUpdate {
	fu.mutation.SetMime(s)
	return fu
}

// SetNillableMime sets the "mime" field if the given value is not nil.
func (fu *FileUpdate) SetNillableMime(s *string) *FileUpdate {
	if s != nil {
		fu.SetMime(*s)
	}
	return fu
}

// SetSize sets the "size" field.
func (fu *FileUpdate) SetSize(i int) *FileUpdate {
	fu.mutation.ResetSize()
	fu.mutation.SetSize(i)
	return fu
}

// AddSize adds i to the "size" field.
func (fu *FileUpdate) AddSize(i int) *FileUpdate {
	fu.mutation.AddSize(i)
	return fu
}

// SetStatus sets the "status" field.
func (fu *FileUpdate) SetStatus(s string) *FileUpdate {
	fu.mutation.SetStatus(s)
	return fu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fu *FileUpdate) SetNillableStatus(s *string) *FileUpdate {
	if s != nil {
		fu.SetStatus(*s)
	}
	return fu
}

// SetNote sets the "note" field.
func (fu *FileUpdate) SetNote(s string) *FileUpdate {
	fu.mutation.SetNote(s)
	return fu
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (fu *FileUpdate) SetNillableNote(s *string) *FileUpdate {
	if s != nil {
		fu.SetNote(*s)
	}
	return fu
}

// ClearNote clears the value of the "note" field.
func (fu *FileUpdate) ClearNote() *FileUpdate {
	fu.mutation.ClearNote()
	return fu
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (fu *FileUpdate) SetProfileID(id uuid.UUID) *FileUpdate {
	fu.mutation.SetProfileID(id)
	return fu
}

// SetProfile sets the "profile" edge to the Profile entity.
func (fu *FileUpdate) SetProfile(p *Profile) *FileUpdate {
	return fu.SetProfileID(p.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (fu *FileUpdate) ClearProfile() *FileUpdate {
	fu.mutation.ClearProfile()
	return fu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	fu.defaults()
	if len(fu.hooks) == 0 {
		if err = fu.check(); err != nil {
			return 0, err
		}
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fu.check(); err != nil {
				return 0, err
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FileUpdate) defaults() {
	if _, ok := fu.mutation.UpdateTime(); !ok {
		v := file.UpdateDefaultUpdateTime()
		fu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FileUpdate) check() error {
	if v, ok := fu.mutation.Filename(); ok {
		if err := file.FilenameValidator(v); err != nil {
			return &ValidationError{Name: "filename", err: fmt.Errorf("ent: validator failed for field \"filename\": %w", err)}
		}
	}
	if v, ok := fu.mutation.Path(); ok {
		if err := file.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf("ent: validator failed for field \"path\": %w", err)}
		}
	}
	if v, ok := fu.mutation.Mime(); ok {
		if err := file.MimeValidator(v); err != nil {
			return &ValidationError{Name: "mime", err: fmt.Errorf("ent: validator failed for field \"mime\": %w", err)}
		}
	}
	if v, ok := fu.mutation.Status(); ok {
		if err := file.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := fu.mutation.Note(); ok {
		if err := file.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if _, ok := fu.mutation.ProfileID(); fu.mutation.ProfileCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"profile\"")
	}
	return nil
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: file.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Sid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: file.FieldSid,
		})
	}
	if value, ok := fu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: file.FieldUpdateTime,
		})
	}
	if value, ok := fu.mutation.Filename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldFilename,
		})
	}
	if value, ok := fu.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldPath,
		})
	}
	if value, ok := fu.mutation.Mime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldMime,
		})
	}
	if value, ok := fu.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if value, ok := fu.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if value, ok := fu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldStatus,
		})
	}
	if value, ok := fu.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldNote,
		})
	}
	if fu.mutation.NoteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: file.FieldNote,
		})
	}
	if fu.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.ProfileTable,
			Columns: []string{file.ProfileColumn},
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
	if nodes := fu.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.ProfileTable,
			Columns: []string{file.ProfileColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileMutation
}

// SetSid sets the "sid" field.
func (fuo *FileUpdateOne) SetSid(s schema.ID) *FileUpdateOne {
	fuo.mutation.SetSid(s)
	return fuo
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableSid(s *schema.ID) *FileUpdateOne {
	if s != nil {
		fuo.SetSid(*s)
	}
	return fuo
}

// SetFilename sets the "filename" field.
func (fuo *FileUpdateOne) SetFilename(s string) *FileUpdateOne {
	fuo.mutation.SetFilename(s)
	return fuo
}

// SetNillableFilename sets the "filename" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableFilename(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetFilename(*s)
	}
	return fuo
}

// SetPath sets the "path" field.
func (fuo *FileUpdateOne) SetPath(s string) *FileUpdateOne {
	fuo.mutation.SetPath(s)
	return fuo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillablePath(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetPath(*s)
	}
	return fuo
}

// SetMime sets the "mime" field.
func (fuo *FileUpdateOne) SetMime(s string) *FileUpdateOne {
	fuo.mutation.SetMime(s)
	return fuo
}

// SetNillableMime sets the "mime" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableMime(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetMime(*s)
	}
	return fuo
}

// SetSize sets the "size" field.
func (fuo *FileUpdateOne) SetSize(i int) *FileUpdateOne {
	fuo.mutation.ResetSize()
	fuo.mutation.SetSize(i)
	return fuo
}

// AddSize adds i to the "size" field.
func (fuo *FileUpdateOne) AddSize(i int) *FileUpdateOne {
	fuo.mutation.AddSize(i)
	return fuo
}

// SetStatus sets the "status" field.
func (fuo *FileUpdateOne) SetStatus(s string) *FileUpdateOne {
	fuo.mutation.SetStatus(s)
	return fuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableStatus(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetStatus(*s)
	}
	return fuo
}

// SetNote sets the "note" field.
func (fuo *FileUpdateOne) SetNote(s string) *FileUpdateOne {
	fuo.mutation.SetNote(s)
	return fuo
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableNote(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetNote(*s)
	}
	return fuo
}

// ClearNote clears the value of the "note" field.
func (fuo *FileUpdateOne) ClearNote() *FileUpdateOne {
	fuo.mutation.ClearNote()
	return fuo
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (fuo *FileUpdateOne) SetProfileID(id uuid.UUID) *FileUpdateOne {
	fuo.mutation.SetProfileID(id)
	return fuo
}

// SetProfile sets the "profile" edge to the Profile entity.
func (fuo *FileUpdateOne) SetProfile(p *Profile) *FileUpdateOne {
	return fuo.SetProfileID(p.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// ClearProfile clears the "profile" edge to the Profile entity.
func (fuo *FileUpdateOne) ClearProfile() *FileUpdateOne {
	fuo.mutation.ClearProfile()
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	var (
		err  error
		node *File
	)
	fuo.defaults()
	if len(fuo.hooks) == 0 {
		if err = fuo.check(); err != nil {
			return nil, err
		}
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fuo.check(); err != nil {
				return nil, err
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FileUpdateOne) defaults() {
	if _, ok := fuo.mutation.UpdateTime(); !ok {
		v := file.UpdateDefaultUpdateTime()
		fuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FileUpdateOne) check() error {
	if v, ok := fuo.mutation.Filename(); ok {
		if err := file.FilenameValidator(v); err != nil {
			return &ValidationError{Name: "filename", err: fmt.Errorf("ent: validator failed for field \"filename\": %w", err)}
		}
	}
	if v, ok := fuo.mutation.Path(); ok {
		if err := file.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf("ent: validator failed for field \"path\": %w", err)}
		}
	}
	if v, ok := fuo.mutation.Mime(); ok {
		if err := file.MimeValidator(v); err != nil {
			return &ValidationError{Name: "mime", err: fmt.Errorf("ent: validator failed for field \"mime\": %w", err)}
		}
	}
	if v, ok := fuo.mutation.Status(); ok {
		if err := file.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf("ent: validator failed for field \"status\": %w", err)}
		}
	}
	if v, ok := fuo.mutation.Note(); ok {
		if err := file.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf("ent: validator failed for field \"note\": %w", err)}
		}
	}
	if _, ok := fuo.mutation.ProfileID(); fuo.mutation.ProfileCleared() && !ok {
		return errors.New("ent: clearing a required unique edge \"profile\"")
	}
	return nil
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (_node *File, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   file.Table,
			Columns: file.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: file.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing File.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for _, f := range fields {
			if !file.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Sid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: file.FieldSid,
		})
	}
	if value, ok := fuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: file.FieldUpdateTime,
		})
	}
	if value, ok := fuo.mutation.Filename(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldFilename,
		})
	}
	if value, ok := fuo.mutation.Path(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldPath,
		})
	}
	if value, ok := fuo.mutation.Mime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldMime,
		})
	}
	if value, ok := fuo.mutation.Size(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if value, ok := fuo.mutation.AddedSize(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
	}
	if value, ok := fuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldStatus,
		})
	}
	if value, ok := fuo.mutation.Note(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldNote,
		})
	}
	if fuo.mutation.NoteCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: file.FieldNote,
		})
	}
	if fuo.mutation.ProfileCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.ProfileTable,
			Columns: []string{file.ProfileColumn},
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
	if nodes := fuo.mutation.ProfileIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.ProfileTable,
			Columns: []string{file.ProfileColumn},
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
	_node = &File{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}