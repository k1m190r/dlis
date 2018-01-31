package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var fname = "TestDataSet.dlis"

//var fname = "n802b_SHELL14.dls"

// DLIS Spec: http://w3.energistics.org/rp66/v1/Toc/main.html

///////////////////////////////////////////////////////////////////////////////
// Logical Format $2.2

// LRSH - Logical Record Segment Header $2.2.2.1 Figure 2-2.
// applies to all segments of LR and must be consistent for all
type LRSH struct {
	Length  uint16 // UNORM, Length, must be even, minimum 16 bytes
	Attribs byte   // Figure 2-3.
	Type    byte   // USHORT. App A.
}

// Logical Record Segment Encryption Packet $2.2.2.2 Figure 2-4.
type LRSEP struct {
	Size           uint16 // UNORM, must be even
	CompanyCode    uint16 // $4.1.9
	EncriptionInfo *byte  // optional, so LRSEP can be 4 bytes
}

// Logical Record Segment Body $2.2.2.3
type LRSB []byte

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
	Body          LRSB
	Trailer       *LRST // optional
}

// Logical Record $2.2.2 can be infinite
// aka Logical Record Body LRB $2.2.2.5
type LR []*LRS

// $2.2.3.1 FHLR is LR with just one segment
type FHLR [1]*LRS

// Logical File $2.2 $2.2.3 = seq of LR, first LR is File Header Logical Record FHLR $5.1 and App A
// LF is terminated when another FHLR is encountered or no more LR are available
type LF []*LR

///////////////////////////////////////////////////////////////////////////////
// Physical Format $2.3

// $2.3.1 Storage Unit (SU) (one tape or one file)
// One SU can contain several LF
// One LF can span multiple SU
// SU starts with SUL and has whole number of LRS
// Termination $2.3.5: run out of data.

// Storage Set ordered set of all SU that have LF that span across SU
// All SU in Storage Set have same Struct and SetID
// Storage set contains single LF.

// Physical Format is intersection of Logical Format, Visible and Invisible envelops
// Invisible does not matter, we dont see it
// Visible Envelop is visible to us on read

// Visible Record
type VR struct {
	Length        uint16 // UNORM, len of whole VR struct, 20 is min
	FormatVersion uint16 // $2.3.6.2 0xFF01, USHORT = 1 - always
	LF            []LRS  // bunch of LRS
}

// read from VR
func (f *VR) Read(b []byte) (n int, err error) {
	// read 8 bytes big-endian

	return 0, nil
}

// write to VR
func (f *VR) Write(b []byte) (n int, err error) {
	return 0, nil
}

// SUL $2.3.2 = storage unit label first 80 bytes (0x50) of Visible Envelop
// Fig 2-7. Only one SUL per SU, before LF.
// whole of SUL is ASCII
type SUL struct {
	SeqNum      string // [4]byte  // sequence number
	DLISVersion string // [5]byte  // "V1.00" - most likely
	Struct      string // [6]byte  // storage unit structure, "RECORD" = Record Storage Unit
	MaxRecLen   string // [5]byte  // maximum record length, applies to Visible Records $2.3.6, $2.3.6.5 abs max is 16,384 (2^14)
	SetID       string // [60]byte // storage set identifier
}

func trimSlice(buf []byte) string {
	return strings.TrimSpace(string(buf))
}

func (s *SUL) Read(f io.Reader) error {
	// SUL is exacly 80 bytes
	var buf = make([]byte, 80)
	n, err := f.Read(buf)
	if err != nil {
		return err
	}
	if n != 80 {
		return errors.New(fmt.Sprintf("expecting len(SUL)==80 bytes, but it is %d", n))
	}

	s.SeqNum = trimSlice(buf[0:4])      // 4
	s.DLISVersion = trimSlice(buf[4:9]) // 5
	s.Struct = trimSlice(buf[9:15])     // 6
	s.MaxRecLen = trimSlice(buf[15:20]) // 5
	s.SetID = trimSlice(buf[20:80])     // rest of it

	return nil
}

type VisibleEnvelop struct {
	Label SUL
	VR    []*VR
}

func main() {

	// open file for read only
	f, err := os.Open(fname)
	if err != nil {
		log.Printf("error opening file %s : %v", fname, err)
		return
	}
	defer dclose(f)

	// buf reader with default 4k buffer
	br := bufio.NewReader(f)

	var sul SUL
	err = sul.Read(br)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v", sul)

	// read
	// need to read an process one VR at a time
	// VR absolute max size is 16,384 (2^14) so this can be a starting point
	// VR actual size is in SUL use that as primary buffer size
	// Read VR header first 8 bytes Len/Format
	// Read rest of VR Len VR-8
	//   Interpret all LRS in VR
	// The reader itself should present a simple range capable interface...

}

func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("error with deferred closing: %v", err)
	}
}
