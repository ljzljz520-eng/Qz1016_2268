package api

import (
	"encoding/json"
	"facility-login/query"
	"facility-login/workflow"
	"net/http"
	"strings"
	"time"
)

type Handler struct{ Engine *workflow.Engine }

func New(e *workflow.Engine) *Handler { return &Handler{Engine: e} }
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) >= 2 && parts[0] == "records" {
		h.record(w, r, parts[1])
		return
	}
	http.Error(w, `{"error":"not found"}`, 404)
}
func (h *Handler) record(w http.ResponseWriter, r *http.Request, id string) {
	switch r.Method {
	case http.MethodPost:
		var in struct {
			Username string `json:"username"`
		}
		if json.NewDecoder(r.Body).Decode(&in) != nil {
			http.Error(w, `{"error":"bad json"}`, 400)
			return
		}
		if h.Engine.Register(id, in.Username, time.Now()) != nil {
			http.Error(w, `{"error":"register failed"}`, 400)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"id": id})
	case http.MethodGet:
		res, e := query.Timeline(h.Engine.Svc.Store, id)
		if e != nil {
			http.Error(w, `{"error":"missing"}`, 404)
			return
		}
		json.NewEncoder(w).Encode(res)
	default:
		http.Error(w, `{"error":"method"}`, 405)
	}
}
func (h *Handler) Routes() *http.ServeMux { m := http.NewServeMux(); m.Handle("/", h); return m }
