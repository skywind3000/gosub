package main

import (
	"fmt"
	"unsafe"

	"github.com/skywind3000/gosub/packet"
)

func test1() {
	b := []byte{}
	b = append(b, 9, 8, 7, 6, 5, 4, 3, 2, 1)
	b = append(b, 0x41, 0x42, 0x43, 0x44, 0x45, 0x46)
	b = append(b, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66)
	s := packet.HexDump(b, true, 0)
	println(s)
	t := &b[2]
	println(*t)
	println(uintptr(unsafe.Pointer(t)))
}

func test2() {
	s := "hello, world"
	fmt.Printf("%T\n", s[1])
	fmt.Printf("%v\n", s[1])
	fmt.Printf("%T\n", s[1:3])
}

func main() {
	test2()
}
