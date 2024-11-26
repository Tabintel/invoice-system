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
        field.String("name"),
        field.String("email"),
        field.String("phone"),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now),
    }
}
