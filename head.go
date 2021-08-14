package klubok

type head arrow

func (h head) isSame(other head) bool {
	return h.toArrow().getPosition() == other.toArrow().getPosition()
}

func (h head) hasTarget(target target) bool {
	return h.entry[targetPosition] == target.toVertex().getPosition()
}

func (h head) setTarget(target target) {
	h.entry[targetPosition] = target.toVertex().getPosition()
}

func (h head) getTarget(vertices nodes) target {
	return vertices.read(h.entry[targetPosition]).toTarget()
}

func (h head) hasPrevious() bool {
	return h.entry[previousHeadPosition] != void
}

func (h head) getPrevious(arrows arrows) head {
	return arrows.read(h.entry[previousHeadPosition]).toHead()
}

func (h head) setPrevious(head head) {
	h.entry[previousHeadPosition] = head.toArrow().getPosition()
}

func (h head) deletePrevious() {
	h.entry[previousHeadPosition] = void
}

func (h head) hasNext() bool {
	return h.entry[nextHeadPosition] != void
}

func (h head) getNext(arrows arrows) head {
	return arrows.read(h.entry[nextHeadPosition]).toHead()
}

func (h head) setNext(head head) {
	h.entry[nextHeadPosition] = head.toArrow().getPosition()
}

func (h head) deleteNext() {
	h.entry[nextHeadPosition] = void
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

func (h head) bindNext(next head, arrows arrows) head {
	h.setNext(next)
	next.setPrevious(h)
	arrows.update(h.toArrow())
	arrows.update(next.toArrow())
	return next
}

func (h head) toArrow() arrow {
	return arrow(h)
}

func (h head) toTail() tail {
	return h.toArrow().toTail()
}
