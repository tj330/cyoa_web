package utils

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ReadArcParam(r *http.Request) (*string, error) {
	arc := chi.URLParam(r, "arc")
	if arc == "" {
		return nil, errors.New("invalid arc parameter")
	}

	return &arc, nil
}
