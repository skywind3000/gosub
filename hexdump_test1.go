package main

import (
	"github.com/skywind3000/gosub/packet"
)

func main() {
	b := []byte{}
	b = append(b, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	b = append(b, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46)
	b = append(b, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66)
	s := packet.HexDump(b, true, 0)
	println(s)
}
