package klubok

const (
	positiveDirection position = 0
	positivePrevious  position = 1
	positiveNext      position = 2
	negativeDirection position = 3
	negativePrevious  position = 4
	negativeNext      position = 5
)

type edge struct {
	position position
	entry    entry
}

func newEdge(position position) edge {
	// the first edge of a vertex is connected to itself
	entry := newVoidEntry()
	entry[positiveDirection] = void
	entry[positivePrevious] = position
	entry[positiveNext] = position
	entry[negativeDirection] = void
	entry[negativePrevious] = position
	entry[negativeNext] = position

	return edge{position: position, entry: entry}
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
