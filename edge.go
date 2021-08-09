package klubok

const (
	positiveDirection position = 0
	positivePrevious  position = 1
	positiveNext      position = 2
	negativeDirection position = 3
	negativePrevious  position = 4
	negativeNext      position = 5
)

type edge entry

func newEdge(e entry) edge {
	return edge(e)
}

func newEmptyEdge() edge {
	return edge(entry{void, void, void, void, void, void})
}

func (e edge) setPositiveVertex(v vertex) {
	e[positiveDirection] = v.getPosition()
}

func (e edge) setNegativeVertex(v vertex) {
	e[negativeDirection] = v.getPosition()
}

func (e edge) update(s storage, p position) {
	s.update(p, entry(e))
}

func (e edge) append(s storage) position {
	return s.append(entry(e))
}
