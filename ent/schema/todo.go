package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// ToDo holds the schema definition for the ToDo entity.
type ToDo struct {
	ent.Schema
}

// Fields of the ToDo.
func (ToDo) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("todo_uuid", uuid.UUID{}).
			Default(uuid.New),
		field.UUID("user_uuid", uuid.UUID{}).
			Default(uuid.New),
		field.String("todo_title"),
		field.String("todo_context"),
		field.Time("updated_at").
			Default(time.Now),
		field.Time("crated_at").
			Default(time.Now),
	}
}

// Edges of the ToDo.
func (ToDo) Edges() []ent.Edge {
	return nil
}
