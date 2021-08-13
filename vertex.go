package klubok

const (
	identifierPosition     position = 0
	previousVertexPosition position = 1
	headEdgeCountPosition  position = 2
	firstEdgeHeadPosition  position = 3
	tailEdgeCountPosition  position = 4
	firstEdgeTailPosition  position = 5
)

type vertex struct {
	position position
	entry    entry
}

func newVertex(newVertex position, previousVertex position) vertex {
	entry := newVoidEntry()
	entry[identifierPosition] = newVertex
	entry[previousVertexPosition] = previousVertex
	entry[headEdgeCountPosition] = 0
	entry[firstEdgeHeadPosition] = void
	entry[tailEdgeCountPosition] = 0
	entry[firstEdgeTailPosition] = void

	return vertex{position: newVertex, entry: entry}
}

func newVertexWithEntry(position position, entry entry) vertex {
	return vertex{position: position, entry: entry}
}

func (v vertex) getPosition() position {
	return v.entry[identifierPosition]
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
