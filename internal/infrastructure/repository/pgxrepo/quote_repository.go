package pgxrepo

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/superdumb33/qoute_service/internal/entities"
)

var (
	ErrNotFound = entities.ErrNotFound
	ErrInternal = entities.ErrInternal
)

type PgxQuoteRepo struct {
	db  *pgxpool.Pool
	log *slog.Logger
}

func NewPgxQuoteRepo(pool *pgxpool.Pool, log *slog.Logger) *PgxQuoteRepo {
	return &PgxQuoteRepo{db: pool, log: log}
}

// CreateQuote will scan created record's ID into passed quote
func (qr *PgxQuoteRepo) CreateQuote(ctx context.Context, quote *entities.Quote) error {
	const op = "repo.CreateQuote"

	query := `INSERT INTO quotes(quote, author) VALUES ($1, $2) RETURNING id`
	if err := qr.db.QueryRow(ctx, query, quote.Quote, quote.Author).Scan(&quote.ID); err != nil {
		qr.log.Error(op, "error", err)
		return fmt.Errorf("%s:%w:%v", op, ErrInternal, err)
	}

	return nil
}

func (qr *PgxQuoteRepo) GetQuotesByAuthor(ctx context.Context, author string) ([]*entities.Quote, error) {
	const op = "repo.GetQuotesByAuthor"
	query := `SELECT id, quote, author, created_at FROM quotes WHERE author = $1 AND deleted_at IS NULL`

	rows, err := qr.db.Query(ctx, query, author)
	if err != nil {
		qr.log.Error(op, "error", err)
		return nil, fmt.Errorf("%s:%w:%v", op, ErrInternal, err)
	}
	defer rows.Close()

	var quotes []*entities.Quote
	for rows.Next() {
		var quote entities.Quote
		if err := rows.Scan(&quote.ID, &quote.Quote, &quote.Author, &quote.CreatedAt); err != nil {
			qr.log.Error(op, "error", err)
			return nil, fmt.Errorf("%s:%w:%v", op, ErrInternal, err)
		}

		quotes = append(quotes, &quote)
	}
	if err := rows.Err(); err != nil {
		qr.log.Error(op, "rows iteration error", err)
		return nil, fmt.Errorf("%s:%w:%v", op, ErrInternal, err)
	}

	if len(quotes) == 0 {
		return nil, fmt.Errorf("%s:%w", op, ErrNotFound)
	}

	return quotes, nil
}

func (qr *PgxQuoteRepo) GetAllQuotes(ctx context.Context) ([]*entities.Quote, error) {
	const op = "repo.GetAllQuotes"
	query := `SELECT id, quote, author, created_at, deleted_at FROM quotes WHERE deleted_at IS NULL`

	rows, err := qr.db.Query(ctx, query)
	if err != nil {
		qr.log.Error(op, "error", err)
		return nil, fmt.Errorf("%s:%w:%v", op, ErrInternal, err)
	}
	defer rows.Close()

	var quotes []*entities.Quote
	for rows.Next() {
		var quote entities.Quote
		if err := rows.Scan(&quote.ID, &quote.Quote, &quote.Author, &quote.CreatedAt, &quote.DeletedAt); err != nil {
			return nil, err
		}

		quotes = append(quotes, &quote)
	}
	if err := rows.Err(); err != nil {
		qr.log.Error(op, "rows iteration error", err)
		return nil, fmt.Errorf("%s:%w:%v", op, ErrInternal, err)
	}

	return quotes, nil
}

func (qr *PgxQuoteRepo) GetRandomQuote(ctx context.Context) (*entities.Quote, error) {
	const op = "repo.GetRandomQuote"
	query := `SELECT id, quote, author, created_at FROM quotes WHERE deleted_at IS NULL ORDER BY random() LIMIT 1`

	var quote entities.Quote
	if err := qr.db.QueryRow(ctx, query).Scan(&quote.ID, &quote.Quote, &quote.Author, &quote.CreatedAt); err != nil {
		if err == pgx.ErrNoRows {
			return nil, fmt.Errorf("%s:%w", op, ErrNotFound)
		}
		qr.log.Error(op, "error", err)
		return nil, fmt.Errorf("%s:%w:%v", op, ErrInternal, err)
	}

	return &quote, nil
}

func (qr *PgxQuoteRepo) DeleteQuoteByID(ctx context.Context, quoteID uuid.UUID) error {
	const op = "repo.DeleteQuoteByID"
	query := `UPDATE quotes SET deleted_at = NOW() WHERE id = $1`

	tag, err := qr.db.Exec(ctx, query, quoteID)
	if err != nil {
		return fmt.Errorf("%s:%w:%v", op, ErrInternal, err)
	}
	if tag.RowsAffected() == 0 {
		return fmt.Errorf("%s:%w", op, ErrNotFound)
	}

	return nil
}
