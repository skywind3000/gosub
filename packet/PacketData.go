package packet

type PacketData struct {
	head int
	tail int
	data []byte
}

func (self *PacketData) NewPacketData(size int, overhead int) {
	var capacity int = size + overhead
	self.head = overhead
	self.tail = overhead
	self.data = make([]byte, capacity)
}

func (self *PacketData) requireHead(size int) bool {
	if self.head < size {
		return false
	}
	return true
}

func (self *PacketData) requireTail(size int) bool {
	if self.tail+size > len(self.data) {
		return false
	}
	return true
}

func (self *PacketData) GetSize() int {
	return self.tail - self.head
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
