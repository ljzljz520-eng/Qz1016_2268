package audit

import (
	"encoding/json"
	"facility-login/model"
	"facility-login/storage"
	"time"
)

type Report struct {
	Total   int
	ByActor map[string]int
	Latest  time.Time
}

func Build(s *storage.Store, id string) (Report, error) {
	xs, e := s.ListAudits(id)
	if e != nil {
		return Report{}, e
	}
	r := Report{ByActor: map[string]int{}}
	for _, a := range xs {
		r.Total++
		r.ByActor[a.Actor]++
		if a.At.After(r.Latest) {
			r.Latest = a.At
		}
	}
	return r, nil
}
func Encode(r Report) ([]byte, error) { return json.Marshal(r) }
func Actions(s *storage.Store, id string) ([]string, error) {
	xs, e := s.ListAudits(id)
	if e != nil {
		return nil, e
	}
	out := []string{}
	for _, a := range xs {
		out = append(out, a.Action)
	}
	return out, nil
}
func Recent(s *storage.Store, id string, since time.Time) ([]model.Audit, error) {
	xs, e := s.ListAudits(id)
	if e != nil {
		return nil, e
	}
	out := []model.Audit{}
	for _, a := range xs {
		if a.At.After(since) {
			out = append(out, a)
		}
	}
	return out, nil
}
