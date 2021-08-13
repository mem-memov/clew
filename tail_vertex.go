package klubok

type tailVertex vertex

func (t tailVertex) hasFirstEdgeTail() bool {
	return t.entry[firstEdgeHeadPosition] != void
}

func (t tailVertex) getFirstEdgeTail(edges edges) edgeTail {
	return edges.read(t.entry[firstEdgeHeadPosition]).toTail()
}

func (t tailVertex) setFirstEdgeTail(edgeTail edgeTail) {
	t.entry[firstEdgeHeadPosition] = edgeTail.toEdge().getPosition()
	t.entry[headEdgeCountPosition]++
}

func (t tailVertex) isFirstEdgeTail(edgeTail edgeTail) bool {
	return t.entry[firstEdgeHeadPosition] == edgeTail.toEdge().getPosition()
}

func (t tailVertex) deleteFirstEdgeTail() {
	t.entry[firstEdgeHeadPosition] = void
}

func (t tailVertex) toVertex() vertex {
	return vertex(t)
}
