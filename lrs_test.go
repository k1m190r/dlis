package dlis

import (
	"fmt"
	"testing"
)

func TestBits(t *testing.T) {
	b1 := byte(0x01)
	for b := range make([]int, 8) {
		fmt.Printf("%b  ", b1<<uint(b))
		fmt.Printf("%v\n", (b1<<uint(b))&(b1<<uint(3)) != 0)
	}

	fmt.Println(((1 << 3) & 0xff) != 0)
}
