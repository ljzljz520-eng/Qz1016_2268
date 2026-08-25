package query

import (
	"facility-login/model"
	"sort"
	"time"
)

type DailyReport struct {
	Date                             time.Time
	Total, Current, Expired, Pending int
}

func MakeReport(rs []model.Record, day time.Time) DailyReport {
	r := DailyReport{Date: day}
	for _, x := range rs {
		if x.CreatedAt.YearDay() != day.YearDay() {
			continue
		}
		r.Total++
		switch x.Status {
		case string(model.StatusCurrent):
			r.Current++
		case string(model.StatusExpired):
			r.Expired++
		default:
			r.Pending++
		}
	}
	return r
}
func TopUsers(rs []model.Record) []string {
	m := map[string]int{}
	for _, r := range rs {
		m[r.Username]++
	}
	type pair struct {
		n string
		c int
	}
	ps := []pair{}
	for n, c := range m {
		ps = append(ps, pair{n, c})
	}
	sort.Slice(ps, func(i, j int) bool { return ps[i].c > ps[j].c })
	out := []string{}
	for _, p := range ps {
		out = append(out, p.n)
	}
	return out
}
func InWindow(r model.Record, start, end time.Time) bool {
	return !r.CreatedAt.Before(start) && r.CreatedAt.Before(end)
}
