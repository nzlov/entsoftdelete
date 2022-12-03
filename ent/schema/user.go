package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("deleted_at").Nillable().Optional(),
		field.String("name").StorageKey("nickname"),
		field.String("test").SchemaType(map[string]string{
			dialect.Postgres: "numeric",
		}),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
