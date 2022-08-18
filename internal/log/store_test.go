package log

/* running tests
Running tool: /usr/local/go/bin/go test -timeout 30s -run ^(TestStoreAppendRead|TestStoreClose)
$ github.com/upalchowdhury/dist-service/internal/log


https://ieftimov.com/posts/testing-in-go-go-test/#:~:text=To%20run%20your%20tests%20in,prints%20the%20complete%20test%20output.


*/

import (
	"io/ioutil"
	"os"
	"testing"

	//"github.com/stretchr/require"
	"github.com/stretchr/testify/require"
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

func testAppend(t *testing.T, s *Store) {
	t.Helper()

	for i := uint64(1); i < 4; i++ {
		n, pos, err := s.Append(write)
		require.NoError(t, err)
		require.Equal(t, pos+n, width*i)

	}
}

func testRead(t *testing.T, s *Store) {
	t.Helper()
	var pos uint64

	for i := uint64(1); i < 4; i++ {
		read, err := s.Read(pos)
		require.NoError(t, err)
		require.Equal(t, write, read)
		pos += width
	}
}

func testReadAt(t *testing.T, s *Store) {
	t.Helper()

	for i, off := uint64(1), int64(0); i < 4; i++ {
		b := make([]byte, lenwidth)

		n, err := s.ReadAt(b, off)

		require.NoError(t, err)

		require.Equal(t, lenwidth, n)
		off += int64(n)

		// size := enc.Uint64(b)
		// b = make([]byte, lenwidth)
		// nu, err := s.ReadAt(b, off)

		// require.NoError(t, err)

		// require.Equal(t, write, b)

		// require.Equal(t, int(size), nu)

		// off += int64(nu)

	}
}

func TestStoreClose(t *testing.T) {

	f, err := ioutil.TempFile("", "store_close_test")
	require.NoError(t, err)
	defer os.Remove(f.Name())
	s, err := newStore(f)

	require.NoError(t, err)

	_, _, err = s.Append(write)
	require.NoError(t, err)

	f, beforeSize, err := openFile(f.Name())
	require.NoError(t, err)

	err = s.Close()
	require.NoError(t, err)

	_, afterSize, err := openFile(f.Name())

	require.NoError(t, err)
	require.True(t, afterSize > beforeSize)
}

func openFile(name string) (file *os.File, size int64, err error) {
	f, err := os.OpenFile(
		name,
		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644,
	)

	if err != nil {
		return nil, 0, err
	}

	fi, err := f.Stat()

	if err != nil {
		return nil, 0, err
	}

	return f, fi.Size(), nil
}
