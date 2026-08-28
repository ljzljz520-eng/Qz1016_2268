package query

import (
	"encoding/json"
	"facility-login/model"
)

func MarshalRecords(rs []model.Record) ([]byte, error) { return json.Marshal(rs) }
func GroupByStatus(rs []model.Record) map[string]int {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Status]++
	}
	return m
}
func HasUser(rs []model.Record, user string) bool {
	for _, r := range rs {
		if r.Username == user {
			return true
		}
	}
	return false
}
