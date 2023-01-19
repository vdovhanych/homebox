// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/hay-kot/homebox/backend/internal/data/ent/attachment"
	"github.com/hay-kot/homebox/backend/internal/data/ent/document"
	"github.com/hay-kot/homebox/backend/internal/data/ent/item"
	"github.com/hay-kot/homebox/backend/internal/data/ent/predicate"
)

// AttachmentUpdate is the builder for updating Attachment entities.
type AttachmentUpdate struct {
	config
	hooks    []Hook
	mutation *AttachmentMutation
}

// Where appends a list predicates to the AttachmentUpdate builder.
func (au *AttachmentUpdate) Where(ps ...predicate.Attachment) *AttachmentUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AttachmentUpdate) SetUpdatedAt(t time.Time) *AttachmentUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetType sets the "type" field.
func (au *AttachmentUpdate) SetType(a attachment.Type) *AttachmentUpdate {
	au.mutation.SetType(a)
	return au
}

// SetNillableType sets the "type" field if the given value is not nil.
func (au *AttachmentUpdate) SetNillableType(a *attachment.Type) *AttachmentUpdate {
	if a != nil {
		au.SetType(*a)
	}
	return au
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (au *AttachmentUpdate) SetItemID(id uuid.UUID) *AttachmentUpdate {
	au.mutation.SetItemID(id)
	return au
}

// SetItem sets the "item" edge to the Item entity.
func (au *AttachmentUpdate) SetItem(i *Item) *AttachmentUpdate {
	return au.SetItemID(i.ID)
}

// SetDocumentID sets the "document" edge to the Document entity by ID.
func (au *AttachmentUpdate) SetDocumentID(id uuid.UUID) *AttachmentUpdate {
	au.mutation.SetDocumentID(id)
	return au
}

// SetDocument sets the "document" edge to the Document entity.
func (au *AttachmentUpdate) SetDocument(d *Document) *AttachmentUpdate {
	return au.SetDocumentID(d.ID)
}

// Mutation returns the AttachmentMutation object of the builder.
func (au *AttachmentUpdate) Mutation() *AttachmentMutation {
	return au.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (au *AttachmentUpdate) ClearItem() *AttachmentUpdate {
	au.mutation.ClearItem()
	return au
}

// ClearDocument clears the "document" edge to the Document entity.
func (au *AttachmentUpdate) ClearDocument() *AttachmentUpdate {
	au.mutation.ClearDocument()
	return au
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AttachmentUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks[int, AttachmentMutation](ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AttachmentUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AttachmentUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AttachmentUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AttachmentUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := attachment.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AttachmentUpdate) check() error {
	if v, ok := au.mutation.GetType(); ok {
		if err := attachment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Attachment.type": %w`, err)}
		}
	}
	if _, ok := au.mutation.ItemID(); au.mutation.ItemCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Attachment.item"`)
	}
	if _, ok := au.mutation.DocumentID(); au.mutation.DocumentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Attachment.document"`)
	}
	return nil
}

func (au *AttachmentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attachment.Table,
			Columns: attachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attachment.FieldID,
			},
		},
	}
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(attachment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.GetType(); ok {
		_spec.SetField(attachment.FieldType, field.TypeEnum, value)
	}
	if au.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.ItemTable,
			Columns: []string{attachment.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.ItemTable,
			Columns: []string{attachment.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if au.mutation.DocumentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.DocumentTable,
			Columns: []string{attachment.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: document.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.DocumentTable,
			Columns: []string{attachment.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: document.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AttachmentUpdateOne is the builder for updating a single Attachment entity.
type AttachmentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AttachmentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AttachmentUpdateOne) SetUpdatedAt(t time.Time) *AttachmentUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetType sets the "type" field.
func (auo *AttachmentUpdateOne) SetType(a attachment.Type) *AttachmentUpdateOne {
	auo.mutation.SetType(a)
	return auo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (auo *AttachmentUpdateOne) SetNillableType(a *attachment.Type) *AttachmentUpdateOne {
	if a != nil {
		auo.SetType(*a)
	}
	return auo
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (auo *AttachmentUpdateOne) SetItemID(id uuid.UUID) *AttachmentUpdateOne {
	auo.mutation.SetItemID(id)
	return auo
}

// SetItem sets the "item" edge to the Item entity.
func (auo *AttachmentUpdateOne) SetItem(i *Item) *AttachmentUpdateOne {
	return auo.SetItemID(i.ID)
}

// SetDocumentID sets the "document" edge to the Document entity by ID.
func (auo *AttachmentUpdateOne) SetDocumentID(id uuid.UUID) *AttachmentUpdateOne {
	auo.mutation.SetDocumentID(id)
	return auo
}

// SetDocument sets the "document" edge to the Document entity.
func (auo *AttachmentUpdateOne) SetDocument(d *Document) *AttachmentUpdateOne {
	return auo.SetDocumentID(d.ID)
}

// Mutation returns the AttachmentMutation object of the builder.
func (auo *AttachmentUpdateOne) Mutation() *AttachmentMutation {
	return auo.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (auo *AttachmentUpdateOne) ClearItem() *AttachmentUpdateOne {
	auo.mutation.ClearItem()
	return auo
}

// ClearDocument clears the "document" edge to the Document entity.
func (auo *AttachmentUpdateOne) ClearDocument() *AttachmentUpdateOne {
	auo.mutation.ClearDocument()
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AttachmentUpdateOne) Select(field string, fields ...string) *AttachmentUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Attachment entity.
func (auo *AttachmentUpdateOne) Save(ctx context.Context) (*Attachment, error) {
	auo.defaults()
	return withHooks[*Attachment, AttachmentMutation](ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AttachmentUpdateOne) SaveX(ctx context.Context) *Attachment {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AttachmentUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AttachmentUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AttachmentUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := attachment.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AttachmentUpdateOne) check() error {
	if v, ok := auo.mutation.GetType(); ok {
		if err := attachment.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Attachment.type": %w`, err)}
		}
	}
	if _, ok := auo.mutation.ItemID(); auo.mutation.ItemCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Attachment.item"`)
	}
	if _, ok := auo.mutation.DocumentID(); auo.mutation.DocumentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Attachment.document"`)
	}
	return nil
}

func (auo *AttachmentUpdateOne) sqlSave(ctx context.Context) (_node *Attachment, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   attachment.Table,
			Columns: attachment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: attachment.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Attachment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, attachment.FieldID)
		for _, f := range fields {
			if !attachment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != attachment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(attachment.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.GetType(); ok {
		_spec.SetField(attachment.FieldType, field.TypeEnum, value)
	}
	if auo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.ItemTable,
			Columns: []string{attachment.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.ItemTable,
			Columns: []string{attachment.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: item.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if auo.mutation.DocumentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.DocumentTable,
			Columns: []string{attachment.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: document.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.DocumentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   attachment.DocumentTable,
			Columns: []string{attachment.DocumentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: document.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Attachment{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{attachment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
