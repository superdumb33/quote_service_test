package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateQuoteRequest struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

type CreateQuoteResponse struct {
	ID uuid.UUID `json:"id"`
}


type GetAllQuotesResponse struct {
	ID     uuid.UUID `json:"id"`
	Quote  string`json:"quote"`
	Author string`json:"author"`
	CreatedAt time.Time`json:"created_at"`
}

type GetRandomQuoteResponse struct{
	ID     uuid.UUID `json:"id"`
	Quote  string`json:"quote"`
	Author string`json:"author"`
	CreatedAt time.Time`json:"created_at"`
}
type GetQuotesByAuthorResponse struct{
	ID     uuid.UUID `json:"id"`
	Quote  string`json:"quote"`
	Author string`json:"author"`
	CreatedAt time.Time`json:"created_at"`
}

type ErrorResponse struct{
	Error string`json:"error"`
}