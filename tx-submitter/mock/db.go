/*
 * @Author: WorldDogs noreply.github.com
 * @Date: 2025-03-12 13:00:15
 * @LastEditors: WorldDogs noreply.github.com
 * @LastEditTime: 2025-03-12 13:04:57
 * @FilePath: /morph/tx-submitter/mock/db.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package mock

import (
	"strconv"
	"sync"

	"morph-l2/tx-submitter/db"
)

type MockDB struct {
	store map[string]string
	m     sync.RWMutex
}

func NewMockDB() *MockDB {
	return &MockDB{
		store: make(map[string]string),
	}
}

func (d *MockDB) GetString(key string) (string, error) {
	d.m.RLock()
	defer d.m.RUnlock()
	if val, ok := d.store[key]; ok {
		return val, nil
	}
	return "", db.ErrKeyNotFound
}

func (d *MockDB) PutString(key, val string) error {
	d.m.Lock()
	defer d.m.Unlock()
	d.store[key] = val
	return nil
}

func (d *MockDB) GetFloat(key string) (float64, error) {
	d.m.RLock()
	defer d.m.RUnlock()
	if val, ok := d.store[key]; ok {
		return strconv.ParseFloat(val, 64)
	}
	return 0, db.ErrKeyNotFound
}

func (d *MockDB) PutFloat(key string, val float64) error {
	d.m.Lock()
	defer d.m.Unlock()
	d.store[key] = strconv.FormatFloat(val, 'f', -1, 64)
	return nil
}

func (d *MockDB) Close() error {
	return nil
}
