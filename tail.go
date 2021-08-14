package clew

type tail arrow

func (t tail) isSame(other tail) bool {
	return t.toArrow().getPosition() == other.toArrow().getPosition()
}

func (t tail) hasSource(source source) bool {
	return t.entry[sourcePosition] == source.toNode().getPosition()
}

func (t tail) setSource(source source) {
	t.entry[sourcePosition] = source.toNode().getPosition()
}

func (t tail) hasNext() bool {
	return t.entry[nextTailPosition] != void
}

func (t tail) getNext(arrows arrows) tail {
	return arrows.read(t.entry[nextTailPosition]).toTail()
}

func (t tail) setNext(arrow tail) {
	t.entry[nextTailPosition] = arrow.toArrow().getPosition()
}

func (t tail) deleteNext() {
	t.entry[nextTailPosition] = void
}

func (t tail) hasPrevious() bool {
	return t.entry[previousTailPosition] != void
}

func (t tail) getPrevious(arrows arrows) tail {
	return arrows.read(t.entry[previousTailPosition]).toTail()
}

func (t tail) setPrevious(tail tail) {
	t.entry[previousTailPosition] = tail.toArrow().getPosition()
}

func (t tail) deletePrevious() {
	t.entry[previousTailPosition] = void
}

func (t tail) isSurrounded() bool {
	return t.entry[previousTailPosition] != void && t.entry[nextTailPosition] != void && t.entry[previousTailPosition] != t.entry[nextTailPosition]
}

func (t tail) isPaired() bool {
	return t.entry[previousTailPosition] != void && t.entry[nextTailPosition] != void && t.entry[previousTailPosition] == t.entry[nextTailPosition]
}

func (t tail) isAlone() bool {
	return t.entry[previousTailPosition] == void && t.entry[nextTailPosition] == void
}

func (t tail) bindNext(next tail, arrows arrows) tail {
	t.setNext(next)
	next.setPrevious(t)
	arrows.update(t.toArrow())
	arrows.update(next.toArrow())
	return next
}

func (t tail) toArrow() arrow {
	return arrow(t)
}

func (t tail) toHead() head {
	return t.toArrow().toHead()
}
