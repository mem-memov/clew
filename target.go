package klubok

type target node

func (t target) hasFirstHead() bool {
	return t.entry[firstHeadPosition] != void
}

func (t target) getFirstHead(arrows arrows) head {
	return arrows.read(t.entry[firstHeadPosition]).toHead()
}

func (t target) setFirstHead(arrowHead head) {
	t.entry[firstHeadPosition] = arrowHead.toArrow().getPosition()
}

func (t target) isFirstHead(arrowHead head) bool {
	return t.entry[firstHeadPosition] == arrowHead.toArrow().getPosition()
}

func (t target) deleteFirstHead() {
	t.entry[firstHeadPosition] = void
}

func (t target) toNode() node {
	return node(t)
}
