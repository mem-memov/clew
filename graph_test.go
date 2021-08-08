package klubok

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()

	want, a := uint(1), g.Create()
	if a != want {
		t.Errorf("want %v, got %v", want, a)
	}

	want, b := uint(2), g.Create()
	if b != want {
		t.Errorf("want %v, got %v", want, b)
	}

	g.Update(a, b)

	want, c := uint(4), g.Create()
	if c != want {
		t.Errorf("want %v, got %v", want, c)
	}

	g.Update(a, c)

	wantHeads, aHeads := []uint{2, 4}, g.Read(a)
	if !reflect.DeepEqual(wantHeads, aHeads) {
		t.Errorf("want %v, got %v", wantHeads, aHeads)
	}
}