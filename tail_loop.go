package klubok

type tailLoop struct {
	vertices vertices
	edges    edges
}

func newNegativeLoop(vertices vertices, edges edges) tailLoop {
	return tailLoop{vertices: vertices, edges: edges}
}

func (n tailLoop) readTails(head vertex) []position {
	tails := make([]position, 0)

	if !head.hasFirstNegativeEdge() {
		return tails
	}

	firstEdge := head.getFirstNegativeEdge(n.edges)
	nextEdge := firstEdge

	tails = append(tails, nextEdge.getPosition())

	for {
		nextEdge = nextEdge.getNextNegativeEdge(n.edges)
		if nextEdge.getPosition() == firstEdge.getPosition() {
			return tails
		}
		tails = append(tails, nextEdge.getPosition())
	}
}

func (t tailLoop) addTail(vertex tailVertex, newEdge edgeTail) {

	if !vertex.hasFirstEdgeTail() {
		vertex.setFirstEdgeTail(newEdge)
		t.vertices.update(vertex.toVertex())
	} else {
		firstEdge := vertex.getFirstEdgeTail(t.edges)
		if firstEdge.hasPreviousEdgeTail() {
			lastEdge := firstEdge.getPreviousEdgeTail(t.edges)
			lastEdge.setNextEdgeTail(newEdge)
			newEdge.setPreviousEdgeTail(lastEdge)
			newEdge.setNextEdgeTail(firstEdge)
			firstEdge.setPreviousEdgeTail(newEdge)
			t.edges.update(lastEdge.toEdge())
		} else {
			firstEdge.setPreviousEdgeTail(newEdge)
			newEdge.setPreviousEdgeTail(firstEdge)
			newEdge.setNextEdgeTail(firstEdge)
		}
		t.edges.update(firstEdge.toEdge())
	}

	t.edges.update(newEdge.toEdge())
}

func (n tailLoop) removeTail(head headVertex, edge edge) {
	if head.isFirstNegative(edge) {
		head.deleteFirstEdgeHead()
		return
	}
}
