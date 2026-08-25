package model

import (
	"testing"
	"time"
)

func TestRecordValidation(t *testing.T) {
	n := time.Now()
	if NewRecord("r", "u", n, n.Add(time.Hour)).Validate(n) != nil {
		t.Fatal()
	}
}
