package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(64).NotEmpty(),
		field.String("call").Unique().MaxLen(64).NotEmpty(),
		field.String("category").MaxLen(16).NotEmpty(),
		field.String("avatar").MaxLen(64).Optional().Nillable(),
	}
}

func (Profile) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		mixin.Time{},
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("profile").Required().Unique(),
		edge.To("episode", Episode.Type),
		edge.To("series", Series.Type),
	}
}

// Indexes of the Profile.
func (Profile) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("call").
			Unique(),
		index.Fields("category"),
	}
}
