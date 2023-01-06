package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("user_uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("user_email").
			Default("null"),
		field.String("user_password").
			Default("null"),
		field.String("user_nickname").
			Default("null"),
		field.Time("created_at").
			Default("null"),
		field.Time("updated_at").
			Default("null"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
