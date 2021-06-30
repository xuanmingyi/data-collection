package schema

import "entgo.io/ent"

type Node struct {
	ent.Schema
}

func (Node) Fields() []ent.Field {
	return nil
}

func (Node) Edges() []ent.Edge {
	return nil
}
