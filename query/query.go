package query

import (
	"facility-login/model"
	"facility-login/storage"
	"sort"
	"strings"
	"time"
)

type Result struct {
	Records []model.Record
	Events  []model.Event
	Audits  []model.Audit
}

func Search(s *storage.Store, term, status string) (Result, error) {
	rs, e := s.ListRecords()
	if e != nil {
		return Result{}, e
	}
	out := []model.Record{}
	for _, r := range rs {
		if (term == "" || strings.Contains(strings.ToLower(r.Username), strings.ToLower(term))) && (status == "" || r.Status == status) {
			out = append(out, r)
		}
	}
	sort.Slice(out, func(i, j int) bool { return out[i].CreatedAt.Before(out[j].CreatedAt) })
	return Result{Records: out}, nil
}
func Timeline(s *storage.Store, id string) (Result, error) {
	e, er := s.ListEvents(id)
	if er != nil {
		return Result{}, er
	}
	a, er := s.ListAudits(id)
	if er != nil {
		return Result{}, er
	}
	r, er := s.GetRecord(id)
	if er != nil {
		return Result{}, er
	}
	return Result{Records: []model.Record{r}, Events: e, Audits: a}, nil
}
func CurrentRecords(s *storage.Store, now time.Time) ([]model.Record, error) {
	rs, e := s.ListRecords()
	if e != nil {
		return nil, e
	}
	out := []model.Record{}
	for _, r := range rs {
		if r.IsCurrent(now) {
			out = append(out, r)
		}
	}
	return out, nil
}
