package klubok

type nodes struct {
	entries    entries
	holes    holes
	lastNode position
}

func newNodes(entries entries, holes holes, lastNode position) nodes {
	return nodes{
		entries:  entries,
		holes:    holes,
		lastNode: lastNode,
	}
}

func (n nodes) produceHole(node node) {
	n.holes.produceHole(node.getPosition())
}

func (n nodes) create() node {

	var vertex node
	if n.holes.exist() {
		vertex = newNode(n.holes.consumeHole(), n.lastNode)
	} else {
		vertex = newNode(n.entries.next(), n.lastNode)
	}

	n.lastNode = vertex.getPosition()

	return vertex
}

func (n nodes) read(position position) node {
	return newVertexWithEntry(position, n.entries.read(position))
}

func (n nodes) update(node node) {
	node.update(n.entries)
}
