/*
 * @Author: WorldDogs noreply.github.com
 * @Date: 2025-02-25 17:01:45
 * @LastEditors: WorldDogs noreply.github.com
 * @LastEditTime: 2025-03-12 14:18:00
 * @FilePath: /morph/tx-submitter/event/storage_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package event

import (
	"morph-l2/tx-submitter/db"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEventInfoStorage(t *testing.T) {

	db, err := db.New("./testleveldb")
	require.NoError(t, err)
	storage := NewEventInfoStorage(db)
	err = storage.Load()
	require.NoError(t, err)

	storage.SetBlockTime(100)
	storage.SetBlockProcessed(100)
	err = storage.Store()
	require.NoError(t, err)

	storage2 := NewEventInfoStorage(db)
	err = storage2.Load()
	require.NoError(t, err)
	require.Equal(t, storage.BlockTime(), storage2.BlockTime())
	require.Equal(t, storage.BlockProcessed(), storage2.BlockProcessed())
}
