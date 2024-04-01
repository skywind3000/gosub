package packet

import (
	"encoding/binary"
	"strings"
)

func Encode8u(p []byte, c byte) []byte {
	p[0] = c
	return p[1:]
}

func Decode8u(p []byte, c *byte) []byte {
	*c = p[0]
	return p[1:]
}

func Encode16u(p []byte, w uint16) []byte {
	binary.LittleEndian.PutUint16(p, w)
	return p[2:]
}

func Decode16u(p []byte, w *uint16) []byte {
	*w = binary.LittleEndian.Uint16(p)
	return p[2:]
}

func Encode32u(p []byte, l uint32) []byte {
	binary.LittleEndian.PutUint32(p, l)
	return p[4:]
}

func Decode32u(p []byte, l *uint32) []byte {
	*l = binary.LittleEndian.Uint32(p)
	return p[4:]
}

func HexDump(p []byte, char_visible bool, limit int) string {
	hex := []rune("0123456789ABCDEF")
	if len(p) > limit && limit > 0 {
		p = p[:limit]
	}
	size := len(p)
	count := (size + 15) / 16
	var buffer []rune = make([]rune, 100)
	offset := 0
	output := []string{}
	for i := 0; i < count; i++ {
		length := min(16, len(p))
		line := buffer[:]
		for j := 0; j < len(line); j++ {
			line[j] = ' '
		}
		line[0] = hex[(offset>>12)&15]
		line[1] = hex[(offset>>8)&15]
		line[2] = hex[(offset>>4)&15]
		line[3] = hex[(offset>>0)&15]
		for j := 0; j < length; j++ {
			start := 6 + j*3
			line[start+0] = hex[(p[j]>>4)&15]
			line[start+1] = hex[(p[j]>>0)&15]
			if j == 8 {
				line[start-1] = '-'
			}
			if char_visible {
				c := '.'
				if p[j] >= 32 && p[j] < 127 {
					c = rune(p[j])
				}
				line[6+16*3+2+j] = c
			}
		}
		if len(p) >= 16 {
			p = p[16:]
		}
		s := string(line)
		s = strings.TrimRight(s, " ")
		output = append(output, s)
	}
	return strings.Join(output, "\n")
}
