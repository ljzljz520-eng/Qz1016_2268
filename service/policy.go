package service

import (
	"errors"
	"facility-login/model"
	"time"
)

func CheckReviewWindow(r model.Record, now time.Time) error {
	if r.ExpiresAt.Before(now.Add(-24 * time.Hour)) {
		return errors.New("review window closed")
	}
	return nil
}
func CanReview(u model.User) bool  { return u.Active && (u.Role == "reviewer" || u.Role == "admin") }
func CanProcess(u model.User) bool { return u.Active && (u.Role == "operator" || u.Role == "admin") }
func ResolveStatus(r model.Record, now time.Time) model.Status {
	if r.ExpiresAt.Before(now) {
		return model.StatusExpired
	}
	if r.Status == string(model.StatusProcessing) {
		return model.StatusProcessing
	}
	return model.StatusCurrent
}
