package klubok

type loops struct {
	vertices vertices
	edges    edges
}

func newLoops(vertices vertices, edges edges) loops {
	return loops{vertices: vertices, edges: edges}
}

func (l loops) createLoop(vertex vertex) loop {
	return newLoop(vertex, l.vertices, l.edges)
}

func (l loops) next() position {

}
