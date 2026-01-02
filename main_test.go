package main

import (
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	expected_output := "You've landed on the home page - there is nothing to see here"
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	home(rec, req)

	if (rec.Body.String() != expected_output) {
		t.Errorf("got %q, want %q", rec.Body.String(), expected_output)	
	}
}
