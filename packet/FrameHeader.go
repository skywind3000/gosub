// =====================================================================
//
// FrameHeader.go -
//
// Last Modified: 2024/04/12 17:18:54
//
// =====================================================================
package packet

import "fmt"

// router to router packet: 20 bytes
// a frame packet may contain multiple atom packets
type FrameHeader struct {
	Mask      uint16 // encryption key = mask + md5(secret_token)
	Signature uint16 // hash(data+cmd+token+ts) for authentication & integrity
	Timestamp uint32 // timestamp in seconds ^ 0xDEADBEEF
	Cmd       uint8  // 4 bits
	Flag      uint8  // 3 bits
	Extension uint8  // 1 bit
	Parity    uint8  // fec index
	Length    uint16 // payload size
	SrcRN     uint32 // source router number
	Seq       uint32 // sequence number
}

func (self *FrameHeader) Marshal(p []byte) {
	p[0] = byte(self.Mask)
	p[1] = byte(self.Mask >> 8)
	p[2] = byte(self.Signature)
	p[3] = byte(self.Signature >> 8)
	p[4] = byte(self.Timestamp)
	p[5] = byte(self.Timestamp >> 8)
	p[6] = byte(self.Timestamp >> 16)
	p[7] = byte(self.Timestamp >> 24)
	p[8] = ((self.Cmd & 0xf) << 4) | ((self.Flag & 0x7) << 1) | (self.Extension & 1)
	p[9] = self.Parity
	p[10] = byte(self.Length)
	p[11] = byte(self.Length >> 8)
	p[12] = byte(self.SrcRN)
	p[13] = byte(self.SrcRN >> 8)
	p[14] = byte(self.SrcRN >> 16)
	p[15] = byte(self.SrcRN >> 24)
	p[16] = byte(self.Seq)
	p[17] = byte(self.Seq >> 8)
	p[18] = byte(self.Seq >> 16)
	p[19] = byte(self.Seq >> 24)
}

func (self *FrameHeader) Unmarshal(p []byte) {
	self.Mask = uint16(p[0]) | (uint16(p[1]) << 8)
	self.Signature = uint16(p[2]) | (uint16(p[3]) << 8)
	self.Timestamp = uint32(p[4]) | (uint32(p[5]) << 8) | (uint32(p[6]) << 16) | (uint32(p[7]) << 24)
	self.Cmd = (p[8] >> 4) & 0xf
	self.Flag = (p[8] >> 1) & 0x7
	self.Extension = p[8] & 1
	self.Parity = p[9]
	self.Length = uint16(p[10]) | (uint16(p[11]) << 8)
	self.SrcRN = uint32(p[12]) | (uint32(p[13]) << 8) | (uint32(p[14]) << 16) | (uint32(p[15]) << 24)
	self.Seq = uint32(p[16]) | (uint32(p[17]) << 8) | (uint32(p[18]) << 16) | (uint32(p[19]) << 24)
}

func (self *FrameHeader) PushHeader(pd *PacketData) bool {
	if !pd.requireHead(20) {
		return false
	}
	pd.head -= 20
	self.Marshal(pd.data[pd.head : pd.head+20])
	return true
}

func (self *FrameHeader) PopHeader(pd *PacketData) bool {
	if pd.GetSize() < 20 {
		return false
	}
	self.Unmarshal(pd.data[pd.head : pd.head+20])
	pd.head += 20
	return true
}

func (self *FrameHeader) Size() int {
	return 20
}

func (self *FrameHeader) String() string {
	return fmt.Sprintf("FrameHeader(Mask: %x, Signature: %x, Timestamp: %d, Cmd: %d, Flag: %d, Extension: %d, Parity: %d, Length: %d, SrcRN: %d, Seq: %d)",
		self.Mask, self.Signature, self.Timestamp, self.Cmd, self.Flag, self.Extension, self.Parity, self.Length, self.SrcRN, self.Seq)
}
