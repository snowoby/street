package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// Series holds the schema definition for the Series entity.
type Series struct {
	ent.Schema
}

func (Series) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		mixin.Time{},
	}
}

// Fields of the Series.
func (Series) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").MaxLen(320).NotEmpty(),
		field.String("type").MaxLen(32).Nillable().Optional(),
	}
}

// Edges of the Series.
func (Series) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("episodes", Episode.Type),
		edge.From("owner", Profile.Type).Ref("series").Unique().Required(),
		edge.From("participant", Profile.Type).Ref("joined_series"),
	}
}
