// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"street/ent/profile"
	"street/ent/series"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Series is the model entity for the Series schema.
type Series struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// CallSign holds the value of the "callSign" field.
	CallSign *string `json:"callSign,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SeriesQuery when eager-loading is set.
	Edges          SeriesEdges `json:"edges"`
	profile_series *uuid.UUID
}

// SeriesEdges holds the relations/edges for other nodes in the graph.
type SeriesEdges struct {
	// Profile holds the value of the profile edge.
	Profile *Profile `json:"profile,omitempty"`
	// Episode holds the value of the episode edge.
	Episode []*Episode `json:"episode,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SeriesEdges) ProfileOrErr() (*Profile, error) {
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

// EpisodeOrErr returns the Episode value or an error if the edge
// was not loaded in eager-loading.
func (e SeriesEdges) EpisodeOrErr() ([]*Episode, error) {
	if e.loadedTypes[1] {
		return e.Episode, nil
	}
	return nil, &NotLoadedError{edge: "episode"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Series) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case series.FieldTitle, series.FieldCallSign, series.FieldContent:
			values[i] = new(sql.NullString)
		case series.FieldCreateTime, series.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case series.FieldID:
			values[i] = new(uuid.UUID)
		case series.ForeignKeys[0]: // profile_series
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Series", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Series fields.
func (s *Series) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case series.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				s.ID = *value
			}
		case series.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				s.CreateTime = value.Time
			}
		case series.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				s.UpdateTime = value.Time
			}
		case series.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case series.FieldCallSign:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field callSign", values[i])
			} else if value.Valid {
				s.CallSign = new(string)
				*s.CallSign = value.String
			}
		case series.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				s.Content = value.String
			}
		case series.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field profile_series", values[i])
			} else if value.Valid {
				s.profile_series = new(uuid.UUID)
				*s.profile_series = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryProfile queries the "profile" edge of the Series entity.
func (s *Series) QueryProfile() *ProfileQuery {
	return (&SeriesClient{config: s.config}).QueryProfile(s)
}

// QueryEpisode queries the "episode" edge of the Series entity.
func (s *Series) QueryEpisode() *EpisodeQuery {
	return (&SeriesClient{config: s.config}).QueryEpisode(s)
}

// Update returns a builder for updating this Series.
// Note that you need to call Series.Unwrap() before calling this method if this Series
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Series) Update() *SeriesUpdateOne {
	return (&SeriesClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Series entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Series) Unwrap() *Series {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Series is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Series) String() string {
	var builder strings.Builder
	builder.WriteString("Series(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(s.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(s.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", title=")
	builder.WriteString(s.Title)
	if v := s.CallSign; v != nil {
		builder.WriteString(", callSign=")
		builder.WriteString(*v)
	}
	builder.WriteString(", content=")
	builder.WriteString(s.Content)
	builder.WriteByte(')')
	return builder.String()
}

// SeriesSlice is a parsable slice of Series.
type SeriesSlice []*Series

func (s SeriesSlice) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
