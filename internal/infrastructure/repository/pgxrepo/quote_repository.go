package pgxrepo

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/superdumb33/qoute_service/internal/entities"
)

type PgxQuoteRepo struct {
	db *pgxpool.Pool
}

func NewPgxQuoteRepo (pool *pgxpool.Pool) *PgxQuoteRepo {
	return &PgxQuoteRepo{db: pool}
}

func (qr *PgxQuoteRepo) CreateQuote (ctx context.Context, quote *entities.Quote) error {
	return nil
}

func (qr *PgxQuoteRepo) GetQuotesByAuthor (ctx context.Context, query string) ([]*entities.Quote, error){
	return nil, nil
}


func (qr *PgxQuoteRepo) DeleteQuoteByID (ctx context.Context, quoteID uuid.UUID) error {
	return nil
}