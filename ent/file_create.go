// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"street/ent/file"
	"street/ent/profile"
	"street/ent/schema"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
}

// SetSid sets the "sid" field.
func (fc *FileCreate) SetSid(s schema.ID) *FileCreate {
	fc.mutation.SetSid(s)
	return fc
}

// SetNillableSid sets the "sid" field if the given value is not nil.
func (fc *FileCreate) SetNillableSid(s *schema.ID) *FileCreate {
	if s != nil {
		fc.SetSid(*s)
	}
	return fc
}

// SetCreateTime sets the "create_time" field.
func (fc *FileCreate) SetCreateTime(t time.Time) *FileCreate {
	fc.mutation.SetCreateTime(t)
	return fc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (fc *FileCreate) SetNillableCreateTime(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetCreateTime(*t)
	}
	return fc
}

// SetUpdateTime sets the "update_time" field.
func (fc *FileCreate) SetUpdateTime(t time.Time) *FileCreate {
	fc.mutation.SetUpdateTime(t)
	return fc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (fc *FileCreate) SetNillableUpdateTime(t *time.Time) *FileCreate {
	if t != nil {
		fc.SetUpdateTime(*t)
	}
	return fc
}

// SetFilename sets the "filename" field.
func (fc *FileCreate) SetFilename(s string) *FileCreate {
	fc.mutation.SetFilename(s)
	return fc
}

// SetNillableFilename sets the "filename" field if the given value is not nil.
func (fc *FileCreate) SetNillableFilename(s *string) *FileCreate {
	if s != nil {
		fc.SetFilename(*s)
	}
	return fc
}

// SetPath sets the "path" field.
func (fc *FileCreate) SetPath(s string) *FileCreate {
	fc.mutation.SetPath(s)
	return fc
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (fc *FileCreate) SetNillablePath(s *string) *FileCreate {
	if s != nil {
		fc.SetPath(*s)
	}
	return fc
}

// SetMime sets the "mime" field.
func (fc *FileCreate) SetMime(s string) *FileCreate {
	fc.mutation.SetMime(s)
	return fc
}

// SetNillableMime sets the "mime" field if the given value is not nil.
func (fc *FileCreate) SetNillableMime(s *string) *FileCreate {
	if s != nil {
		fc.SetMime(*s)
	}
	return fc
}

// SetSize sets the "size" field.
func (fc *FileCreate) SetSize(i int) *FileCreate {
	fc.mutation.SetSize(i)
	return fc
}

// SetStatus sets the "status" field.
func (fc *FileCreate) SetStatus(s string) *FileCreate {
	fc.mutation.SetStatus(s)
	return fc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (fc *FileCreate) SetNillableStatus(s *string) *FileCreate {
	if s != nil {
		fc.SetStatus(*s)
	}
	return fc
}

// SetNote sets the "note" field.
func (fc *FileCreate) SetNote(s string) *FileCreate {
	fc.mutation.SetNote(s)
	return fc
}

// SetNillableNote sets the "note" field if the given value is not nil.
func (fc *FileCreate) SetNillableNote(s *string) *FileCreate {
	if s != nil {
		fc.SetNote(*s)
	}
	return fc
}

// SetID sets the "id" field.
func (fc *FileCreate) SetID(u uuid.UUID) *FileCreate {
	fc.mutation.SetID(u)
	return fc
}

// SetProfileID sets the "profile" edge to the Profile entity by ID.
func (fc *FileCreate) SetProfileID(id uuid.UUID) *FileCreate {
	fc.mutation.SetProfileID(id)
	return fc
}

// SetProfile sets the "profile" edge to the Profile entity.
func (fc *FileCreate) SetProfile(p *Profile) *FileCreate {
	return fc.SetProfileID(p.ID)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	var (
		err  error
		node *File
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			if node, err = fc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			if fc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FileCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FileCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() {
	if _, ok := fc.mutation.Sid(); !ok {
		v := file.DefaultSid()
		fc.mutation.SetSid(v)
	}
	if _, ok := fc.mutation.CreateTime(); !ok {
		v := file.DefaultCreateTime()
		fc.mutation.SetCreateTime(v)
	}
	if _, ok := fc.mutation.UpdateTime(); !ok {
		v := file.DefaultUpdateTime()
		fc.mutation.SetUpdateTime(v)
	}
	if _, ok := fc.mutation.Filename(); !ok {
		v := file.DefaultFilename
		fc.mutation.SetFilename(v)
	}
	if _, ok := fc.mutation.Path(); !ok {
		v := file.DefaultPath
		fc.mutation.SetPath(v)
	}
	if _, ok := fc.mutation.Mime(); !ok {
		v := file.DefaultMime
		fc.mutation.SetMime(v)
	}
	if _, ok := fc.mutation.Status(); !ok {
		v := file.DefaultStatus
		fc.mutation.SetStatus(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := file.DefaultID()
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.Sid(); !ok {
		return &ValidationError{Name: "sid", err: errors.New(`ent: missing required field "sid"`)}
	}
	if _, ok := fc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := fc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if _, ok := fc.mutation.Filename(); !ok {
		return &ValidationError{Name: "filename", err: errors.New(`ent: missing required field "filename"`)}
	}
	if v, ok := fc.mutation.Filename(); ok {
		if err := file.FilenameValidator(v); err != nil {
			return &ValidationError{Name: "filename", err: fmt.Errorf(`ent: validator failed for field "filename": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`ent: missing required field "path"`)}
	}
	if v, ok := fc.mutation.Path(); ok {
		if err := file.PathValidator(v); err != nil {
			return &ValidationError{Name: "path", err: fmt.Errorf(`ent: validator failed for field "path": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Mime(); !ok {
		return &ValidationError{Name: "mime", err: errors.New(`ent: missing required field "mime"`)}
	}
	if v, ok := fc.mutation.Mime(); ok {
		if err := file.MimeValidator(v); err != nil {
			return &ValidationError{Name: "mime", err: fmt.Errorf(`ent: validator failed for field "mime": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "size"`)}
	}
	if _, ok := fc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "status"`)}
	}
	if v, ok := fc.mutation.Status(); ok {
		if err := file.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "status": %w`, err)}
		}
	}
	if v, ok := fc.mutation.Note(); ok {
		if err := file.NoteValidator(v); err != nil {
			return &ValidationError{Name: "note", err: fmt.Errorf(`ent: validator failed for field "note": %w`, err)}
		}
	}
	if _, ok := fc.mutation.ProfileID(); !ok {
		return &ValidationError{Name: "profile", err: errors.New("ent: missing required edge \"profile\"")}
	}
	return nil
}

func (fc *FileCreate) sqlSave(ctx context.Context) (*File, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
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

func (fc *FileCreate) createSpec() (*File, *sqlgraph.CreateSpec) {
	var (
		_node = &File{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: file.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: file.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.Sid(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: file.FieldSid,
		})
		_node.Sid = value
	}
	if value, ok := fc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: file.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := fc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: file.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := fc.mutation.Filename(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldFilename,
		})
		_node.Filename = value
	}
	if value, ok := fc.mutation.Path(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldPath,
		})
		_node.Path = value
	}
	if value, ok := fc.mutation.Mime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldMime,
		})
		_node.Mime = value
	}
	if value, ok := fc.mutation.Size(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: file.FieldSize,
		})
		_node.Size = value
	}
	if value, ok := fc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldStatus,
		})
		_node.Status = value
	}
	if value, ok := fc.mutation.Note(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: file.FieldNote,
		})
		_node.Note = value
	}
	if nodes := fc.mutation.ProfileIDs(); len(nodes) > 0 {
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
		_node.profile_file = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	builders []*FileCreate
}

// Save creates the File entities in the database.
func (fcb *FileCreateBulk) Save(ctx context.Context) ([]*File, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*File, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FileMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FileCreateBulk) SaveX(ctx context.Context) []*File {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FileCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FileCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
