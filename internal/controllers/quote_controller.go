package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/superdumb33/qoute_service/internal/services"
)

type QuoteController struct {
	service *services.QuoteService
}

func NewQuoteController (service *services.QuoteService) *QuoteController {
	return &QuoteController{service: service}
}

func (qc *QuoteController) RegisterRoutes (r *mux.Router) {
	r.HandleFunc("/quotes", qc.CreateQuote).Methods(http.MethodPost)
}


func (qc *QuoteController) CreateQuote (w http.ResponseWriter, r *http.Request) {
	
}