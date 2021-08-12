package klubok

const (
	identifier     position = 0
	previousVertex position = 1
	positiveCount  position = 2
	firstPositive  position = 3
	negativeCount  position = 4
	firstNegative  position = 5
)

type vertex struct {
	position position
	entry    entry
}

func newVertex(position position, previousVertexPosition position) vertex {
	entry := newVoidEntry()
	entry[identifier] = position
	entry[previousVertex] = previousVertexPosition
	entry[positiveCount] = 0
	entry[firstPositive] = void
	entry[negativeCount] = 0
	entry[firstNegative] = void

	return vertex{position: position, entry: entry}
}

func newVertexWithEntry(position position, entry entry) vertex {
	return vertex{position: position, entry: entry}
}

func (v vertex) getPosition() position {
	return v.entry[identifier]
}

func (v vertex) toTail() tailVertex {
	return tailVertex(v)
}

func (v vertex) toHead() headVertex {
	return headVertex(v)
}

func (v vertex) update(entries entries) {
	entries.update(v.position, v.entry)
}

func (v vertex) append(entries entries) {
	entries.append(v.entry)
}
