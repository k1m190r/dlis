package dlis

import (
	"fmt"
	"testing"
)

func _TestEFLR(t *testing.T) {
	// t.Log(EFLR)
	for i, r := range eflr {
		fmt.Print(i, " ")
		for _, v := range r {
			switch v1 := v.(type) {
			case string:
				fmt.Print(v1, " ")
			case []interface{}:
				fmt.Print(v1, " ")
			}

		}
		fmt.Println()
	}
}
