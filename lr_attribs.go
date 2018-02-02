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
	abyte byte
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
	setf(a.NotFirst, "Not First", "Is First")
	setf(a.NotLast, "Not Last", "Is Last")
	setf(a.Encrypted, "Encrypted", "Not Encrypted")
	setf(a.HasEncryptPacket, "Has EncryptionPacket", "No EncryptionPacket")
	setf(a.HasChecksum, "Has Checksum", "No Checksum")
	setf(a.HasTrailingLen, "Has TrailingLen", "No TrailingLen")
	setf(a.HasPadding, "Has Padding", "No Padding")

	return fmt.Sprintf("[%b], %v\n", a.abyte, aa)
}

func (a *LRAttribs) Parse(b byte) {
	a.abyte = b
	a.Explicit = ((1 << 7) & b) != 0 // gimmick to check if bit 7 (most significant) is set
	a.NotFirst = ((1 << 6) & b) != 0 // using (1 << 6) to make bit number explicit
	a.NotLast = ((1 << 5) & b) != 0
	a.Encrypted = ((1 << 4) & b) != 0
	a.HasEncryptPacket = ((1 << 3) & b) != 0
	a.HasChecksum = ((1 << 2) & b) != 0
	a.HasTrailingLen = ((1 << 1) & b) != 0
	a.HasPadding = ((1 << 0) & b) != 0
}
