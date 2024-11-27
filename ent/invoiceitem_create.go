// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Tabintel/invoice-system/ent/invoiceitem"
)

// InvoiceItemCreate is the builder for creating a InvoiceItem entity.
type InvoiceItemCreate struct {
	config
	mutation *InvoiceItemMutation
	hooks    []Hook
}

// Mutation returns the InvoiceItemMutation object of the builder.
func (iic *InvoiceItemCreate) Mutation() *InvoiceItemMutation {
	return iic.mutation
}

// Save creates the InvoiceItem in the database.
func (iic *InvoiceItemCreate) Save(ctx context.Context) (*InvoiceItem, error) {
	return withHooks(ctx, iic.sqlSave, iic.mutation, iic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (iic *InvoiceItemCreate) SaveX(ctx context.Context) *InvoiceItem {
	v, err := iic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iic *InvoiceItemCreate) Exec(ctx context.Context) error {
	_, err := iic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iic *InvoiceItemCreate) ExecX(ctx context.Context) {
	if err := iic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iic *InvoiceItemCreate) check() error {
	return nil
}

func (iic *InvoiceItemCreate) sqlSave(ctx context.Context) (*InvoiceItem, error) {
	if err := iic.check(); err != nil {
		return nil, err
	}
	_node, _spec := iic.createSpec()
	if err := sqlgraph.CreateNode(ctx, iic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	iic.mutation.id = &_node.ID
	iic.mutation.done = true
	return _node, nil
}

func (iic *InvoiceItemCreate) createSpec() (*InvoiceItem, *sqlgraph.CreateSpec) {
	var (
		_node = &InvoiceItem{config: iic.config}
		_spec = sqlgraph.NewCreateSpec(invoiceitem.Table, sqlgraph.NewFieldSpec(invoiceitem.FieldID, field.TypeInt))
	)
	return _node, _spec
}

// InvoiceItemCreateBulk is the builder for creating many InvoiceItem entities in bulk.
type InvoiceItemCreateBulk struct {
	config
	err      error
	builders []*InvoiceItemCreate
}

// Save creates the InvoiceItem entities in the database.
func (iicb *InvoiceItemCreateBulk) Save(ctx context.Context) ([]*InvoiceItem, error) {
	if iicb.err != nil {
		return nil, iicb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(iicb.builders))
	nodes := make([]*InvoiceItem, len(iicb.builders))
	mutators := make([]Mutator, len(iicb.builders))
	for i := range iicb.builders {
		func(i int, root context.Context) {
			builder := iicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*InvoiceItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, iicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, iicb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, iicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (iicb *InvoiceItemCreateBulk) SaveX(ctx context.Context) []*InvoiceItem {
	v, err := iicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (iicb *InvoiceItemCreateBulk) Exec(ctx context.Context) error {
	_, err := iicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iicb *InvoiceItemCreateBulk) ExecX(ctx context.Context) {
	if err := iicb.Exec(ctx); err != nil {
		panic(err)
	}
}