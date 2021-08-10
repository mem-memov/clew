package klubok

const (
	identifier     position = 0
	previousVertex position = 1
	firstPositive  position = 2
	lastPositive   position = 3
	firstNegative  position = 4
	lastNegative   position = 5
)

type vertex struct {
	position position
	entry    entry
}

func newVertex(position position, previousVertexPosition position) vertex {
	entry := newEntry()
	entry[identifier] = position
	entry[previousVertex] = previousVertexPosition
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

func (v vertex) getFirstPositiveEdge(edges edges) edge {
	return edges.read(v.entry[firstPositive])
}

func (v vertex) hasFirstNegativeEdge() bool {
	return v.entry[firstNegative] != void
}

func (v vertex) getFirstNegativeEdge(edges edges) edge {
	return edges.read(v.entry[firstNegative])
}

func (v vertex) setPreviousPositiveEdge(edgePosition position) {
	v.entry[positivePrevious] = edgePosition
}

func (v vertex) setNextPositiveEdgeIfEmpty(edgePosition position) {
	if v.entry[positiveNext] == void {
		v.entry[positiveNext] = edgePosition
	}
}

func (v vertex) setPreviousNegativeEdge(edgePosition position) {
	v.entry[negativePrevious] = edgePosition
}

func (v vertex) setNextNegativeEdgeIfEmpty(edgePosition position) {
	if v.entry[negativeNext] == void {
		v.entry[negativeNext] = edgePosition
	}
}

func (v vertex) update(s entries) {
	s.update(v.entry[identifier], v.entry)
}
