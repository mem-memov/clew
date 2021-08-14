package clew

const (
	identifierPosition   position = 0
	previousNodePosition position = 1
	headCountPosition    position = 2
	firstHeadPosition    position = 3
	tailCountPosition    position = 4
	firstTailPosition    position = 5
)

type node struct {
	position position
	entry    entry
}

func newNode(new position, previous position) node {
	entry := newVoidEntry()
	entry[identifierPosition] = new
	entry[previousNodePosition] = previous
	entry[headCountPosition] = 0
	entry[firstHeadPosition] = void
	entry[tailCountPosition] = 0
	entry[firstTailPosition] = void

	return node{position: new, entry: entry}
}

func newNodeWithEntry(position position, entry entry) node {
	return node{position: position, entry: entry}
}

func (n node) isValid() bool {
	return n.entry[identifierPosition] == n.position
}

func (n node) isSame(node node) bool {
	return n.entry[identifierPosition] == node.getPosition()
}

func (n node) getPosition() position {
	return n.entry[identifierPosition]
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
