package dlis

import (
	"log"

	"gopkg.in/yaml.v2"
)

// http://w3.energistics.org/rp66/V1/rp66v1_appa.html
// A.2 Explicitly Formatted Logical Record

var eflr [][]interface{}

//  Code inplied by the index of the array
// Code, Type, Description, AllowableSetTypes []

var raweflr = []byte(`[

  [FHLR, "File Header", [FILE-HEADER]],  #0
  [OLR, Origin, [ORIGIN, WELL-REFERENCE]], #1
  [AXIS, "Coordinate Axis", [AXIS]], #2
  [CHANNL, "Channel-related information", [CHANNEL]], #3
  [FRAME, "Frame Data", [FRAME, PATH]], #4
  [STATIC, "Static Data", [CALIBRATION, CALIBRATION-COEFFICIENT, 
    CALIBRATION-MEASUREMENT, COMPUTATION, EQUIPMENT, GROUP, PARAMETER, 
    PROCESS, SPICE, TOOL, ZONE]], #5
  [SCRIPT, "Textual Data", [COMMENT]], #6
  [UPDATE, "Update Data", [UPDATE]], #7
  [UDI, "Unformatted Data Identifier", [NO-FORMAT]], #8
  [LNAME, "Long Name", [LONG-NAME]], #9
  [SPEC, "Specification", [ATTRIBUTE, CODE, EFLR, IFLR, OBJECT-TYPE, 
    REPRESENTATION-CODE, SPECIFICATION, UNIT-SYMBOL]], #10
  [DICT, "Dictionary", [BASE-DICTIONARY, IDENTIFIER, LEXICON, OPTION]], #11

]`)

func init() {
	// f, err := os.Open("eflr_type.yaml")
	// if err != nil {
	// 	log.Fatal("OH OH error loading the eflr_type.yaml")
	// }
	// defer dclose(f)
	// raw, err := ioutil.ReadAll(f)

	if err := yaml.Unmarshal(raweflr, &eflr); err != nil {
		log.Fatal("error loading eflr.yaml")
	}
}

func EFLRType(code byte) []interface{} {
	// 12 onwards is reserved
	if code > 11 {
		return []interface{}{"RESERVED", "RESERVED", []string{"RESERVED"}}
	}
	return eflr[int(code)]
}
