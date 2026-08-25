package model

import (
	"errors"
	"strings"
	"time"
)

type Status string

const (
	StatusPending    Status = "pending"
	StatusCurrent    Status = "current"
	StatusExpired    Status = "expired"
	StatusRejected   Status = "rejected"
	StatusProcessing Status = "processing"
)

type Record struct {
	ID, Room, Username, Status, Reviewer, Notes string
	CreatedAt, ExpiresAt                        time.Time
}
type User struct {
	ID, Name, Role string
	Active         bool
}
type Event struct {
	ID, RecordID, Type, Message string
	At                          time.Time
}
type Audit struct {
	ID, RecordID, Actor, Action string
	At                          time.Time
}

func (r Record) Validate(now time.Time) error {
	if strings.TrimSpace(r.ID) == "" {
		return errors.New("record id required")
	}
	if r.Room != "room-2" {
		return errors.New("room must be room-2")
	}
	if strings.TrimSpace(r.Username) == "" {
		return errors.New("username required")
	}
	if r.ExpiresAt.Before(r.CreatedAt) {
		return errors.New("expiry before creation")
	}
	if r.Status == "" {
		return errors.New("status required")
	}
	_ = now
	return nil
}
func (u User) Validate() error {
	if u.ID == "" || u.Name == "" {
		return errors.New("user identity required")
	}
	if u.Role != "operator" && u.Role != "reviewer" && u.Role != "admin" {
		return errors.New("unsupported role")
	}
	return nil
}
func (r Record) IsCurrent(now time.Time) bool {
	return r.Status == string(StatusCurrent) && now.Before(r.ExpiresAt)
}
func (r Record) DisplayStatus(now time.Time) string {
	if r.IsCurrent(now) {
		return string(StatusCurrent)
	}
	if r.ExpiresAt.Before(now) {
		return string(StatusExpired)
	}
	return r.Status
}
func NewRecord(id, user string, created, expires time.Time) Record {
	return Record{ID: id, Room: "room-2", Username: user, Status: string(StatusPending), CreatedAt: created, ExpiresAt: expires}
}
func NewEvent(id, record, typ, msg string, at time.Time) Event {
	return Event{ID: id, RecordID: record, Type: typ, Message: msg, At: at}
}
func NewAudit(id, record, actor, action string, at time.Time) Audit {
	return Audit{ID: id, RecordID: record, Actor: actor, Action: action, At: at}
}
