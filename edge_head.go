package klubok

type edgeHead edge

func (e edgeHead) hasHeadVertex(headVertex headVertex) bool {
	return e.entry[headVertexPosition] == headVertex.toVertex().getPosition()
}

func (e edgeHead) setHeadVertex(headVertex headVertex) {
	e.entry[headVertexPosition] = headVertex.toVertex().getPosition()
}

func (e edgeHead) getHeadVertex(vertices vertices) headVertex {
	return vertices.read(e.entry[headVertexPosition]).toHead()
}

func (e edgeHead) getNextEdgeHead(edges edges) edgeHead {
	return edges.read(e.entry[nextEdgeHeadPosition]).toHead()
}

func (e edgeHead) getPreviousEdgeHead(edges edges) edgeHead {
	return edges.read(e.entry[previousEdgeHeadPosition]).toHead()
}

func (e edgeHead) hasPreviousEdgeHead() bool {
	return e.entry[previousEdgeHeadPosition] != void
}

func (e edgeHead) setPreviousEdgeHead(edgeHead edgeHead) {
	e.entry[previousEdgeHeadPosition] = edgeHead.toEdge().getPosition()
}

func (e edgeHead) setNextEdgeHead(edgeHead edgeHead) {
	e.entry[nextEdgeHeadPosition] = edgeHead.toEdge().getPosition()
}

func (e edgeHead) toEdge() edge {
	return edge(e)
}
