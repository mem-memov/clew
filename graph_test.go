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
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
	}) {
		t.Error(s)
	}
}

func TestGraph_Create3(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	g.Create()
	g.Create()
	g.Create()

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
		{2, 0, 0, 0, 0, 0}, // 2 b
		{3, 0, 0, 0, 0, 0}, // 3 c
	}) {
		t.Error(s)
	}
}

func TestGraph_Delete1(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
	}

	err := g.Delete(1)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}) {
		t.Error(s)
	}
}

func TestGraph_Delete3(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0},
		{2, 1, 0, 0, 0, 0},
		{3, 2, 0, 0, 0, 0},
	}

	err := g.Delete(1)

	if err != nil {
		t.Fail()
	}

	err = g.Delete(2)

	if err != nil {
		t.Fail()
	}

	err = g.Delete(3)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{0, 0, 0, 0, 0, 0}, // 1 hole
		{1, 0, 0, 0, 0, 0}, // 2 hole
		{2, 0, 0, 0, 0, 0}, // 3 hole
	}) {
		t.Error(s)
	}
}

func TestGraph_Connect1To1(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
		{2, 0, 0, 0, 0, 0}, // 2 b
	}

	err := g.Connect(1, 2)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 1, 0, 3}, // 1 a
		{2, 0, 1, 0, 3, 0}, // 2 b
		{2, 0, 0, 1, 0, 0}, // 3 a -> b
	}) {
		t.Error(s)
	}
}

func TestGraph_Connect1ToItself(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
	}

	err := g.Connect(1, 1)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 1, 1, 2, 2}, // 1 a
		{1, 0, 0, 1, 0, 0}, // 2 a -> a
	}) {
		t.Error(s)
	}
}

func TestGraph_Delete1With1Connection(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 1, 0, 3}, // 1 a
		{2, 0, 1, 0, 3, 0}, // 2 b
		{2, 0, 0, 1, 0, 0}, // 3 a -> b
	}

	err := g.Delete(1)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{3, 0, 0, 0, 0, 0}, // 1 hole
		{2, 0, 0, 0, 0, 0}, // 2 b
		{0, 0, 0, 0, 0, 0}, // 3 hole
	}) {
		t.Error(s)
	}
}

func TestGraph_Delete1WithIAndOutConnections(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 1, 0, 4}, // 1 a
		{2, 0, 1, 1, 4, 5}, // 2 b
		{3, 0, 1, 0, 5, 0}, // 3 c
		{2, 0, 0, 1, 0, 0}, // 4 a -> b
		{3, 0, 0, 2, 0, 0}, // 5 b -> c
	}

	err := g.Delete(2)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
		{4, 0, 0, 0, 0, 0}, // 2 hole
		{3, 0, 0, 0, 0, 0}, // 3 c
		{5, 0, 0, 0, 0, 0}, // 4 hole
		{0, 0, 0, 0, 0, 0}, // 5 hole
	}) {
		t.Error(s)
	}
}

func TestGraph_Connect1To2Different(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
		{2, 0, 0, 0, 0, 0}, // 2 b
		{3, 0, 0, 0, 0, 0}, // 3 c
	}

	err := g.Connect(1, 2)

	if err != nil {
		t.Fail()
	}

	err = g.Connect(1, 3)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 2, 0, 4}, // 1 a
		{2, 0, 1, 0, 4, 0}, // 2 b
		{3, 0, 1, 0, 5, 0}, // 3 c
		{2, 0, 0, 1, 5, 5}, // 4 a -> b
		{3, 0, 0, 1, 4, 4}, // 5 a -> c
	}) {
		t.Error(s)
	}
}

func TestGraph_Connect1InAndOut(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
		{2, 0, 0, 0, 0, 0}, // 2 b
		{3, 0, 0, 0, 0, 0}, // 3 c
	}

	err := g.Connect(1, 2)

	if err != nil {
		t.Fail()
	}

	err = g.Connect(2, 3)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 1, 0, 4}, // 1 a
		{2, 0, 1, 1, 4, 5}, // 2 b
		{3, 0, 1, 0, 5, 0}, // 3 c
		{2, 0, 0, 1, 0, 0}, // 4 a -> b
		{3, 0, 0, 2, 0, 0}, // 5 b -> c
	}) {
		t.Error(s)
	}
}

func TestGraph_Connect2DifferentTo1(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
		{2, 1, 0, 0, 0, 0}, // 2 b
		{3, 2, 0, 0, 0, 0}, // 3 c
	}

	err := g.Connect(2, 1)

	if err != nil {
		t.Fail()
	}

	err = g.Connect(3, 1)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 2, 0, 4, 0}, // 1 a
		{2, 1, 0, 1, 0, 4}, // 2 b
		{3, 2, 0, 1, 0, 5}, // 3 c
		{1, 5, 5, 2, 0, 0}, // 4 b -> a
		{1, 4, 4, 3, 0, 0}, // 5 c -> a
	}) {
		t.Error(s)
	}
}

func TestGraph_Connect1To3Different(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
		{2, 1, 0, 0, 0, 0}, // 2 b
		{3, 2, 0, 0, 0, 0}, // 3 c
		{4, 3, 0, 0, 0, 0}, // 4 d
	}

	err := g.Connect(1, 2)

	if err != nil {
		t.Fail()
	}

	err = g.Connect(1, 3)

	if err != nil {
		t.Fail()
	}

	err = g.Connect(1, 4)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 3, 0, 5}, // 1 a
		{2, 1, 1, 0, 5, 0}, // 2 b
		{3, 2, 1, 0, 6, 0}, // 3 c
		{4, 3, 1, 0, 7, 0}, // 4 d
		{2, 0, 0, 1, 7, 6}, // 5 a -> b
		{3, 0, 0, 1, 5, 7}, // 6 a -> c
		{4, 0, 0, 1, 6, 5}, // 7 a -> d
	}) {
		t.Error(s)
	}
}

func TestGraph_Disconnect(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 1, 0, 3}, // 1 a
		{2, 0, 1, 0, 3, 0}, // 2 b
		{2, 0, 0, 1, 0, 0}, // 3 a -> b
	}

	err := g.Disconnect(1, 2)

	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(s.entries, [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
		{2, 0, 0, 0, 0, 0}, // 2 b
		{0, 0, 0, 0, 0, 0}, // hole
	}) {
		t.Error(s)
	}
}

func TestGraph_ConnectMutually(t *testing.T) {
	s := NewSliceStorage()
	g, _ := NewGraph(s)

	s.entries = [][6]uint{
		{0, 0, 0, 0, 0, 0}, // 0
		{1, 0, 0, 0, 0, 0}, // 1 a
		{2, 0, 0, 0, 0, 0}, // 2 b
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
		{1, 0, 1, 1, 4, 3}, // 1 a
		{2, 0, 1, 1, 3, 4}, // 2 b
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

	//aHeads, _ = g.ReadSources(a)
	//if !reflect.DeepEqual([]uint{}, aHeads) {
	//	t.Errorf("want %v, got %v", []uint{}, aHeads)
	//}

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

	g.Delete(a)

	cHeads, _ = g.ReadSources(c)
	if !reflect.DeepEqual([]uint{}, cHeads) {
		t.Errorf("want %v, got %v", []uint{}, cHeads)
	}
}
