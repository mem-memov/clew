package klubok

type nodes struct {
	entries  entries
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

	var node node
	if n.holes.exist() {
		node = newNode(n.holes.consumeHole(), n.lastNode)
	} else {
		node = newNode(n.entries.next(), n.lastNode)
	}

	n.lastNode = node.getPosition()
	n.append(node)

	return node
}

func (n nodes) read(position position) node {
	return newNodeWithEntry(position, n.entries.read(position))
}

func (n nodes) append(node node) {
	node.append(n.entries)
}

func (n nodes) update(node node) {
	node.update(n.entries)
}
