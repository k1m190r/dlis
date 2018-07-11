package dlis

import (
	"testing"
)

func _TestFN(t *testing.T) {
	var v1 = S("hello")

	if v1.IsErr() {
		t.Log(v1)
		return
	}
	t.Logf("val: %v", v1)
}
