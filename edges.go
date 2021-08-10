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

func (e edges) read(position position) edge {
	return newEdge(position, e.entries.read(position))
}

func (e edges) append(edge edge) position {
	return edge.append(e.entries)
}
