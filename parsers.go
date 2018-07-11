package dlis

import (
	"bytes"
	"errors"
	"strconv"
)

// P is list of parsers
var P *V

// PFn actual list of parsers
var PFn []FN

func init() {
	P = NewFn(make([]FN, 100))
	PFn = P.Fs()

	// 0: return 0
	PFn[0] = func(v *V) *V {
		return I(0)
	}

	// 50: 2.3.2 Storage Unit Label (SUL)
	PFn[50] = func(v *V) *V {
		// assume to receive at least 80 bytes in []byte
		inbuf := v.B()
		if inbuf == nil || len(inbuf) < 80 {
			return E(errors.New("expected []byte len(80) but see nil"))
		}
		// return [4]*V
		val := NewV(make([]*V, 5))
		vs := val.V() // value struct

		// Seq Number as int
		seqNumStr := string(bytes.TrimSpace(inbuf[0:4])) // len == 4
		seqNum, err := strconv.Atoi(seqNumStr)
		if err != nil {
			val.AddE(err)
		}
		vs[0] = I(seqNum)

		// DLIS Version as string
		ver := string(bytes.TrimSpace(inbuf[4:9])) // 5
		vs[1] = S(ver)

		// Structure as string - likely "RECORD"
		struc := string(bytes.TrimSpace(inbuf[9:15])) // 6
		vs[2] = S(struc)

		// Max Rec Length non negative int
		recLenStr := string(bytes.TrimSpace(inbuf[15:20])) // 5
		recLen, err := strconv.Atoi(recLenStr)
		if err != nil {
			val.AddE(err)
		}
		if recLen < 0 { // max rec len cannot be negative
			val.AddE(errors.New("SUL record length cannot be negative"))
		}
		vs[3] = I(recLen)

		// Storage Set ID
		storeSetID := string(bytes.TrimSpace(inbuf[20:80])) // 60
		vs[4] = S(storeSetID)

		return val
	}
}
