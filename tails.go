package klubok

type tails struct {
	vertices nodes
	arrows   arrows
}

func newTails(vertices nodes, arrows arrows) tails {
	return tails{vertices: vertices, arrows: arrows}
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
		if next.toArrow().getPosition() == first.toArrow().getPosition() {
			return tails
		}
		tails = append(tails, next.toArrow().getPosition())
	}
}

func (t tails) addTail(source source, new tail) {

	if !source.hasFirstTail() {

		source.setFirstTail(new)
		t.vertices.update(source.toVertex())
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

func (t tails) removeTail(source source, arrow tail) {
	if source.isFirstTail(arrow) {
		source.deleteFirstTail()
		return
	}

	//firstPositiveArrow := source.getFirstTail(t.arrows)
}
