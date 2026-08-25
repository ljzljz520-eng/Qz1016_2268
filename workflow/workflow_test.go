package workflow

import (
	"facility-login/storage"
	"testing"
	"time"
)

func TestWorkflowOne(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	e := New(s)
	if e.Register("r", "u", time.Now()) != nil {
		t.Fatal()
	}
}
