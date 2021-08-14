package klubok

import (
	//"reflect"
	"reflect"
	"testing"
)

func TestGraph(t *testing.T) {
	g := NewGraph(NewSliceStorage())

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

	g.Connect(a, b)

	// a b a->b c

	want, c := uint(4), g.Create()
	if c != want {
		t.Errorf("want %v, got %v", want, c)
	}

	// a b a->b c a->c

	g.Connect(a, c)

	wantHeads, aHeads := []uint{2, 4}, g.ReadTargets(a)
	if !reflect.DeepEqual(wantHeads, aHeads) {
		t.Errorf("want %v, got %v", wantHeads, aHeads)
	}

	//wantHeads, aHeads = []uint{}, g.ReadSources(a)
	//if !reflect.DeepEqual(wantHeads, aHeads) {
	//	t.Errorf("want %v, got %v", wantHeads, aHeads)
	//}
	//
	//wantHeads, bHeads := []uint{}, g.ReadTargets(b)
	//if !reflect.DeepEqual(wantHeads, bHeads) {
	//	t.Errorf("want %v, got %v", wantHeads, bHeads)
	//}
	//
	//wantHeads, bHeads = []uint{1}, g.ReadSources(b)
	//if !reflect.DeepEqual(wantHeads, bHeads) {
	//	t.Errorf("want %v, got %v", wantHeads, bHeads)
	//}
	//
	//wantHeads, cHeads := []uint{}, g.ReadTargets(c)
	//if !reflect.DeepEqual(wantHeads, cHeads) {
	//	t.Errorf("want %v, got %v", wantHeads, cHeads)
	//}
	//
	//wantHeads, cHeads = []uint{1}, g.ReadSources(c)
	//if !reflect.DeepEqual(wantHeads, cHeads) {
	//	t.Errorf("want %v, got %v", wantHeads, cHeads)
	//}

	// b c

	//g.Delete(a)
	//
	//wantHeads, cHeads = []uint{}, g.ReadNegative(c)
	//if !reflect.DeepEqual(wantHeads, cHeads) {
	//	t.Errorf("want %v, got %v", wantHeads, cHeads)
	//}
}
