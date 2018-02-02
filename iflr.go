package dlis

import (
	"log"

	yaml "gopkg.in/yaml.v2"
)

var iflr [][]interface{}

//  Code inplied by the index of the array
// Code, Type, Description, AllowableSetTypes []

var rawiflr = []byte(`[

  [FDATA, "Frame Data", [FRAME]],  #0
  [NOFORMAT, "Unformatted Data", [NO-FORMAT]], #1

]`)

func init() {
	// f, err := os.Open("eflr.yaml")
	// if err != nil {
	// 	log.Fatal("OH OH error loading the eflr.yaml")
	// }
	// defer dclose(f)
	// raw, err := ioutil.ReadAll(f)

	if err := yaml.Unmarshal(rawiflr, &iflr); err != nil {
		log.Fatal("error loading eflr.yaml")
	}
}

func IFLRType(code byte) []interface{} {
	if code == 127 {
		return []interface{}{"EOD", "End of Data", []string{}}
	}
	if code > 1 {
		return []interface{}{"RESERVED", "RESERVED", []string{"RESERVED"}}
	}
	return iflr[int(code)]
}
