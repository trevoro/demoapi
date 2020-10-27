package http

import (
	"net/http/httptest"
	"testing"
)

func TestPingHandler(t *testing.T) {
	s := &Server{}
	handler := s.handlePing()
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)
	if rec.Body.String() != "OK\n" {
		t.Fatal("error")
	}
}
