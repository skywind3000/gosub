package packet

import "testing"

func TestSerialize(t *testing.T) {
	b := []byte{}
	b = append(b, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	b = append(b, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46)
	b = append(b, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66)
	s1 := "0000  09 08 07 06 05 04 03 02-01 41 42 43 44 45 46 61   .........ABCDEFa"
	s2 := "0000  62 63 64 65 66                                    bcdef"
	ss := s1 + "\n" + s2
	s := HexDump(b, true, 0)
	if s != ss {
		t.Fatalf("expect %s but got %s", ss, s)
	}
}
