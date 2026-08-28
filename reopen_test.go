package main

import (
	"facility-login/model"
	"facility-login/storage"
	"testing"
	"time"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := t.TempDir() + "/persist.db"
	s, e := storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	n := time.Now()
	if e = s.SaveRecord(model.NewRecord("persist", "u", n, n.Add(time.Hour))); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = storage.Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	r, e := s.GetRecord("persist")
	if e != nil || r.ID != "persist" {
		t.Fatal(r, e)
	}
}
