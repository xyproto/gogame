package gogame

import (
	"testing"
)

func TestText(tst *testing.T) {
	tst.Log("hi")
	t := New()
	if t.x != 0 {
		tst.Errorf("Text.x != 0\n")
	}
}
