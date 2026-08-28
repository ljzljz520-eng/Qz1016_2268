package model

import "time"

func NormalizeStatus(s string) Status {
	switch s {
	case "pending", "current", "expired", "rejected", "processing":
		return Status(s)
	default:
		return StatusPending
	}
}
func StatusForExpiry(expires, now time.Time) Status {
	if expires.Before(now) {
		return StatusExpired
	}
	return StatusCurrent
}
func ValidTransition(from, to Status) bool {
	if from == to {
		return true
	}
	switch from {
	case StatusPending:
		return to == StatusCurrent || to == StatusRejected
	case StatusCurrent:
		return to == StatusProcessing || to == StatusExpired
	case StatusProcessing:
		return to == StatusCurrent || to == StatusExpired
	case StatusExpired, StatusRejected:
		return false
	}
	return false
}
