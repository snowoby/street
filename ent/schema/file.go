package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("filename").MaxLen(320).NotEmpty().Default("file"),
		field.String("path").MaxLen(64).NotEmpty().Default("media"),
		field.String("mime").MaxLen(320).NotEmpty().Default("application/octet-stream"),
		field.Int("size"),
		field.String("status").MaxLen(16).Default("created"),
		field.String("note").MaxLen(128).Optional(),
	}
}

func (File) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
		mixin.Time{},
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", Account.Type).Ref("file").Required().Unique(),
	}
}
