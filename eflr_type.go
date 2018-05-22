package dlis

// http://w3.energistics.org/rp66/V1/rp66v1_appa.html
// A.2 Explicitly Formatted Logical Record

// LRType logical record type
type LRType struct {
	Type              string
	Description       string
	AllowableSetTypes []string
}

// Code inplied by the index of the array
// Code, Type, Description, AllowableSetTypes []

var eflr = []LRType{
	{"FHLR", "File Header", []string{"FILE-HEADER"}},               // 0
	{"OLR", "Origin", []string{"ORIGIN", "WELL-REFERENCE"}},        // 1
	{"AXIS", "Coordinate Axis", []string{"AXIS"}},                  // 2
	{"CHANNL", "Channel-related information", []string{"CHANNEL"}}, // 3
	{"FRAME", "Frame Data", []string{"FRAME", "PATH"}},             // 4
	{"STATIC", "Static Data", []string{"CALIBRATION", "CALIBRATION-COEFFICIENT",
		"CALIBRATION-MEASUREMENT", "COMPUTATION", "EQUIPMENT", "GROUP", "PARAMETER",
		"PROCESS", "SPICE", "TOOL", "ZONE"}}, // 5
	{"SCRIPT", "Textual Data", []string{"COMMENT"}},               // 6
	{"UPDATE", "Update Data", []string{"UPDATE"}},                 // 7
	{"UDI", "Unformatted Data Identifier", []string{"NO-FORMAT"}}, // 8
	{"LNAME", "Long Name", []string{"LONG-NAME"}},                 // 9
	{"SPEC", "Specification", []string{"ATTRIBUTE", "CODE", "EFLR", "IFLR", "OBJECT-TYPE",
		"REPRESENTATION-CODE", "SPECIFICATION", "UNIT-SYMBOL"}}, // 10
	{"DICT", "Dictionary", []string{"BASE-DICTIONARY", "IDENTIFIER", "LEXICON", "OPTION"}}, // 11
}

func EFLRType(code byte) LRType {
	// 12 onwards is reserved
	if code > 11 {
		return LRType{"RESERVED", "RESERVED", []string{"RESERVED"}}
	}
	return eflr[int(code)]
}
