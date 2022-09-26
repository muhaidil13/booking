package main

import (
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var mh myHandler
	h := NoSurf(&mh)
	switch h.(type) {
	case http.Handler:
		//
	default:
		t.Error("error")
	}
}
func TestSessionLoad(t *testing.T) {
	var mh myHandler
	h := SessionLoad(&mh)
	switch h.(type) {
	case http.Handler:
		//
	default:
		t.Error("error")
	}
}
