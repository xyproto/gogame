package gogame

import (
	"testing"
)

func TestApp(t *testing.T) {
	t.Log("hi")
	a := New()
	if a.x != 0 {
		t.Errorf("App.x != 0\n")
	}
	a.Run()
	//t.Errorf("%s\n", s)
	//const in, out = 4, 2
	//if x := Sqrt(in); x != out {
	//	t.Errorf("Sqrt(%v) = %v, want %v", in, x, out)
	//}
}
