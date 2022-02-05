// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"street/ent/migrate"

	"street/ent/account"
	"street/ent/episode"
	"street/ent/file"
	"street/ent/profile"
	"street/ent/series"
	"street/ent/token"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Account is the client for interacting with the Account builders.
	Account *AccountClient
	// Episode is the client for interacting with the Episode builders.
	Episode *EpisodeClient
	// File is the client for interacting with the File builders.
	File *FileClient
	// Profile is the client for interacting with the Profile builders.
	Profile *ProfileClient
	// Series is the client for interacting with the Series builders.
	Series *SeriesClient
	// Token is the client for interacting with the Token builders.
	Token *TokenClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Account = NewAccountClient(c.config)
	c.Episode = NewEpisodeClient(c.config)
	c.File = NewFileClient(c.config)
	c.Profile = NewProfileClient(c.config)
	c.Series = NewSeriesClient(c.config)
	c.Token = NewTokenClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Account: NewAccountClient(cfg),
		Episode: NewEpisodeClient(cfg),
		File:    NewFileClient(cfg),
		Profile: NewProfileClient(cfg),
		Series:  NewSeriesClient(cfg),
		Token:   NewTokenClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		config:  cfg,
		Account: NewAccountClient(cfg),
		Episode: NewEpisodeClient(cfg),
		File:    NewFileClient(cfg),
		Profile: NewProfileClient(cfg),
		Series:  NewSeriesClient(cfg),
		Token:   NewTokenClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Account.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Account.Use(hooks...)
	c.Episode.Use(hooks...)
	c.File.Use(hooks...)
	c.Profile.Use(hooks...)
	c.Series.Use(hooks...)
	c.Token.Use(hooks...)
}

// AccountClient is a client for the Account schema.
type AccountClient struct {
	config
}

// NewAccountClient returns a client for the Account from the given config.
func NewAccountClient(c config) *AccountClient {
	return &AccountClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `account.Hooks(f(g(h())))`.
func (c *AccountClient) Use(hooks ...Hook) {
	c.hooks.Account = append(c.hooks.Account, hooks...)
}

// Create returns a create builder for Account.
func (c *AccountClient) Create() *AccountCreate {
	mutation := newAccountMutation(c.config, OpCreate)
	return &AccountCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Account entities.
func (c *AccountClient) CreateBulk(builders ...*AccountCreate) *AccountCreateBulk {
	return &AccountCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Account.
func (c *AccountClient) Update() *AccountUpdate {
	mutation := newAccountMutation(c.config, OpUpdate)
	return &AccountUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AccountClient) UpdateOne(a *Account) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccount(a))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AccountClient) UpdateOneID(id uuid.UUID) *AccountUpdateOne {
	mutation := newAccountMutation(c.config, OpUpdateOne, withAccountID(id))
	return &AccountUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Account.
func (c *AccountClient) Delete() *AccountDelete {
	mutation := newAccountMutation(c.config, OpDelete)
	return &AccountDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *AccountClient) DeleteOne(a *Account) *AccountDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *AccountClient) DeleteOneID(id uuid.UUID) *AccountDeleteOne {
	builder := c.Delete().Where(account.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AccountDeleteOne{builder}
}

// Query returns a query builder for Account.
func (c *AccountClient) Query() *AccountQuery {
	return &AccountQuery{
		config: c.config,
	}
}

// Get returns a Account entity by its id.
func (c *AccountClient) Get(ctx context.Context, id uuid.UUID) (*Account, error) {
	return c.Query().Where(account.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AccountClient) GetX(ctx context.Context, id uuid.UUID) *Account {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryToken queries the token edge of a Account.
func (c *AccountClient) QueryToken(a *Account) *TokenQuery {
	query := &TokenQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(token.Table, token.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.TokenTable, account.TokenColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProfile queries the profile edge of a Account.
func (c *AccountClient) QueryProfile(a *Account) *ProfileQuery {
	query := &ProfileQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.ProfileTable, account.ProfileColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryFile queries the file edge of a Account.
func (c *AccountClient) QueryFile(a *Account) *FileQuery {
	query := &FileQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(account.Table, account.FieldID, id),
			sqlgraph.To(file.Table, file.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, account.FileTable, account.FileColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AccountClient) Hooks() []Hook {
	return c.hooks.Account
}

// EpisodeClient is a client for the Episode schema.
type EpisodeClient struct {
	config
}

// NewEpisodeClient returns a client for the Episode from the given config.
func NewEpisodeClient(c config) *EpisodeClient {
	return &EpisodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `episode.Hooks(f(g(h())))`.
func (c *EpisodeClient) Use(hooks ...Hook) {
	c.hooks.Episode = append(c.hooks.Episode, hooks...)
}

// Create returns a create builder for Episode.
func (c *EpisodeClient) Create() *EpisodeCreate {
	mutation := newEpisodeMutation(c.config, OpCreate)
	return &EpisodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Episode entities.
func (c *EpisodeClient) CreateBulk(builders ...*EpisodeCreate) *EpisodeCreateBulk {
	return &EpisodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Episode.
func (c *EpisodeClient) Update() *EpisodeUpdate {
	mutation := newEpisodeMutation(c.config, OpUpdate)
	return &EpisodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EpisodeClient) UpdateOne(e *Episode) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisode(e))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EpisodeClient) UpdateOneID(id uuid.UUID) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisodeID(id))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Episode.
func (c *EpisodeClient) Delete() *EpisodeDelete {
	mutation := newEpisodeMutation(c.config, OpDelete)
	return &EpisodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *EpisodeClient) DeleteOne(e *Episode) *EpisodeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *EpisodeClient) DeleteOneID(id uuid.UUID) *EpisodeDeleteOne {
	builder := c.Delete().Where(episode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EpisodeDeleteOne{builder}
}

// Query returns a query builder for Episode.
func (c *EpisodeClient) Query() *EpisodeQuery {
	return &EpisodeQuery{
		config: c.config,
	}
}

// Get returns a Episode entity by its id.
func (c *EpisodeClient) Get(ctx context.Context, id uuid.UUID) (*Episode, error) {
	return c.Query().Where(episode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EpisodeClient) GetX(ctx context.Context, id uuid.UUID) *Episode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProfile queries the profile edge of a Episode.
func (c *EpisodeClient) QueryProfile(e *Episode) *ProfileQuery {
	query := &ProfileQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, id),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, episode.ProfileTable, episode.ProfileColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySeries queries the series edge of a Episode.
func (c *EpisodeClient) QuerySeries(e *Episode) *SeriesQuery {
	query := &SeriesQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, id),
			sqlgraph.To(series.Table, series.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, episode.SeriesTable, episode.SeriesColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EpisodeClient) Hooks() []Hook {
	return c.hooks.Episode
}

// FileClient is a client for the File schema.
type FileClient struct {
	config
}

// NewFileClient returns a client for the File from the given config.
func NewFileClient(c config) *FileClient {
	return &FileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `file.Hooks(f(g(h())))`.
func (c *FileClient) Use(hooks ...Hook) {
	c.hooks.File = append(c.hooks.File, hooks...)
}

// Create returns a create builder for File.
func (c *FileClient) Create() *FileCreate {
	mutation := newFileMutation(c.config, OpCreate)
	return &FileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of File entities.
func (c *FileClient) CreateBulk(builders ...*FileCreate) *FileCreateBulk {
	return &FileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for File.
func (c *FileClient) Update() *FileUpdate {
	mutation := newFileMutation(c.config, OpUpdate)
	return &FileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *FileClient) UpdateOne(f *File) *FileUpdateOne {
	mutation := newFileMutation(c.config, OpUpdateOne, withFile(f))
	return &FileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *FileClient) UpdateOneID(id uuid.UUID) *FileUpdateOne {
	mutation := newFileMutation(c.config, OpUpdateOne, withFileID(id))
	return &FileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for File.
func (c *FileClient) Delete() *FileDelete {
	mutation := newFileMutation(c.config, OpDelete)
	return &FileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FileClient) DeleteOne(f *File) *FileDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FileClient) DeleteOneID(id uuid.UUID) *FileDeleteOne {
	builder := c.Delete().Where(file.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &FileDeleteOne{builder}
}

// Query returns a query builder for File.
func (c *FileClient) Query() *FileQuery {
	return &FileQuery{
		config: c.config,
	}
}

// Get returns a File entity by its id.
func (c *FileClient) Get(ctx context.Context, id uuid.UUID) (*File, error) {
	return c.Query().Where(file.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FileClient) GetX(ctx context.Context, id uuid.UUID) *File {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccount queries the account edge of a File.
func (c *FileClient) QueryAccount(f *File) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := f.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(file.Table, file.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, file.AccountTable, file.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(f.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *FileClient) Hooks() []Hook {
	return c.hooks.File
}

// ProfileClient is a client for the Profile schema.
type ProfileClient struct {
	config
}

// NewProfileClient returns a client for the Profile from the given config.
func NewProfileClient(c config) *ProfileClient {
	return &ProfileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `profile.Hooks(f(g(h())))`.
func (c *ProfileClient) Use(hooks ...Hook) {
	c.hooks.Profile = append(c.hooks.Profile, hooks...)
}

// Create returns a create builder for Profile.
func (c *ProfileClient) Create() *ProfileCreate {
	mutation := newProfileMutation(c.config, OpCreate)
	return &ProfileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Profile entities.
func (c *ProfileClient) CreateBulk(builders ...*ProfileCreate) *ProfileCreateBulk {
	return &ProfileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Profile.
func (c *ProfileClient) Update() *ProfileUpdate {
	mutation := newProfileMutation(c.config, OpUpdate)
	return &ProfileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ProfileClient) UpdateOne(pr *Profile) *ProfileUpdateOne {
	mutation := newProfileMutation(c.config, OpUpdateOne, withProfile(pr))
	return &ProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ProfileClient) UpdateOneID(id uuid.UUID) *ProfileUpdateOne {
	mutation := newProfileMutation(c.config, OpUpdateOne, withProfileID(id))
	return &ProfileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Profile.
func (c *ProfileClient) Delete() *ProfileDelete {
	mutation := newProfileMutation(c.config, OpDelete)
	return &ProfileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ProfileClient) DeleteOne(pr *Profile) *ProfileDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ProfileClient) DeleteOneID(id uuid.UUID) *ProfileDeleteOne {
	builder := c.Delete().Where(profile.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ProfileDeleteOne{builder}
}

// Query returns a query builder for Profile.
func (c *ProfileClient) Query() *ProfileQuery {
	return &ProfileQuery{
		config: c.config,
	}
}

// Get returns a Profile entity by its id.
func (c *ProfileClient) Get(ctx context.Context, id uuid.UUID) (*Profile, error) {
	return c.Query().Where(profile.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ProfileClient) GetX(ctx context.Context, id uuid.UUID) *Profile {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccount queries the account edge of a Profile.
func (c *ProfileClient) QueryAccount(pr *Profile) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, profile.AccountTable, profile.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEpisode queries the episode edge of a Profile.
func (c *ProfileClient) QueryEpisode(pr *Profile) *EpisodeQuery {
	query := &EpisodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, id),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, profile.EpisodeTable, profile.EpisodeColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySeries queries the series edge of a Profile.
func (c *ProfileClient) QuerySeries(pr *Profile) *SeriesQuery {
	query := &SeriesQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(profile.Table, profile.FieldID, id),
			sqlgraph.To(series.Table, series.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, profile.SeriesTable, profile.SeriesColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ProfileClient) Hooks() []Hook {
	return c.hooks.Profile
}

// SeriesClient is a client for the Series schema.
type SeriesClient struct {
	config
}

// NewSeriesClient returns a client for the Series from the given config.
func NewSeriesClient(c config) *SeriesClient {
	return &SeriesClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `series.Hooks(f(g(h())))`.
func (c *SeriesClient) Use(hooks ...Hook) {
	c.hooks.Series = append(c.hooks.Series, hooks...)
}

// Create returns a create builder for Series.
func (c *SeriesClient) Create() *SeriesCreate {
	mutation := newSeriesMutation(c.config, OpCreate)
	return &SeriesCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Series entities.
func (c *SeriesClient) CreateBulk(builders ...*SeriesCreate) *SeriesCreateBulk {
	return &SeriesCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Series.
func (c *SeriesClient) Update() *SeriesUpdate {
	mutation := newSeriesMutation(c.config, OpUpdate)
	return &SeriesUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SeriesClient) UpdateOne(s *Series) *SeriesUpdateOne {
	mutation := newSeriesMutation(c.config, OpUpdateOne, withSeries(s))
	return &SeriesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SeriesClient) UpdateOneID(id uuid.UUID) *SeriesUpdateOne {
	mutation := newSeriesMutation(c.config, OpUpdateOne, withSeriesID(id))
	return &SeriesUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Series.
func (c *SeriesClient) Delete() *SeriesDelete {
	mutation := newSeriesMutation(c.config, OpDelete)
	return &SeriesDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *SeriesClient) DeleteOne(s *Series) *SeriesDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *SeriesClient) DeleteOneID(id uuid.UUID) *SeriesDeleteOne {
	builder := c.Delete().Where(series.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SeriesDeleteOne{builder}
}

// Query returns a query builder for Series.
func (c *SeriesClient) Query() *SeriesQuery {
	return &SeriesQuery{
		config: c.config,
	}
}

// Get returns a Series entity by its id.
func (c *SeriesClient) Get(ctx context.Context, id uuid.UUID) (*Series, error) {
	return c.Query().Where(series.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SeriesClient) GetX(ctx context.Context, id uuid.UUID) *Series {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryProfile queries the profile edge of a Series.
func (c *SeriesClient) QueryProfile(s *Series) *ProfileQuery {
	query := &ProfileQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(series.Table, series.FieldID, id),
			sqlgraph.To(profile.Table, profile.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, series.ProfileTable, series.ProfileColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryEpisode queries the episode edge of a Series.
func (c *SeriesClient) QueryEpisode(s *Series) *EpisodeQuery {
	query := &EpisodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(series.Table, series.FieldID, id),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, series.EpisodeTable, series.EpisodeColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SeriesClient) Hooks() []Hook {
	return c.hooks.Series
}

// TokenClient is a client for the Token schema.
type TokenClient struct {
	config
}

// NewTokenClient returns a client for the Token from the given config.
func NewTokenClient(c config) *TokenClient {
	return &TokenClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `token.Hooks(f(g(h())))`.
func (c *TokenClient) Use(hooks ...Hook) {
	c.hooks.Token = append(c.hooks.Token, hooks...)
}

// Create returns a create builder for Token.
func (c *TokenClient) Create() *TokenCreate {
	mutation := newTokenMutation(c.config, OpCreate)
	return &TokenCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Token entities.
func (c *TokenClient) CreateBulk(builders ...*TokenCreate) *TokenCreateBulk {
	return &TokenCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Token.
func (c *TokenClient) Update() *TokenUpdate {
	mutation := newTokenMutation(c.config, OpUpdate)
	return &TokenUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TokenClient) UpdateOne(t *Token) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withToken(t))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TokenClient) UpdateOneID(id uuid.UUID) *TokenUpdateOne {
	mutation := newTokenMutation(c.config, OpUpdateOne, withTokenID(id))
	return &TokenUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Token.
func (c *TokenClient) Delete() *TokenDelete {
	mutation := newTokenMutation(c.config, OpDelete)
	return &TokenDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TokenClient) DeleteOne(t *Token) *TokenDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TokenClient) DeleteOneID(id uuid.UUID) *TokenDeleteOne {
	builder := c.Delete().Where(token.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TokenDeleteOne{builder}
}

// Query returns a query builder for Token.
func (c *TokenClient) Query() *TokenQuery {
	return &TokenQuery{
		config: c.config,
	}
}

// Get returns a Token entity by its id.
func (c *TokenClient) Get(ctx context.Context, id uuid.UUID) (*Token, error) {
	return c.Query().Where(token.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TokenClient) GetX(ctx context.Context, id uuid.UUID) *Token {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryAccount queries the account edge of a Token.
func (c *TokenClient) QueryAccount(t *Token) *AccountQuery {
	query := &AccountQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(token.Table, token.FieldID, id),
			sqlgraph.To(account.Table, account.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, token.AccountTable, token.AccountColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TokenClient) Hooks() []Hook {
	return c.hooks.Token
}
