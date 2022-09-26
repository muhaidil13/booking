package main

import (
	"testing"

	"github.com/go-chi/chi"
	"github.com/muhaidil13/booking/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.Appconfig
	mux := routes(&app)
	switch mux.(type) {
	case *chi.Mux:
		// do anything
	default:
		t.Error("Type Not Equals *chi.Mux")
	}
}
