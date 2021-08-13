package klubok

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

func newArrow(position position) arrow {
	// the first arrow of a node is connected to itself
	entry := newVoidEntry()
	entry[targetPosition] = void
	entry[previousHeadPosition] = position
	entry[nextHeadPosition] = position
	entry[sourcePosition] = void
	entry[previousTailPosition] = position
	entry[nextTailPosition] = position

	return arrow{position: position, entry: entry}
}

func existingArrow(position position, entry entry) arrow {
	return arrow{position: position, entry: entry}
}

func (e arrow) getPosition() position {
	return e.position
}

func (e arrow) toTail() tail {
	return tail(e)
}

func (e arrow) toHead() head {
	return head(e)
}

func (e arrow) update(entries entries) {
	entries.update(e.position, e.entry)
}

func (e arrow) append(entries entries) {
	entries.append(e.entry)
}
