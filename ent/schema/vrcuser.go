package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// VRCUser holds the schema definition for the VRCUser entity.
type VRCUser struct {
	ent.Schema
}

// Fields of the VRCUser.
func (VRCUser) Fields() []ent.Field {
	return []ent.Field{
		// id is a ID of a VRChat user
		field.String("id").NotEmpty().Unique(),
	}
}

// Edges of the VRCUser.
func (VRCUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Required().Unique(),
		edge.From("discord", DiscordUser.Type).Ref("vrc").Unique(),
	}
}
