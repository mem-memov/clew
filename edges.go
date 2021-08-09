package klubok

type edges struct {
	entries storage
}

func newEdges(s storage) edges {
	return edges{
		entries: s,
	}
}

func (e edges) append(edge edge) position {
	return edge.append(e.entries)
}