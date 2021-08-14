package clew

type heads struct {
	nodes  *nodes
	arrows *arrows
}

func newHeads(nodes *nodes, arrows *arrows) *heads {
	return &heads{nodes: nodes, arrows: arrows}
}

func (h *heads) readHeads(target target) []position {
	heads := make([]position, 0)

	if !target.hasFirstHead() {
		return heads
	}

	first := target.getFirstHead(h.arrows)
	next := first

	heads = append(heads, next.toArrow().getPosition())

	for {
		next = next.getNext(h.arrows)
		if next.isSame(first) {
			return heads
		}
		heads = append(heads, next.toArrow().getPosition())
	}
}

func (h *heads) addHead(target target, new head) {

	if !target.hasFirstHead() {

		target.setFirstHead(new)
		h.nodes.update(target.toNode())
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

func (h *heads) removeHead(target target, removed head) {

	first := target.getFirstHead(h.arrows)
	if first.isSame(removed) {
		if first.isSurrounded() {
			next := first.getPrevious(h.arrows).bindNext(first.getNext(h.arrows), h.arrows)
			target.setFirstHead(next)
		} else if first.isPaired() {
			second := first.getNext(h.arrows)
			second.deletePrevious()
			second.deleteNext()
			target.setFirstHead(second)
		} else if first.isAlone() {
			target.deleteFirstHead()
		}
		h.nodes.update(target.toNode())
		return
	}

	previous := first

	for {
		current := previous.getNext(h.arrows)
		if current.isSame(first) {
			return
		}
		if current.isSame(removed) {
			previous.bindNext(current.getNext(h.arrows), h.arrows)
			return
		}
		previous = current
	}
}

func (h *heads) deleteTarget(target target) {
	if !target.hasFirstHead() {
		return
	}

	first := target.getFirstHead(h.arrows)
	next := first
	h.arrows.produceHole(next.toArrow())

	for {
		next = next.getNext(h.arrows)
		if next.isSame(first) {
			return
		}
		h.arrows.produceHole(next.toArrow())
	}
}
