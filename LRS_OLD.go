// Logical Record Segment
package dlis

// http://w3.energistics.org/rp66/v1/rp66v1_sec3.html

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

// OBNAME: k&j&n'a..x : Origin k : Copy Number j : INDENT
// e.g.: 1&0&5'Depth

type OBNAME struct {
	Origin   byte
	CopyNum  byte
	ObjectID IDENT
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

type Component struct {
	Descriptor Descriptor
}

// LRSB interpretation as EFLR
type EFLR struct {
	Descriptor byte // $3.2.2.1 Bits 1-3 Role, 4-Type (Objects in the Set), 5-Name
}
