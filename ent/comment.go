// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"street/ent/comment"
	"street/ent/episode"
	"street/ent/profile"
	"street/ent/schema"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Sid holds the value of the "sid" field.
	Sid schema.ID `json:"sid,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Content holds the value of the "content" field.
	Content string `json:"content,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges             CommentEdges `json:"edges"`
	episode_comments  *uuid.UUID
	profile_commenter *uuid.UUID
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Episode holds the value of the episode edge.
	Episode *Episode `json:"episode,omitempty"`
	// Author holds the value of the author edge.
	Author *Profile `json:"author,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// EpisodeOrErr returns the Episode value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) EpisodeOrErr() (*Episode, error) {
	if e.loadedTypes[0] {
		if e.Episode == nil {
			// The edge episode was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: episode.Label}
		}
		return e.Episode, nil
	}
	return nil, &NotLoadedError{edge: "episode"}
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) AuthorOrErr() (*Profile, error) {
	if e.loadedTypes[1] {
		if e.Author == nil {
			// The edge author was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: profile.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldSid:
			values[i] = new(schema.ID)
		case comment.FieldContent:
			values[i] = new(sql.NullString)
		case comment.FieldCreateTime, comment.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case comment.FieldID:
			values[i] = new(uuid.UUID)
		case comment.ForeignKeys[0]: // episode_comments
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case comment.ForeignKeys[1]: // profile_commenter
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Comment", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (c *Comment) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				c.ID = *value
			}
		case comment.FieldSid:
			if value, ok := values[i].(*schema.ID); !ok {
				return fmt.Errorf("unexpected type %T for field sid", values[i])
			} else if value != nil {
				c.Sid = *value
			}
		case comment.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				c.CreateTime = value.Time
			}
		case comment.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				c.UpdateTime = value.Time
			}
		case comment.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				c.Content = value.String
			}
		case comment.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field episode_comments", values[i])
			} else if value.Valid {
				c.episode_comments = new(uuid.UUID)
				*c.episode_comments = *value.S.(*uuid.UUID)
			}
		case comment.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field profile_commenter", values[i])
			} else if value.Valid {
				c.profile_commenter = new(uuid.UUID)
				*c.profile_commenter = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryEpisode queries the "episode" edge of the Comment entity.
func (c *Comment) QueryEpisode() *EpisodeQuery {
	return (&CommentClient{config: c.config}).QueryEpisode(c)
}

// QueryAuthor queries the "author" edge of the Comment entity.
func (c *Comment) QueryAuthor() *ProfileQuery {
	return (&CommentClient{config: c.config}).QueryAuthor(c)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return (&CommentClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", sid=")
	builder.WriteString(fmt.Sprintf("%v", c.Sid))
	builder.WriteString(", create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", content=")
	builder.WriteString(c.Content)
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment

func (c Comments) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}