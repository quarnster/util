package binary_test

import (
	"bytes"
	"github.com/quarnster/util/encoding/binary"
	"io"
	"testing"
)

func TestBitReader(t *testing.T) {
	buf := bytes.NewReader([]byte{9, 129})
	exp := []bool{
		false, false, false, false, true, false, false, true,
		true, false, false, false, false, false, false, true,
	}

	br := binary.BitReader{Inner: buf}
	for i, e := range exp {
		b, err := br.ReadBit()
		if err != nil {
			t.Errorf("Test %d errored: %v", i, err)
		} else if b != e {
			t.Errorf("Test %d expected %v, but got %v", i, e, b)
		}
	}
	if _, err := br.ReadBit(); err != io.EOF {
		t.Errorf("Expected an EOF error but didn't get one")
	}

}

func TestBitReaderBits(t *testing.T) {
	buf := bytes.NewReader([]byte{9, 129, 128 + 32})
	exp := int64((9 << 11) | (129 << 3) | 5)

	br := binary.BitReader{Inner: buf}
	i, err := br.ReadBits(19)
	if err != nil {
		t.Error(err)
	} else if i != exp {
		t.Errorf("Expected %v, but got %v", exp, i)
	}
}
