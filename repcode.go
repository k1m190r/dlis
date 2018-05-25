package dlis

import (
	"encoding/binary"
	"fmt"
	"math"
)

// http://w3.energistics.org/rp66/v1/rp66v1_appb.html

// Val is "universal value"
type Val struct {
	// payload with a value
	s *string
	i *int
	f *float64
	v *Val

	c int // count
	e error
}

// Funcs

// FSINGL RepCode 2
func FSINGL(in []byte) *Val {
	if len(in) < 4 {
		err := fmt.Errorf("length of input slice %v < 4", len(in))
		return &Val{e: err}
	}
	f := float64(math.Float32frombits(binary.BigEndian.Uint32(in[:4])))
	return &Val{f: &f, c: 4}
}

// FDOUBL RepCode 7
func FDOUBL(in []byte) *Val {
	if len(in) < 8 {
		err := fmt.Errorf("length of input slice %v < 4", len(in))
		return &Val{e: err}
	}
	f := math.Float64frombits(binary.BigEndian.Uint64(in[:8]))
	return &Val{f: &f, c: 8}
}

// USHORT RepCode 15
func USHORT(in []byte) *Val {
	if len(in) < 1 {
		err := fmt.Errorf("length of input slice %v < 1", len(in))
		return &Val{e: err}
	}
	i := int(in[0])
	return &Val{i: &i, c: 1}
}

// UNORM RepCode 16
func UNORM(in []byte) *Val {
	i := int(binary.BigEndian.Uint16(in[:2]))
	return &Val{i: &i, c: 2}
}

// ULONG RepCode 17
func ULONG(in []byte) *Val {
	i := int(binary.BigEndian.Uint32(in[:4]))
	return &Val{i: &i, c: 4}
}

// UVARI RepCode 18
func UVARI(in []byte) *Val {
	b1 := in[0]
	if checkBit(b1, 7) { //
		if checkBit(b1, 6) { // 4 bytes
			tmp := [4]byte{b1 & 0x3F} // first byte with mask 0011_1111
			copy(tmp[1:], in[1:4])    // remaining 3 bytes
			i := int(binary.BigEndian.Uint32(tmp[:]))
			return &Val{i: &i, c: 4}
		}
		// 2 bytes
		tmp := [2]byte{b1 & 0x3F} // first byte with mask 0011_1111
		tmp[1] = in[1]
		i := int(binary.BigEndian.Uint16(tmp[:]))
		return &Val{i: &i, c: 2}

	}
	// single byte
	i := int(b1)
	return &Val{i: &i, c: 1} // bit 7 is 0
}

// IDENT RepCode 19
func IDENT(in []byte) *Val {
	v := USHORT(in)
	ln := *v.i

	if ln == 0 {
		s := ""
		return &Val{s: &s}
	}

	// only allowed 33-96, 123-126
	// TODO check for allowed
	s := string(in[1:(1 + ln)])
	return &Val{s: &s, c: (1 + ln)}
}

// ASCII RepCode 20
func ASCII(in []byte) *Val {
	v := UVARI(in)
	idlen := v.c
	asciilen := *v.i

	if idlen == 0 {
		s := ""
		return &Val{s: &s, c: (idlen + asciilen)}
	}

	s := string(in[idlen : idlen+asciilen])
	return &Val{s: &s, c: (idlen + asciilen)}
}

// ORIGIN RepCode 22 equivalent to UVARY
var ORIGIN = UVARI

// OBNAME RepCode 23
// ORIGIN, USHORT, IDENT
func OBNAME(in []byte) *Val {
	// ORIGIN
	v := ORIGIN(in)
	olen := v.c
	origin := *v.i

	// COPY
	vc := USHORT(in[olen:])
	v.v = vc // chain vc into v
	clen := vc.c
	copy := *vc.i

	// IDENT
	vi := IDENT(in[(olen + clen):])
	vc.v = vi // chain vi into vc
	ilen := vi.c
	ident := *vi.s

	return v
}

// RepCode holds all the information about REPCODE, most importantly it has
// Read function that reads actual repcode
var RepCode = []struct {
	// Code is index
	Name        string
	Size        int // # of bytes
	Descirption string

	Read func([]byte) *Val
}{
	{}, // 0 is not present
	{"FSHORT", 2, "Low precision floating point", nil}, // 1

	{"FSINGL", 4, "IEEE single precision floating point", FSINGL}, // 2

	{"FSING1", 8, "Validated single precision floating point", nil},          // 3
	{"FSING2", 12, "Two-way validated single precision floating point", nil}, // 4
	{"ISINGL", 4, "IBM single precision floating point", nil},                // 5
	{"VSINGL", 4, "VAX single precision floating point", nil},                // 6

	{"FDOUBL", 8, "IEEE double precision floating point", FDOUBL}, // 7

	{"FDOUB1", 16, "Validated double precision floating point", nil},         // 8
	{"FDOUB2", 24, "Two-way validated double precision floating point", nil}, // 9
	{"CSINGL", 8, "Single precision complex", nil},                           // 10
	{"CDOUBL", 16, "Double precision complex", nil},                          // 11
	{"SSHORT", 1, "Short signed integer", nil},                               // 12
	{"SNORM", 2, "Normal signed integer", nil},                               // 13
	{"SLONG", 4, "Long signed integer", nil},                                 // 14

	{"USHORT", 1, "Short unsigned integer", USHORT},                    // 15
	{"UNORM", 2, "Normal unsigned integer", UNORM},                     // 16
	{"ULONG", 4, "Long unsigned integer", ULONG},                       // 17
	{"UVARI", 0, "Variable-length unsigned integer 1, 2, or 4", UVARI}, // 18
	{"IDENT", 0, "Variable-length identifier", IDENT},                  // 19
	{"ASCII", 0, "Variable-length ASCII character string", ASCII},      // 20

	{"DTIME", 8, "Date and time", nil}, // 21

	// 23 http://www.energistics.org/geosciences/geology-standards/rp66-organization-codes
	{"ORIGIN", 0, "Origin reference", ORIGIN}, // 22
	{"OBNAME", 0, "Object name", OBNAME},      // 23

	{"OBJREF", 0, "Object reference", nil},    // 24
	{"ATTREF", 0, "Attribute reference", nil}, // 25
	{"STATUS", 1, "Boolean status", nil},      // 26
	{"UNITS", 0, "Units expression",
		func(in []byte) (interface{}, int) {
			ln := in[0]
			if ln == 0 {
				return "", 0
			}
			// only allowed
			// lower case letters [a, b, c, ..., z]
			// upper case letters [A, B, C, ..., Z]
			// digits [0, 1, 2, ..., 9]
			// blank [ ]
			// hyphen or minus sign [-]
			// dot or period [.]
			// slash [/]
			// parentheses [(, )]

			// lotsa other rules

			// TODO check for allowed
			return string(in[1 : 1+ln]), int(1 + ln)
		}}, // 27

	// rp66.v2 has up to repcode 42, we'll keep up to 50 reserved
	{}, // 28
	{}, // 29
	{}, // 30
	{}, // 31
	{}, // 32
	{}, // 33
	{}, // 34
	{}, // 35
	{}, // 36
	{}, // 37
	{}, // 38
	{}, // 39
	{}, // 40
	{}, // 41
	{}, // 42
	{}, // 43
	{}, // 44
	{}, // 45
	{}, // 46
	{}, // 47
	{}, // 48
	{}, // 49
	{}, // 50

	{}, // 51 SUL
	{}, // 52
	{}, // 53

}
