package transaction

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Transactioner interface {
	WithTx(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)
}

var txKey = struct{}{}

func WithTx(ctx context.Context, db *sqlx.DB) (*sqlx.Tx, error) {
	tx, ok := getTx(ctx)
	if ok {
		return tx, nil
	}

	tx, err := db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func getTx(ctx context.Context) (*sqlx.Tx, bool) {
	tx, ok := ctx.Value(&txKey).(*sqlx.Tx)
	return tx, ok
}

func withTx(ctx context.Context, tx *sqlx.Tx) context.Context {
	return context.WithValue(ctx, &txKey, tx)
}

type transactioner struct {
	db *sqlx.DB
}

func NewTransactioner(db *sqlx.DB) Transactioner {
	return &transactioner{db: db}
}

func (t transactioner) WithTx(ctx context.Context, fn func(context.Context) (interface{}, error)) (interface{}, error) {
	tx, err := t.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	ctx = withTx(ctx, tx)

	v, err := fn(ctx)
	if err != nil {
		err = fmt.Errorf("transactional operation failed: %w", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return nil, fmt.Errorf("rollback failed: %w: %v", err, rollbackErr)
		}
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		err = fmt.Errorf("commit failed: %w", err)
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return nil, fmt.Errorf("rollback failed: %w: %v", err, rollbackErr)
		}
		return nil, err
	}

	return v, nil
}
