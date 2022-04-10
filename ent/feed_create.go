// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"polunzh/my-feed/ent/feed"
	"polunzh/my-feed/ent/group"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FeedCreate is the builder for creating a Feed entity.
type FeedCreate struct {
	config
	mutation *FeedMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (fc *FeedCreate) SetName(s string) *FeedCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fc *FeedCreate) SetNillableName(s *string) *FeedCreate {
	if s != nil {
		fc.SetName(*s)
	}
	return fc
}

// SetURL sets the "url" field.
func (fc *FeedCreate) SetURL(s string) *FeedCreate {
	fc.mutation.SetURL(s)
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FeedCreate) SetUpdatedAt(t time.Time) *FeedCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FeedCreate) SetNillableUpdatedAt(t *time.Time) *FeedCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FeedCreate) SetCreatedAt(t time.Time) *FeedCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FeedCreate) SetNillableCreatedAt(t *time.Time) *FeedCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// AddGroupIDs adds the "group" edge to the Group entity by IDs.
func (fc *FeedCreate) AddGroupIDs(ids ...int) *FeedCreate {
	fc.mutation.AddGroupIDs(ids...)
	return fc
}

// AddGroup adds the "group" edges to the Group entity.
func (fc *FeedCreate) AddGroup(g ...*Group) *FeedCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return fc.AddGroupIDs(ids...)
}

// Mutation returns the FeedMutation object of the builder.
func (fc *FeedCreate) Mutation() *FeedMutation {
	return fc.mutation
}

// Save creates the Feed in the database.
func (fc *FeedCreate) Save(ctx context.Context) (*Feed, error) {
	var (
		err  error
		node *Feed
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FeedMutation)
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
func (fc *FeedCreate) SaveX(ctx context.Context) *Feed {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FeedCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FeedCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FeedCreate) defaults() {
	if _, ok := fc.mutation.Name(); !ok {
		v := feed.DefaultName
		fc.mutation.SetName(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := feed.DefaultUpdatedAt
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := feed.DefaultCreatedAt
		fc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FeedCreate) check() error {
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Feed.name"`)}
	}
	if _, ok := fc.mutation.URL(); !ok {
		return &ValidationError{Name: "url", err: errors.New(`ent: missing required field "Feed.url"`)}
	}
	if v, ok := fc.mutation.URL(); ok {
		if err := feed.URLValidator(v); err != nil {
			return &ValidationError{Name: "url", err: fmt.Errorf(`ent: validator failed for field "Feed.url": %w`, err)}
		}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Feed.updated_at"`)}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Feed.created_at"`)}
	}
	return nil
}

func (fc *FeedCreate) sqlSave(ctx context.Context) (*Feed, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fc *FeedCreate) createSpec() (*Feed, *sqlgraph.CreateSpec) {
	var (
		_node = &Feed{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: feed.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: feed.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldName,
		})
		_node.Name = value
	}
	if value, ok := fc.mutation.URL(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: feed.FieldURL,
		})
		_node.URL = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: feed.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if nodes := fc.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feed.GroupTable,
			Columns: []string{feed.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
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

// FeedCreateBulk is the builder for creating many Feed entities in bulk.
type FeedCreateBulk struct {
	config
	builders []*FeedCreate
}

// Save creates the Feed entities in the database.
func (fcb *FeedCreateBulk) Save(ctx context.Context) ([]*Feed, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Feed, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeedMutation)
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
func (fcb *FeedCreateBulk) SaveX(ctx context.Context) []*Feed {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FeedCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FeedCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
