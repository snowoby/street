// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"street/ent/episode"
	"street/ent/profile"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Episode is the model entity for the Episode schema.
type Episode struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EpisodeQuery when eager-loading is set.
	Edges           EpisodeEdges `json:"edges"`
	profile_episode *uuid.UUID
}

// EpisodeEdges holds the relations/edges for other nodes in the graph.
type EpisodeEdges struct {
	// Profile holds the value of the profile edge.
	Profile *Profile `json:"profile,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EpisodeEdges) ProfileOrErr() (*Profile, error) {
	if e.loadedTypes[0] {
		if e.Profile == nil {
			// The edge profile was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Profile, nil
	}
	return nil, &NotLoadedError{edge: "profile"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Episode) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case episode.FieldTitle, episode.FieldContent:
			values[i] = new(sql.NullString)
		case episode.FieldCreateTime, episode.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case episode.FieldID:
			values[i] = new(uuid.UUID)
		case episode.ForeignKeys[0]: // profile_episode
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Episode", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Episode fields.
func (e *Episode) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case episode.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case episode.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				e.CreateTime = value.Time
			}
		case episode.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				e.UpdateTime = value.Time
			}
		case episode.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				e.Title = value.String
			}
		case episode.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				e.Content = value.String
			}
		case episode.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field profile_episode", values[i])
			} else if value.Valid {
				e.profile_episode = new(uuid.UUID)
				*e.profile_episode = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryProfile queries the "profile" edge of the Episode entity.
func (e *Episode) QueryProfile() *ProfileQuery {
	return (&EpisodeClient{config: e.config}).QueryProfile(e)
}

// Update returns a builder for updating this Episode.
// Note that you need to call Episode.Unwrap() before calling this method if this Episode
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Episode) Update() *EpisodeUpdateOne {
	return (&EpisodeClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Episode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Episode) Unwrap() *Episode {
	tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Episode is not a transactional entity")
	}
	e.config.driver = tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Episode) String() string {
	var builder strings.Builder
	builder.WriteString("Episode(")
	builder.WriteString(fmt.Sprintf("id=%v", e.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(e.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(e.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", title=")
	builder.WriteString(e.Title)
	builder.WriteString(", content=")
	builder.WriteString(e.Content)
	builder.WriteByte(')')
	return builder.String()
}

// Episodes is a parsable slice of Episode.
type Episodes []*Episode

func (e Episodes) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}