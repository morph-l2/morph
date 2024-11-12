package db

import (
	"fmt"
	"strconv"

	"morph-l2/tx-submitter/utils"

	"github.com/morph-l2/go-ethereum/ethdb/leveldb"
)

var (
	ErrKeyNotFound = fmt.Errorf("not found")
)

type Db struct {
	db *leveldb.Database
}

func New(pathname string) (*Db, error) {
	// leveldb
	ldb, err := leveldb.New(pathname, 0, 0, "tx-submitter", false)
	if err != nil {
		return nil, fmt.Errorf("failed to create leveldb: %w", err)
	}
	return &Db{db: ldb}, nil
}
func (d *Db) GetFloat(key string) (float64, error) {
	v, err := d.db.Get([]byte(key))
	if err != nil {
		return 0, fmt.Errorf("failed get key from leveldb %s: %w", key, err)
	}
	res, err := utils.ParseStringToType[float64](string(v))
	if err != nil {
		return 0, fmt.Errorf("failed to parse string to float64 %s", err)
	}
	return res, nil
}
func (d *Db) PutFloat(key string, val float64) error {
	valStr := strconv.FormatFloat(val, 'f', -1, 64)
	err := d.db.Put([]byte(key), []byte(valStr))
	if err != nil {
		return fmt.Errorf("failed to put key into leveldb %w", err)
	}
	return nil
}
func (d *Db) Close() error {
	return d.db.Close()
}
