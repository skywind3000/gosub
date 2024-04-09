// =====================================================================
//
// AtomPacket.go -
//
// Created by skywind on 2024/04/07
// Last Modified: 2024/04/07 15:20:14
//
// =====================================================================
package packet

type AtomPacket struct {
	Header AtomHeader
	Data   *PacketData
}

func NewAtomPacket(size int, overhead int) *AtomPacket {
	self := new(AtomPacket)
	self.Header = AtomHeader{}
	self.Data = NewPacketData(size, overhead)
	return self
}

func (self *AtomPacket) Release() {
	self.Data.Release()
	self.Data = nil
}

func (self *AtomPacket) GetSize() int {
	return self.Data.GetSize()
}

func (self *AtomPacket) PopHeader() bool {
	return self.Header.PopHeader(self.Data)
}

func (self *AtomPacket) PushHeader() bool {
	return self.Header.PushHeader(self.Data)
}
