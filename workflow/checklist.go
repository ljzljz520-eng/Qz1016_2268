package workflow

import (
	"facility-login/model"
	"time"
)

type Checklist struct {
	Received, Validated, Saved, Displayed bool
	UpdatedAt                             time.Time
}

func NewChecklist(now time.Time) Checklist { return Checklist{UpdatedAt: now} }
func (c *Checklist) Receive(now time.Time) { c.Received = true; c.UpdatedAt = now }
func (c *Checklist) Validate(now time.Time) {
	if c.Received {
		c.Validated = true
		c.UpdatedAt = now
	}
}
func (c *Checklist) Save(now time.Time) {
	if c.Validated {
		c.Saved = true
		c.UpdatedAt = now
	}
}
func (c *Checklist) Display(now time.Time) {
	if c.Saved {
		c.Displayed = true
		c.UpdatedAt = now
	}
}
func (c Checklist) Complete() bool { return c.Received && c.Validated && c.Saved && c.Displayed }
func BuildReviewChecklist(r model.Record, now time.Time) Checklist {
	c := NewChecklist(now)
	c.Receive(now)
	if r.ID != "" {
		c.Validate(now)
	}
	return c
}
