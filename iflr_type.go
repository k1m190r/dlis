package dlis

import (
	"fmt"
)

var iflr = []LRType{
	{"FDATA", "Frame Data", []string{"FRAME"}},              // 0
	{"NOFORMAT", "Unformatted Data", []string{"NO-FORMAT"}}, // 1
}

// Code inplied by the index of the array
// Code, Type, Description, AllowableSetTypes []

func IFLRType(code byte) LRType {
	if code == 127 {
		return LRType{"EOD", "End of Data", []string{}}
	}
	if code > 1 {
		return LRType{"RESERVED", "RESERVED", []string{"RESERVED"}}
	}
	return iflr[int(code)]
}

func ParseIFLR(s *LRS) {
	fmt.Println("MAKE THIS IFLR WORK!!!")
}
