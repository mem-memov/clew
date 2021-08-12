package klubok

type tailVertex vertex

func (t tailVertex) hasFirstEdgeTail() bool {
	return t.entry[firstPositive] != void
}

func (t tailVertex) getFirstEdgeTail(edges edges) edgeTail {
	return edges.read(t.entry[firstPositive]).toTail()
}

func (t tailVertex) setFirstEdgeTail(edgeTail edgeTail) {
	t.entry[firstPositive] = edgeTail.toEdge().getPosition()
	t.entry[positiveCount]++
}

func (t tailVertex) isFirstEdgeTail(edgeTail edgeTail) bool {
	return t.entry[firstPositive] == edgeTail.toEdge().getPosition()
}

func (t tailVertex) deleteFirstEdgeTail() {
	t.entry[firstPositive] = void
}

func (t tailVertex) toVertex() vertex {
	return vertex(t)
}