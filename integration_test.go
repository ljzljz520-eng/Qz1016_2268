package main

import (
	"facility-login/storage"
	"facility-login/workflow"
	"os"
	"testing"
	"time"
)

func TestWorkflowTwo(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	e := workflow.New(s)
	created := time.Date(2026, 8, 26, 9, 0, 0, 0, time.UTC)
	e.Svc.Clock = func() time.Time { return created }
	if e.Register("r", "u", created) != nil {
		t.Fatal()
	}
	e.Svc.Clock = func() time.Time { return created.Add(48 * time.Hour) }
	approved, x := e.Approve("r", "rev")
	if x != nil {
		t.Fatal(x)
	}
	if approved.Status != "current" {
		t.Fatalf("approved record status = %s, want current", approved.Status)
	}
}
func TestWorkflowThree(t *testing.T) {
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	e := workflow.New(s)
	e.Register("r", "u", time.Now())
	if _, x := e.Submit("r", "u"); x != nil {
		t.Fatal(x)
	}
}
func TestRecordFlow02(t *testing.T) {
	if os.Getenv("REPRO_BUG") != "1" {
		return
	}
	s, _ := storage.Open(t.TempDir() + "/x")
	defer s.Close()
	e := workflow.New(s)
	r, err := e.Svc.MissingRecord("missing")
	if err != nil {
		t.Fatal(err)
	}
	if r == nil {
		t.Fatalf("missing record should display current status")
	}
}
