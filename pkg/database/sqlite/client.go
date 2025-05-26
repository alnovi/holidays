package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/georgysavva/scany/v2/sqlscan"
	_ "github.com/mattn/go-sqlite3"
)

type key string

const txKey key = "tx"

type Option func(c *Client) error

func WithLogger(logger *slog.Logger) Option {
	return func(c *Client) error {
		if logger != nil {
			c.logger = logger
		}
		return nil
	}
}

type Client struct {
	master *sql.DB
	logger *slog.Logger
}

func NewClient(db string, opts ...Option) (*Client, error) {
	master, err := sql.Open("sqlite3", db)
	if err != nil {
		return nil, err
	}

	c := &Client{
		master: master,
		logger: slog.New(slog.DiscardHandler),
	}

	for _, opt := range opts {
		if err = opt(c); err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) Master() *sql.DB {
	return c.master
}

func (c *Client) Ping(ctx context.Context) error {
	return c.master.PingContext(ctx)
}

func (c *Client) Exec(ctx context.Context, query string, args ...any) (sql.Result, error) {
	c.logger.Debug(query, logArgs(args)...)
	if tx, ok := ctx.Value(txKey).(sql.Tx); ok {
		return tx.ExecContext(ctx, query, args...)
	}
	return c.master.ExecContext(ctx, query, args...)
}

func (c *Client) Query(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	c.logger.Debug(query, logArgs(args)...)
	if tx, ok := ctx.Value(txKey).(sql.Tx); ok {
		return tx.QueryContext(ctx, query, args...)
	}
	return c.master.QueryContext(ctx, query, args...)
}

func (c *Client) QueryRow(ctx context.Context, query string, args ...any) *sql.Row {
	c.logger.Debug(query, logArgs(args)...)
	if tx, ok := ctx.Value(txKey).(sql.Tx); ok {
		return tx.QueryRowContext(ctx, query, args...)
	}
	return c.master.QueryRowContext(ctx, query, args...)
}

func (c *Client) ScanQuery(ctx context.Context, dst any, query string, args ...any) error {
	c.logger.Debug(query, logArgs(args)...)
	rows, err := c.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	defer func() {
		_ = rows.Close()
	}()
	return sqlscan.ScanAll(dst, rows)
}

func (c *Client) ScanQueryRow(ctx context.Context, dst any, query string, args ...any) error {
	c.logger.Debug(query, logArgs(args)...)
	rows, err := c.Query(ctx, query, args...)
	if err != nil {
		return err
	}
	defer func() {
		_ = rows.Close()
	}()
	return sqlscan.ScanOne(dst, rows)
}

func (c *Client) Close(_ context.Context) error {
	return c.master.Close()
}

func logArgs(args []any) []any {
	attr := make([]any, 0, len(args)*2) //nolint:mnd
	for i, arg := range args {
		k := fmt.Sprintf("$%d", i+1)
		attr = append(attr, slog.Any(k, arg))
	}
	return attr
}
