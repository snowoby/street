package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Token holds the schema definition for the Token entity.
type Token struct {
	ent.Schema
}

// Fields of the Token.
func (Token) Fields() []ent.Field {
	return []ent.Field{
		field.String("body").Unique().MaxLen(320).NotEmpty(),
		field.String("type").MaxLen(16).NotEmpty(),
		field.Time("expire_time"),
	}
}

func (Token) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		mixin.Time{},
	}
}

// Edges of the Token.
func (Token) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("token").Required().Unique(),
	}
}

func (Token) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("body").
			Unique(),
	}
}
