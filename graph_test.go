package clew

import (
	"reflect"
	"testing"
)

func TestGraph_oneNode(t *testing.T) {
	g, _ := NewGraph(NewSliceStorage())

	a, _ := g.Create()
	if a != uint(1) {
		t.Errorf("want %v, got %v", uint(1), a)
	}

	aHeads, _ := g.ReadTargets(a)
	if !reflect.DeepEqual([]uint{}, aHeads) {
		t.Errorf("want %v, got %v", []uint{}, aHeads)
	}

	aTails, _ := g.ReadSources(a)
	if !reflect.DeepEqual([]uint{}, aTails) {
		t.Errorf("want %v, got %v", []uint{}, aTails)
	}

	_ = g.Delete(a)
}

func TestGraph(t *testing.T) {
	g, _ := NewGraph(NewSliceStorage())

	// a

	a, _ := g.Create()
	if a != uint(1) {
		t.Errorf("want %v, got %v", uint(1), a)
	}

	// a b

	b, _ := g.Create()
	if b != uint(2) {
		t.Errorf("want %v, got %v", uint(2), b)
	}

	// a b a->b

	_ = g.Connect(a, b)

	// a b a->b c

	c, _ := g.Create()
	if c != uint(4) {
		t.Errorf("want %v, got %v", uint(4), c)
	}

	// a b a->b c a->c

	_ = g.Connect(a, c)

	aHeads, _ := g.ReadTargets(a)
	if !reflect.DeepEqual([]uint{2, 4}, aHeads) {
		t.Errorf("want %v, got %v", []uint{2, 4}, aHeads)
	}

	aHeads, _ = g.ReadSources(a)
	if !reflect.DeepEqual([]uint{}, aHeads) {
		t.Errorf("want %v, got %v", []uint{}, aHeads)
	}

	bHeads, _ := g.ReadTargets(b)
	if !reflect.DeepEqual([]uint{}, bHeads) {
		t.Errorf("want %v, got %v", []uint{}, bHeads)
	}

	bHeads, _ = g.ReadSources(b)
	if !reflect.DeepEqual([]uint{1}, bHeads) {
		t.Errorf("want %v, got %v", []uint{1}, bHeads)
	}

	cHeads, _ := g.ReadTargets(c)
	if !reflect.DeepEqual([]uint{}, cHeads) {
		t.Errorf("want %v, got %v", []uint{}, cHeads)
	}

	cHeads, _ = g.ReadSources(c)
	if !reflect.DeepEqual([]uint{1}, cHeads) {
		t.Errorf("want %v, got %v", []uint{1}, cHeads)
	}

	// b c

	//g.Disconnect(a, b)
	//g.Disconnect(a, c)
	//
	//wantHeads, cHeads = []uint{}, g.ReadSources(c)
	//if !reflect.DeepEqual(wantHeads, cHeads) {
	//	t.Errorf("want %v, got %v", wantHeads, cHeads)
	//}
}
