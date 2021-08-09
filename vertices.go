package klubok

type vertices struct {
	lastVertex position
	entries storage
}

func newVertices(s storage, l position) vertices {
	return vertices{
		lastVertex: l,
		entries: s,
	}
}

func (v vertices) create(p position) vertex {
	vertex := newVertex(p)
	return vertex
}

func (v vertices) read(p position) vertex {
	return newVertex(v.entries.read(p))
}

func (v vertices) update(vx vertex) {
	vx.update(v.entries)
}