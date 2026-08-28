package service

import (
	"facility-login/model"
	"facility-login/storage"
	"testing"
	"time"
)

func TestServiceRegister(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	x := New(s)
	n := time.Now()
	if x.Register(model.NewRecord("r", "u", n, n.Add(time.Hour)), model.User{ID: "u", Name: "U", Role: "operator", Active: true}) != nil {
		t.Fatal()
	}
}
