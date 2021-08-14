package clew

type arrows struct {
	entries *entries
	holes   *holes
}

func newArrows(entries *entries, holes *holes) *arrows {
	return &arrows{
		entries: entries,
		holes:   holes,
	}
}

func (a *arrows) produceHole(arrow arrow) {
	a.holes.produceHole(arrow.getPosition())
}

func (a *arrows) create(source source, target target) arrow {

	var arrow arrow
	if a.holes.exist() {
		arrow = newArrow(a.holes.consumeHole(), source, target)
		arrow.update(a.entries)
	} else {
		arrow = newArrow(a.entries.next(), source, target)
		arrow.append(a.entries)
	}

	return arrow
}

func (a *arrows) read(position position) arrow {
	return existingArrow(position, a.entries.read(position))
}

func (a *arrows) update(arrow arrow) {
	arrow.update(a.entries)
}
