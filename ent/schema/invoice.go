package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/edge"
	"time"
)

// Invoice holds the schema definition for the Invoice entity.
type Invoice struct {
	ent.Schema
}

// Fields of the Invoice.
func (Invoice) Fields() []ent.Field {
	return []ent.Field{
		field.String("reference_number").Unique(),
		field.Float("total_amount"),
		field.String("status").Default("draft"),
		field.Time("issue_date"),
		field.Time("due_date"),
		field.String("currency").Default("USD"),
		field.Time("created_at").Default(time.Now),
		field.String("share_token").
			Optional(),
		field.Time("share_expiry").
			Optional(),
	}
}
// Edges of the Invoice.
func (Invoice) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("creator", User.Type).Ref("invoices").Unique(),
		edge.To("items", InvoiceItem.Type),
		edge.To("payments", Payment.Type),
		edge.To("customer", Customer.Type).Unique(),
	}
}
