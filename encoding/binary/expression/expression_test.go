// Copyright 2013 Fredrik Ehnbom
// Use of this source code is governed by a 2-clause
// BSD-style license that can be found in the LICENSE file.

package expression

import (
	"testing"
)

func TestParser(t *testing.T) {
	tests := [][]string{{"(MyMask & (Test >> 3)) << 0x2", `0-29: "EXPRESSION"
	0-29: "ShiftLeft"
		1-20: "Mask"
			1-7: "DotIdentifier"
				1-7: "Identifier" - Data: "MyMask"
			11-20: "ShiftRight"
				11-15: "DotIdentifier"
					11-15: "Identifier" - Data: "Test"
				19-20: "Constant" - Data: "3"
		26-29: "Constant" - Data: "0x2"
	29-29: "EndOfFile" - Data: ""
`},
		{"Length-1", `0-8: "EXPRESSION"
	0-8: "Sub"
		0-6: "DotIdentifier"
			0-6: "Identifier" - Data: "Length"
		7-8: "Constant" - Data: "1"
	8-8: "EndOfFile" - Data: ""
`},
		{"A.B.C", `0-5: "EXPRESSION"
	0-5: "DotIdentifier"
		0-1: "Identifier" - Data: "A"
		2-3: "Identifier" - Data: "B"
		4-5: "Identifier" - Data: "C"
	5-5: "EndOfFile" - Data: ""
`},
	}
	var p EXPRESSION

	for _, test := range tests {
		if !p.Parse(test[0]) {
			t.Errorf("Didn't parse correctly: %s\n", p.Error())
		} else {
			root := p.RootNode()
			if root.Range.End() != p.ParserData.Len() {
				t.Errorf("Parsing didn't finish: %v\n%s", root, p.Error())
			} else if root.String() != test[1] {
				t.Error("Output differs\n", root.String(), "\n", test[1])
			}
		}
	}
}
