package dlis

import (
	"fmt"
	"io"
	"log"
)

type Reader struct {
	FileName string
	Label    SUL
	VRCount  int // VR records read
	Err      []error
	r        io.Reader
}

// NewDLISReader reads the SUL and preps the rest of reading
func NewDLISReader(r io.Reader) (ret *Reader) {
	// get the label
	ret = new(Reader)

	// read Storage Unit Label
	err := ret.Label.Read(r)
	if err != nil {
		log.Printf("error reading Storage Label: %v", err)
		ret.Err = append(ret.Err)
		return
	}

	// buffered reader
	ret.r = r // bufio.NewReaderSize(r, ret.Label.MaxRecLen)

	return
}

// Read reads next VR from the dlis
func (r *Reader) Read() (vr *VR) {
	vr = NewVR(r.r)
	r.VRCount++
	return
}

func (r *Reader) String() string {
	return fmt.Sprintf("DLIS Reader\n%s\nVRCount:%d\nErr: %v\n",
		r.Label.String(), r.VRCount, r.Err)
}

func dclose(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Printf("error with deferred closing: %v", err)
	}
}
