package main

import (
	"testing"
)

func TestRun(t *testing.T) {
	Err := Run()
	if Err != nil {
		t.Error(Err)
	}
}
