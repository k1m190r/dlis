package dlis

import (
	"reflect"
	"testing"
)

func TestRepCode(t *testing.T) {
	// for each RepCode get test cases with expected values
	for repCode, expRepVal := range RepCodeTest {

		readfn := RepCode[repCode].Read
		if readfn == nil {
			t.Logf("no function to test for REPCODE: %v", repCode)
			continue
		}

		if (expRepVal.In == nil) || (len(expRepVal.In) == 0) {
			t.Logf("no test cases for REPCODE: %v", repCode)
			continue
		}

		if (len(expRepVal.In) != len(expRepVal.Ret)) || (len(expRepVal.In) != len(expRepVal.Len)) {
			t.Logf("REPCODE: %v length of test cases is unequal.", repCode)
			continue
		}

		// for each test case
		for tstidx := range expRepVal.In {
			in := expRepVal.In[tstidx]

			exTp := expRepVal.Type
			exLn := expRepVal.Len[tstidx]
			exRet := expRepVal.Ret[tstidx]

			// run
			retVal, retLen := readfn(in)

			// check type
			if exTp != reflect.TypeOf(retVal).Kind() {
				t.Logf("REPCODE: %v return type expected %v but found %v\n",
					repCode, exTp, reflect.TypeOf(retVal).Kind())
				t.Fail()
				continue
			}

			// check length
			if exLn != retLen {
				t.Logf("REPCODE: %v return length expected %v but found %v\n",
					repCode, exLn, retLen)
				t.Fail()
				continue
			}

			// check value
			switch rv := retVal.(type) {
			case int:
				ev, ok := exRet.(int)
				if !ok {
					t.Logf("REPCODE: %v expected: %v of type %T, but found: %v of type %T",
						repCode, exRet, exRet, retVal, retVal)
					t.Fail()
					continue
				}

				if ev != rv {
					t.Logf("REPCODE: %v expected %v != returned %v",
						repCode, ev, rv)
					t.Fail()
					continue
				}

			case float32:
				ev, ok := exRet.(float32)
				if !ok {
					t.Logf("REPCODE: %v expected: %v of type %T, but found: %v of type %T",
						repCode, exRet, exRet, retVal, retVal)
					t.Fail()
					continue
				}

				if ev != rv {
					t.Logf("REPCODE: %v expected %v != returned %v", repCode, ev, rv)
					t.Fail()
					continue
				}

			case float64:
				ev, ok := exRet.(float64)
				if !ok {
					t.Logf("REPCODE: %v expected: %v of type %T, but found: %v of type %T",
						repCode, exRet, exRet, retVal, retVal)
					t.Fail()
					continue
				}

				if ev != rv {
					t.Logf("REPCODE: %v expected %v != returned %v", repCode, ev, rv)
					t.Fail()
					continue
				}

			case OBNAME:
				t.Log("repcode_test.go OBNAME to be done")
				continue

			default:
				t.Logf("Unknown type %v\n", reflect.TypeOf(rv))
				t.Fail()
			}
		}
	}
}

var RepCodeTest = []struct {
	In   [][]byte      // test cases
	Ret  []interface{} // value
	Type reflect.Kind  // type, this doesn't change
	Len  []int         // length

}{
	{}, // 0 is not present

	// {"FSHORT", 2, "Low precision floating point", nil}, // 1
	{},

	// {"FSINGL", 4, "IEEE single precision floating point", func}, // 2
	{
		In: [][]byte{
			[]byte("FILE"), // case 0
			[]byte("NAME"), // case 1
		},
		Ret: []interface{}{
			float32(12883.0673828125), // case 0
			float32(810766656),        // case 1
		},
		Type: reflect.Float32,
		Len:  []int{4, 4},
	},

	// {"FSING1", 8, "Validated single precision floating point", nil}, // 3
	{},

	// {"FSING2", 12, "Two-way validated single precision floating point", nil}, // 4
	{},

	// {"ISINGL", 4, "IBM single precision floating point", nil},                // 5
	{},

	// {"VSINGL", 4, "VAX single precision floating point", nil},                // 6
	{},

	// {"FDOUBL", 8, "IEEE double precision floating point",func }, // 7
	{
		In: [][]byte{
			[]byte("CREATION"), // case 0
			[]byte("PRODUCER"), // case 1
		},
		Ret: []interface{}{
			float64(2.0570785880292664E16), // case 0
			float64(8.480444221488923E78),  // case 1
		},
		Type: reflect.Float64,
		Len:  []int{8, 8},
	},

	// {"FDOUB1", 16, "Validated double precision floating point", nil},         // 8
	{},

	// {"FDOUB2", 24, "Two-way validated double precision floating point", nil}, // 9
	{},

	// {"CSINGL", 8, "Single precision complex", nil},                           // 10
	{},

	// {"CDOUBL", 16, "Double precision complex", nil},                          // 11
	{},

	// {"SSHORT", 1, "Short signed integer", nil},                               // 12
	{},

	// {"SNORM", 2, "Normal signed integer", nil},                               // 13
	{},

	// {"SLONG", 4, "Long signed integer", nil},                                 // 14
	{},

	// {"USHORT", 1, "Short unsigned integer", func}, // 15
	{
		In: [][]byte{
			[]byte("a"),
			[]byte("b"),
			[]byte("c"),
		},
		Ret: []interface{}{
			int(97),
			int(98),
			int(99),
		},
		Type: reflect.Int,
		Len:  []int{1, 1, 1},
	},

	// {"UNORM", 2, "Normal unsigned integer", func}, // 16
	{
		In: [][]byte{
			[]byte("ab"),
			[]byte("cd"),
			[]byte("ef"),
		},
		Ret: []interface{}{
			int(24930),
			int(25444),
			int(25958),
		},
		Type: reflect.Int,
		Len:  []int{2, 2, 2},
	},

	// {"ULONG", 4, "Long unsigned integer", func}, // 17
	{
		In: [][]byte{
			[]byte("abcd"),
			[]byte("xyz "),
			[]byte("1234"),
		},
		Ret: []interface{}{
			int(1633837924),
			int(2021227040),
			int(825373492),
		},
		Type: reflect.Int,
		Len:  []int{4, 4, 4},
	},

	// {"UVARI", 0, "Variable-length unsigned integer 1, 2, or 4", func}, // 18
	{
		In: [][]byte{
			[]byte("a"),
			[]byte{0xAA, 0xAA},             // A = 1010 -> 2 = 0010
			[]byte{0xCC, 0xCC, 0xCC, 0xCC}, // C = 1100 -> 0 = 0000
		},
		Ret: []interface{}{
			int(97),
			int(10922),
			int(214748364),
		},
		Type: reflect.Int,
		Len:  []int{1, 2, 4},
	},
	/*
		{"IDENT", 0, "Variable-length identifier",
			func(in []byte) (interface{}, int) {
				ln := in[0]
				if ln == 0 {
					return "", 0
				}
				// only allowed 33-96, 123-126
				// TODO check for allowed
				return string(in[1 : 1+ln]), int(1 + ln)
			}}, // 19

		{"ASCII", 0, "Variable-length ASCII character string",
			func(in []byte) (interface{}, int) {
				b1 := in[0]

				if b1 == 0 {
					return "", 1
				}

				var idlen, asciilen int
				if checkBit(7, uint(b1)) { //
					if checkBit(6, uint(b1)) { // 4 bytes
						tmp := [4]byte{b1 & 0x3F} // first byte with mask 0011_1111
						copy(tmp[1:], in[1:4])    // remaining 3 bytes
						idlen = 4
						asciilen = int(binary.BigEndian.Uint32(tmp[:]))
					} else { // 2 bytes
						tmp := [2]byte{b1 & 0x3F} // first byte with mask 0011_1111
						tmp[1] = in[1]
						idlen = 2
						asciilen = int(binary.BigEndian.Uint16(tmp[:]))
					}
				} else {
					// single byte
					idlen = 1
					asciilen = int(b1) // bit 7 is 0
				}

				return string(in[idlen : idlen+asciilen]), (idlen + asciilen)
			}}, // 20

		// {"DTIME", 8, "Date and time", nil}, // 21

		{"ORIGIN", 0, "Origin reference",
			func(in []byte) (interface{}, int) {
				b1 := in[0]
				if checkBit(7, uint(b1)) { //
					if checkBit(6, uint(b1)) { // 4 bytes
						tmp := [4]byte{b1 & 0x3F} // first byte with mask 0011_1111
						copy(tmp[1:], in[1:4])    // remaining 3 bytes
						return int(binary.BigEndian.Uint32(tmp[:])), 4
					} else { // 2 bytes
						tmp := [2]byte{b1 & 0x3F} // first byte with mask 0011_1111
						tmp[1] = in[1]
						return int(binary.BigEndian.Uint16(tmp[:])), 2
					}
				}
				// single byte
				return int(b1), 1 // bit 7 is 0
			}}, // 22

		{"OBNAME", 0, "Object name",
			func(in []byte) (interface{}, int) {

				// ORIGIN
				b1 := in[0]
				var olen, origin int
				if checkBit(7, uint(b1)) { //
					if checkBit(6, uint(b1)) { // 4 bytes
						tmp := [4]byte{b1 & 0x3F} // first byte with mask 0011_1111
						copy(tmp[1:], in[1:4])    // remaining 3 bytes
						olen = 4
						origin = int(binary.BigEndian.Uint32(tmp[:]))
					} else { // 2 bytes
						tmp := [2]byte{b1 & 0x3F} // first byte with mask 0011_1111
						tmp[1] = in[1]
						olen = 2
						origin = int(binary.BigEndian.Uint16(tmp[:]))
					}
				} else {
					// single byte
					olen = 1
					origin = int(b1) // bit 7 is 0
				}

				// COPY
				copy := int(in[olen])

				// IDENT
				ln := int(in[olen+1])
				ident := ""
				if ln == 0 {
					ident = ""
				} else {
					ident = string(in[olen+2 : olen+2+ln])
				}
				// only allowed 33-96, 123-126
				// TODO check for allowed

				return struct {
					Origin, Copy int
					Ident        string
				}{
					origin, copy, ident,
				}, int(olen + 2 + ln)
			}}, // 23

		// {"OBJREF", 0, "Object reference", nil},    // 24
		// {"ATTREF", 0, "Attribute reference", nil}, // 25
		// {"STATUS", 1, "Boolean status", nil},      // 26
		// {"UNITS", 0, "Units expression", nil},     // 27
	*/
}
