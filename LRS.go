// Logical Record Segment
package main

// http://w3.energistics.org/rp66/v1/rp66v1_sec3.html

// $3.2 Explicitly Formatted Logical Record (EFLR)
// Template for columns/ attributes, and the their characteristics
// Table of information
//   Rows are Objects
//   Columns are Attributes of Objects

// $3.2.2 EFLR Component

// Notation

// IDENT: n'a..x
// e.g: 5'Hello ; 6'Origin

// OBNAME: k&j&n'a..x : Origin k : Copy Number j : INDENT
// e.g.: 1&0&5'Depth

// "null" REPCODE len bytes all 0

// 0' null ASCII or IDENT, zero length string

// "reserved" bit is zero

// $3.2.2.1 Descriptor

// First byte is Descriptor
// Bits 1-3 Role Fig 3-2
// Format Fig 3-3, 3-4, 3-5

// LRSB interpretation as EFLR
type EFLR struct {
	Descriptor byte // $3.2.2.1 Bits 1-3 Role, 4-Type (Objects in the Set), 5-Name
}
