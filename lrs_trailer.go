package dlis

import (
	"encoding/binary"
	"fmt"
	"log"
)

// LRST Logical Record Segment Trailer $2.2.2.4
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

func (t *LRST) String() string {
	msg := []string{}
	if t.PadBytes != nil {
		msg = append(msg, fmt.Sprintf("PadBytes, last is Count: %v", t.PadBytes))
	}
	if t.CheckSum != nil {
		msg = append(msg, fmt.Sprintf("CheckSum: %d", *t.CheckSum))
	}
	if t.Length != nil {
		msg = append(msg, fmt.Sprintf("Traling Length: %d", *t.Length))
	}
	return fmt.Sprintf("Trailer:%v", msg)
}

// ParseLRSTrailer parses trailer backwards
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
			s.Err = append(s.Err, fmt.Errorf(
				"checksum failed. expected: %X found %X",
				t.CheckSum, cs))
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
	// http://w3.energistics.org/rp66/v1/rp66v1_appe.html

	log.Fatal("checkSum in lrs_trailer is not implemented.")
	// return the sum and if it is failed...

	// Sum includes everything in the LRS which precedes the checksum value.
	// The checksum value itself and the Trailing Length, if present, are not included.
	// need to merge the header, encryption packet and pad bytes

	// When encryption is used, the checksum is computed after the
	// Logical Record Segment Body and Pad Bytes have been encrypted.

	/* assume even number of bytes
	   1) c=0 initialize 16-bit checksum to zero
	   2) loop i=1,n,2 loop over the data two bytes at a time
	   3) t=byte(i+1)*256+byte(i) compute a 16-bit addend by concatenating the next two bytes of data
	   4) c=c+t add the addend to the checksum
	   5) if carry c=c+1 add carry to checksum
	   6) c=c*2 left shift checksum
	   7) if carry c=c+1 add carry to checksum
	   8) endloop
	*/

	return 0, true
}
