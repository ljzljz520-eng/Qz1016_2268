package api

import (
	"net/http"
	"strings"
)

func PathID(r *http.Request) string {
	p := strings.Trim(r.URL.Path, "/")
	parts := strings.Split(p, "/")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}
func MethodAllowed(r *http.Request, methods ...string) bool {
	for _, m := range methods {
		if r.Method == m {
			return true
		}
	}
	return false
}
func WriteError(w http.ResponseWriter, status int, msg string) {
	http.Error(w, `{"error":"`+msg+`"}`, status)
}
