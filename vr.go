package dlis

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

// VR is Visible Record
type VR struct {
	Length        int // unit16 UNORM, len of whole VR struct, 20 is min
	FormatVersion int // unit16 $2.3.6.2 0xFF01, USHORT = 1 - always
	LRSCount      int // Count of LRS records read
	Err           []error
	body          []byte // this is reader for the body
	bodyOffset    int
}

func (vr *VR) String() string {
	return fmt.Sprintf(
		"Visible Record: Len: %d; FormatVer: %X; LRSCount: %d\nErrors: %v",
		vr.Length, vr.FormatVersion, vr.LRSCount, vr.Err)
}

// NewVR makes new Visual Record and returns it's address
// Stop if VR.Err != nil
func NewVR(r io.Reader) (vr *VR) {
	vr = new(VR)

	buff := make([]byte, 2)

	// read 2 bytes of Length
	n, err := io.ReadFull(r, buff) // r.Read(buff)
	if err != nil {
		vr.Err = append(vr.Err, err)
		return
	}
	if n != 2 { // should read 2 bytes
		vr.Err = append(vr.Err,
			errors.New("Visible Record error reading Lenth"))
		return
	}
	vr.Length = int(binary.BigEndian.Uint16(buff))

	// read next 2 bytes of FormatVersion
	n, err = r.Read(buff)
	if err != nil {
		vr.Err = append(vr.Err, err)
		return
	}
	if n != 2 { // should read 2 bytes
		vr.Err = append(vr.Err,
			errors.New("Visible Record error reading FormatVersion"))
		return
	}
	vr.FormatVersion = int(binary.BigEndian.Uint16(buff))

	// check that it is 0xFF01
	if vr.FormatVersion != 0xFF01 {
		vr.Err = append(vr.Err, fmt.Errorf(
			"Expected Visible Record FormatVersion to be 0xFF01 but it is %d",
			vr.FormatVersion))
	}

	// read the rest of the VR in to the temp buff
	restLen := vr.Length - 4 // -4 bytes for Lenght and FormatVersion
	vr.body = make([]byte, restLen)
	n, err = r.Read(vr.body)
	if err != nil {
		vr.Err = append(vr.Err, err)
		return
	}
	if n != int(restLen) { // should read restLen number of bytes
		vr.Err = append(vr.Err, fmt.Errorf(
			"visible record error reading the body of record. Expected %d, actual %d",
			restLen, n))
		return
	}

	return
}

// ReadLRS returns next LRS
func (vr *VR) ReadLRS() (l *LRS) {
	if vr.bodyOffset >= (vr.Length - 4) {
		// the VR is exausted return nil
		return nil
	}
	l = NewLRS(vr.body[vr.bodyOffset:])
	vr.bodyOffset += l.Header.Length
	vr.LRSCount++
	return
}
