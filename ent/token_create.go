// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"street/ent/account"
	"street/ent/schema"
	"street/ent/token"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TokenCreate is the builder for creating a Token entity.
type TokenCreate struct {
	config
	mutation *TokenMutation
	hooks    []Hook
}

// SetSID sets the "SID" field.
func (tc *TokenCreate) SetSID(s schema.ID) *TokenCreate {
	tc.mutation.SetSID(s)
	return tc
}

// SetNillableSID sets the "SID" field if the given value is not nil.
func (tc *TokenCreate) SetNillableSID(s *schema.ID) *TokenCreate {
	if s != nil {
		tc.SetSID(*s)
	}
	return tc
}

// SetCreateTime sets the "create_time" field.
func (tc *TokenCreate) SetCreateTime(t time.Time) *TokenCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *TokenCreate) SetNillableCreateTime(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetUpdateTime sets the "update_time" field.
func (tc *TokenCreate) SetUpdateTime(t time.Time) *TokenCreate {
	tc.mutation.SetUpdateTime(t)
	return tc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (tc *TokenCreate) SetNillableUpdateTime(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetUpdateTime(*t)
	}
	return tc
}

// SetBody sets the "body" field.
func (tc *TokenCreate) SetBody(s string) *TokenCreate {
	tc.mutation.SetBody(s)
	return tc
}

// SetType sets the "type" field.
func (tc *TokenCreate) SetType(s string) *TokenCreate {
	tc.mutation.SetType(s)
	return tc
}

// SetExpireTime sets the "expire_time" field.
func (tc *TokenCreate) SetExpireTime(t time.Time) *TokenCreate {
	tc.mutation.SetExpireTime(t)
	return tc
}

// SetID sets the "id" field.
func (tc *TokenCreate) SetID(u uuid.UUID) *TokenCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetAccountID sets the "account" edge to the Account entity by ID.
func (tc *TokenCreate) SetAccountID(id uuid.UUID) *TokenCreate {
	tc.mutation.SetAccountID(id)
	return tc
}

// SetAccount sets the "account" edge to the Account entity.
func (tc *TokenCreate) SetAccount(a *Account) *TokenCreate {
	return tc.SetAccountID(a.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tc *TokenCreate) Mutation() *TokenMutation {
	return tc.mutation
}

// Save creates the Token in the database.
func (tc *TokenCreate) Save(ctx context.Context) (*Token, error) {
	var (
		err  error
		node *Token
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokenCreate) SaveX(ctx context.Context) *Token {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TokenCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TokenCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TokenCreate) defaults() {
	if _, ok := tc.mutation.SID(); !ok {
		v := token.DefaultSID()
		tc.mutation.SetSID(v)
	}
	if _, ok := tc.mutation.CreateTime(); !ok {
		v := token.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		v := token.DefaultUpdateTime()
		tc.mutation.SetUpdateTime(v)
	}
	if _, ok := tc.mutation.ID(); !ok {
		v := token.DefaultID()
		tc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TokenCreate) check() error {
	if v, ok := tc.mutation.SID(); ok {
		if err := token.SIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "SID", err: fmt.Errorf(`ent: validator failed for field "SID": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "create_time"`)}
	}
	if _, ok := tc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "update_time"`)}
	}
	if _, ok := tc.mutation.Body(); !ok {
		return &ValidationError{Name: "body", err: errors.New(`ent: missing required field "body"`)}
	}
	if v, ok := tc.mutation.Body(); ok {
		if err := token.BodyValidator(v); err != nil {
			return &ValidationError{Name: "body", err: fmt.Errorf(`ent: validator failed for field "body": %w`, err)}
		}
	}
	if _, ok := tc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "type"`)}
	}
	if v, ok := tc.mutation.GetType(); ok {
		if err := token.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "type": %w`, err)}
		}
	}
	if _, ok := tc.mutation.ExpireTime(); !ok {
		return &ValidationError{Name: "expire_time", err: errors.New(`ent: missing required field "expire_time"`)}
	}
	if _, ok := tc.mutation.AccountID(); !ok {
		return &ValidationError{Name: "account", err: errors.New("ent: missing required edge \"account\"")}
	}
	return nil
}

func (tc *TokenCreate) sqlSave(ctx context.Context) (*Token, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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

func (tc *TokenCreate) createSpec() (*Token, *sqlgraph.CreateSpec) {
	var (
		_node = &Token{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: token.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: token.FieldID,
			},
		}
	)
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := tc.mutation.SID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldSID,
		})
		_node.SID = &value
	}
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: token.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := tc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: token.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := tc.mutation.Body(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldBody,
		})
		_node.Body = value
	}
	if value, ok := tc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: token.FieldType,
		})
		_node.Type = value
	}
	if value, ok := tc.mutation.ExpireTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: token.FieldExpireTime,
		})
		_node.ExpireTime = value
	}
	if nodes := tc.mutation.AccountIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.AccountTable,
			Columns: []string{token.AccountColumn},
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
		_node.account_token = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TokenCreateBulk is the builder for creating many Token entities in bulk.
type TokenCreateBulk struct {
	config
	builders []*TokenCreate
}

// Save creates the Token entities in the database.
func (tcb *TokenCreateBulk) Save(ctx context.Context) ([]*Token, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Token, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokenMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TokenCreateBulk) SaveX(ctx context.Context) []*Token {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TokenCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TokenCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}
