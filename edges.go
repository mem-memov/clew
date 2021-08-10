package klubok

type edges struct {
	entries entries
}

func newEdges(entries entries) edges {
	return edges{
		entries: entries,
	}
}

func (e edges) read(position position) edge {
	return newEdge(position, e.entries.read(position))
}

func (e edges) append(edge edge) position {
	return edge.append(e.entries)
}
