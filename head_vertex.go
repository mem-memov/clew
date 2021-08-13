package klubok

type headVertex vertex

func (h headVertex) hasFirstEdgeHead() bool {
	return h.entry[firstEdgeTailPosition] != void
}

func (h headVertex) getFirstEdgeHead(edges edges) edgeHead {
	return edges.read(h.entry[firstEdgeTailPosition]).toHead()
}

func (h headVertex) setFirstEdgeHead(edgeHead edgeHead) {
	h.entry[firstEdgeTailPosition] = edgeHead.toEdge().getPosition()
	h.entry[tailEdgeCountPosition]++
}

func (h headVertex) isFirstEdgeHead(edgeHead edgeHead) bool {
	return h.entry[firstEdgeTailPosition] == edgeHead.toEdge().getPosition()
}

func (h headVertex) deleteFirstEdgeHead() {
	h.entry[firstEdgeTailPosition] = void
}

func (h headVertex) toVertex() vertex {
	return vertex(h)
}
