package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// InvoiceItem holds the schema definition for the InvoiceItem entity.
type InvoiceItem struct {
    ent.Schema
}

func (InvoiceItem) Fields() []ent.Field {
    return []ent.Field{
        field.String("name"),
        field.String("description"),
        field.Int("quantity"),
        field.Float("rate"),
        field.Float("total"),
    }
}

func (InvoiceItem) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("invoice", Invoice.Type).
            Ref("items").
            Unique(),
    }
}
