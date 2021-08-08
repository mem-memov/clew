package klubok

import (
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph()

	// a

	want, a := uint(1), g.Create()
	if a != want {
		t.Errorf("want %v, got %v", want, a)
	}

	// a b

	want, b := uint(2), g.Create()
	if b != want {
		t.Errorf("want %v, got %v", want, b)
	}

	// a b a->b

	g.Update(a, b)

	// a b a->b c

	want, c := uint(4), g.Create()
	if c != want {
		t.Errorf("want %v, got %v", want, c)
	}

	// a b a->b c a->c

	g.Update(a, c)

	wantHeads, aHeads := []uint{2, 4}, g.ReadPositive(a)
	if !reflect.DeepEqual(wantHeads, aHeads) {
		t.Errorf("want %v, got %v", wantHeads, aHeads)
	}

	wantHeads, aHeads = []uint{}, g.ReadNegative(a)
	if !reflect.DeepEqual(wantHeads, aHeads) {
		t.Errorf("want %v, got %v", wantHeads, aHeads)
	}

	wantHeads, bHeads := []uint{}, g.ReadPositive(b)
	if !reflect.DeepEqual(wantHeads, bHeads) {
		t.Errorf("want %v, got %v", wantHeads, bHeads)
	}

	wantHeads, bHeads = []uint{1}, g.ReadNegative(b)
	if !reflect.DeepEqual(wantHeads, bHeads) {
		t.Errorf("want %v, got %v", wantHeads, bHeads)
	}

	wantHeads, cHeads := []uint{}, g.ReadPositive(c)
	if !reflect.DeepEqual(wantHeads, cHeads) {
		t.Errorf("want %v, got %v", wantHeads, cHeads)
	}

	wantHeads, cHeads = []uint{1}, g.ReadNegative(c)
	if !reflect.DeepEqual(wantHeads, cHeads) {
		t.Errorf("want %v, got %v", wantHeads, cHeads)
	}

	// b c

	//g.Delete(a)
	//
	//wantHeads, cHeads = []uint{}, g.ReadNegative(c)
	//if !reflect.DeepEqual(wantHeads, cHeads) {
	//	t.Errorf("want %v, got %v", wantHeads, cHeads)
	//}
}