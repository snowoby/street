// Code generated by entc, DO NOT EDIT.

package episode

import (
	"street/ent/schema"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the episode type in the database.
	Label = "episode"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSid holds the string denoting the sid field in the database.
	FieldSid = "sid"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldExtra holds the string denoting the extra field in the database.
	FieldExtra = "extra"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// EdgeSeries holds the string denoting the series edge name in mutations.
	EdgeSeries = "series"
	// Table holds the table name of the episode in the database.
	Table = "episodes"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "episodes"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "profile_episode"
	// SeriesTable is the table that holds the series relation/edge.
	SeriesTable = "episodes"
	// SeriesInverseTable is the table name for the Series entity.
	// It exists in this package in order to avoid circular dependency with the "series" package.
	SeriesInverseTable = "series"
	// SeriesColumn is the table column denoting the series relation/edge.
	SeriesColumn = "series_episode"
)

// Columns holds all SQL columns for episode fields.
var Columns = []string{
	FieldID,
	FieldSid,
	FieldCreateTime,
	FieldUpdateTime,
	FieldTitle,
	FieldContent,
	FieldExtra,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "episodes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"profile_episode",
	"series_episode",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultSid holds the default value on creation for the "sid" field.
	DefaultSid func() schema.ID
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
	// DefaultExtra holds the default value on creation for the "extra" field.
	DefaultExtra func() schema.EpisodeExtra
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
