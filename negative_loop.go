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
