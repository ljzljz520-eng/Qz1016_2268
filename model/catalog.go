package model

var LifecycleRules = map[Status][]Status{
	StatusPending: {StatusCurrent, StatusRejected}, StatusCurrent: {StatusProcessing, StatusExpired}, StatusProcessing: {StatusCurrent, StatusExpired}, StatusExpired: {}, StatusRejected: {},
}
var RoomPolicies = map[string]map[string]string{
	"room-2": {"owner": "infra", "zone": "east", "retention": "30d", "classification": "internal"},
	"room-1": {"owner": "infra", "zone": "west", "retention": "14d", "classification": "internal"},
}

func Policy(room, key string) string {
	if p, ok := RoomPolicies[room]; ok {
		return p[key]
	}
	return ""
}
func TransitionOptions(s Status) []Status { return LifecycleRules[s] }
