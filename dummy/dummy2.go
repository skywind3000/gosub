package dummy2

import (
	"https://github.com/skywind3000/gosub/packet"
)

func DummyAdd2(x int, y int) int {
	packet := packet.AtomHeader{}
	return x + y
}


