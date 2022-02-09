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
		field.String("title").MaxLen(64).NotEmpty(),
		field.Text("content").NotEmpty(),
		field.Bytes("extra").GoType(EpisodeExtra{}).DefaultFunc(NewEpisodeExtra),
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
		edge.From("series", Series.Type).Ref("episode").Unique(),
	}
}
