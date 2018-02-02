package dlis

import (
	"encoding/binary"
	"errors"
	"fmt"
	"log"
)

// Logical Record Segment Trailer $2.2.2.4
type LRST struct {
	// LRST is meant to be read backwards from end to front
	// this is allow backwards traversal

	// padding is to achive minimum size of LRS and Even LRS length
	PadBytes []byte // null if not present, Pad Count and Pad Bytes
	PadCount *int   // USHORT, uint8, can be one indicating self
	CheckSum *int   // UNORM, uint16 optional, see LRSH.Attribs bit 6, App E
	Length   *int   // UNORM, uint16, optional, UNORM, Trailing Length,
	// Copy from LRS Header Len, every LRS or none, to allow traversal backwards
}

func ParseLRSTrailer(s *LRS) {
	// parse backward from Len -> CheckSum -> PadCount -> PadBytes

	var t LRST
	s.Trailer = &t
	ats := s.Header.Attribs

	if ats.HasTrailingLen {
		// last 2 bytes of the body
		st := len(s.body) - 2
		tlen := int(binary.BigEndian.Uint16(s.body[st:]))
		t.Length = &tlen

		// body now is but last 2 bytes
		s.body = s.body[:st]
	}

	if ats.HasChecksum {
		// last 2 bytes of the body
		st := len(s.body) - 2
		tcs := int(binary.BigEndian.Uint16(s.body[st:]))
		t.CheckSum = &tcs
		if cs, ok := checkSum(s); !ok { // checksum failed
			s.Err = append(s.Err, errors.New(fmt.Sprintf(
				"checksum failed. expected: %X found %X",
				t.CheckSum, cs)))
		}
		// body now is but last 2 bytes
		s.body = s.body[:st]
	}

	if ats.HasPadding {
		st := len(s.body) - 1 // last byte
		pc := int(s.body[st])
		t.PadCount = &pc

		// st-pc pad count 1 or more
		st = len(s.body) - pc
		t.PadBytes = s.body[st:]
		// body now is but last pc bytes
		s.body = s.body[:st]
	}
}

func checkSum(s *LRS) (int, bool) {
	log.Fatal("OH OH Need to check this!! checksum")
	// return the sum and if it is failed...
	return 0, true
}
