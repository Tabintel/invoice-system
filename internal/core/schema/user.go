package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
)

type User struct {
    ent.Schema
}

func (User) Fields() []ent.Field {
    return []ent.Field{
        field.String("name"),
        field.String("email").Unique(),
        field.String("phone"),
        field.String("company_name"),
        field.String("company_logo"),
        field.String("location"),
        field.Enum("role").Values("admin", "user"),
        field.Time("created_at").Default(time.Now),
        field.Time("updated_at").Default(time.Now),
    }
}


