package gogame

import (
	"testing"
)

func TestWindow(t *testing.T) {
	t.Log("hi")
	w := New()
	if w.x != 0 {
		t.Errorf("Window.x != 0\n")
	}
	//t.Errorf("%s\n", s)
	//const in, out = 4, 2
	//if x := Sqrt(in); x != out {
	//	t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	//}
}
