package dlis

import (
	"bufio"
	"fmt"
	"io"
	"log"
)

// LR logical record, keeps all the LRS parsing results
// Sets, Templates, Objects
type LR []*LRS

// LF logical file is set of LR
// LF starts with File Header LR and ends on the next FHLR
type LF []LR

// Reader is dlis.Reader that does all the reading
type Reader struct {
	FileName string
	Label    SULT
	VRCount  int // VR records read
	Err      []error

	// reader for the underlying file
	r io.Reader

	// LogFiles set of all LF from DLIS file
	LogFiles []LF
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
	ret.r = bufio.NewReaderSize(r, ret.Label.MaxRecLen)

	return
}

// ReadAll reads the whole DLIS

// ReadVR reads next VR from the dlis
func (r *Reader) ReadVR() (vr *VR) {
	vr = NewVR(r.r)
	r.VRCount++
	return
}

func (r *Reader) String() string {
	return fmt.Sprintf("DLIS Reader\n%s\nVRCount:%d\nErr: %v\n",
		r.Label.String(), r.VRCount, r.Err)
}
