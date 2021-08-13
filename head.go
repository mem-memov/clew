package klubok

type head arrow

func (h head) hasTarget(target target) bool {
	return h.entry[targetPosition] == target.toVertex().getPosition()
}

func (h head) setTarget(target target) {
	h.entry[targetPosition] = target.toVertex().getPosition()
}

func (h head) getTarget(vertices nodes) target {
	return vertices.read(h.entry[targetPosition]).toTarget()
}

func (h head) getNext(arrows arrows) head {
	return arrows.read(h.entry[nextHeadPosition]).toHead()
}

func (h head) getPrevious(arrows arrows) head {
	return arrows.read(h.entry[previousHeadPosition]).toHead()
}

func (h head) hasPrevious() bool {
	return h.entry[previousHeadPosition] != void
}

func (h head) setPrevious(head head) {
	h.entry[previousHeadPosition] = head.toArrow().getPosition()
}

func (h head) setNext(head head) {
	h.entry[nextHeadPosition] = head.toArrow().getPosition()
}

func (h head) toArrow() arrow {
	return arrow(h)
}
