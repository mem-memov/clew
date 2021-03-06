package clew

type holes struct {
	lastHolePosition position
	entries          *entries
}

func newHoles(entries *entries, lastHole position) *holes {
	return &holes{
		lastHolePosition: lastHole,
		entries:          entries,
	}
}

func (h *holes) exist() bool {
	return h.lastHolePosition != void
}

func (h *holes) consumeHole() (position, error) {
	entry, err := h.entries.read(h.lastHolePosition)
	if err != nil {
		return void, err
	}
	lastHole := existingHole(h.lastHolePosition, entry)
	h.lastHolePosition = lastHole.getPreviousHolePosition()
	return lastHole.getPosition(), nil
}

func (h *holes) produceHole(position position) error {
	newHole := newHole(position, h.lastHolePosition)

	err := newHole.update(h.entries)
	if err != nil {
		return err
	}

	h.lastHolePosition = position

	return nil
}
