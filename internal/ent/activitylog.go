package schema

import "entgo.io/ent"

// ActivityLog holds the schema definition for the ActivityLog entity.
type ActivityLog struct {
	ent.Schema
}

// Fields of the ActivityLog.
func (ActivityLog) Fields() []ent.Field {
	return nil
}

// Edges of the ActivityLog.
func (ActivityLog) Edges() []ent.Edge {
	return nil
}
