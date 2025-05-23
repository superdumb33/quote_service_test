package controllers

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/superdumb33/qoute_service/internal/dto"
	"github.com/superdumb33/qoute_service/internal/entities"
	"github.com/superdumb33/qoute_service/internal/services"
)

type QuoteController struct {
	service *services.QuoteService
	log     *slog.Logger
}

func NewQuoteController(service *services.QuoteService, log *slog.Logger) *QuoteController {
	return &QuoteController{service: service, log: log}
}
func (qc *QuoteController) RegisterRoutes(r *mux.Router) {
	api := r.PathPrefix("/quotes").Subrouter()

	api.HandleFunc("", qc.CreateQuote).Methods(http.MethodPost)
	api.HandleFunc("", qc.GetQuotesByAuthor).Methods(http.MethodGet).Queries("author", "{author}")
	api.HandleFunc("", qc.GetAllQuotes).Methods(http.MethodGet)
	api.HandleFunc("/random", qc.GetRandomQuote).Methods(http.MethodGet)
	api.HandleFunc("/{id}", qc.DeleteQuoteByID).Methods(http.MethodDelete)
}
// @Summary      Create a new quote
// @Description  Creates a new record with provided quote
// @Tags         quotes
// @Accept       json
// @Produce      json
// @Param        payload  body      dto.CreateQuoteRequest  true  "Quote payload"
// @Success      201      {object}  dto.CreateQuoteResponse
// @Failure      400      {object}  dto.ErrorResponse
// @Failure      500      {object}  dto.ErrorResponse
// @Router       /quotes [post]
func (qc *QuoteController) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateQuoteRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request: invalid JSON", http.StatusBadRequest)
		return
	}
	if req.Author == "" || req.Quote == "" {
		http.Error(w, "Bad request: author and quote must be non-empty", http.StatusBadRequest)
		return
	}

	quote := &entities.Quote{Author: req.Author, Quote: req.Quote}
	if err := qc.service.CreateQuote(r.Context(), quote); err != nil {
		qc.log.Error("CreateQuote failed", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp := &dto.CreateQuoteResponse{
		ID: quote.ID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
// @Summary      Get quotes
// @Description  Returns all of the quotes or filters by author if query is provided
// @Tags         quotes
// @Produce      json
// @Param        author  query     string  false  "Quote author"
// @Success      200     {array}   dto.GetAllQuotesResponse
// @Failure      500     {object}  dto.ErrorResponse
// @Router       /quotes [get]
func (qc *QuoteController) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	quotes, err := qc.service.GetAllQuotes(r.Context())
	if err != nil {
		qc.log.Error("GetAllQuotes failed", "error", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp := make([]dto.GetAllQuotesResponse, 0, len(quotes))
	for _, q := range quotes {
		resp = append(resp, dto.GetAllQuotesResponse{
			ID:        q.ID,
			Author:    q.Author,
			Quote:     q.Quote,
			CreatedAt: q.CreatedAt,
		})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Summary      Get a random quote
// @Description  Returs random quote
// @Tags         quotes
// @Produce      json
// @Success      200  {object}  dto.GetRandomQuoteResponse
// @Failure      404  {object}  dto.ErrorResponse  "if there's no stored qoutes"
// @Router       /quotes/random [get]
func (qc *QuoteController) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote, err := qc.service.GetRandomQuote(r.Context())
	if err != nil {
		http.Error(w, err.Error(), statusCodeFromError(err))
		return
	}
	resp := &dto.GetRandomQuoteResponse{
		ID:        quote.ID,
		Quote:     quote.Quote,
		Author:    quote.Author,
		CreatedAt: quote.CreatedAt,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (qc *QuoteController) GetQuotesByAuthor(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	quotes, err := qc.service.GetQuotesByAuthor(r.Context(), author)
	if err != nil {
		http.Error(w, err.Error(), statusCodeFromError(err))
		return
	}
	resp := make([]dto.GetAllQuotesResponse, 0, len(quotes))
	for _, q := range quotes {
		resp = append(resp, dto.GetAllQuotesResponse{
			ID:        q.ID,
			Author:    q.Author,
			Quote:     q.Quote,
			CreatedAt: q.CreatedAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
// @Summary      Delete quote by ID
// @Description  Deletes a quote by it's UUID
// @Tags         quotes
// @Param        id   path      string  true  "Quote ID"
// @Success      204  "No Content"
// @Failure      400  {object}  dto.ErrorResponse  "invalid ID"
// @Failure      404  {object}  dto.ErrorResponse  "not found"
// @Router       /quotes/{id} [delete]
func (qc *QuoteController) DeleteQuoteByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, "Bad request: invalid ID", http.StatusBadRequest)
		return
	}
	if err := qc.service.DeleteQuoteByID(r.Context(), id); err != nil {
		http.Error(w, err.Error(), statusCodeFromError(err))
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func statusCodeFromError(err error) int {
	switch {
	case errors.Is(err, entities.ErrNotFound):
		return 404
	default:
		return 500
	}
}
