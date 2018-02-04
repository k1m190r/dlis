package dlis

import "fmt"

type LRAttribs struct { // byte Figure 2-3.
	Explicit, // or Indirect
	NotFirst, // Not First
	NotLast, // Not Last
	Encrypted,
	HasEncryptPacket,
	HasChecksum, // in LRS Trailer
	HasTrailingLen, // in LRS Trailer
	HasPadding bool // in LRS Trailer
	obyte byte // orignal byte
}

func (a *LRAttribs) String() string {
	aa := []string{}

	// c conditional, t,f = true false
	setf := func(c bool, t, f string) {
		if c {
			aa = append(aa, t)
		} else {
			aa = append(aa, f)
		}
	}

	setf(a.Explicit, "Explicit", "Implicit")
	setf(a.NotFirst, "Not First", "")
	setf(a.NotLast, "Not Last", "")
	setf(a.Encrypted, "Encrypted", "")
	setf(a.HasEncryptPacket, "Has EncryptionPacket", "")
	setf(a.HasChecksum, "Has Checksum", "")
	setf(a.HasTrailingLen, "Has TrailingLen", "")
	setf(a.HasPadding, "Has Padding", "")

	return fmt.Sprintf("[%b], %v", a.obyte, aa)
}

func (a *LRAttribs) Parse(b byte) {
	a.obyte = b
	a.Explicit = checkBit(b, 7)
	a.NotFirst = checkBit(b, 6)
	a.NotLast = checkBit(b, 5)
	a.Encrypted = checkBit(b, 4)
	a.HasEncryptPacket = checkBit(b, 3)
	a.HasChecksum = checkBit(b, 2)
	a.HasTrailingLen = checkBit(b, 1)
	a.HasPadding = checkBit(b, 0)
}

func checkBit(b byte, bit uint) bool {
	// gimmick to check if bit 7 (most significant) is set
	// using (1 << 6) to make bit number explicit
	return ((1 << bit) & b) != 0
}
