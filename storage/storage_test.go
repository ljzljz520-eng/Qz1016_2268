package storage

import (
	"facility-login/model"
	"testing"
	"time"
)

func TestStoreRoundTrip(t *testing.T) {
	p := t.TempDir() + "/x.db"
	s, _ := Open(p)
	r := model.NewRecord("r", "u", time.Now(), time.Now().Add(time.Hour))
	if s.SaveRecord(r) != nil {
		t.Fatal()
	}
	s.Close()
	s, _ = Open(p)
	if _, e := s.GetRecord("r"); e != nil {
		t.Fatal(e)
	}
	s.Close()
}
