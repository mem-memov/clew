package clew

const (
	previousHolePosition position = 0
)

type hole struct {
	position position
	entry    entry
}

func newHole(position position, previous position) hole {
	entry := newVoidEntry()
	entry[previousHolePosition] = previous
	return hole{position: position, entry: newVoidEntry()}
}

func existingHole(position position, entry entry) hole {
	return hole{position: position, entry: entry}
}

func (h hole) getPosition() position {
	return h.entry[previousHolePosition]
}

func (h hole) getPreviousHolePosition() position {
	return h.entry[previousHolePosition]
}

func (h hole) update(entries *entries) {
	entries.update(h.position, h.entry)
}
