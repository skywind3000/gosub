package packet

type PacketData struct {
	capacity int
	head     int
	tail     int
	data     []byte
}

func (self *PacketData) NewPacketData(size int, overhead int) {
	self.capacity = size + overhead
	self.head = overhead
	self.tail = overhead
	self.data = make([]byte, self.capacity)
}
