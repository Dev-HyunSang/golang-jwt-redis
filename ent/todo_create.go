// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/dev-hyunsang/golang-jwt-redis/ent/todo"
	"github.com/google/uuid"
)

// ToDoCreate is the builder for creating a ToDo entity.
type ToDoCreate struct {
	config
	mutation *ToDoMutation
	hooks    []Hook
}

// SetTodoUUID sets the "todo_uuid" field.
func (tdc *ToDoCreate) SetTodoUUID(u uuid.UUID) *ToDoCreate {
	tdc.mutation.SetTodoUUID(u)
	return tdc
}

// SetNillableTodoUUID sets the "todo_uuid" field if the given value is not nil.
func (tdc *ToDoCreate) SetNillableTodoUUID(u *uuid.UUID) *ToDoCreate {
	if u != nil {
		tdc.SetTodoUUID(*u)
	}
	return tdc
}

// SetUserUUID sets the "user_uuid" field.
func (tdc *ToDoCreate) SetUserUUID(u uuid.UUID) *ToDoCreate {
	tdc.mutation.SetUserUUID(u)
	return tdc
}

// SetNillableUserUUID sets the "user_uuid" field if the given value is not nil.
func (tdc *ToDoCreate) SetNillableUserUUID(u *uuid.UUID) *ToDoCreate {
	if u != nil {
		tdc.SetUserUUID(*u)
	}
	return tdc
}

// SetTodoTitle sets the "todo_title" field.
func (tdc *ToDoCreate) SetTodoTitle(s string) *ToDoCreate {
	tdc.mutation.SetTodoTitle(s)
	return tdc
}

// SetTodoContext sets the "todo_context" field.
func (tdc *ToDoCreate) SetTodoContext(s string) *ToDoCreate {
	tdc.mutation.SetTodoContext(s)
	return tdc
}

// SetUpdatedAt sets the "updated_at" field.
func (tdc *ToDoCreate) SetUpdatedAt(t time.Time) *ToDoCreate {
	tdc.mutation.SetUpdatedAt(t)
	return tdc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tdc *ToDoCreate) SetNillableUpdatedAt(t *time.Time) *ToDoCreate {
	if t != nil {
		tdc.SetUpdatedAt(*t)
	}
	return tdc
}

// SetCratedAt sets the "crated_at" field.
func (tdc *ToDoCreate) SetCratedAt(t time.Time) *ToDoCreate {
	tdc.mutation.SetCratedAt(t)
	return tdc
}

// SetNillableCratedAt sets the "crated_at" field if the given value is not nil.
func (tdc *ToDoCreate) SetNillableCratedAt(t *time.Time) *ToDoCreate {
	if t != nil {
		tdc.SetCratedAt(*t)
	}
	return tdc
}

// Mutation returns the ToDoMutation object of the builder.
func (tdc *ToDoCreate) Mutation() *ToDoMutation {
	return tdc.mutation
}

// Save creates the ToDo in the database.
func (tdc *ToDoCreate) Save(ctx context.Context) (*ToDo, error) {
	var (
		err  error
		node *ToDo
	)
	tdc.defaults()
	if len(tdc.hooks) == 0 {
		if err = tdc.check(); err != nil {
			return nil, err
		}
		node, err = tdc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ToDoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tdc.check(); err != nil {
				return nil, err
			}
			tdc.mutation = mutation
			if node, err = tdc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tdc.hooks) - 1; i >= 0; i-- {
			if tdc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tdc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tdc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ToDo)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ToDoMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tdc *ToDoCreate) SaveX(ctx context.Context) *ToDo {
	v, err := tdc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tdc *ToDoCreate) Exec(ctx context.Context) error {
	_, err := tdc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdc *ToDoCreate) ExecX(ctx context.Context) {
	if err := tdc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tdc *ToDoCreate) defaults() {
	if _, ok := tdc.mutation.TodoUUID(); !ok {
		v := todo.DefaultTodoUUID()
		tdc.mutation.SetTodoUUID(v)
	}
	if _, ok := tdc.mutation.UserUUID(); !ok {
		v := todo.DefaultUserUUID()
		tdc.mutation.SetUserUUID(v)
	}
	if _, ok := tdc.mutation.UpdatedAt(); !ok {
		v := todo.DefaultUpdatedAt()
		tdc.mutation.SetUpdatedAt(v)
	}
	if _, ok := tdc.mutation.CratedAt(); !ok {
		v := todo.DefaultCratedAt()
		tdc.mutation.SetCratedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tdc *ToDoCreate) check() error {
	if _, ok := tdc.mutation.TodoUUID(); !ok {
		return &ValidationError{Name: "todo_uuid", err: errors.New(`ent: missing required field "ToDo.todo_uuid"`)}
	}
	if _, ok := tdc.mutation.UserUUID(); !ok {
		return &ValidationError{Name: "user_uuid", err: errors.New(`ent: missing required field "ToDo.user_uuid"`)}
	}
	if _, ok := tdc.mutation.TodoTitle(); !ok {
		return &ValidationError{Name: "todo_title", err: errors.New(`ent: missing required field "ToDo.todo_title"`)}
	}
	if _, ok := tdc.mutation.TodoContext(); !ok {
		return &ValidationError{Name: "todo_context", err: errors.New(`ent: missing required field "ToDo.todo_context"`)}
	}
	if _, ok := tdc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "ToDo.updated_at"`)}
	}
	if _, ok := tdc.mutation.CratedAt(); !ok {
		return &ValidationError{Name: "crated_at", err: errors.New(`ent: missing required field "ToDo.crated_at"`)}
	}
	return nil
}

func (tdc *ToDoCreate) sqlSave(ctx context.Context) (*ToDo, error) {
	_node, _spec := tdc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tdc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tdc *ToDoCreate) createSpec() (*ToDo, *sqlgraph.CreateSpec) {
	var (
		_node = &ToDo{config: tdc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: todo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: todo.FieldID,
			},
		}
	)
	if value, ok := tdc.mutation.TodoUUID(); ok {
		_spec.SetField(todo.FieldTodoUUID, field.TypeUUID, value)
		_node.TodoUUID = value
	}
	if value, ok := tdc.mutation.UserUUID(); ok {
		_spec.SetField(todo.FieldUserUUID, field.TypeUUID, value)
		_node.UserUUID = value
	}
	if value, ok := tdc.mutation.TodoTitle(); ok {
		_spec.SetField(todo.FieldTodoTitle, field.TypeString, value)
		_node.TodoTitle = value
	}
	if value, ok := tdc.mutation.TodoContext(); ok {
		_spec.SetField(todo.FieldTodoContext, field.TypeString, value)
		_node.TodoContext = value
	}
	if value, ok := tdc.mutation.UpdatedAt(); ok {
		_spec.SetField(todo.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := tdc.mutation.CratedAt(); ok {
		_spec.SetField(todo.FieldCratedAt, field.TypeTime, value)
		_node.CratedAt = value
	}
	return _node, _spec
}

// ToDoCreateBulk is the builder for creating many ToDo entities in bulk.
type ToDoCreateBulk struct {
	config
	builders []*ToDoCreate
}

// Save creates the ToDo entities in the database.
func (tdcb *ToDoCreateBulk) Save(ctx context.Context) ([]*ToDo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tdcb.builders))
	nodes := make([]*ToDo, len(tdcb.builders))
	mutators := make([]Mutator, len(tdcb.builders))
	for i := range tdcb.builders {
		func(i int, root context.Context) {
			builder := tdcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ToDoMutation)
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
					_, err = mutators[i+1].Mutate(root, tdcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tdcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
		if _, err := mutators[0].Mutate(ctx, tdcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tdcb *ToDoCreateBulk) SaveX(ctx context.Context) []*ToDo {
	v, err := tdcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tdcb *ToDoCreateBulk) Exec(ctx context.Context) error {
	_, err := tdcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tdcb *ToDoCreateBulk) ExecX(ctx context.Context) {
	if err := tdcb.Exec(ctx); err != nil {
		panic(err)
	}
}
