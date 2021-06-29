package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

type File struct {
	ent.Schema
}

func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id"),
		field.String("path"),
		field.String("status"),

		field.Time("created_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
		field.Time("updated_at").
			Default(time.Now).SchemaType(map[string]string{
			dialect.MySQL: "datetime",
		}),
	}
}

func (File) Edges() []ent.Edge {
	return []ent.Edge{}
}
