package dlis

import (
	"strings"
	"testing"
)

var sultxt = "   1V1.00RECORD 8192Default Storage Set                                         "
var sul SUL

func _TestRead(t *testing.T) {
	err := sul.Read(strings.NewReader(sultxt))
	if err != nil {
		t.Error(err)
	}
	t.Log(sul)
}
