package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewRouter(handler *http.Handler) http.Handler {
	r := chi.NewRouter()
	return r
}
