package klubok

type headLoop struct {
	vertices vertices
	edges    edges
}

func newPositiveLoop(vertices vertices, edges edges) headLoop {
	return headLoop{vertices: vertices, edges: edges}
}

func (p headLoop) readHeads(tail tailVertex) []position {
	heads := make([]position, 0)

	if !tail.hasFirstEdgeTail() {
		return heads
	}

	firstEdge := tail.getFirstEdgeTail(p.edges).toHead()
	nextEdge := firstEdge

	heads = append(heads, nextEdge.toEddge().getPosition())

	for {
		nextEdge = nextEdge.getNextPositiveEdge(p.edges).toHead()
		if nextEdge.toEddge().getPosition() == firstEdge.toEddge().getPosition() {
			return heads
		}
		heads = append(heads, nextEdge.toEddge().getPosition())
	}
}

func (p headLoop) addHead(tail tailVertex, edge edgeTail) {

	if !head.hasFirstEdgeHead() {
		head.setFirstEdgeHead(edge)
		n.vertices.update(head.toVertex())
	} else {
		firstNegativeEdge := head.getFirstEdgeHead(n.edges)
		if firstNegativeEdge.hasPreviousNegativeEdge() {
			lastNegativeEdge := firstNegativeEdge.getPreviousNegativeEdge(n.edges)
			lastNegativeEdge.setNextNegativeEdge(edge)
			edge.setPreviousNegativeEdge(lastNegativeEdge)
			edge.setNextNegativeEdge(firstNegativeEdge)
			firstNegativeEdge.setPreviousNegativeEdge(edge)
			n.edges.update(lastNegativeEdge)
		} else {
			firstNegativeEdge.setPreviousNegativeEdge(edge)
			edge.setPreviousNegativeEdge(firstNegativeEdge)
			edge.setNextNegativeEdge(firstNegativeEdge)
		}
		n.edges.update(firstNegativeEdge)
	}

	n.edges.update(edge)
}

func (p headLoop) removeHead(tail tailVertex, edge edge) {
	if tail.isFirstEdgeTail(edge) {
		tail.deleteFirstEdgeTail()
		return
	}

	firstPositiveEdge := tail.getFirstEdgeTail(p.edges)
}
