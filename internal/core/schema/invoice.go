package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

type Invoice struct {
    ent.Schema
}

func (Invoice) Fields() []ent.Field {
    return []ent.Field{
        field.String("reference_number").Unique(),
        field.Float("amount"),
        field.String("currency").Default("USD"),
        field.Time("issue_date"),
        field.Time("due_date"),
        field.Enum("status").Values("paid", "unpaid", "overdue", "draft", "pending"),
        field.String("notes").Optional(),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now),
    }
}

func (Invoice) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("customer", Customer.Type).Ref("invoices").Unique(),
        edge.From("sender", User.Type).Ref("sent_invoices").Unique(),
        edge.To("payments", Payment.Type),
        edge.To("activities", Activity.Type),
    }
}
