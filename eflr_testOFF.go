package dlis

import (
	"fmt"
	"testing"
)

func TestEFLR(t *testing.T) {
	// t.Log(EFLR)
	for i, r := range eflr {
		fmt.Print(i, r)
		fmt.Println()
	}
}
