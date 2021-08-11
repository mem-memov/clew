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

func (e edge) hasPositiveVertex(v vertex) bool {
	return e.entry[positiveDirection] == v.getPosition()
}

func (e edge) setPositiveVertex(v vertex) {
	e.entry[positiveDirection] = v.getPosition()
}

func (e edge) setNegativeVertex(v vertex) {
	e.entry[negativeDirection] = v.getPosition()
}

func (e edge) getPositiveVertex(vertices vertices) vertex {
	return vertices.read(e.entry[positiveDirection])
}

func (e edge) getNextPositiveEdge(edges edges) edge {
	return edges.read(e.entry[positiveNext])
}

func (e edge) getNextNegativeEdge(edges edges) edge {
	return edges.read(e.entry[negativeNext])
}

func (e edge) getPreviousPositiveEdge(edges edges) edge {
	return edges.read(e.entry[positivePrevious])
}

func (e edge) getPreviousNegativeEdge(edges edges) edge {
	return edges.read(e.entry[negativePrevious])
}

func (e edge) hasPreviousPositiveEdge() bool {
	return e.entry[positivePrevious] != void
}

func (e edge) hasPreviousNegativeEdge() bool {
	return e.entry[negativePrevious] != void
}

func (e edge) setPreviousPositiveEdge(edge edge) {
	e.entry[positivePrevious] = edge.getPosition()
}

func (e edge) setPreviousNegativeEdge(edge edge) {
	e.entry[negativePrevious] = edge.getPosition()
}

func (e edge) setNextPositiveEdge(edge edge) {
	e.entry[positiveNext] = edge.getPosition()
}

func (e edge) setNextNegativeEdge(edge edge) {
	e.entry[negativeNext] = edge.getPosition()
}

func (e edge) update(entries entries) {
	entries.update(e.position, e.entry)
}

func (e edge) append(entries entries) {
	entries.append(e.entry)
}
