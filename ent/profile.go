// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"street/ent/account"
	"street/ent/profile"
	"street/ent/schema"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Profile is the model entity for the Profile schema.
type Profile struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Sid holds the value of the "sid" field.
	Sid schema.ID `json:"sid,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Call holds the value of the "call" field.
	Call string `json:"call,omitempty"`
	// Category holds the value of the "category" field.
	Category string `json:"category,omitempty"`
	// Avatar holds the value of the "avatar" field.
	Avatar *string `json:"avatar,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProfileQuery when eager-loading is set.
	Edges           ProfileEdges `json:"edges"`
	account_profile *uuid.UUID
}

// ProfileEdges holds the relations/edges for other nodes in the graph.
type ProfileEdges struct {
	// Account holds the value of the account edge.
	Account *Account `json:"account,omitempty"`
	// Episode holds the value of the episode edge.
	Episode []*Episode `json:"episode,omitempty"`
	// Commenter holds the value of the commenter edge.
	Commenter []*Comment `json:"commenter,omitempty"`
	// Series holds the value of the series edge.
	Series []*Series `json:"series,omitempty"`
	// JoinedSeries holds the value of the joined_series edge.
	JoinedSeries []*Series `json:"joined_series,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// AccountOrErr returns the Account value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ProfileEdges) AccountOrErr() (*Account, error) {
	if e.loadedTypes[0] {
		if e.Account == nil {
			// The edge account was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: account.Label}
		}
		return e.Account, nil
	}
	return nil, &NotLoadedError{edge: "account"}
}

// EpisodeOrErr returns the Episode value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) EpisodeOrErr() ([]*Episode, error) {
	if e.loadedTypes[1] {
		return e.Episode, nil
	}
	return nil, &NotLoadedError{edge: "episode"}
}

// CommenterOrErr returns the Commenter value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) CommenterOrErr() ([]*Comment, error) {
	if e.loadedTypes[2] {
		return e.Commenter, nil
	}
	return nil, &NotLoadedError{edge: "commenter"}
}

// SeriesOrErr returns the Series value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) SeriesOrErr() ([]*Series, error) {
	if e.loadedTypes[3] {
		return e.Series, nil
	}
	return nil, &NotLoadedError{edge: "series"}
}

// JoinedSeriesOrErr returns the JoinedSeries value or an error if the edge
// was not loaded in eager-loading.
func (e ProfileEdges) JoinedSeriesOrErr() ([]*Series, error) {
	if e.loadedTypes[4] {
		return e.JoinedSeries, nil
	}
	return nil, &NotLoadedError{edge: "joined_series"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Profile) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case profile.FieldSid:
			values[i] = new(schema.ID)
		case profile.FieldTitle, profile.FieldCall, profile.FieldCategory, profile.FieldAvatar:
			values[i] = new(sql.NullString)
		case profile.FieldCreateTime, profile.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case profile.FieldID:
			values[i] = new(uuid.UUID)
		case profile.ForeignKeys[0]: // account_profile
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Profile", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Profile fields.
func (pr *Profile) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case profile.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				pr.ID = *value
			}
		case profile.FieldSid:
			if value, ok := values[i].(*schema.ID); !ok {
				return fmt.Errorf("unexpected type %T for field sid", values[i])
			} else if value != nil {
				pr.Sid = *value
			}
		case profile.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				pr.CreateTime = value.Time
			}
		case profile.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				pr.UpdateTime = value.Time
			}
		case profile.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pr.Title = value.String
			}
		case profile.FieldCall:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field call", values[i])
			} else if value.Valid {
				pr.Call = value.String
			}
		case profile.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				pr.Category = value.String
			}
		case profile.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				pr.Avatar = new(string)
				*pr.Avatar = value.String
			}
		case profile.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field account_profile", values[i])
			} else if value.Valid {
				pr.account_profile = new(uuid.UUID)
				*pr.account_profile = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryAccount queries the "account" edge of the Profile entity.
func (pr *Profile) QueryAccount() *AccountQuery {
	return (&ProfileClient{config: pr.config}).QueryAccount(pr)
}

// QueryEpisode queries the "episode" edge of the Profile entity.
func (pr *Profile) QueryEpisode() *EpisodeQuery {
	return (&ProfileClient{config: pr.config}).QueryEpisode(pr)
}

// QueryCommenter queries the "commenter" edge of the Profile entity.
func (pr *Profile) QueryCommenter() *CommentQuery {
	return (&ProfileClient{config: pr.config}).QueryCommenter(pr)
}

// QuerySeries queries the "series" edge of the Profile entity.
func (pr *Profile) QuerySeries() *SeriesQuery {
	return (&ProfileClient{config: pr.config}).QuerySeries(pr)
}

// QueryJoinedSeries queries the "joined_series" edge of the Profile entity.
func (pr *Profile) QueryJoinedSeries() *SeriesQuery {
	return (&ProfileClient{config: pr.config}).QueryJoinedSeries(pr)
}

// Update returns a builder for updating this Profile.
// Note that you need to call Profile.Unwrap() before calling this method if this Profile
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Profile) Update() *ProfileUpdateOne {
	return (&ProfileClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the Profile entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Profile) Unwrap() *Profile {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Profile is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Profile) String() string {
	var builder strings.Builder
	builder.WriteString("Profile(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", sid=")
	builder.WriteString(fmt.Sprintf("%v", pr.Sid))
	builder.WriteString(", create_time=")
	builder.WriteString(pr.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(pr.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", title=")
	builder.WriteString(pr.Title)
	builder.WriteString(", call=")
	builder.WriteString(pr.Call)
	builder.WriteString(", category=")
	builder.WriteString(pr.Category)
	if v := pr.Avatar; v != nil {
		builder.WriteString(", avatar=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Profiles is a parsable slice of Profile.
type Profiles []*Profile

func (pr Profiles) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
