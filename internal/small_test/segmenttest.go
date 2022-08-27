package small_test

import (
	"fmt"
	"io/ioutil"

	api "github.com/upalchowdhury/dist-service/api/v1"
)

/*
// Proto file fields
type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value  []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Offset uint64 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

*/

func TestSegment() {
	dir, _ := ioutil.TempDir("", "segment_test")

	// defer os.Close(dir)

	want := &api.Record{Value: []byte("hello world")}

	//var baseOffset uint64
	s, err := newSegment(dir, 16)
	if err != nil {
		panic(err)
	}

	// require.NoError(t, err)

	// require.Equal(t, uint64(16), s.nextOffset, s.nextOffset)

	for i := uint64(0); i < 3; i++ {
		off, err := s.Append(want)
		if err != nil {
			panic(err)
		}
		fmt.Println(off)
		// require.NoError(t, err)
		// require.Equal(t, 16+i, off)

		got, err := s.Read(off)
		if err != nil {
			panic(err)
		}
		fmt.Println(got)
		// require.NoError(t, err)
		// require.Equal(t, want.Value, got.Value)

	}
}
