package service

import (
	"facility-login/model"
	"fmt"
	"time"
)

type Notification struct {
	RecordID, Channel, Message string
	At                         time.Time
}

func BuildNotification(r model.Record, channel string, now time.Time) Notification {
	return Notification{r.ID, channel, fmt.Sprintf("%s status %s", r.Username, r.Status), now}
}
func RenderNotification(n Notification) string { return n.Channel + ":" + n.Message }
func ShouldNotify(r model.Record) bool {
	return r.Status == string(model.StatusExpired) || r.Status == string(model.StatusProcessing)
}
