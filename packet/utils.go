package packet

func LuaHashRange(seed uint32, p []byte, start int, end int) uint32 {
	var size int = end - start
	var step int = (size >> 5) + 1
	var h uint32 = seed ^ uint32(size)
	for i := start; i < end; i += step {
		h = h ^ ((h << 5) + (h >> 2) + uint32(p[i]))
	}
	return h
}

func LuaHashString(seed uint32, s string) uint32 {
	var size int = len(s)
	var step int = (size >> 5) + 1
	var h uint32 = seed ^ uint32(size)
	for i := 0; i < size; i += step {
		h = h ^ ((h << 5) + (h >> 2) + uint32(s[i]))
	}
	return h
}

func LuaHashSlice(seed uint32, p []byte) uint32 {
	return LuaHashRange(seed, p, 0, len(p))
}

func LuaHashUInt8(seed uint32, x uint8) uint32 {
	var h uint32 = seed ^ 1
	h = h ^ ((h << 5) + (h >> 2) + uint32(x))
	return h
}

func LuaHashUInt16(seed uint32, x uint16) uint32 {
	var h uint32 = seed ^ 2
	h = h ^ ((h << 5) + (h >> 2) + uint32(x&0xFF))
	h = h ^ ((h << 5) + (h >> 2) + uint32((x>>8)&0xFF))
	return h
}

func LuaHashUInt32(seed uint32, x uint32) uint32 {
	var h uint32 = seed ^ 4
	h = h ^ ((h << 5) + (h >> 2) + ((x >> 0) & 0xFF))
	h = h ^ ((h << 5) + (h >> 2) + ((x >> 8) & 0xFF))
	h = h ^ ((h << 5) + (h >> 2) + ((x >> 16) & 0xFF))
	h = h ^ ((h << 5) + (h >> 2) + ((x >> 24) & 0xFF))
	return h
}

func LuaHashUInt64(seed uint32, x uint64) uint32 {
	var h uint32 = seed ^ 8
	h = h ^ ((h << 5) + (h >> 2) + uint32(x&0xFF))
	h = h ^ ((h << 5) + (h >> 2) + uint32((x>>8)&0xFF))
	h = h ^ ((h << 5) + (h >> 2) + uint32((x>>16)&0xFF))
	h = h ^ ((h << 5) + (h >> 2) + uint32((x>>24)&0xFF))
	h = h ^ ((h << 5) + (h >> 2) + uint32((x>>32)&0xFF))
	h = h ^ ((h << 5) + (h >> 2) + uint32((x>>40)&0xFF))
	h = h ^ ((h << 5) + (h >> 2) + uint32((x>>48)&0xFF))
	h = h ^ ((h << 5) + (h >> 2) + uint32((x>>56)&0xFF))
	return h
}
