package klubok

type tailLoop struct {
	vertices vertices
	edges    edges
}

func newTailLoop(vertices vertices, edges edges) tailLoop {
	return tailLoop{vertices: vertices, edges: edges}
}

func (t tailLoop) readTails(vertex tailVertex) []position {
	tails := make([]position, 0)

	if !vertex.hasFirstEdgeTail() {
		return tails
	}

	firstEdge := vertex.getFirstEdgeTail(t.edges)
	nextEdge := firstEdge

	tails = append(tails, nextEdge.toEdge().getPosition())

	for {
		nextEdge = nextEdge.getNextEdgeTail(t.edges)
		if nextEdge.toEdge().getPosition() == firstEdge.toEdge().getPosition() {
			return tails
		}
		tails = append(tails, nextEdge.toEdge().getPosition())
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

func (t tailLoop) removeTail(tail tailVertex, edge edgeTail) {
	if tail.isFirstEdgeTail(edge) {
		tail.deleteFirstEdgeTail()
		return
	}

	//firstPositiveEdge := tail.getFirstEdgeTail(t.edges)
}
