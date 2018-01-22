package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

// var fname = "TestDataSet.dlis"
var fname = "n802b_SHELL14.dls"

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
	LF            LF     // bunch of LRS
}

// SUL $2.3.2 = storage unit label first 80 bytes (0x50) of Visible Envelop
// Fig 2-7. Only one SUL per SU, before LF.
type SUL struct {
	SeqNum      [4]byte  // sequence number
	DLISVersion [5]byte  // "V1.00" - most likely
	Struct      [6]byte  // storage unit structure, "RECORD" = Record Storage Unit
	MaxRecLen   [5]byte  // maximum record length, applies to Visible Records $2.3.6, $2.3.6.5 abs max is 16,384 (2^14)
	SetID       [60]byte // storage set identifier
}

func (s SUL) String() string {
	return fmt.Sprintf("Sequence Number: %s; DLISVersion: %s; Structure: %s; MaxiRecLen: %s, SetID: %s",
		string(s.SeqNum[:]), string(s.DLISVersion[:]),
		string(s.Struct[:]), string(s.MaxRecLen[:]),
		string(s.SetID[:]))
}

type VisibleEnvelop struct {
	SUL SUL
}

func readManual() (env VisibleEnvelop) {
	f, err := os.Open(fname)
	if err != nil {
		log.Printf("error opening file %s : %v", fname, err)
		return
	}
	defer dclose(f)

	_, err = f.Read(env.SUL.SeqNum[:])
	_, err = f.Read(env.SUL.DLISVersion[:])
	_, err = f.Read(env.SUL.Struct[:])
	_, err = f.Read(env.SUL.MaxRecLen[:])
	_, err = f.Read(env.SUL.SetID[:])

	return
}
func readReflect() (env VisibleEnvelop) {
	f, err := os.Open(fname)
	if err != nil {
		log.Printf("error opening file %s : %v", fname, err)
		return
	}
	defer dclose(f)

	err = binary.Read(f, binary.BigEndian, &env)

	return
}

func main() {

}

func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("error with deferred closing: %v", err)
	}
}
