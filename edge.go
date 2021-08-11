package klubok
// first edge is connected to itself
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

func newVoidEdge() edge {
	return edge{position: void, entry: entry{void, void, void, void, void, void}}
}

func (e edge) atPosition(position position) bool {
	return e.position == position
}

func (e edge) getPosition() position {
	return e.position
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

func (e edge) sendPositiveVertex(headStream chan<- uint) {
	headStream <- uint(e.entry[positiveDirection])
}

func (e edge) sendNegativeVertex(tailStream chan<- uint) {
	tailStream <- uint(e.entry[negativeDirection])
}

func (e edge) getNextPositiveEdge(edges edges) edge {
	return edges.read(e.entry[positiveNext])
}

func (e edge) getNextNegativeEdge(edges edges) edge {
	return edges.read(e.entry[negativeNext])
}

func (e edge) hasDifferentPreviousPositiveEdge() bool {
	return e.entry[positivePrevious] != e.position
}

func (e edge) hasDifferentPreviousNegativeEdge() bool {
	return e.entry[negativePrevious] != e.position
}

func (e edge) getPreviousPositiveEdge(edges edges) edge {
	return edges.read(e.entry[positivePrevious])
}

func (e edge) getPreviousNegativeEdge(edges edges) edge {
	return edges.read(e.entry[negativePrevious])
}

func (e edge) update(entries entries) {
	entries.update(e.position, e.entry)
}

func (e edge) append(entries entries) {
	entries.append(e.entry)
}
