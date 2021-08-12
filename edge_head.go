package klubok

type edgeHead edge

func (e edgeHead) hasHeadVertex(headVertex headVertex) bool {
	return e.entry[positiveDirection] == headVertex.toVertex().getPosition()
}

func (e edgeHead) setHeadVertex(headVertex headVertex) {
	e.entry[positiveDirection] = headVertex.toVertex().getPosition()
}

func (e edgeHead) getHeadVertex(vertices vertices) headVertex {
	return vertices.read(e.entry[positiveDirection]).toHead()
}

func (e edgeHead) getNextEdgeHead(edges edges) edgeHead {
	return edges.read(e.entry[positiveNext]).toHead()
}

func (e edgeHead) getPreviousEdgeHead(edges edges) edgeHead {
	return edges.read(e.entry[positivePrevious]).toHead()
}

func (e edgeHead) hasPreviousEdgeHead() bool {
	return e.entry[positivePrevious] != void
}

func (e edgeHead) setPreviousEdgeHead(edgeHead edgeHead) {
	e.entry[positivePrevious] = edgeHead.toEdge().getPosition()
}

func (e edgeHead) setNextEdgeHead(edgeHead edgeHead) {
	e.entry[positiveNext] = edgeHead.toEdge().getPosition()
}

func (e edgeHead) toEdge() edge {
	return edge(e)
}