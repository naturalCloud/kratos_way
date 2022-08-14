// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"
	"user/internal/data/ent/users"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UsersCreate is the builder for creating a Users entity.
type UsersCreate struct {
	config
	mutation *UsersMutation
	hooks    []Hook
}

// SetMobile sets the "mobile" field.
func (uc *UsersCreate) SetMobile(s string) *UsersCreate {
	uc.mutation.SetMobile(s)
	return uc
}

// SetPassword sets the "password" field.
func (uc *UsersCreate) SetPassword(s string) *UsersCreate {
	uc.mutation.SetPassword(s)
	return uc
}

// SetNickName sets the "nick_name" field.
func (uc *UsersCreate) SetNickName(s string) *UsersCreate {
	uc.mutation.SetNickName(s)
	return uc
}

// SetNillableNickName sets the "nick_name" field if the given value is not nil.
func (uc *UsersCreate) SetNillableNickName(s *string) *UsersCreate {
	if s != nil {
		uc.SetNickName(*s)
	}
	return uc
}

// SetBirthday sets the "birthday" field.
func (uc *UsersCreate) SetBirthday(t time.Time) *UsersCreate {
	uc.mutation.SetBirthday(t)
	return uc
}

// SetNillableBirthday sets the "birthday" field if the given value is not nil.
func (uc *UsersCreate) SetNillableBirthday(t *time.Time) *UsersCreate {
	if t != nil {
		uc.SetBirthday(*t)
	}
	return uc
}

// SetGender sets the "gender" field.
func (uc *UsersCreate) SetGender(s string) *UsersCreate {
	uc.mutation.SetGender(s)
	return uc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (uc *UsersCreate) SetNillableGender(s *string) *UsersCreate {
	if s != nil {
		uc.SetGender(*s)
	}
	return uc
}

// SetRole sets the "role" field.
func (uc *UsersCreate) SetRole(i int32) *UsersCreate {
	uc.mutation.SetRole(i)
	return uc
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uc *UsersCreate) SetNillableRole(i *int32) *UsersCreate {
	if i != nil {
		uc.SetRole(*i)
	}
	return uc
}

// SetAddTime sets the "add_time" field.
func (uc *UsersCreate) SetAddTime(t time.Time) *UsersCreate {
	uc.mutation.SetAddTime(t)
	return uc
}

// SetNillableAddTime sets the "add_time" field if the given value is not nil.
func (uc *UsersCreate) SetNillableAddTime(t *time.Time) *UsersCreate {
	if t != nil {
		uc.SetAddTime(*t)
	}
	return uc
}

// SetUpdateTime sets the "update_time" field.
func (uc *UsersCreate) SetUpdateTime(t time.Time) *UsersCreate {
	uc.mutation.SetUpdateTime(t)
	return uc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (uc *UsersCreate) SetNillableUpdateTime(t *time.Time) *UsersCreate {
	if t != nil {
		uc.SetUpdateTime(*t)
	}
	return uc
}

// SetDeletedAt sets the "deleted_at" field.
func (uc *UsersCreate) SetDeletedAt(t time.Time) *UsersCreate {
	uc.mutation.SetDeletedAt(t)
	return uc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uc *UsersCreate) SetNillableDeletedAt(t *time.Time) *UsersCreate {
	if t != nil {
		uc.SetDeletedAt(*t)
	}
	return uc
}

// SetIsDeletedAt sets the "is_deleted_at" field.
func (uc *UsersCreate) SetIsDeletedAt(i int8) *UsersCreate {
	uc.mutation.SetIsDeletedAt(i)
	return uc
}

// SetNillableIsDeletedAt sets the "is_deleted_at" field if the given value is not nil.
func (uc *UsersCreate) SetNillableIsDeletedAt(i *int8) *UsersCreate {
	if i != nil {
		uc.SetIsDeletedAt(*i)
	}
	return uc
}

// SetID sets the "id" field.
func (uc *UsersCreate) SetID(i int64) *UsersCreate {
	uc.mutation.SetID(i)
	return uc
}

// Mutation returns the UsersMutation object of the builder.
func (uc *UsersCreate) Mutation() *UsersMutation {
	return uc.mutation
}

// Save creates the Users in the database.
func (uc *UsersCreate) Save(ctx context.Context) (*Users, error) {
	var (
		err  error
		node *Users
	)
	uc.defaults()
	if len(uc.hooks) == 0 {
		if err = uc.check(); err != nil {
			return nil, err
		}
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UsersMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = uc.check(); err != nil {
				return nil, err
			}
			uc.mutation = mutation
			if node, err = uc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			if uc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Users)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UsersMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UsersCreate) SaveX(ctx context.Context) *Users {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UsersCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UsersCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UsersCreate) defaults() {
	if _, ok := uc.mutation.Gender(); !ok {
		v := users.DefaultGender
		uc.mutation.SetGender(v)
	}
	if _, ok := uc.mutation.Role(); !ok {
		v := users.DefaultRole
		uc.mutation.SetRole(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uc *UsersCreate) check() error {
	if _, ok := uc.mutation.Mobile(); !ok {
		return &ValidationError{Name: "mobile", err: errors.New(`ent: missing required field "Users.mobile"`)}
	}
	if _, ok := uc.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Users.password"`)}
	}
	return nil
}

func (uc *UsersCreate) sqlSave(ctx context.Context) (*Users, error) {
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int64(id)
	}
	return _node, nil
}

func (uc *UsersCreate) createSpec() (*Users, *sqlgraph.CreateSpec) {
	var (
		_node = &Users{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: users.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: users.FieldID,
			},
		}
	)
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.Mobile(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldMobile,
		})
		_node.Mobile = value
	}
	if value, ok := uc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldPassword,
		})
		_node.Password = value
	}
	if value, ok := uc.mutation.NickName(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldNickName,
		})
		_node.NickName = value
	}
	if value, ok := uc.mutation.Birthday(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: users.FieldBirthday,
		})
		_node.Birthday = value
	}
	if value, ok := uc.mutation.Gender(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: users.FieldGender,
		})
		_node.Gender = value
	}
	if value, ok := uc.mutation.Role(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt32,
			Value:  value,
			Column: users.FieldRole,
		})
		_node.Role = value
	}
	if value, ok := uc.mutation.AddTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: users.FieldAddTime,
		})
		_node.AddTime = value
	}
	if value, ok := uc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: users.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := uc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: users.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if value, ok := uc.mutation.IsDeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt8,
			Value:  value,
			Column: users.FieldIsDeletedAt,
		})
		_node.IsDeletedAt = value
	}
	return _node, _spec
}

// UsersCreateBulk is the builder for creating many Users entities in bulk.
type UsersCreateBulk struct {
	config
	builders []*UsersCreate
}

// Save creates the Users entities in the database.
func (ucb *UsersCreateBulk) Save(ctx context.Context) ([]*Users, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*Users, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UsersMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UsersCreateBulk) SaveX(ctx context.Context) []*Users {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UsersCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UsersCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}
