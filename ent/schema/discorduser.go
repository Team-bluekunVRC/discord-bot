package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// DiscordUser holds the schema definition for the DiscordUser entity.
type DiscordUser struct {
	ent.Schema
}

// Fields of the DiscordUser.
func (DiscordUser) Fields() []ent.Field {
	return []ent.Field{
		// id is a ID of a discord user
		field.String("id").NotEmpty().Unique(),
	}
}

// Edges of the DiscordUser.
func (DiscordUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required(),
		edge.To("vrc", VRCUser.Type).Required(),
	}
}
