package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique().MaxLen(320).NotEmpty(),
		field.String("password").Sensitive(),
	}
}

func (Account) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		mixin.Time{},
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("token", Token.Type),
		edge.To("profile", Profile.Type),
		edge.To("file", File.Type),
	}
}

// Indexes of the Account.
func (Account) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("email").
			Unique(),
	}
}
