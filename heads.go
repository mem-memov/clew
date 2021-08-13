package klubok

type heads struct {
	vertices nodes
	arrows   arrows
}

func newHeads(vertices nodes, arrows arrows) heads {
	return heads{vertices: vertices, arrows: arrows}
}

func (h heads) readHeads(target target) []position {
	heads := make([]position, 0)

	if !target.hasFirstHead() {
		return heads
	}

	first := target.getFirstHead(h.arrows)
	next := first

	heads = append(heads, next.toArrow().getPosition())

	for {
		next = next.getNext(h.arrows)
		if next.toArrow().getPosition() == first.toArrow().getPosition() {
			return heads
		}
		heads = append(heads, next.toArrow().getPosition())
	}
}

func (h heads) addHead(target target, new head) {

	if !target.hasFirstHead() {

		target.setFirstHead(new)
		h.vertices.update(target.toVertex())
	} else {

		first := target.getFirstHead(h.arrows)
		if first.hasPrevious() {

			last := first.getPrevious(h.arrows)
			last.setNext(new)
			new.setPrevious(last)
			new.setNext(first)
			first.setPrevious(new)
			h.arrows.update(last.toArrow())
		} else {

			first.setPrevious(new)
			new.setPrevious(first)
			new.setNext(first)
		}
		h.arrows.update(first.toArrow())
	}

	h.arrows.update(new.toArrow())
}

func (h heads) removeHead(head target, arrow head) {
	if head.isFirstHead(arrow) {
		head.deleteFirstHead()
		return
	}
}
