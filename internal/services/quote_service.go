package services

import (
	"context"
	"log/slog"

	"github.com/google/uuid"
	"github.com/superdumb33/qoute_service/internal/entities"
)

type QuoteRepo interface {
	CreateQuote(ctx context.Context, quote *entities.Quote) error
	GetAllQuotes(ctx context.Context) ([]*entities.Quote, error)
	GetRandomQuote(ctx context.Context) (*entities.Quote, error)
	GetQuotesByAuthor(ctx context.Context, author string) ([]*entities.Quote, error)
	DeleteQuoteByID(ctx context.Context, quoteID uuid.UUID) error
}

type QuoteService struct {
	repo QuoteRepo
	log  *slog.Logger
}

func NewQuoteService(repo QuoteRepo, log *slog.Logger) *QuoteService {
	return &QuoteService{repo: repo, log: log}
}

func (qs *QuoteService) CreateQuote(ctx context.Context, quote *entities.Quote) error {
	const op = "service.GetAllQuotes"
	if err := qs.repo.CreateQuote(ctx, quote); err != nil {
		return err
	}

	return nil
}

func (qs *QuoteService) GetAllQuotes(ctx context.Context) ([]*entities.Quote, error) {
	const op = "service.GetAllQuotes"
	quotes, err := qs.repo.GetAllQuotes(ctx)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func (qs *QuoteService) GetRandomQuote(ctx context.Context) (*entities.Quote, error) {
	quote, err := qs.repo.GetRandomQuote(ctx)
	if err != nil {
		return nil, err
	}

	return quote, nil
}

func (qs *QuoteService) GetQuotesByAuthor(ctx context.Context, author string) ([]*entities.Quote, error) {
	quotes, err := qs.repo.GetQuotesByAuthor(ctx, author)
	if err != nil {
		return nil, err
	}

	return quotes, nil
}

func (qs *QuoteService) DeleteQuoteByID(ctx context.Context, quoteID uuid.UUID) error {
	if err := qs.repo.DeleteQuoteByID(ctx, quoteID); err != nil {
		return err
	}
	
	return nil
}
