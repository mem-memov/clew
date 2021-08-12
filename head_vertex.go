package klubok

type headVertex vertex

func (h headVertex) hasFirstEdgeHead() bool {
	return h.entry[firstNegative] != void
}

func (h headVertex) getFirstEdgeHead(edges edges) edgeHead {
	return edges.read(h.entry[firstNegative]).toHead()
}

func (h headVertex) setFirstEdgeHead(edgeHead edgeHead) {
	h.entry[firstNegative] = edgeHead.toEdge().getPosition()
	h.entry[negativeCount]++
}

func (h headVertex) isFirstEdgeHead(edgeHead edgeHead) bool {
	return h.entry[firstNegative] == edgeHead.toEdge().getPosition()
}

func (h headVertex) deleteFirstEdgeHead() {
	h.entry[firstNegative] = void
}

func (h headVertex) toVertex() vertex {
	return vertex(h)
}
