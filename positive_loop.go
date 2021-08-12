package klubok

type positiveLoop struct {
	vertices vertices
	edges    edges
}

func newPositiveLoop(vertices vertices, edges edges) positiveLoop {
	return positiveLoop{vertices: vertices, edges: edges}
}

func (p positiveLoop) readHeads(tail vertex) []position {
	heads := make([]position, 0)

	if !tail.hasFirstPositiveEdge() {
		return heads
	}

	firstEdge := tail.getFirstPositiveEdge(p.edges)
	nextEdge := firstEdge

	heads = append(heads, nextEdge.getPosition())

	for {
		nextEdge = nextEdge.getNextPositiveEdge(p.edges)
		if nextEdge.getPosition() == firstEdge.getPosition() {
			return heads
		}
		heads = append(heads, nextEdge.getPosition())
	}
}

func (p positiveLoop) addHead(tail vertex, edge edge) {

	if !tail.hasFirstPositiveEdge() {
		tail.setFirstPositiveEdge(edge)
		p.vertices.update(tail)
	} else {
		firstPositiveEdge := tail.getFirstPositiveEdge(p.edges)
		if firstPositiveEdge.hasPreviousPositiveEdge() {
			lastPositiveEdge := firstPositiveEdge.getPreviousPositiveEdge(p.edges)
			lastPositiveEdge.setNextPositiveEdge(edge)
			edge.setPreviousPositiveEdge(lastPositiveEdge)
			edge.setNextPositiveEdge(firstPositiveEdge)
			firstPositiveEdge.setPreviousPositiveEdge(edge)
			p.edges.update(lastPositiveEdge)
		} else {
			firstPositiveEdge.setPreviousPositiveEdge(edge)
			edge.setPreviousPositiveEdge(firstPositiveEdge)
			edge.setNextPositiveEdge(firstPositiveEdge)
		}
		p.edges.update(firstPositiveEdge)
	}
}

func (p positiveLoop) removeHead(tail vertex, edge edge) {

}
