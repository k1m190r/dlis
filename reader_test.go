package dlis

import (
	"log"
	"os"
	"testing"
)

var fname = "test/TestDataSet.dlis"

//var fname = "test/n802b.dls"

// TestNewDLISReader tests dlis reader from
func TestNewDLISReader(t *testing.T) {

	// open file for read only
	f, err := os.Open(fname)
	if err != nil {
		log.Printf("error opening file %s : %v", fname, err)
		return
	}
	defer dclose(f)

	// dlis reader, read SUL
	dlisread := NewDLISReader(f)

	// read VR records
	// vr1 := dlisread.Read()
	// for {
	// 	lrs := vr1.Read()
	// 	if lrs == nil {
	// 		break
	// 	}
	// 	t.Log(lrs)
	// }
	// t.Log(vr1)

	// vr2 := dlisread.Read()
	// for {
	// 	lrs := vr2.Read()
	// 	if lrs == nil {
	// 		break
	// 	}
	// 	t.Log(lrs)
	// }

	// t.Log(vr2)

	for i := 0; i < 2; i++ {
		vr := dlisread.Read()
		if len(vr.Err) != 0 {
			t.Log(vr)
			break
		}

		for {
			lrs := vr.Read()
			if lrs == nil {
				break
			}
			t.Log(lrs)
		}

		t.Log(vr)
	}

	t.Log(dlisread)
}
