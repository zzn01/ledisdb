package store

import (
	"fmt"
	"os"
	"path"

	"github.com/zzn01/ledisdb/config"
	"github.com/zzn01/ledisdb/store/driver"

	_ "github.com/zzn01/ledisdb/store/boltdb"
	_ "github.com/zzn01/ledisdb/store/goleveldb"
	_ "github.com/zzn01/ledisdb/store/leveldb"
	_ "github.com/zzn01/ledisdb/store/mdb"
	_ "github.com/zzn01/ledisdb/store/rocksdb"
)

func getStorePath(cfg *config.Config) string {
	if len(cfg.DBPath) > 0 {
		return cfg.DBPath
	} else {
		return path.Join(cfg.DataDir, fmt.Sprintf("%s_data", cfg.DBName))
	}
}

func Open(cfg *config.Config) (*DB, error) {
	s, err := driver.GetStore(cfg)
	if err != nil {
		return nil, err
	}

	path := getStorePath(cfg)

	if err := os.MkdirAll(path, 0755); err != nil {
		return nil, err
	}

	idb, err := s.Open(path, cfg)
	if err != nil {
		return nil, err
	}

	db := new(DB)
	db.db = idb
	db.name = s.String()
	db.st = &Stat{}
	db.cfg = cfg

	return db, nil
}

func Repair(cfg *config.Config) error {
	s, err := driver.GetStore(cfg)
	if err != nil {
		return err
	}

	path := getStorePath(cfg)

	return s.Repair(path, cfg)
}

func init() {
}
