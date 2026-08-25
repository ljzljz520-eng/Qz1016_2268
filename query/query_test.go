package query

import (
	"facility-login/model"
	"facility-login/storage"
	"testing"
	"time"
)

func TestSearch(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	n := time.Now()
	s.SaveRecord(model.NewRecord("r", "alice", n, n.Add(time.Hour)))
	x, e := Search(s, "ali", "")
	if e != nil || len(x.Records) != 1 {
		t.Fatal(x, e)
	}
}
