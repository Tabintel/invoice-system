package schema

import "entgo.io/ent"

// PaymentDetails holds the schema definition for the PaymentDetails entity.
type PaymentDetails struct {
	ent.Schema
}

// Fields of the PaymentDetails.
func (PaymentDetails) Fields() []ent.Field {
	return nil
}

// Edges of the PaymentDetails.
func (PaymentDetails) Edges() []ent.Edge {
	return nil
}
