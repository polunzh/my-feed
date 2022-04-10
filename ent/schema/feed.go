package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Feed holds the schema definition for the Feed entity.
type Feed struct {
	ent.Schema
}

// Fields of the Feed.
func (Feed) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default(""),
		field.String("url").NotEmpty().Unique(),
		field.Time("updated_at").Default(time.Now()),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Feed.
func (Feed) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("group", Group.Type),
	}
}
