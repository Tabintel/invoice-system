package schema

import (
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

// PaymentDetails holds the schema definition for the PaymentDetails entity.
type PaymentDetails struct {
    ent.Schema
}

func (PaymentDetails) Fields() []ent.Field {
    return []ent.Field{
        field.String("account_name"),
        field.String("account_number"),
        field.String("routing_number"),
        field.String("bank_name"),
        field.String("bank_address"),
    }
}

func (PaymentDetails) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("invoice", Invoice.Type).
            Ref("payment_details").
            Unique(),
    }
}
