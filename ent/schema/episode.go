package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Episode holds the schema definition for the Episode entity.
type Episode struct {
	ent.Schema
}

// Fields of the Episode.
func (Episode) Fields() []ent.Field {
	return []ent.Field{
		field.String("cover").MaxLen(320).Nillable().Optional(),
		field.String("title").MaxLen(320).NotEmpty(),
		field.Text("content").NotEmpty(),
	}
}

func (Episode) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		mixin.Time{},
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Profile.Type).Ref("episode").Required().Unique(),
		edge.To("comments", Comment.Type),
	}
}
