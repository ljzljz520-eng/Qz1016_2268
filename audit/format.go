package audit

import (
	"facility-login/model"
	"fmt"
	"time"
)

func Format(a model.Audit) string {
	return fmt.Sprintf("%s %s %s", a.At.Format(time.RFC3339), a.Actor, a.Action)
}
func Sort(xs []model.Audit) []model.Audit {
	for i := 0; i < len(xs); i++ {
		for j := i + 1; j < len(xs); j++ {
			if xs[j].At.Before(xs[i].At) {
				xs[i], xs[j] = xs[j], xs[i]
			}
		}
	}
	return xs
}
func CountAction(xs []model.Audit, action string) int {
	n := 0
	for _, a := range xs {
		if a.Action == action {
			n++
		}
	}
	return n
}
