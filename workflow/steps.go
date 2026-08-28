package workflow

import (
	"facility-login/model"
	"time"
)

type Step struct {
	Name string
	Done bool
	At   time.Time
}

func BuildSteps(names []string, now time.Time) []Step {
	out := make([]Step, len(names))
	for i, n := range names {
		out[i] = Step{Name: n, At: now}
	}
	return out
}
func CompleteStep(xs []Step, name string, now time.Time) []Step {
	for i := range xs {
		if xs[i].Name == name {
			xs[i].Done = true
			xs[i].At = now
		}
	}
	return xs
}
func WorkflowStatus(r model.Record) string {
	if r.Status == string(model.StatusCurrent) {
		return "ready"
	}
	return "attention"
}
