package klubok

type headLoop struct {
	vertices vertices
	edges    edges
}

func newHeadLoop(vertices vertices, edges edges) headLoop {
	return headLoop{vertices: vertices, edges: edges}
}

func (h headLoop) readHeads(vertex headVertex) []position {
	heads := make([]position, 0)

	if !vertex.hasFirstEdgeHead() {
		return heads
	}

	firstEdge := vertex.getFirstEdgeHead(h.edges)
	nextEdge := firstEdge

	heads = append(heads, nextEdge.toEdge().getPosition())

	for {
		nextEdge = nextEdge.getNextEdgeHead(h.edges)
		if nextEdge.toEdge().getPosition() == firstEdge.toEdge().getPosition() {
			return heads
		}
		heads = append(heads, nextEdge.toEdge().getPosition())
	}
}

func (h headLoop) addHead(vertex headVertex, newEdge edgeHead) {

	if !vertex.hasFirstEdgeHead() {
		vertex.setFirstEdgeHead(newEdge)
		h.vertices.update(vertex.toVertex())
	} else {
		firstEdge := vertex.getFirstEdgeHead(h.edges)
		if firstEdge.hasPreviousEdgeHead() {
			lastEdge := firstEdge.getPreviousEdgeHead(h.edges)
			lastEdge.setNextEdgeHead(newEdge)
			newEdge.setPreviousEdgeHead(lastEdge)
			newEdge.setNextEdgeHead(firstEdge)
			firstEdge.setPreviousEdgeHead(newEdge)
			h.edges.update(lastEdge.toEdge())
		} else {
			firstEdge.setPreviousEdgeHead(newEdge)
			newEdge.setPreviousEdgeHead(firstEdge)
			newEdge.setNextEdgeHead(firstEdge)
		}
		h.edges.update(firstEdge.toEdge())
	}

	h.edges.update(newEdge.toEdge())
}

func (h headLoop) removeHead(head headVertex, edge edgeHead) {
	if head.isFirstEdgeHead(edge) {
		head.deleteFirstEdgeHead()
		return
	}
}
