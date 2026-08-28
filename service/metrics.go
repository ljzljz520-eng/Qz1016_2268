package service

import (
	"facility-login/model"
	"sync"
	"time"
)

type Metrics struct {
	mu                                        sync.RWMutex
	Registered, Reviewed, Processed, Rejected int
	Last                                      time.Time
}

func (m *Metrics) RecordRegistration() {
	m.mu.Lock()
	m.Registered++
	m.Last = time.Now()
	m.mu.Unlock()
}
func (m *Metrics) RecordReview(ok bool) {
	m.mu.Lock()
	m.Reviewed++
	if !ok {
		m.Rejected++
	}
	m.Last = time.Now()
	m.mu.Unlock()
}
func (m *Metrics) RecordProcess()    { m.mu.Lock(); m.Processed++; m.Last = time.Now(); m.mu.Unlock() }
func (m *Metrics) Snapshot() Metrics { m.mu.RLock(); defer m.mu.RUnlock(); return *m }
func (m Metrics) Rate() float64 {
	if m.Registered == 0 {
		return 0
	}
	return float64(m.Processed) / float64(m.Registered)
}
func Describe(r model.Record) string {
	if r.Status == string(model.StatusCurrent) {
		return "当前有效"
	}
	if r.Status == string(model.StatusExpired) {
		return "已过期"
	}
	return "待处理"
}
