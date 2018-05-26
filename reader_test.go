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

	for i := 0; i < 1; i++ {
		vr := dlisread.ReadVR()
		if len(vr.Err) != 0 {
			t.Log(vr)
			break
		}

		for {
			lrs := vr.ReadLRS()
			if lrs == nil {
				break
			}
			t.Log(lrs)
		}

		t.Log(vr)
	}

	t.Log(dlisread)
}
