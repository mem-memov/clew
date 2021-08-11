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

func (p positiveLoop) addHead(tail vertex, head vertex) {

	edge := p.edges.create()
	edge.setPositiveVertex(head)
	edge.setNegativeVertex(tail)

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

	if !head.hasFirstNegativeEdge() {
		head.setFirstNegativeEdge(edge)
		p.vertices.update(tail)
	} else {
		firstNegativeEdge := tail.getFirstNegativeEdge(p.edges)
		if firstNegativeEdge.hasPreviousNegativeEdge() {
			lastNegativeEdge := firstNegativeEdge.getPreviousNegativeEdge(p.edges)
			lastNegativeEdge.setNextNegativeEdge(edge)
			edge.setPreviousNegativeEdge(lastNegativeEdge)
			edge.setNextNegativeEdge(firstNegativeEdge)
			firstNegativeEdge.setPreviousNegativeEdge(edge)
			p.edges.update(lastNegativeEdge)
		} else {
			firstNegativeEdge.setPreviousNegativeEdge(edge)
			edge.setPreviousNegativeEdge(firstNegativeEdge)
			edge.setNextNegativeEdge(firstNegativeEdge)
		}
		p.edges.update(firstNegativeEdge)
	}

	p.edges.update(edge)
}

func (p positiveLoop) removeHead(tail vertex, head vertex) {
	if !tail.hasFirstPositiveEdge() {
		return
	}

	firstEdge := tail.getFirstPositiveEdge(p.edges)
	nextEdge := firstEdge

	if nextEdge.hasPositiveVertex(head) {
		nextEdge.
	}

	for {
		nextEdge = nextEdge.getNextPositiveEdge(p.edges)
		if nextEdge.getPosition() == firstEdge.getPosition() {
			return
		}
		if nextEdge.hasPositiveVertex(head) {

		}
	}
}
