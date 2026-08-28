package storage

import (
	"facility-login/model"
	"sort"
	"time"
)

func SortRecords(rs []model.Record) []model.Record {
	sort.Slice(rs, func(i, j int) bool { return rs[i].CreatedAt.Before(rs[j].CreatedAt) })
	return rs
}
func FilterRoom(rs []model.Record, room string) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.Room == room {
			out = append(out, r)
		}
	}
	return out
}
func Expiring(rs []model.Record, now time.Time) []model.Record {
	out := []model.Record{}
	for _, r := range rs {
		if r.ExpiresAt.After(now) && r.ExpiresAt.Before(now.Add(24*time.Hour)) {
			out = append(out, r)
		}
	}
	return out
}
