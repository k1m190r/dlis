package dlis

import (
	"log"
	"os"
	"testing"
)

var fname = "test/TestDataSet.dlis"

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
	t.Log(dlisread)

	// read VR records
	vr1 := dlisread.Read()
	for {
		lrs := vr1.Read()
		if lrs == nil {
			break
		}
		t.Log(lrs.Header.String())
	}

	vr2 := dlisread.Read()
	t.Log(vr2)

	// for {
	// 	vr := dlisread.Read()
	// 	if len(vr.Err) != 0 {
	// 		t.Log(vr)
	// 		break
	// 	}
	// 	t.Log(vr)
	// }
}
