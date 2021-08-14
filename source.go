package klubok

type source node

func (s source) hasFirstTail() bool {
	return s.entry[firstTailPosition] != void
}

func (s source) getFirstTail(arrows arrows) tail {
	return arrows.read(s.entry[firstTailPosition]).toTail()
}

func (s source) setFirstTail(tail tail) {
	s.entry[firstTailPosition] = tail.toArrow().getPosition()
}

func (s source) isFirstTail(tail tail) bool {
	return s.entry[firstTailPosition] == tail.toArrow().getPosition()
}

func (s source) deleteFirstTail() {
	s.entry[firstTailPosition] = void
}

func (s source) toNode() node {
	return node(s)
}
