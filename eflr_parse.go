package dlis

import (
	"fmt"
)

// 3 - Logical Record Syntax
// http://w3.energistics.org/rp66/v1/rp66v1_sec3.html

// Roles holds roles for the Components
var Roles = []struct {
	Role string
	Type string
}{
	{"ABSATR", "Absent Attribute"}, // 000 0

	{"ATTRIB", "Attribute"},           // 001 1
	{"INVATR", "Invariant Attribute"}, // 010 2

	{"OBJECT", "Object"}, // 011 3

	{"reserved", "-"}, // 100 4

	{"RDSET", "Redundant Set"},  // 101 5
	{"RSET", "Replacement Set"}, // 110 6
	{"SET", "Set"},              // 111 7

}

// SetChars is set Characteristics
var SetChars = []struct {
	Chars   string
	RepCode int
	Default interface{}
}{
	{}, {}, {},

	{"Name", 19, byte(0)}, // 3

	{"Type", 19, nil}, // 4 19 IDENT

	{}, {}, {}, {},
}

func parseSet(s *LRS) {
	fmt.Print("\nS: ")

	if len(s.body) == 0 {
		fmt.Println("End of LRS body")
		return
	}

	// get byte one
	b1 := s.body[0]

	// restart body from 1+
	s.body = s.body[1:]

	if checkBit(b1, 4) { // Type
		repc := SetChars[4].RepCode
		val, ln := RepCode[repc].Read(s.body[:])
		s.body = s.body[ln:]
		fmt.Print("  Type:", val)
	}

	if checkBit(b1, 3) { // Name
		repc := SetChars[3].RepCode
		val, ln := RepCode[repc].Read(s.body[:])
		s.body = s.body[ln:]
		fmt.Print("  Name:", val)
	}

}

// ObjectChars is set Characteristics
var ObjectChars = []struct {
	Chars   string
	RepCode int
	Default interface{}
}{
	{}, {}, {}, {},

	{"Name", 23, nil}, // 4

	{}, {}, {}, {},
}

func parseObject(s *LRS) {
	fmt.Print("\nO: ")

	if len(s.body) == 0 {
		fmt.Println("End of LRS body")
		return
	}

	// get byte one
	b1 := s.body[0]

	// restart body from 1+
	s.body = s.body[1:]

	if checkBit(b1, 4) { // Name
		repc := ObjectChars[4].RepCode
		val, ln := RepCode[repc].Read(s.body[:])
		s.body = s.body[ln:]
		fmt.Print(" Name:", val)
	}
}

// AttribChars is set Characteristics
var AttribChars = []struct {
	Chars   string
	RepCode int
	Default interface{}
}{
	{"Value", 19, nil}, // 0: 0 here means it's repcode is defined by the REPCODE 19
	// Value is defined by Count of REPCODE type Units
	// If count is 0 the value is "undefined"
	{"Units", 27, byte(0)}, // 1: 27 UNITS
	{"REPCODE", 15, 19},    // 2: 15 USHORT
	{"Count", 18, 1},       // 3: 18 UVARI
	{"Label", 19, byte(0)}, // 4: 19 IDENT

	{}, {}, {},
}

func parseAttrib(s *LRS) {
	fmt.Print("\nA: ")

	if len(s.body) == 0 {
		fmt.Println("End of LRS body")
		return
	}

	// get byte one
	b1 := s.body[0]

	// restart body from 1+
	s.body = s.body[1:]

	if checkBit(b1, 4) { // Label
		repc := AttribChars[4].RepCode
		val, ln := RepCode[repc].Read(s.body[:])
		s.body = s.body[ln:]
		fmt.Print(" Label:", val)
	}

	if checkBit(b1, 3) { // Count
		repc := AttribChars[3].RepCode
		val, ln := RepCode[repc].Read(s.body[:])
		s.body = s.body[ln:]
		fmt.Print("  Count:", val)
	}

	if checkBit(b1, 2) { // REPCODE
		repc := AttribChars[2].RepCode
		val, ln := RepCode[repc].Read(s.body[:])
		s.body = s.body[ln:]
		fmt.Print(" Repcode:", val)
	}

	if checkBit(b1, 1) { // Units
		repc := AttribChars[1].RepCode
		val, ln := RepCode[repc].Read(s.body[:])
		s.body = s.body[ln:]
		fmt.Print(" Units:", val)
	}

	if checkBit(b1, 0) { // Value
		// check the value of REPCODE otherwise default to 19
		// value of the REPCODE in the template
		repc := AttribChars[0].RepCode
		val, ln := RepCode[repc].Read(s.body[:])
		s.body = s.body[ln:]
		fmt.Print(" Value:", val)
	}

}

// ParseEFLR parses the LRS body into Components
// TODO need to decide what happens to Components
func ParseEFLR(s *LRS) {
	for {

		if len(s.body) == 0 {
			fmt.Println("\nEnd of LRS body")
			return
		}

		b := s.body[0]
		role := b >> 5 // first 3 bits
		switch role {
		case 5, 6, 7: // Set Roles
			parseSet(s)
		case 3: // Object role
			parseObject(s)
		case 1, 2: // Attribute roles
			parseAttrib(s)
		case 0: // Absent
			fmt.Println("Absent argument")
			fmt.Println("Something is wrong...?")
			return
		}
	}
}

// $3.2 Explicitly Formatted Logical Record (EFLR)
// Template for columns/ attributes, and the their characteristics
// Table of information
//   Rows are Objects
//   Columns are Attributes of Objects
// Alternatively viewed as Set of Objects, of the Type Defined by Template

// Each EFLR contains one and only one Set.
// Set maybe of several different types implied by the EFLR Type.

// Set is 1+ Object of same type, preceded by Template.
// Each Object has 1+ Attributes.
// Sets, Objects and Attributes have Characteristics

// $3.2.2 EFLR Component

// Notation

// IDENT: n'a..x
// e.g: 5'Hello ; 6'Origin
type IDENT struct {
	Size byte
	Name string
}

// "null" REPCODE len bytes all 0

// 0' null ASCII or IDENT, zero length string, 1 byte = 0

// "reserved" bit is zero

// $3.2.2.1 Descriptor

// First byte of Component is Descriptor
// Bits 1-3 Role Fig 3-2
// Format Fig 3-3, 3-4, 3-5

type Descriptor struct {
	Role   byte // bits 1-3, Fig 3-2
	Format byte
	// Role Set (101, 110, 111): Fig 3-3, bit 4 - Type IDENT, 5 - Name IDENT
	//   defaults: Type - not defined, Name - 0'
	// Role Obj (011): Fig - 3-4, bit 4 - Name OBNAME
	// Role Attrib (001, 010): Fig 3-5, Label, Count, RepCode, Units, Value
	//   Value 0+ Elements of RepCode with Units, # of Elements is Count
	//   if Count==0 Value is undef ie Absent Value
}

// Component describe entities: Set, Objects, Attributes
type Component struct {
	Descriptor Descriptor // first byte
}

///////////
type Character struct {
}

type Set struct {
	Character Character
}

type Attribute struct {
	Character Character
}

type Object struct {
	Attributes []Attribute
	Character  Character
}

// Template specify: presence, order and default Character
// of the Attributes in the Objects in the Set
type Template struct {
}

// LRSB interpretation as EFLR
type EFLR struct {
	Set      Set
	Template Template
	Objects  []Object
	// Descriptor byte // $3.2.2.1 Bits 1-3 Role, 4-Type (Objects in the Set), 5-Name
}
