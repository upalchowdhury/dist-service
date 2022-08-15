package log

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/require"
)

var (
	write = []byte("hello world")
	width = uint64(len(write)) + lenwidth
)

func TestStoreAppendRead(t *testing.T) {
	f, err := ioutil.TempFile("", "store_append_read_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())

	s, err := newStore(f)
	require.NoError(t, err)

	testAppend(t, s)
	testRead(t, s)
	testReadAt(t, s)

	s, err = newStore(f)
	require.NoError(t, err)
	testRead(t, s)
}
