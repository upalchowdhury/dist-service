package small_test

import (
	"bufio"
	"encoding/binary"
	"os"
	"sync"
)

/* endianness (byte order). This binary library is used for encode/decode such as streaming data using
 io.Write or
 bufio.NewWriter


Text protocols, like JSON, use only the printable set of characters in ASCII or Unicode to communicate. For example, the number “26” is represented using the “2” and “6” bytes because those are printable characters. This is great for humans to read but slow for computers to read.
With a binary protocol, the number “26” can be represented using a single byte — 0x1A in hexadecimal. That’s a 50% reduction in space and it’s already in the computer’s native binary format so it doesn’t need to be parsed. This performance difference looks insignificant for a single number
but it adds up when processing millions or billions of numbers.


In binary encoding, you do have to make this choice with the order you write bytes. There are two kinds of endianness — big endian and little endian. Big endian is when you write your most significant byte first and little endian is when you write your least significant byte first.

For example, let’s take the decimal number 287,454,020 which is 0x11223344 in hexidecimal. The most significant byte is 0x11 and the least significant byte is 0x44.

Encoding this in big endian looks like:

11 22 33 44
and encoding in little endian looks like:

44 33 22 11

Big endian also has an interesting property that it is lexicographically sortable. That means that you can compare two binary-encoded numbers starting from the first byte and moving to the last byte. That’s how bytes.Equal() and bytes.Compare() work. This is because the most significant bytes come first in big endian encoding.








*/
var (
	enc = binary.BigEndian
)

const (
	lenwidth = 8
)

type Store struct {
	*os.File
	mu   sync.Mutex
	buf  *bufio.Writer
	size uint64
}

func newStore(f *os.File) (*Store, error) {

	fi, err := os.Stat(f.Name())

	if err != nil {
		return nil, err
	}

	size := uint64(fi.Size())

	return &Store{
		File: f,
		size: size,
		buf:  bufio.NewWriter(f),
	}, nil
}

func (s *Store) Append(p []byte) (n uint64, pos uint64, err error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	pos = s.size
	if err := binary.Write(s.buf, enc, uint64(len(p))); err != nil {
		return 0, 0, err
	}

	w, err := s.buf.Write(p)
	if err != nil {
		return 0, 0, err
	}
	w += lenwidth
	s.size += uint64(w)

	return uint64(w), pos, nil
}

func (s *Store) Read(pos uint64) ([]byte, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
		return nil, err
	}

	size := make([]byte, lenwidth)
	// returns length of bytes in the file starting from byte offset off
	if _, err := s.File.ReadAt(size, int64(pos)); err != nil {
		return nil, err
	}
	b := make([]byte, enc.Uint64(size))

	if _, err := s.File.ReadAt(b, int64(pos+lenwidth)); err != nil {
		return nil, err
	}

	return b, nil
}

func (s *Store) ReadAt(p []byte, off int64) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if err := s.buf.Flush(); err != nil {
		return 0, err
	}

	return s.File.ReadAt(p, off)
}

func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.buf.Flush()

	if err != nil {
		return err
	}

	return s.File.Close()
}
