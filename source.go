package clew

type source node

func (s source) hasFirstTail() bool {
	return s.entry[firstTailPosition] != void
}

func (s source) getFirstTail(arrows *arrows) tail {
	return arrows.read(s.entry[firstTailPosition]).toTail()
}

func (s source) setFirstTail(tail tail) source {
	s.entry[firstTailPosition] = tail.toArrow().getPosition()
	return s
}

func (s source) isFirstTail(tail tail) bool {
	return s.entry[firstTailPosition] == tail.toArrow().getPosition()
}

func (s source) deleteFirstTail() source {
	s.entry[firstTailPosition] = void
	return s
}

func (s source) toNode() node {
	return node(s)
}
