package storage

import (
	"encoding/json"
	"facility-login/model"
	"fmt"
	"go.etcd.io/bbolt"
	"os"
	"sync"
	"time"
)

var buckets = [][]byte{[]byte("records"), []byte("users"), []byte("events"), []byte("audits")}

type Store struct {
	db   *bbolt.DB
	mu   sync.RWMutex
	path string
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db, path: path}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.db == nil {
		return nil
	}
	err := s.db.Close()
	s.db = nil
	return err
}
func (s *Store) Path() string { return s.path }
func (s *Store) put(bucket, key string, v any) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket([]byte(bucket)).Put([]byte(key), data) })
}
func (s *Store) get(bucket, key string, out any) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket([]byte(bucket)).Get([]byte(key))
		if v == nil {
			return os.ErrNotExist
		}
		return json.Unmarshal(v, out)
	})
}
func (s *Store) SaveRecord(r model.Record) error { return s.put("records", r.ID, r) }
func (s *Store) GetRecord(id string) (model.Record, error) {
	var r model.Record
	e := s.get("records", id, &r)
	return r, e
}
func (s *Store) SaveUser(u model.User) error { return s.put("users", u.ID, u) }
func (s *Store) GetUser(id string) (model.User, error) {
	var u model.User
	e := s.get("users", id, &u)
	return u, e
}
func (s *Store) SaveEvent(e model.Event) error { return s.put("events", e.ID, e) }
func (s *Store) SaveAudit(a model.Audit) error { return s.put("audits", a.ID, a) }
func (s *Store) ListRecords() ([]model.Record, error) {
	out := []model.Record{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("records")).ForEach(func(_, v []byte) error {
			var r model.Record
			if e := json.Unmarshal(v, &r); e != nil {
				return e
			}
			out = append(out, r)
			return nil
		})
	})
	return out, e
}
func (s *Store) ListEvents(recordID string) ([]model.Event, error) {
	out := []model.Event{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("events")).ForEach(func(_, v []byte) error {
			var x model.Event
			if e := json.Unmarshal(v, &x); e != nil {
				return e
			}
			if recordID == "" || x.RecordID == recordID {
				out = append(out, x)
			}
			return nil
		})
	})
	return out, e
}
func (s *Store) ListAudits(recordID string) ([]model.Audit, error) {
	out := []model.Audit{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte("audits")).ForEach(func(_, v []byte) error {
			var x model.Audit
			if e := json.Unmarshal(v, &x); e != nil {
				return e
			}
			if recordID == "" || x.RecordID == recordID {
				out = append(out, x)
			}
			return nil
		})
	})
	return out, e
}
func (s *Store) Health() error {
	if s.db == nil {
		return fmt.Errorf("store closed")
	}
	return s.db.View(func(*bbolt.Tx) error { return nil })
}
func (s *Store) SaveBundle(r model.Record, e model.Event, a model.Audit) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		for _, item := range []struct {
			b, k string
			v    any
		}{{"records", r.ID, r}, {"events", e.ID, e}, {"audits", a.ID, a}} {
			d, er := json.Marshal(item.v)
			if er != nil {
				return er
			}
			if er = tx.Bucket([]byte(item.b)).Put([]byte(item.k), d); er != nil {
				return er
			}
		}
		return nil
	})
}
func (s *Store) Touch(id string, at time.Time) error {
	r, e := s.GetRecord(id)
	if e != nil {
		return e
	}
	r.Notes = at.UTC().Format(time.RFC3339)
	return s.SaveRecord(r)
}
