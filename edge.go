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

func newEdge(position position, entry entry) edge {
	return edge{position: position, entry: entry}
}

func newEmptyEdge() edge {
	return edge{position: void, entry: entry{void, void, void, void, void, void}}
}

func (e edge) setPositiveVertex(v vertex) {
	e.entry[positiveDirection] = v.getPosition()
}

func (e edge) getPositiveVertex(vertices vertices) vertex {
	return vertices.read(e.entry[positiveDirection])
}

func (e edge) sendPositiveVertex(c chan<- uint) {
	c <- uint(e.entry[positiveDirection])
}

func (e edge) setNegativeVertex(v vertex) {
	e.entry[negativeDirection] = v.getPosition()
}

func (e edge) hasNextPositiveEdge() bool {
	return e.entry[positiveNext] != void
}

func (e edge) getNextPositiveEdge(edges edges) edge {
	return edges.read(e.entry[positiveNext])
}

func (e edge) hasNextNegativeEdge() bool {
	return e.entry[negativeNext] != void
}

func (e edge) getNextNegativeEdge(edges edges) edge {
	return edges.read(e.entry[negativeNext])
}

func (e edge) update(s entries, p position) {
	s.update(p, e.entry)
}

func (e edge) append(s entries) position {
	return s.append(e.entry)
}
