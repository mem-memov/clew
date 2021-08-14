package clew

import (
	"reflect"
	"testing"
)

func TestGraph_NewGraph(t *testing.T) {
	s := NewSliceStorage()
	_, err := NewGraph(s)

	if err != nil {
		t.Fail()
	}

	if len(s.entries) != 1 {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0},
	}) {
		t.Fail()
	}
}

func TestGraph_Create(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	id, err := g.Create()

	if err != nil {
		t.Fail()
	}

	if id != uint(1) {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
	}) {
		t.Fail()
	}
}

func TestGraph_Create3(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	g.Create()
	g.Create()
	g.Create()

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
		{3, 2, 0, 0, 0, 0},
	}) {
		t.Fail()
	}
}

func TestGraph_Connect(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
	}

	err := g.Connect(1, 2)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 3},
		{2, 1, 0, 3, 0, 0},
		{2, 0, 0, 1, 0, 0},
	}) {
		t.Error(s)
	}
}

func TestGraph_Disconnect(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 3},
		{2, 1, 0, 3, 0, 0},
		{2, 0, 0, 1, 0, 0},
	}

	err := g.Disconnect(1, 2)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}) {
		t.Error(s)
	}
}

func TestGraph_ConnectMutually(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
	}

	err := g.Connect(1, 2)

	if err != nil {
		t.Fail()
	}

	err = g.Connect(2, 1)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 4, 0, 3}, // 1 a
		{2, 1, 0, 3, 0, 4}, // 2 b
		{2, 0, 0, 1, 0, 0}, // 3 a -> b
		{1, 0, 0, 2, 0, 0}, // 4 b -> a
	}) {
		t.Error(s)
	}
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
