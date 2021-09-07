package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Role holds the schema definition for the Role entity.
type Role struct {
	ent.Schema
}

// Fields of the Role.
func (Role) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),

		// name is the name of this role, should be user friendly
		field.String("name").NotEmpty(),
	}
}

// Edges of the Role.
func (Role) Edges() []ent.Edge {
	return nil
}
