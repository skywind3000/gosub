// =====================================================================
//
// AtomHeader.go -
//
// Last Modified: 2024/04/07 15:33:57
//
// =====================================================================
package packet

import "fmt"

// AtomPacket Header: 16 bytes
type AtomHeader struct {
	Cmd        uint8  // 4 bits: push, sr_push, icmp
	Cls        uint8  // 3 bits
	Extension  uint8  // 1 bit
	TtlOrCount uint8  // ttl(PUSH) or address list count (SR_PUSH)
	Length     uint16 // payload size
	SrcRN      uint32 // source router number
	DstRN      uint32 // destination router number
	SrcAP      uint16 // source access point
	DstAP      uint16 // destination access point
}

func (self *AtomHeader) Marshal(p []byte) {
	p[0] = ((self.Cmd & 0xf) << 4) | ((self.Cls & 0x3) << 1)
	p[0] = p[0] | (self.Extension & 1)
	p[1] = self.TtlOrCount
	p[2] = byte(self.Length)
	p[3] = byte(self.Length >> 8)
	p[4] = byte(self.SrcRN)
	p[5] = byte(self.SrcRN >> 8)
	p[6] = byte(self.SrcRN >> 16)
	p[7] = byte(self.SrcRN >> 24)
	p[8] = byte(self.DstRN)
	p[9] = byte(self.DstRN >> 8)
	p[10] = byte(self.DstRN >> 16)
	p[11] = byte(self.DstRN >> 24)
	p[12] = byte(self.SrcAP)
	p[13] = byte(self.SrcAP >> 8)
	p[14] = byte(self.DstAP)
	p[15] = byte(self.DstAP >> 8)
}

func (self *AtomHeader) Unmarshal(p []byte) {
	self.Cmd = (p[0] >> 4) & 0xf
	self.Cls = (p[0] >> 1) & 0x3
	self.Extension = p[0] & 1
	self.TtlOrCount = p[1]
	self.Length = uint16(p[2]) | (uint16(p[3]) << 8)
	self.SrcRN = uint32(p[4]) | (uint32(p[5]) << 8) | (uint32(p[6]) << 16) | (uint32(p[7]) << 24)
	self.DstRN = uint32(p[8]) | (uint32(p[9]) << 8) | (uint32(p[10]) << 16) | (uint32(p[11]) << 24)
	self.SrcAP = uint16(p[12]) | (uint16(p[13]) << 8)
	self.DstAP = uint16(p[14]) | (uint16(p[15]) << 8)
}

func (self *AtomHeader) PushHeader(pd *PacketData) bool {
	if !pd.requireHead(16) {
		return false
	}
	pd.head -= 16
	self.Marshal(pd.data[pd.head : pd.head+16])
	return true
}

func (self *AtomHeader) PopHeader(pd *PacketData) bool {
	if pd.GetSize() < 16 {
		return false
	}
	self.Unmarshal(pd.data[pd.head : pd.head+16])
	pd.head += 16
	return true
}

func (self *AtomHeader) Size() int {
	return 16
}

func (self *AtomHeader) String() string {
	return fmt.Sprintf("AtomHeader(Cmd: %d, Cls: %d, Extension: %d, TtlOrCount: %d, Length: %d, SrcRN: %d, DstRN: %d, SrcAP: %d, DstAP: %d)",
		self.Cmd, self.Cls, self.Extension, self.TtlOrCount, self.Length, self.SrcRN, self.DstRN, self.SrcAP, self.DstAP)
}
