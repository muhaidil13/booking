package main

import (
	"net/http"
	"os"
	"testing"
)

type myHandler struct{}

func TestMain(t *testing.M) {
	os.Exit(t.Run())
}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
