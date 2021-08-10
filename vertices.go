package klubok

type vertices struct {
	lastVertex position
	entries    entries
}

func newVertices(entries entries, lastVertex position) vertices {
	return vertices{
		lastVertex: lastVertex,
		entries:    entries,
	}
}

func (v vertices) create(position position) vertex {
	vertex := newVertex(position, v.lastVertex)
	v.lastVertex = position

	return vertex
}

func (v vertices) read(position position) vertex {
	return newVertexWithEntry(position, v.entries.read(position))
}

func (v vertices) update(vx vertex) {
	vx.update(v.entries)
}
