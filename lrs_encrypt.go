package dlis

import "log"

// Logical Record Segment Encryption Packet $2.2.2.2 Figure 2-4.
type LRSEP struct {
	Size           int    // uint16, UNORM, must be even
	CompanyCode    int    // uint16 // $4.1.9
	EncriptionInfo *byte  // optional, so LRSEP can be 4 bytes
	bytes          []byte // keep the original bytes for checksum
}

func ParseEncryption(s *LRS) {
	log.Fatal("OH OH, not implemented!!!")
}
