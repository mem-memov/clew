package klubok

const (
	headVertexPosition       position = 0
	previousEdgeHeadPosition position = 1
	nextEdgeHeadPosition     position = 2
	tailVertexPosition       position = 3
	previousEdgeTailPosition position = 4
	nextEdgeTailPosition     position = 5
)

type edge struct {
	position position
	entry    entry
}

func newEdge(newEdge position) edge {
	// the first edge of a vertex is connected to itself
	entry := newVoidEntry()
	entry[headVertexPosition] = void
	entry[previousEdgeHeadPosition] = newEdge
	entry[nextEdgeHeadPosition] = newEdge
	entry[tailVertexPosition] = void
	entry[previousEdgeTailPosition] = newEdge
	entry[nextEdgeTailPosition] = newEdge

	return edge{position: newEdge, entry: entry}
}

func existingEdge(position position, entry entry) edge {
	return edge{position: position, entry: entry}
}

func (e edge) getPosition() position {
	return e.position
}

func (e edge) toTail() edgeTail {
	return edgeTail(e)
}

func (e edge) toHead() edgeHead {
	return edgeHead(e)
}

func (e edge) update(entries entries) {
	entries.update(e.position, e.entry)
}

func (e edge) append(entries entries) {
	entries.append(e.entry)
}
