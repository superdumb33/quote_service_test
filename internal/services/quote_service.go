package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/superdumb33/qoute_service/internal/entities"
)

type QuoteRepo interface {
	CreateQuote(ctx context.Context, quote *entities.Quote) error
	GetQuotesByAuthor (ctx context.Context, query string) ([]*entities.Quote, error)
	DeleteQuoteByID (ctx context.Context, quoteID uuid.UUID) error
}

type QuoteService struct {
	repo QuoteRepo
}

func NewQuoteService (repo QuoteRepo) *QuoteService {
	return &QuoteService{repo: repo}
}