package clew

const (
	previousNodePosition position = 0
	nextNodePosition     position = 1
	headCountPosition    position = 2
	tailCountPosition    position = 3
	firstHeadPosition    position = 4
	firstTailPosition    position = 5
)

type node struct {
	position position
	entry    entry
}

func newNode(position position) node {
	entry := newVoidEntry()
	entry[previousNodePosition] = 0
	entry[nextNodePosition] = 0
	entry[headCountPosition] = 0
	entry[tailCountPosition] = 0
	entry[firstHeadPosition] = void
	entry[firstTailPosition] = void

	return node{position: position, entry: entry}
}

func newNodeWithEntry(position position, entry entry) node {
	return node{position: position, entry: entry}
}

func (n node) isValid() bool {
	return n.position == n.position
}

func (n node) isSame(node node) bool {
	return n.position == node.getPosition()
}

func (n node) getPosition() position {
	return n.position
}

func (n node) setPreviousNode(node node) node {
	n.entry[previousNodePosition] = node.getPosition()
	return n
}

func (n node) setNextNode(node node) node {
	n.entry[nextNodePosition] = node.getPosition()
	return n
}

func (n node) getReference(nodes *nodes) (node, node, error) {

	previousNode, err := nodes.read(n.entry[previousNodePosition])
	if err != nil {
		return node{}, node{}, err
	}

	nextNode, err := nodes.read(n.entry[nextNodePosition])
	if err != nil {
		return node{}, node{}, err
	}

	return previousNode,nextNode, nil
}

func (n node) deletePreviousReference() node {
	n.entry[previousNodePosition] = void
	return n
}

func (n node) deleteNextReference() node {
	n.entry[nextNodePosition] = void
	return n
}

func (n node) toSource() source {
	return source(n)
}

func (n node) toTarget() target {
	return target(n)
}

func (n node) update(entries *entries) error {
	return entries.update(n.position, n.entry)
}
