package clew

type head arrow

func (h head) isSame(other head) bool {
	return h.toArrow().getPosition() == other.toArrow().getPosition()
}

func (h head) hasTarget(target target) bool {
	return h.entry[targetPosition] == target.toNode().getPosition()
}

func (h head) getTarget(nodes *nodes) target {
	return nodes.read(h.entry[targetPosition]).toTarget()
}

func (h head) getTargetPosition() position {
	return h.entry[targetPosition]
}

func (h head) hasPrevious() bool {
	return h.entry[previousHeadPosition] != void
}

func (h head) getPrevious(arrows *arrows) head {
	return arrows.read(h.entry[previousHeadPosition]).toHead()
}

func (h head) setPrevious(head head) head {
	h.entry[previousHeadPosition] = head.toArrow().getPosition()
	return h
}

func (h head) deletePrevious() head {
	h.entry[previousHeadPosition] = void
	return h
}

func (h head) hasNext() bool {
	return h.entry[nextHeadPosition] != void
}

func (h head) getNext(arrows *arrows) head {
	return arrows.read(h.entry[nextHeadPosition]).toHead()
}

func (h head) setNext(head head) head {
	h.entry[nextHeadPosition] = head.toArrow().getPosition()
	return h
}

func (h head) deleteNext() head {
	h.entry[nextHeadPosition] = void
	return h
}

func (h head) isSurrounded() bool {
	return h.entry[previousHeadPosition] != void && h.entry[nextHeadPosition] != void && h.entry[previousHeadPosition] != h.entry[nextHeadPosition]
}

func (h head) isPaired() bool {
	return h.entry[previousHeadPosition] != void && h.entry[nextHeadPosition] != void && h.entry[previousHeadPosition] == h.entry[nextHeadPosition]
}

func (h head) isAlone() bool {
	return h.entry[previousHeadPosition] == void && h.entry[nextHeadPosition] == void
}

func (h head) bindNext(next head, arrows *arrows) head {
	current := h.setNext(next)
	next = next.setPrevious(current)
	arrows.update(current.toArrow())
	arrows.update(next.toArrow())
	return next
}

func (h head) toArrow() arrow {
	return arrow(h)
}

func (h head) toTail() tail {
	return h.toArrow().toTail()
}
