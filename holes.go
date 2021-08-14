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

func (h holes) exist() bool {
	return h.lastHolePosition != void
}

func (h holes) consumeHole() position {
	lastHole := existingHole(h.lastHolePosition, h.entries.read(h.lastHolePosition))
	h.lastHolePosition = lastHole.getPreviousHolePosition()
	return lastHole.getPosition()
}

func (h holes) produceHole(position position) {
	newHole := newHole(position, h.lastHolePosition)
	newHole.update(h.entries)
	h.lastHolePosition = position
}
