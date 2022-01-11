package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("title").MaxLen(64).NotEmpty(),
		field.String("callSign").MaxLen(32).Unique().Nillable().Optional(),
		field.Text("content").Default(""),
	}
}

// Edges of the Series.
func (Series) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("profile", Profile.Type).Ref("series").Required().Unique(),
		edge.To("episode", Episode.Type),
	}
}

// Indexes of the Series.
func (Series) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("callSign").
			Unique(),
	}
}
