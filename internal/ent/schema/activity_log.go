package schema

import (
    "time"
    "entgo.io/ent"
    "entgo.io/ent/schema/field"
    "entgo.io/ent/schema/edge"
)

type ActivityLog struct {
    ent.Schema
}

func (ActivityLog) Fields() []ent.Field {
    return []ent.Field{
        field.String("action_type"),
        field.String("description"),
        field.Time("timestamp").Default(time.Now),
        field.String("performed_by"),
    }
}

func (ActivityLog) Edges() []ent.Edge {
    return []ent.Edge{
        edge.From("invoice", Invoice.Type).Ref("activities").Unique(),
    }
}
