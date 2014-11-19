package binary

import "io"

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
