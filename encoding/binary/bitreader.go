package binary

import (
	"fmt"
	"io"
)

type BitReader struct {
	Inner    io.Reader
	currByte uint8
	currPos  uint8
}

func (b *BitReader) ReadBit() (bool, error) {
	if b.currPos == 0 {
		b.currPos = 8
		buf := make([]byte, 1)
		n, err := b.Inner.Read(buf)
		if err != nil {
			return false, err
		} else if n != 1 {
			panic("n isn't 1 (shouldn't happen..")
		}
		b.currByte = buf[0]
	}
	b.currPos--
	r := (b.currByte & (1 << b.currPos)) != 0
	return r, nil
}

func (b *BitReader) ReadBits(count int) (int64, error) {
	if count > 64 || count < 0 {
		return 0, fmt.Errorf("count out of range: %d", count)
	}
	var ret int64
	for ; count > 0; count-- {
		if bi, err := b.ReadBit(); err != nil {
			return 0, err
		} else if bi {
			ret |= 1 << uint64(count-1)
		}
	}
	return ret, nil
}
