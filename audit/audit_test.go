package audit

import (
	"facility-login/model"
	"facility-login/storage"
	"testing"
	"time"
)

func TestBuild(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	s.SaveAudit(model.NewAudit("a", "r", "u", "review", time.Now()))
	r, e := Build(s, "r")
	if e != nil || r.Total != 1 {
		t.Fatal(r, e)
	}
}
