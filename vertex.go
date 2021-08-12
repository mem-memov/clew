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

func (v vertex) hasFirstPositiveEdge() bool {
	return v.entry[firstPositive] != void
}

func (v vertex) hasFirstNegativeEdge() bool {
	return v.entry[firstNegative] != void
}

func (v vertex) getFirstPositiveEdge(edges edges) edge {
	return edges.read(v.entry[firstPositive])
}

func (v vertex) getFirstNegativeEdge(edges edges) edge {
	return edges.read(v.entry[firstNegative])
}

func (v vertex) setFirstPositiveEdge(edge edge) {
	v.entry[firstPositive] = edge.getPosition()
	v.entry[positiveCount]++
}

func (v vertex) setFirstNegativeEdge(edge edge) {
	v.entry[firstNegative] = edge.getPosition()
	v.entry[negativeCount]++
}

func (v vertex) deleteFirstPositiveEdge() {
	v.entry[firstNegative] = void
}

func (v vertex) update(entries entries) {
	entries.update(v.position, v.entry)
}

func (v vertex) append(entries entries) {
	entries.append(v.entry)
}
