package klubok

type nodes struct {
	entries    entries
	holes      holes
	lastVertex position
}

func newNodes(entries entries, holes holes, lastVertex position) nodes {
	return nodes{
		entries:    entries,
		holes:      holes,
		lastVertex: lastVertex,
	}
}

func (n nodes) create() node {

	var vertex node
	if n.holes.exist() {
		vertex = newNode(n.holes.last(), n.lastVertex)
		n.holes.consumeHole(vertex)
	} else {
		vertex = newNode(n.entries.next(), n.lastVertex)
	}

	n.lastVertex = vertex.getPosition()

	return vertex
}

func (n nodes) produceHole(vertex node) {
	n.holes.produceHole(vertex)
}

func (n nodes) read(position position) node {
	return newVertexWithEntry(position, n.entries.read(position))
}

func (n nodes) update(vertex node) {
	vertex.update(n.entries)
}
