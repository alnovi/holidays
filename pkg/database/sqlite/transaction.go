package sqlite

import (
	"context"
	"database/sql"
	"fmt"
)

type Transaction struct {
	db *sql.DB
}

func NewTransaction(db *sql.DB) *Transaction {
	return &Transaction{db: db}
}

func (t *Transaction) ReadCommitted(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.transaction(ctx, fn)
}

func (t *Transaction) transaction(ctx context.Context, fn func(ctx context.Context) error) error {
	if _, ok := ctx.Value(txKey).(*sql.Tx); ok {
		return fn(ctx)
	}

	tx, err := t.db.Begin()
	if err != nil {
		return fmt.Errorf("can't begin transaction: %w", err)
	}

	defer func() {
		_ = tx.Rollback()
	}()

	err = fn(context.WithValue(ctx, txKey, tx))
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("can't commit transaction: %w", err)
	}

	return nil
}
