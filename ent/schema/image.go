package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Image holds the schema definition for the Image entity.
type Image struct {
	ent.Schema
}

// Fields of the Image.
func (Image) Fields() []ent.Field {
	return []ent.Field{
		// id is the ID of this image
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),

		// storage_path is the location that this image is stored at.
		field.String("storage_path").NotEmpty(),

		// type is the type of image this is
		field.Enum("type").Values("flier", "poster"),
	}
}

// Edges of the Image.
func (Image) Edges() []ent.Edge {
	return []ent.Edge{}
}
