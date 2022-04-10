package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default(""),
		field.String("url").NotEmpty().Unique(),
		field.Time("updated_at").Default(time.Now()),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("group", Group.Type),
	}
}
