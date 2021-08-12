package klubok

type edgeTail edge

func (e edgeTail) setTailVertex(tailVertex tailVertex) {
	e.entry[negativeDirection] = tailVertex.toVertex().getPosition()
}

func (e edgeTail) getNextEdgeTail(edges edges) edgeTail {
	return edges.read(e.entry[negativeNext]).toTail()
}

func (e edgeTail) getPreviousEdgeTail(edges edges) edgeTail {
	return edges.read(e.entry[negativePrevious]).toTail()
}

func (e edgeTail) hasPreviousEdgeTail() bool {
	return e.entry[negativePrevious] != void
}

func (e edgeTail) setPreviousEdgeTail(edge edgeTail) {
	e.entry[negativePrevious] = edge.toEdge().getPosition()
}

func (e edgeTail) setNextEdgeTail(edge edgeTail) {
	e.entry[negativeNext] = edge.toEdge().getPosition()
}

func (e edgeTail) toEdge() edge {
	return edge(e)
}