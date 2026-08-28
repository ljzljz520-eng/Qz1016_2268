package api

import (
	"encoding/json"
	"facility-login/model"
	"net/http"
)

type Envelope struct {
	OK    bool
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Envelope{OK: status < 400, Data: v})
}
func WriteRecord(w http.ResponseWriter, r model.Record)     { WriteJSON(w, http.StatusOK, r) }
func WriteRecords(w http.ResponseWriter, rs []model.Record) { WriteJSON(w, http.StatusOK, rs) }
func ParseStatus(v string) model.Status                     { return model.NormalizeStatus(v) }
