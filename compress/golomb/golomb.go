// See http://en.wikipedia.org/wiki/Golomb_coding
package golomb

import "github.com/quarnster/util/encoding/binary"
import "log"

type Golomb int

func (M Golomb) QR(value int) (q, r int) {
	return value / int(M), value % int(M)
}

func bitLen(val int) int {
	var ret int
	for val > (1 << uint32(ret)) {
		ret++
	}
	return ret
}

func (M Golomb) Read(br *binary.BitReader) (int, error) {
	q, r := 0, 0
	for {
		b, err := br.ReadBit()
		if err != nil {
			return 0, err
		}
		if b {
			q++
		} else {
			break
		}
	}
	bl := bitLen(int(M))
	cutoff := (2 << uint32(bl)) - int(M)

	for i := 0; i < bl-1; i++ {
		r <<= 1
		b, err := br.ReadBit()
		if err != nil {
			return 0, err
		}
		if b {
			r++
		}
	}
	if r >= cutoff {
		r <<= 1
		b, err := br.ReadBit()
		if err != nil {
			return 0, err
		}
		if b {
			r++
		}
	}
	log.Println(q, r)
	return q*int(M) + r, nil
}

/*
func (M Golomb) Encode(q, r int) int {
}
*/
