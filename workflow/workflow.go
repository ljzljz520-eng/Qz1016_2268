package workflow

import (
	"facility-login/model"
	"facility-login/service"
	"facility-login/storage"
	"fmt"
	"time"
)

type Engine struct{ Svc *service.Service }

func New(s *storage.Store) *Engine { return &Engine{Svc: service.New(s)} }
func (e *Engine) Register(id, user string, now time.Time) error {
	r := model.NewRecord(id, user, now, now.Add(24*time.Hour))
	u := model.User{ID: user, Name: user, Role: "operator", Active: true}
	return e.Svc.Register(r, u)
}
func (e *Engine) Approve(id, reviewer string) (model.Record, error) {
	u := model.User{ID: reviewer, Name: reviewer, Role: "reviewer", Active: true}
	if !service.CanReview(u) {
		return model.Record{}, fmt.Errorf("reviewer denied")
	}
	return e.Svc.Review(id, reviewer, true)
}
func (e *Engine) Submit(id, actor string) (model.Record, error) {
	u := model.User{ID: actor, Name: actor, Role: "operator", Active: true}
	if !service.CanProcess(u) {
		return model.Record{}, fmt.Errorf("operator denied")
	}
	return e.Svc.Process(id, actor)
}
func (e *Engine) Track(id string) ([]model.Event, error) { return e.Svc.Store.ListEvents(id) }
func (e *Engine) Archive(id, actor string) error {
	r, err := e.Svc.Current(id)
	if err != nil {
		return err
	}
	return e.Svc.Store.SaveAudit(model.NewAudit(id+"-archive", r.ID, actor, "archive", e.Svc.Clock()))
}
