package klubok

type vertices struct {
	entries    entries
	holes      holes
	lastVertex position
}

func newVertices(entries entries, holes holes, lastVertex position) vertices {
	return vertices{
		entries:    entries,
		holes:      holes,
		lastVertex: lastVertex,
	}
}

func (v vertices) create() vertex {

	var vertex vertex
	if v.holes.exist() {
		vertex = newVertex(v.holes.last(), v.lastVertex)
		v.holes.consumeHole(vertex)
	} else {
		vertex = newVertex(v.entries.next(), v.lastVertex)
	}

	v.lastVertex = vertex.getPosition()

	return vertex
}

func (v vertices) read(position position) vertex {
	return newVertexWithEntry(position, v.entries.read(position))
}

func (v vertices) update(vx vertex) {
	vx.update(v.entries)
}
