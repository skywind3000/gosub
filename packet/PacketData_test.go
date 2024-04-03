package packet

import "testing"

func TestPacketData(t *testing.T) {
	pd := NewPacketData(64, 32)
	if pd.GetSize() != 0 {
		t.Error("size error")
	}
	if pd.head != 32 || pd.tail != 32 {
		t.Error("head or tail error")
	}
	pd.PushHeadUInt8(1)
	pd.PushHeadUInt16(0x2233)
	pd.PushHeadUInt32(0x44556677)
	pd.PushTailUInt8(5)
	pd.PushTailUInt16(0x6677)
	pd.PushTailUInt32(0x8899aabb)
	if pd.GetSize() != 14 {
		t.Error("GetSize() error")
	}
	if pd.head != 32-7 {
		t.Error("head position error")
	}
	if pd.tail != 32+7 {
		t.Error("tail position error")
	}
	if pd.PopHeadUInt32() != 0x44556677 {
		t.Error("PopHeadUInt32() error")
	}
	if pd.PopHeadUInt16() != 0x2233 {
		t.Error("PopHeadUInt16() error")
	}
	if pd.PopHeadUInt8() != 1 {
		t.Error("PopHeadUInt8() error")
	}
	if pd.GetError() != nil {
		t.Error("GetError() error")
	}
	if pd.head != 32 {
		t.Error("head position error")
	}
	if pd.PopTailUInt32() != 0x8899aabb {
		t.Error("PopTailUInt32() error")
	}
	if pd.PopTailUInt16() != 0x6677 {
		t.Error("PopTailUInt16() error")
	}
	if pd.PopTailUInt8() != 5 {
		t.Error("PopTailUInt8() error")
	}
	if pd.tail != 32 {
		t.Error("tail position error")
	}
}
