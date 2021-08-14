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

func (a *arrows) produceHole(arrow arrow) error {
	return a.holes.produceHole(arrow.getPosition())
}

func (a *arrows) create(source source, target target) (arrow, error) {

	if a.holes.exist() {
		position, err := a.holes.consumeHole()
		if err != nil {
			return arrow{}, err
		}

		arrow := newArrow(position, source, target)

		err = arrow.update(a.entries)
		if err != nil {
			return arrow, err
		}

		return arrow, nil
	}

	position, err := a.entries.create()
	if err != nil {
		return arrow{}, err
	}

	arrow := newArrow(position, source, target)

	err = arrow.update(a.entries)
	if err != nil {
		return arrow, err
	}

	return arrow, nil
}

func (a *arrows) read(position position) (arrow, error) {
	entry, err := a.entries.read(position)
	if err != nil {
		return arrow{}, nil
	}

	return existingArrow(position, entry), nil
}

func (a *arrows) update(arrow arrow) error {
	return arrow.update(a.entries)
}
