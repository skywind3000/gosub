// =====================================================================
//
// AccessHeader.go -
//
// Last Modified: 2024/04/12 17:27:39
//
// =====================================================================
package packet

import "fmt"

type AccessHeader struct {
	Mask      uint16 // encryption key = mask + md5(secret_token)
	Signature uint16 // hash(data+cmd+token+ts) for authentication & integrity
	Timestamp uint32 // timestamp in seconds ^ 0xDEADBEEF
	Cmd       uint8  //
	Reserved  uint8  // reserved
	Length    uint16 // payload size
	SrcAP     uint16 // source access point
}

func (self *AccessHeader) Marshal(p []byte) {
	p[0] = byte(self.Mask)
	p[1] = byte(self.Mask >> 8)
	p[2] = byte(self.Signature)
	p[3] = byte(self.Signature >> 8)
	p[4] = byte(self.Timestamp)
	p[5] = byte(self.Timestamp >> 8)
	p[6] = byte(self.Timestamp >> 16)
	p[7] = byte(self.Timestamp >> 24)
	p[8] = self.Cmd
	p[9] = self.Reserved
	p[10] = byte(self.Length)
	p[11] = byte(self.Length >> 8)
	p[12] = byte(self.SrcAP)
	p[13] = byte(self.SrcAP >> 8)
}

func (self *AccessHeader) Unmarshal(p []byte) {
	self.Mask = uint16(p[0]) | (uint16(p[1]) << 8)
	self.Signature = uint16(p[2]) | (uint16(p[3]) << 8)
	self.Timestamp = uint32(p[4]) | (uint32(p[5]) << 8) | (uint32(p[6]) << 16) | (uint32(p[7]) << 24)
	self.Cmd = p[8]
	self.Reserved = p[9]
	self.Length = uint16(p[10]) | (uint16(p[11]) << 8)
	self.SrcAP = uint16(p[12]) | (uint16(p[13]) << 8)
}

func (self *AccessHeader) PushHeader(pd *PacketData) bool {
	if !pd.requireHead(14) {
		return false
	}
	pd.head -= 14
	self.Marshal(pd.data[pd.head:])
	return true
}

func (self *AccessHeader) PopHeader(pd *PacketData) bool {
	if !pd.requireHead(14) {
		return false
	}
	self.Unmarshal(pd.data[pd.head:])
	pd.head += 14
	return true
}

func (self *AccessHeader) Size() int {
	return 14
}

func (self *AccessHeader) String() string {
	return fmt.Sprintf("Mask:%d Signature:%d Timestamp:%d Cmd:%d Reserved:%d Length:%d SrcAP:%d",
		self.Mask, self.Signature, self.Timestamp, self.Cmd, self.Reserved, self.Length, self.SrcAP)
}
