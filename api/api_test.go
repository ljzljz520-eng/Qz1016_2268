package api

import (
	"facility-login/storage"
	"facility-login/workflow"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	h := New(workflow.New(s))
	r := httptest.NewRequest("GET", "/missing/r", nil)
	w := httptest.NewRecorder()
	h.ServeHTTP(w, r)
	if w.Code != 404 {
		t.Fatal(w.Code)
	}
}
