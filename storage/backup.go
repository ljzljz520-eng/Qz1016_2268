package storage

import (
	"fmt"
	"go.etcd.io/bbolt"
	"os"
	"path/filepath"
)

func (s *Store) Backup(dst string) error {
	if s.db == nil {
		return fmt.Errorf("closed")
	}
	f, e := os.Create(filepath.Clean(dst))
	if e != nil {
		return e
	}
	defer f.Close()
	return s.db.View(func(tx *bbolt.Tx) error { return tx.Copy(f) })
}
func EnsureDir(path string) error { return os.MkdirAll(filepath.Dir(path), 0755) }
func FileExists(path string) bool { _, e := os.Stat(path); return e == nil }
