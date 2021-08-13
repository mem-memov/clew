package klubok

type arrows struct {
	entries entries
	holes   holes
}

func newArrows(entries entries, holes holes) arrows {
	return arrows{
		entries: entries,
		holes:   holes,
	}
}

func (e arrows) produceHole(arrow arrow) {
	e.holes.produceHole(arrow)
}

func (e arrows) create() arrow {

	var arrow arrow
	if e.holes.exist() {
		arrow = newArrow(e.holes.last())
		e.holes.consumeHole(arrow)
	} else {
		arrow = newArrow(e.entries.next())
	}

	return arrow
}

func (e arrows) read(position position) arrow {
	return existingArrow(position, e.entries.read(position))
}

func (e arrows) append(arrow arrow) {
	arrow.append(e.entries)
}

func (e arrows) update(arrow arrow) {
	arrow.update(e.entries)
}
