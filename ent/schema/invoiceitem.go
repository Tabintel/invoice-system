package schema

import "entgo.io/ent"

// InvoiceItem holds the schema definition for the InvoiceItem entity.
type InvoiceItem struct {
	ent.Schema
}

// Fields of the InvoiceItem.
func (InvoiceItem) Fields() []ent.Field {
	return nil
}

// Edges of the InvoiceItem.
func (InvoiceItem) Edges() []ent.Edge {
	return nil
}
