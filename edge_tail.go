package klubok

type edgeTail edge

func (e edgeTail) hasTailVertex(tailVertex tailVertex) bool {
	return e.entry[tailVertexPosition] == tailVertex.toVertex().getPosition()
}

func (e edgeTail) setTailVertex(tailVertex tailVertex) {
	e.entry[tailVertexPosition] = tailVertex.toVertex().getPosition()
}

func (e edgeTail) getNextEdgeTail(edges edges) edgeTail {
	return edges.read(e.entry[nextEdgeTailPosition]).toTail()
}

func (e edgeTail) getPreviousEdgeTail(edges edges) edgeTail {
	return edges.read(e.entry[previousEdgeTailPosition]).toTail()
}

func (e edgeTail) hasPreviousEdgeTail() bool {
	return e.entry[previousEdgeTailPosition] != void
}

func (e edgeTail) setPreviousEdgeTail(edge edgeTail) {
	e.entry[previousEdgeTailPosition] = edge.toEdge().getPosition()
}

func (e edgeTail) setNextEdgeTail(edge edgeTail) {
	e.entry[nextEdgeTailPosition] = edge.toEdge().getPosition()
}

func (e edgeTail) toEdge() edge {
	return edge(e)
}
