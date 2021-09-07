package schema

import (
	"net/mail"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User is the user schema definition. Stored here is all information
// related to a given user.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// id is the ID of this user, we use UUIDs so it needs to
		// be manually set here.
		field.UUID("id", uuid.UUID{}).Unique().Default(uuid.New),

		// email is the email of this User.
		field.String("email").Validate(func(s string) error {
			_, err := mail.ParseAddress(s)
			return err
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("images", Image.Type),
		edge.To("audio", AudioClip.Type),
		edge.To("role", Role.Type),
	}
}
