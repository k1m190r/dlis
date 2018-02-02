package dlis

import "fmt"

// LRSH - Logical Record Segment Header $2.2.2.1 Figure 2-2.
// applies to all segments of LR and must be consistent for all
type LRSH struct {
	Length  int // uint16 UNORM, Length, must be even, minimum 16 bytes
	Attribs LRAttribs
	Type    byte // USHORT. App A.
	bytes   []byte
}

func (h *LRSH) String() string {
	var lrtype []interface{}
	if h.Attribs.Explicit {
		lrtype = EFLRType(h.Type)
	} else {
		lrtype = IFLRType(h.Type)
	}
	return fmt.Sprintf(
		"Header: Len: %d; Attribs: %s; Type: %v\n",
		h.Length, h.Attribs.String(), lrtype)
}
