package clew

const (
	targetPosition       position = 0
	previousHeadPosition position = 1
	nextHeadPosition     position = 2
	sourcePosition       position = 3
	previousTailPosition position = 4
	nextTailPosition     position = 5
)

type arrow struct {
	position position
	entry    entry
}

func newArrow(position position, source source, target target) arrow {
	entry := newVoidEntry()
	entry[targetPosition] = target.toNode().getPosition()
	entry[sourcePosition] = source.toNode().getPosition()
	return arrow{position: position, entry: entry}
}

func existingArrow(position position, entry entry) arrow {
	return arrow{position: position, entry: entry}
}

func (a arrow) isValid() bool {
	return a.entry[targetPosition] != a.position
}

func (a arrow) getPosition() position {
	return a.position
}

func (a arrow) toTail() tail {
	return tail(a)
}

func (a arrow) toHead() head {
	return head(a)
}

func (a arrow) update(entries *entries) error {
	return entries.update(a.position, a.entry)
}
