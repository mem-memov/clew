package klubok

type tail arrow

func (t tail) hasSource(source source) bool {
	return t.entry[sourcePosition] == source.toVertex().getPosition()
}

func (t tail) setSource(source source) {
	t.entry[sourcePosition] = source.toVertex().getPosition()
}

func (t tail) getNext(arrows arrows) tail {
	return arrows.read(t.entry[nextTailPosition]).toTail()
}

func (t tail) getPrevious(arrows arrows) tail {
	return arrows.read(t.entry[previousTailPosition]).toTail()
}

func (t tail) hasPrevious() bool {
	return t.entry[previousTailPosition] != void
}

func (t tail) setPrevious(tail tail) {
	t.entry[previousTailPosition] = tail.toArrow().getPosition()
}

func (t tail) setNext(arrow tail) {
	t.entry[nextTailPosition] = arrow.toArrow().getPosition()
}

func (t tail) toArrow() arrow {
	return arrow(t)
}
