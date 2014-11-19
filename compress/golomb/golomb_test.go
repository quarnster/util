package golomb

import (
	"bytes"
	"github.com/quarnster/util/encoding/binary"
	"testing"
)

func TestBitLen(t *testing.T) {
	for i := 0; i < 32; i++ {
		if r := bitLen(int(1 << uint32(i))); r != i {
			t.Errorf("Test %d: %d != %d", i, r, i)
		}
	}
}

func TestDecode(t *testing.T) {
	g := Golomb(10)
	buf := bytes.NewBuffer([]byte{0xf2})
	br := &binary.BitReader{Inner: buf}
	v, err := g.Read(br)
	if err != nil {
		t.Errorf("%v", err)
	} else if v != 42 {
		t.Logf("Expected 42, got %d", v)
	}
}
