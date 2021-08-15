package clew

type tail arrow

func (t tail) isSame(other tail) bool {
	return t.toArrow().getPosition() == other.toArrow().getPosition()
}

func (t tail) hasSource(source source) bool {
	return t.entry[sourcePosition] == source.toNode().getPosition()
}

func (t tail) getSourcePosition() position {
	return t.entry[sourcePosition]
}

func (t tail) getSource(nodes *nodes) (source, error) {
	node, err := nodes.read(t.entry[sourcePosition])
	if err != nil {
		return source{}, err
	}
	return node.toSource(), nil
}

func (t tail) hasNext() bool {
	return t.entry[nextTailPosition] != void
}

func (t tail) getNext(arrows *arrows) (tail, error) {
	arrow, err := arrows.read(t.entry[nextTailPosition])
	if err != nil {
		return tail{}, err
	}

	return arrow.toTail(), nil
}

func (t tail) setNext(arrow tail) tail {
	t.entry[nextTailPosition] = arrow.toArrow().getPosition()
	return t
}

func (t tail) deleteNext() tail {
	t.entry[nextTailPosition] = void
	return t
}

func (t tail) hasPrevious() bool {
	return t.entry[previousTailPosition] != void
}

func (t tail) getPrevious(arrows *arrows) (tail, error) {
	arrow, err := arrows.read(t.entry[previousTailPosition])
	if err != nil {
		return tail{}, err
	}

	return arrow.toTail(), nil
}

func (t tail) setPrevious(tail tail) tail {
	t.entry[previousTailPosition] = tail.toArrow().getPosition()
	return t
}

func (t tail) deletePrevious() tail {
	t.entry[previousTailPosition] = void
	return t
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

func (t tail) bindNext(next tail, arrows *arrows) (tail, error) {
	current := t.setNext(next)
	next = next.setPrevious(current)

	err := arrows.update(current.toArrow())
	if err != nil {
		return tail{}, err
	}

	err = arrows.update(next.toArrow())
	if err != nil {
		return tail{}, err
	}

	return next, nil
}

func (t tail) toArrow() arrow {
	return arrow(t)
}

func (t tail) toHead() head {
	return t.toArrow().toHead()
}
