package router

// Code generated by openapi-generator-go DO NOT EDIT., don't modify it manually

import (
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
)

// NewRouter creates a new router for the spec and the given handlers.
// Tunnelbot - OpenAPI 3.0
//
// This is the API specification for the tunnelbot backend.
//
// 1
func NewRouter() http.Handler {

	r := chi.NewRouter()

	return r
}

func optionsHandlerFunc(allowedMethods ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Allow", strings.Join(allowedMethods, ", "))
	}
}
