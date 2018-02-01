package dlis

import (
	"encoding/binary"
	"fmt"
)

///////////////////////////////////////////////////////////////////////////////
// Logical Format $2.2

// LRSH - Logical Record Segment Header $2.2.2.1 Figure 2-2.
// applies to all segments of LR and must be consistent for all
type LRSH struct {
	Length  int  // uint16 UNORM, Length, must be even, minimum 16 bytes
	Attribs byte // Figure 2-3.
	Type    byte // USHORT. App A.
}

func (h *LRSH) String() string {
	return fmt.Sprintf(
		"Header: Len: %d; Attribs: %b; Type: %d\n",
		h.Length, h.Attribs, h.Type)
}

// Logical Record Segment Encryption Packet $2.2.2.2 Figure 2-4.
type LRSEP struct {
	Size           uint16 // UNORM, must be even
	CompanyCode    uint16 // $4.1.9
	EncriptionInfo *byte  // optional, so LRSEP can be 4 bytes
}

// Logical Record Segment Trailer $2.2.2.4
type LRST struct {
	Padding  []byte
	CheckSum *uint16 // optional, see LRSH.Attribs bit 6, App E
	Length   uint16  // UNORM, Trailing Length
}

// Logical Record Segment is interface between LF and Physical Format
// it applies to whole of LR not LRS, redundancy is intentional
type LRS struct {
	Header        LRSH
	EncryptPacket *LRSEP // optional
	body          []byte // LRS Body $2.2.2.3
	Trailer       *LRST  // optional

	Err []error
}

func (s *LRS) String() string {
	return fmt.Sprintf(
		"Logical Record Segment\n%sBody Len:%d\nErr: %v\n",
		s.Header.String(), len(s.body), s.Err)
}

func NewLRS(b []byte) (s *LRS) {
	s = new(LRS)

	// read 2 bytes of Length
	s.Header.Length = int(binary.BigEndian.Uint16(b[:2]))

	// read next 2 bytes Attribs and Type
	s.Header.Attribs = b[2] // 3rd
	s.Header.Type = b[3]    // 4th

	// Read rest of the segment body
	s.body = b[4:s.Header.Length]

	// Check Attribs
	// Read Encryption Packet and Trailer
	// defintion of LRS body will change depending on
	// Attrib flags, Presence of Encryption Packet and Trailer

	return
}
