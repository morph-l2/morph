package db

import (
	"fmt"
	"strconv"
	"sync"

	"morph-l2/tx-submitter/utils"

	"github.com/morph-l2/go-ethereum/ethdb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/errors"
)

var (
	ErrKeyNotFound = errors.ErrNotFound
)

type Db struct {
	db *leveldb.Database
	m  sync.Mutex
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
	d.m.Lock()
	defer d.m.Unlock()
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
	d.m.Lock()
	defer d.m.Unlock()
	valStr := strconv.FormatFloat(val, 'f', -1, 64)
	err := d.db.Put([]byte(key), []byte(valStr))
	if err != nil {
		return fmt.Errorf("failed to put key into leveldb %w", err)
	}
	return nil
}
func (d *Db) GetString(key string) (string, error) {
	d.m.Lock()
	defer d.m.Unlock()
	v, err := d.db.Get([]byte(key))
	if err != nil {
		if err == errors.ErrNotFound {
			return "", ErrKeyNotFound
		}
		return "", fmt.Errorf("failed to get key from leveldb %w", err)
	}
	return string(v), nil
}
func (d *Db) PutString(key, val string) error {
	d.m.Lock()
	defer d.m.Unlock()
	return d.db.Put([]byte(key), []byte(val))
}
func (d *Db) GetBytes(key []byte) ([]byte, error) {
	d.m.Lock()
	defer d.m.Unlock()
	v, err := d.db.Get(key)
	if err != nil {
		if err == errors.ErrNotFound {
			return nil, ErrKeyNotFound
		}
		return nil, fmt.Errorf("failed to get key from leveldb: %w", err)
	}
	return v, nil
}
func (d *Db) PutBytes(key, val []byte) error {
	d.m.Lock()
	defer d.m.Unlock()
	return d.db.Put(key, val)
}
func (d *Db) Delete(key []byte) error {
	d.m.Lock()
	defer d.m.Unlock()
	return d.db.Delete(key)
}
func (d *Db) Close() error {
	return d.db.Close()
}
