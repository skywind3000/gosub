// =====================================================================
//
// # GPacket.go - General Packet
//
// Last Modified: 2024/04/12 19:04:21
//
// =====================================================================
package packet

type GPacket struct {
	Atom   AtomHeader
	Frame  FrameHeader
	Access AccessHeader
	Data   PacketData
}

func NewGPacket(size int) *GPacket {
	pkt := &GPacket{
		Frame:  FrameHeader{},
		Atom:   AtomHeader{},
		Access: AccessHeader{},
	}
	pkt.Data.Init(size, 64)
	return pkt
}

func (self *GPacket) Release() {
	self.Data.Release()
}
