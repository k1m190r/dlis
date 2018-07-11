package dlis

import (
	"log"
	"os"
	"testing"
)

func TestSUL(t *testing.T) {

	var fname = "test/TestDataSet.dlis"

	// open file for read only
	f, err := os.Open(fname)
	if err != nil {
		log.Printf("error opening file %s : %v", fname, err)
		return
	}
	defer dclose(f)

	buf := make([]byte, 80)

	_, err = f.Read(buf)
	if err != nil {
		t.Log(err)
	}

	sulTest := NewB(buf)
	SULP := PFn[50]
	res := SULP(sulTest)
	t.Log(res)
}
