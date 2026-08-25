package model

import "encoding/json"

func EncodeRecord(r Record) ([]byte, error) { return json.Marshal(r) }
func DecodeRecord(b []byte) (Record, error) { var r Record; e := json.Unmarshal(b, &r); return r, e }
func EncodeEvent(e Event) ([]byte, error)   { return json.Marshal(e) }
func DecodeEvent(b []byte) (Event, error)   { var e Event; x := json.Unmarshal(b, &e); return e, x }
func EncodeAudit(a Audit) ([]byte, error)   { return json.Marshal(a) }
func DecodeAudit(b []byte) (Audit, error)   { var a Audit; e := json.Unmarshal(b, &a); return a, e }
