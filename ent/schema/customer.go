package schema

import (
	"time"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Customer struct {
	ent.Schema
}

func (Customer) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			Positive(),
		field.String("name"),
		field.String("email"),
		field.String("phone"),
		field.String("address"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Customer.
func (Customer) Edges() []ent.Edge {
	return nil
}
