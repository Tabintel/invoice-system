package schema

import "entgo.io/ent"

// Invoice holds the schema definition for the Invoice entity.
type Invoice struct {
	ent.Schema
}

// Fields of the Invoice.
func (Invoice) Fields() []ent.Field {
	return nil
}

// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return nil
}
