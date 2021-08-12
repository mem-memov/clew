package klubok

type negativeLoop struct {
	vertices vertices
	edges    edges
}

func newNegativeLoop(vertices vertices, edges edges) negativeLoop {
	return negativeLoop{vertices: vertices, edges: edges}
}

func (n negativeLoop) readTails(head vertex) []position {
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

func (n negativeLoop) addTail(head vertex, edge edge) {

	if !head.hasFirstNegativeEdge() {
		head.setFirstNegativeEdge(edge)
		n.vertices.update(head)
	} else {
		firstNegativeEdge := head.getFirstNegativeEdge(n.edges)
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
}

func (n negativeLoop) removeTail(head vertex, edge edge) {

}
