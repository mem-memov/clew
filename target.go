package clew

type target node

func (t target) hasFirstHead() bool {
	return t.entry[firstHeadPosition] != void
}

func (t target) getFirstHead(arrows *arrows) (head, error) {
	node, err := arrows.read(t.entry[firstHeadPosition])
	if err != nil {
		return head{}, err
	}

	return node.toHead(), nil
}

func (t target) setFirstHead(arrowHead head) target {
	t.entry[firstHeadPosition] = arrowHead.toArrow().getPosition()
	return t
}

func (t target) isFirstHead(arrowHead head) bool {
	return t.entry[firstHeadPosition] == arrowHead.toArrow().getPosition()
}

func (t target) deleteFirstHead() target {
	t.entry[firstHeadPosition] = void
	return t
}

func (t target) incrementHeadCount() target {
	t.entry[headCountPosition]++
	return t
}

func (t target) decrementHeadCount() target {
	t.entry[headCountPosition]--
	return t
}

func (t target) toNode() node {
	return node(t)
}
