package dlis

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

// SUL $2.3.2 = storage unit label first 80 bytes (0x50) of Visible Envelop
// Fig 2-7. Only one SUL per SU, before LF.
// whole of SUL is ASCII
type SUL struct {
	SeqNum      string // 4 sequence number
	DLISVersion string // 5 "V1.00" - most likely
	Struct      string // 6 storage unit structure, "RECORD" = Record Storage Unit
	MaxRecLen   int    // 5 max rec length, applies to Visible Records $2.3.6, $2.3.6.5 abs max is 16,384 (2^14)
	SetID       string // 60 storage set identifier
}

func trimSlice(buf []byte) string {
	return strings.TrimSpace(string(buf))
}

func (s *SUL) String() string {
	return fmt.Sprintf(
		"Storage Unit Label\nSeqNum: %s; DLISVersion: %s; Struct: %s; MaxRecLen: %d;\nSetID: %s\n",
		s.SeqNum, s.DLISVersion, s.Struct, s.MaxRecLen, s.SetID)
}

func (s *SUL) Read(f io.Reader) error {
	// SUL is exacly 80 bytes
	var buf = make([]byte, 80)
	n, err := f.Read(buf)
	if err != nil {
		return err
	}
	if n != 80 {
		return fmt.Errorf("expecting len(SUL)==80 bytes, but it is %d", n)
	}

	s.SeqNum = trimSlice(buf[0:4])      // 4
	s.DLISVersion = trimSlice(buf[4:9]) // 5
	s.Struct = trimSlice(buf[9:15])     // 6
	s.SetID = trimSlice(buf[20:80])     // rest of it

	s.MaxRecLen, err = strconv.Atoi(trimSlice(buf[15:20])) // 5
	if err != nil {
		return err
	}

	return nil
}
