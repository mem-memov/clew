package klubok

const (
	identifier     position = 0
	previousVertex position = 1
	firstPositive  position = 2
	lastPositive   position = 3
	firstNegative  position = 4
	lastNegative   position = 5
)

type vertex entry

func newVertex(e entry) vertex {
	return vertex(e)
}

func (v vertex) getPosition() position {
	return v[identifier]
}

func (v vertex) setPreviousPositiveEdge(edgePosition position) {
	v[positivePrevious] = edgePosition
}

func (v vertex) setNextPositiveEdgeIfEmpty(edgePosition position) {
	if v[positiveNext] == void {
		v[positiveNext] = edgePosition
	}
}

func (v vertex) setPreviousNegativeEdge(edgePosition position) {
	v[negativePrevious] = edgePosition
}

func (v vertex) setNextNegativeEdgeIfEmpty(edgePosition position) {
	if v[negativeNext] == void {
		v[negativeNext] = edgePosition
	}
}

func (v vertex) update(s storage) {
	s.update(v[identifier], entry(v))
}
