package localpool

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func getTestJournal() *Journal {

	jn := New("journal.rlp")
	err := jn.rm()
	if err != nil {
		panic(err)
	}
	err = jn.Init()
	if err != nil {
		panic(err)
	}
	return jn
}

func TestAddToFileStart(t *testing.T) {
	jn := getTestJournal()
	err := jn.AddToFileStart("hello")
	require.NoError(t, err)
	err = jn.AddToFileStart("hello")
	require.NoError(t, err)
	fl, err := jn.GetFirstLine()
	require.NoError(t, err)
	require.Equal(t, "hello", fl)
}
func TestAddToFileEnd(t *testing.T) {
	jn := getTestJournal()
	err := jn.AddToFileEnd("hello")
	require.NoError(t, err)
	err = jn.AddToFileEnd("hello")
	require.NoError(t, err)
	err = jn.AddToFileEnd("hello")
	require.NoError(t, err)
	err = jn.AddToFileEnd("hello")
	require.NoError(t, err)
	ll, err := jn.GetLastLine()
	require.NoError(t, err)
	require.Equal(t, "hello", ll)

}
