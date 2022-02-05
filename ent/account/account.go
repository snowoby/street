// Code generated by entc, DO NOT EDIT.

package account

import (
	"street/ent/schema"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the account type in the database.
	Label = "account"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSid holds the string denoting the sid field in the database.
	FieldSid = "sid"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// EdgeToken holds the string denoting the token edge name in mutations.
	EdgeToken = "token"
	// EdgeProfile holds the string denoting the profile edge name in mutations.
	EdgeProfile = "profile"
	// EdgeFile holds the string denoting the file edge name in mutations.
	EdgeFile = "file"
	// Table holds the table name of the account in the database.
	Table = "accounts"
	// TokenTable is the table that holds the token relation/edge.
	TokenTable = "tokens"
	// TokenInverseTable is the table name for the Token entity.
	// It exists in this package in order to avoid circular dependency with the "token" package.
	TokenInverseTable = "tokens"
	// TokenColumn is the table column denoting the token relation/edge.
	TokenColumn = "account_token"
	// ProfileTable is the table that holds the profile relation/edge.
	ProfileTable = "profiles"
	// ProfileInverseTable is the table name for the Profile entity.
	// It exists in this package in order to avoid circular dependency with the "profile" package.
	ProfileInverseTable = "profiles"
	// ProfileColumn is the table column denoting the profile relation/edge.
	ProfileColumn = "account_profile"
	// FileTable is the table that holds the file relation/edge.
	FileTable = "files"
	// FileInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FileInverseTable = "files"
	// FileColumn is the table column denoting the file relation/edge.
	FileColumn = "account_file"
)

// Columns holds all SQL columns for account fields.
var Columns = []string{
	FieldID,
	FieldSid,
	FieldCreateTime,
	FieldUpdateTime,
	FieldEmail,
	FieldPassword,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// EmailValidator is a validator for the "email" field. It is called by the builders before save.
	EmailValidator func(string) error
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)
