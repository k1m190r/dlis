package dlis

import (
	"encoding/binary"
	"fmt"
)

///////////////////////////////////////////////////////////////////////////////
// Logical Format $2.2

// LRS Logical Record Segment is interface between LF and Physical Format
// it applies to whole of LR not LRS, redundancy is intentional
type LRS struct {
	Header        LRSH
	EncryptPacket *LRSEP // optional

	body    []byte // LRS Body $2.2.2.3
	Trailer *LRST  // optional

	// Parse whatever it present here
	Set               Set      // must, pass it along with template to next
	RedundantSetCount int      // 0 means none
	ReplacementSets   []Set    // 0 means never
	Template          Template // must
	Objects           []Object // object types are restricted by the Set, at least 1

	Err []error
}

func (s *LRS) String() string {
	trailer := ""
	if s.Trailer != nil {
		trailer = s.Trailer.String()
	}
	return fmt.Sprintf(
		"Logical Record Segment\n%sBody Len:%d\n%s\nErr: %v\n",
		s.Header.String(), len(s.body), trailer, s.Err)
}

// NewLRS constructs new LRS from []byte slice
// b - is remainder of the VR body
func NewLRS(b []byte) (s *LRS) {
	s = new(LRS)

	// read 2 bytes of Length
	s.Header.Length = int(binary.BigEndian.Uint16(b[:2]))
	// Check this must be even and min 16

	// read next 2 bytes Attribs and Type
	s.Header.Attribs.Parse(b[2]) // 3rd
	s.Header.Type = b[3]         // 4th
	s.Header.bytes = b[:4]       // keep header raw bytes for later checksum

	// Read rest of the segment body
	s.body = b[4:s.Header.Length]

	s.parse()

	return
}

func (s *LRS) parse() {
	ats := s.Header.Attribs
	if ats.Encrypted || ats.HasEncryptPacket {
		ParseEncryption(s)
	}

	if ats.HasChecksum || ats.HasTrailingLen || ats.HasPadding {
		ParseLRSTrailer(s)
	}

	if ats.Explicit {
		ParseEFLR(s)

	} else {
		ParseIFLR(s)
	}
}
