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
	if e.Register("r", "u", time.Now()) != nil {
		t.Fatal()
	}
	if _, x := e.Approve("r", "rev"); x != nil {
		t.Fatal(x)
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
