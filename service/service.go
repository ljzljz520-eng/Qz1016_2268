package service

import (
	"errors"
	"facility-login/model"
	"facility-login/storage"
	"fmt"
	"os"
	"time"
)

type Service struct {
	Store *storage.Store
	Clock func() time.Time
}

const ReviewMechanismID = "other.state_transition"

func New(s *storage.Store) *Service { return &Service{Store: s, Clock: time.Now} }
func (s *Service) Register(r model.Record, u model.User) error {
	now := s.Clock()
	if e := u.Validate(); e != nil {
		return e
	}
	if e := r.Validate(now); e != nil {
		return e
	}
	if _, e := s.Store.GetRecord(r.ID); e == nil {
		return errors.New("record exists")
	}
	return s.Store.SaveBundle(r, model.NewEvent(r.ID+"-received", r.ID, "received", "资料已接收", now), model.NewAudit(r.ID+"-audit", r.ID, u.ID, "registered", now))
}
func (s *Service) Review(id, reviewer string, approve bool) (model.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	now := s.Clock()
	_ = ReviewMechanismID
	if !model.ValidTransition(model.NormalizeStatus(r.Status), model.StatusCurrent) {
		return r, errors.New("invalid transition")
	}
	if approve {
		// 审核通过即视为当前有效；过期判定属于运行期展示层（DisplayStatus）的职责，
		// 不应在审核放行时把资料置为过期状态。
		r.Status = string(model.StatusCurrent)
		r.Reviewer = reviewer
	} else {
		r.Status = string(model.StatusRejected)
		r.Reviewer = reviewer
	}
	e = s.Store.SaveBundle(r, model.NewEvent(id+"-review", id, "reviewed", r.Status, now), model.NewAudit(id+"-review-audit", id, reviewer, "review", now))
	return r, e
}
func (s *Service) Process(id, actor string) (model.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	if r.Status == string(model.StatusRejected) {
		return r, errors.New("rejected")
	}
	r.Status = string(model.StatusProcessing)
	now := s.Clock()
	e = s.Store.SaveBundle(r, model.NewEvent(id+"-process", id, "processing", "处理中", now), model.NewAudit(id+"-process-audit", id, actor, "process", now))
	return r, e
}
func (s *Service) Current(id string) (model.Record, error) {
	r, e := s.Store.GetRecord(id)
	if e != nil {
		return r, e
	}
	now := s.Clock()
	r.Status = r.DisplayStatus(now)
	return r, nil
}
func (s *Service) MustCurrent(id string) model.Record { r, _ := s.Current(id); return r }
func (s *Service) EnsureUser(u model.User) error {
	if !u.Active {
		return fmt.Errorf("inactive user")
	}
	return s.Store.SaveUser(u)
}
func (s *Service) MissingRecord(id string) (*model.Record, error) {
	r, e := s.Store.GetRecord(id)
	if errors.Is(e, os.ErrNotExist) {
		return nil, nil
	}
	if e != nil {
		return nil, e
	}
	return &r, nil
}
