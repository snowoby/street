// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AccountsColumns holds the columns for the "accounts" table.
	AccountsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "sid", Type: field.TypeInt64, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "email", Type: field.TypeString, Unique: true, Size: 320},
		{Name: "password", Type: field.TypeString},
	}
	// AccountsTable holds the schema information for the "accounts" table.
	AccountsTable = &schema.Table{
		Name:       "accounts",
		Columns:    AccountsColumns,
		PrimaryKey: []*schema.Column{AccountsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "account_email",
				Unique:  true,
				Columns: []*schema.Column{AccountsColumns[4]},
			},
		},
	}
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "sid", Type: field.TypeInt64, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "episode_comments", Type: field.TypeUUID, Nullable: true},
		{Name: "profile_commenter", Type: field.TypeUUID, Nullable: true},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "comments_episodes_comments",
				Columns:    []*schema.Column{CommentsColumns[5]},
				RefColumns: []*schema.Column{EpisodesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "comments_profiles_commenter",
				Columns:    []*schema.Column{CommentsColumns[6]},
				RefColumns: []*schema.Column{ProfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EpisodesColumns holds the columns for the "episodes" table.
	EpisodesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "sid", Type: field.TypeInt64, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "cover", Type: field.TypeString, Nullable: true, Size: 320},
		{Name: "title", Type: field.TypeString, Size: 320},
		{Name: "content", Type: field.TypeString, Size: 2147483647},
		{Name: "profile_episode", Type: field.TypeUUID, Nullable: true},
		{Name: "series_episodes", Type: field.TypeUUID, Nullable: true},
	}
	// EpisodesTable holds the schema information for the "episodes" table.
	EpisodesTable = &schema.Table{
		Name:       "episodes",
		Columns:    EpisodesColumns,
		PrimaryKey: []*schema.Column{EpisodesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "episodes_profiles_episode",
				Columns:    []*schema.Column{EpisodesColumns[7]},
				RefColumns: []*schema.Column{ProfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "episodes_series_episodes",
				Columns:    []*schema.Column{EpisodesColumns[8]},
				RefColumns: []*schema.Column{SeriesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FilesColumns holds the columns for the "files" table.
	FilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "sid", Type: field.TypeInt64, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "filename", Type: field.TypeString, Size: 320, Default: "file"},
		{Name: "path", Type: field.TypeString, Size: 64, Default: "media"},
		{Name: "mime", Type: field.TypeString, Size: 320, Default: "application/octet-stream"},
		{Name: "size", Type: field.TypeInt},
		{Name: "status", Type: field.TypeString, Size: 16, Default: "created"},
		{Name: "note", Type: field.TypeString, Nullable: true, Size: 128},
		{Name: "account_file", Type: field.TypeUUID, Nullable: true},
	}
	// FilesTable holds the schema information for the "files" table.
	FilesTable = &schema.Table{
		Name:       "files",
		Columns:    FilesColumns,
		PrimaryKey: []*schema.Column{FilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "files_accounts_file",
				Columns:    []*schema.Column{FilesColumns[10]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProfilesColumns holds the columns for the "profiles" table.
	ProfilesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "sid", Type: field.TypeInt64, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 64},
		{Name: "call", Type: field.TypeString, Unique: true, Size: 64},
		{Name: "category", Type: field.TypeString, Size: 16},
		{Name: "avatar", Type: field.TypeString, Nullable: true, Size: 64},
		{Name: "account_profile", Type: field.TypeUUID, Nullable: true},
	}
	// ProfilesTable holds the schema information for the "profiles" table.
	ProfilesTable = &schema.Table{
		Name:       "profiles",
		Columns:    ProfilesColumns,
		PrimaryKey: []*schema.Column{ProfilesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "profiles_accounts_profile",
				Columns:    []*schema.Column{ProfilesColumns[8]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "profile_call",
				Unique:  true,
				Columns: []*schema.Column{ProfilesColumns[5]},
			},
			{
				Name:    "profile_category",
				Unique:  false,
				Columns: []*schema.Column{ProfilesColumns[6]},
			},
		},
	}
	// SeriesColumns holds the columns for the "series" table.
	SeriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "sid", Type: field.TypeInt64, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "title", Type: field.TypeString, Size: 320},
		{Name: "type", Type: field.TypeString, Nullable: true, Size: 32},
		{Name: "profile_series", Type: field.TypeUUID, Nullable: true},
	}
	// SeriesTable holds the schema information for the "series" table.
	SeriesTable = &schema.Table{
		Name:       "series",
		Columns:    SeriesColumns,
		PrimaryKey: []*schema.Column{SeriesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "series_profiles_series",
				Columns:    []*schema.Column{SeriesColumns[6]},
				RefColumns: []*schema.Column{ProfilesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// TokensColumns holds the columns for the "tokens" table.
	TokensColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "sid", Type: field.TypeInt64, Unique: true},
		{Name: "create_time", Type: field.TypeTime},
		{Name: "update_time", Type: field.TypeTime},
		{Name: "body", Type: field.TypeString, Unique: true, Size: 320},
		{Name: "type", Type: field.TypeString, Size: 16},
		{Name: "expire", Type: field.TypeTime},
		{Name: "account_token", Type: field.TypeUUID, Nullable: true},
	}
	// TokensTable holds the schema information for the "tokens" table.
	TokensTable = &schema.Table{
		Name:       "tokens",
		Columns:    TokensColumns,
		PrimaryKey: []*schema.Column{TokensColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tokens_accounts_token",
				Columns:    []*schema.Column{TokensColumns[7]},
				RefColumns: []*schema.Column{AccountsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "token_body",
				Unique:  true,
				Columns: []*schema.Column{TokensColumns[4]},
			},
		},
	}
	// ProfileJoinedSeriesColumns holds the columns for the "profile_joined_series" table.
	ProfileJoinedSeriesColumns = []*schema.Column{
		{Name: "profile_id", Type: field.TypeUUID},
		{Name: "series_id", Type: field.TypeUUID},
	}
	// ProfileJoinedSeriesTable holds the schema information for the "profile_joined_series" table.
	ProfileJoinedSeriesTable = &schema.Table{
		Name:       "profile_joined_series",
		Columns:    ProfileJoinedSeriesColumns,
		PrimaryKey: []*schema.Column{ProfileJoinedSeriesColumns[0], ProfileJoinedSeriesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "profile_joined_series_profile_id",
				Columns:    []*schema.Column{ProfileJoinedSeriesColumns[0]},
				RefColumns: []*schema.Column{ProfilesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "profile_joined_series_series_id",
				Columns:    []*schema.Column{ProfileJoinedSeriesColumns[1]},
				RefColumns: []*schema.Column{SeriesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AccountsTable,
		CommentsTable,
		EpisodesTable,
		FilesTable,
		ProfilesTable,
		SeriesTable,
		TokensTable,
		ProfileJoinedSeriesTable,
	}
)

func init() {
	CommentsTable.ForeignKeys[0].RefTable = EpisodesTable
	CommentsTable.ForeignKeys[1].RefTable = ProfilesTable
	EpisodesTable.ForeignKeys[0].RefTable = ProfilesTable
	EpisodesTable.ForeignKeys[1].RefTable = SeriesTable
	FilesTable.ForeignKeys[0].RefTable = AccountsTable
	ProfilesTable.ForeignKeys[0].RefTable = AccountsTable
	SeriesTable.ForeignKeys[0].RefTable = ProfilesTable
	TokensTable.ForeignKeys[0].RefTable = AccountsTable
	ProfileJoinedSeriesTable.ForeignKeys[0].RefTable = ProfilesTable
	ProfileJoinedSeriesTable.ForeignKeys[1].RefTable = SeriesTable
}
