// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"street/ent/file"
	"street/ent/profile"
	"street/ent/schema"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// File is the model entity for the File schema.
type File struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// Sid holds the value of the "sid" field.
	Sid schema.ID `json:"sid,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Filename holds the value of the "filename" field.
	Filename string `json:"filename,omitempty"`
	// Path holds the value of the "path" field.
	Path string `json:"path,omitempty"`
	// Mime holds the value of the "mime" field.
	Mime string `json:"mime,omitempty"`
	// Size holds the value of the "size" field.
	Size int `json:"size,omitempty"`
	// Status holds the value of the "status" field.
	Status string `json:"status,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FileQuery when eager-loading is set.
	Edges        FileEdges `json:"edges"`
	profile_file *uuid.UUID
}

// FileEdges holds the relations/edges for other nodes in the graph.
type FileEdges struct {
	// Profile holds the value of the profile edge.
	Profile *Profile `json:"profile,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ProfileOrErr returns the Profile value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FileEdges) ProfileOrErr() (*Profile, error) {
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
func (*File) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case file.FieldSid:
			values[i] = new(schema.ID)
		case file.FieldSize:
			values[i] = new(sql.NullInt64)
		case file.FieldFilename, file.FieldPath, file.FieldMime, file.FieldStatus, file.FieldNote:
			values[i] = new(sql.NullString)
		case file.FieldCreateTime, file.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		case file.FieldID:
			values[i] = new(uuid.UUID)
		case file.ForeignKeys[0]: // profile_file
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type File", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the File fields.
func (f *File) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case file.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				f.ID = *value
			}
		case file.FieldSid:
			if value, ok := values[i].(*schema.ID); !ok {
				return fmt.Errorf("unexpected type %T for field sid", values[i])
			} else if value != nil {
				f.Sid = *value
			}
		case file.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				f.CreateTime = value.Time
			}
		case file.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				f.UpdateTime = value.Time
			}
		case file.FieldFilename:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field filename", values[i])
			} else if value.Valid {
				f.Filename = value.String
			}
		case file.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				f.Path = value.String
			}
		case file.FieldMime:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mime", values[i])
			} else if value.Valid {
				f.Mime = value.String
			}
		case file.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				f.Size = int(value.Int64)
			}
		case file.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				f.Status = value.String
			}
		case file.FieldNote:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field note", values[i])
			} else if value.Valid {
				f.Note = value.String
			}
		case file.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field profile_file", values[i])
			} else if value.Valid {
				f.profile_file = new(uuid.UUID)
				*f.profile_file = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryProfile queries the "profile" edge of the File entity.
func (f *File) QueryProfile() *ProfileQuery {
	return (&FileClient{config: f.config}).QueryProfile(f)
}

// Update returns a builder for updating this File.
// Note that you need to call File.Unwrap() before calling this method if this File
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *File) Update() *FileUpdateOne {
	return (&FileClient{config: f.config}).UpdateOne(f)
}

// Unwrap unwraps the File entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *File) Unwrap() *File {
	tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: File is not a transactional entity")
	}
	f.config.driver = tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *File) String() string {
	var builder strings.Builder
	builder.WriteString("File(")
	builder.WriteString(fmt.Sprintf("id=%v", f.ID))
	builder.WriteString(", sid=")
	builder.WriteString(fmt.Sprintf("%v", f.Sid))
	builder.WriteString(", create_time=")
	builder.WriteString(f.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(f.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", filename=")
	builder.WriteString(f.Filename)
	builder.WriteString(", path=")
	builder.WriteString(f.Path)
	builder.WriteString(", mime=")
	builder.WriteString(f.Mime)
	builder.WriteString(", size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", status=")
	builder.WriteString(f.Status)
	builder.WriteString(", note=")
	builder.WriteString(f.Note)
	builder.WriteByte(')')
	return builder.String()
}

// Files is a parsable slice of File.
type Files []*File

func (f Files) config(cfg config) {
	for _i := range f {
		f[_i].config = cfg
	}
}
