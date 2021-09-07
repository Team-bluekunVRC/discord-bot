package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// AudioClip holds the schema definition for the AudioClip entity.
type AudioClip struct {
	ent.Schema
}

// Fields of the AudioClip.
func (AudioClip) Fields() []ent.Field {
	return []ent.Field{
		// id is the ID of this image
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),

		// storage_path is the location that this image is stored at.
		field.String("storage_path").NotEmpty(),

		// type is the type of audio clip this is
		field.Enum("type").Values("intro"),
	}
}

// Edges of the AudioClip.
func (AudioClip) Edges() []ent.Edge {
	return nil
}
