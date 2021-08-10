package klubok

type loop struct {
	vertex   vertex
	vertices vertices
	edges    edges
}

func newLoop(vertex vertex, vertices vertices, edges edges) loop {
	return loop{vertex: vertex, vertices: vertices, edges: edges}
}

func (l loop) getPosition() position {
	return l.vertex.getPosition()
}
