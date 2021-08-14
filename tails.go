package klubok

type tails struct {
	nodes  nodes
	arrows arrows
}

func newTails(vertices nodes, arrows arrows) tails {
	return tails{nodes: vertices, arrows: arrows}
}

func (t tails) readTails(source source) []position {
	tails := make([]position, 0)

	if !source.hasFirstTail() {
		return tails
	}

	first := source.getFirstTail(t.arrows)
	next := first

	tails = append(tails, next.toArrow().getPosition())

	for {
		next = next.getNext(t.arrows)
		if next.isSame(first) {
			return tails
		}
		tails = append(tails, next.toArrow().getPosition())
	}
}

func (t tails) addTail(source source, new tail) {

	if !source.hasFirstTail() {

		source.setFirstTail(new)
		t.nodes.update(source.toVertex())
	} else {

		first := source.getFirstTail(t.arrows)
		if first.hasPrevious() {

			last := first.getPrevious(t.arrows)
			last.setNext(new)
			new.setPrevious(last)
			new.setNext(first)
			first.setPrevious(new)
			t.arrows.update(last.toArrow())
		} else {

			first.setPrevious(new)
			new.setPrevious(first)
			new.setNext(first)
		}
		t.arrows.update(first.toArrow())
	}

	t.arrows.update(new.toArrow())
}

func (t tails) removeTail(source source, removed tail) {

	first := source.getFirstTail(t.arrows)
	if first.isSame(removed) {
		if first.isSurrounded() {
			next := first.getPrevious(t.arrows).bindNext(first.getNext(t.arrows), t.arrows)
			source.setFirstTail(next)
		} else if first.isPaired() {
			second := first.getNext(t.arrows)
			second.deletePrevious()
			second.deleteNext()
			source.setFirstTail(second)
		} else if first.isAlone() {
			source.deleteFirstTail()
		}
		t.nodes.update(source.toVertex())
		return
	}

	previous := first

	for {
		current := previous.getNext(t.arrows)
		if current.isSame(first) {
			return
		}
		if current.isSame(removed) {
			previous.bindNext(current.getNext(t.arrows), t.arrows)
			return
		}
		previous = current
	}
}

func (t tails) deleteSource(source source) {
	if !source.hasFirstTail() {
		return
	}

	first := source.getFirstTail(t.arrows)
	next := first
	t.arrows.produceHole(next.toArrow())

	for {
		next = next.getNext(t.arrows)
		if next.isSame(first) {
			return
		}
		t.arrows.produceHole(next.toArrow())
	}
}
