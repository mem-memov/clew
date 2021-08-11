package klubok

type edges struct {
	entries entries
	holes   holes
}

func newEdges(entries entries, holes holes) edges {
	return edges{
		entries: entries,
		holes:   holes,
	}
}

func (e edges) produceHole(edge edge) {
	e.holes.produceHole(edge)
}

func (e edges) create() edge {

	var edge edge
	if e.holes.exist() {
		edge = newEdge(e.holes.last())
		e.holes.consumeHole(edge)
	} else {
		edge = newEdge(e.entries.next())
	}

	return edge
}

func (e edges) read(position position) edge {
	return existingEdge(position, e.entries.read(position))
}

func (e edges) append(edge edge) {
	edge.append(e.entries)
}

func (e edges) update(edge edge) {
	edge.update(e.entries)
}
