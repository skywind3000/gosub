package packet

import (
	"fmt"
)

// ---------------------------------------------------------------------
// PacketData
// ---------------------------------------------------------------------
type PacketData struct {
	head int
	tail int
	err  error
	data []byte
}

// ---------------------------------------------------------------------
// errors
// ---------------------------------------------------------------------
var (
	ErrRequireHeadSpace error = fmt.Errorf("require more head space")
	ErrRequireTailSpace error = fmt.Errorf("require more tail space")
	ErrRequireDataSize  error = fmt.Errorf("require more data size")
)

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// constructor
func NewPacketData(size int, overhead int) *PacketData {
	var capacity int = size + overhead
	self := new(PacketData)
	self.head = overhead
	self.tail = overhead
	self.err = nil
	self.data = make([]byte, capacity)
	return self
}

func (self *PacketData) Release() {
	self.head = 0
	self.tail = 0
	self.data = nil
}

func (self *PacketData) GetSize() int {
	return self.tail - self.head
}

func (self *PacketData) GetError() error {
	return self.err
}

func (self *PacketData) ClearError() {
	self.err = nil
}

func (self *PacketData) String() string {
	t := fmt.Sprintf("PacketData(%d/%d)", self.GetSize(), len(self.data))
	return t
}

func (self *PacketData) DumpHex(limit int) string {
	data := self.data[self.head:self.tail]
	str := HexDump(data, true, limit)
	return str
}

// For short term usage. Don't keep the reference to the data
func (self *PacketData) GetData() []byte {
	return self.data[self.head:self.tail]
}

func (self *PacketData) requireHead(size int) bool {
	if self.head < size {
		self.err = ErrRequireHeadSpace
		return false
	}
	return true
}

func (self *PacketData) requireTail(size int) bool {
	if self.tail+size > len(self.data) {
		self.err = ErrRequireTailSpace
		return false
	}
	return true
}

func (self *PacketData) requireData(size int) bool {
	if self.GetSize() < size {
		self.err = ErrRequireDataSize
		return false
	}
	return true
}

func (self *PacketData) PushHead(p []byte) bool {
	if !self.requireHead(len(p)) {
		return false
	}
	self.head -= len(p)
	copy(self.data[self.head:], p)
	return true
}

func (self *PacketData) PushTail(p []byte) bool {
	if !self.requireTail(len(p)) {
		return false
	}
	copy(self.data[self.tail:], p)
	self.tail += len(p)
	return true
}

func (self *PacketData) PopHead(p []byte) bool {
	if self.GetSize() < len(p) {
		return false
	}
	copy(p, self.data[self.head:])
	self.head += len(p)
	return true
}

func (self *PacketData) PopTail(p []byte) bool {
	if self.GetSize() < len(p) {
		return false
	}
	self.tail -= len(p)
	copy(p, self.data[self.tail:])
	return true
}

func (self *PacketData) PushHeadUInt8(b uint8) bool {
	if !self.requireHead(1) {
		return false
	}
	self.head--
	self.data[self.head] = b
	return true
}

func (self *PacketData) PushTailUInt8(b uint8) bool {
	if !self.requireTail(1) {
		return false
	}
	self.data[self.tail] = b
	self.tail++
	return true
}

func (self *PacketData) PopHeadUInt8() byte {
	if !self.requireData(1) {
		return 0
	}
	var b byte = self.data[self.head]
	self.head++
	return b
}

func (self *PacketData) PopTailUInt8() byte {
	if !self.requireData(1) {
		return 0
	}
	self.tail--
	var b byte = self.data[self.tail]
	return b
}

func (self *PacketData) PushHeadUInt16(w uint16) bool {
	if !self.requireHead(2) {
		return false
	}
	self.head -= 2
	self.data[self.head+1] = byte(w >> 8)
	self.data[self.head] = byte(w)
	return true
}

func (self *PacketData) PushTailUInt16(w uint16) bool {
	if !self.requireTail(2) {
		return false
	}
	self.data[self.tail] = byte(w)
	self.data[self.tail+1] = byte(w >> 8)
	self.tail += 2
	return true
}

func (self *PacketData) PopHeadUInt16() uint16 {
	if !self.requireData(2) {
		return 0
	}
	var w1 uint16 = uint16(self.data[self.head+0])
	var w2 uint16 = uint16(self.data[self.head+1])
	var w uint16 = w1 | (w2 << 8)
	self.head += 2
	return w
}

func (self *PacketData) PopTailUInt16() uint16 {
	if !self.requireData(2) {
		return 0
	}
	self.tail -= 2
	var w1 uint16 = uint16(self.data[self.tail+0])
	var w2 uint16 = uint16(self.data[self.tail+1])
	var w uint16 = w1 | (w2 << 8)
	return w
}

func (self *PacketData) PushHeadUInt32(l uint32) bool {
	if !self.requireHead(4) {
		return false
	}
	self.head -= 4
	self.data[self.head] = byte(l)
	self.data[self.head+1] = byte(l >> 8)
	self.data[self.head+2] = byte(l >> 16)
	self.data[self.head+3] = byte(l >> 24)
	return true
}

func (self *PacketData) PushTailUInt32(l uint32) bool {
	if !self.requireTail(4) {
		return false
	}
	self.data[self.tail] = byte(l)
	self.data[self.tail+1] = byte(l >> 8)
	self.data[self.tail+2] = byte(l >> 16)
	self.data[self.tail+3] = byte(l >> 24)
	self.tail += 4
	return true
}

func (self *PacketData) PopHeadUInt32() uint32 {
	if !self.requireData(4) {
		return 0
	}
	var l1 uint32 = uint32(self.data[self.head+0])
	var l2 uint32 = uint32(self.data[self.head+1])
	var l3 uint32 = uint32(self.data[self.head+2])
	var l4 uint32 = uint32(self.data[self.head+3])
	var l uint32 = l1 | (l2 << 8) | (l3 << 16) | (l4 << 24)
	self.head += 4
	return l
}

func (self *PacketData) PopTailUInt32() uint32 {
	if !self.requireData(4) {
		return 0
	}
	self.tail -= 4
	var l1 uint32 = uint32(self.data[self.tail+0])
	var l2 uint32 = uint32(self.data[self.tail+1])
	var l3 uint32 = uint32(self.data[self.tail+2])
	var l4 uint32 = uint32(self.data[self.tail+3])
	var l uint32 = l1 | (l2 << 8) | (l3 << 16) | (l4 << 24)
	return l
}
