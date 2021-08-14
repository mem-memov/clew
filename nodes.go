package clew

import "fmt"

type nodes struct {
	entries  *entries
	holes    *holes
	lastNode position
}

func newNodes(entries *entries, holes *holes, lastNode position) *nodes {
	return &nodes{
		entries:  entries,
		holes:    holes,
		lastNode: lastNode,
	}
}

func (n *nodes) produceHole(node node) {
	n.holes.produceHole(node.getPosition())
}

func (n *nodes) create() (node, error) {

	if n.holes.exist() {
		position, err := n.holes.consumeHole()
		if err != nil {
			return node{}, err
		}

		node := newNode(position, n.lastNode)

		err = node.update(n.entries)
		if err != nil {
			return node, err
		}

		n.lastNode = node.getPosition()

		return node, nil
	}

	position, err := n.entries.next()
	if err != nil {
		return node{}, err
	}

	node := newNode(position, n.lastNode)

	err = node.append(n.entries)
	if err != nil {
		return node, err
	}

	n.lastNode = node.getPosition()

	return node, nil
}

func (n *nodes) read(position position) (node, error) {
	entry, err := n.entries.read(position)
	if err != nil {
		return node{}, err
	}

	node := newNodeWithEntry(position, entry)

	if !node.isValid() {
		return node, fmt.Errorf("node invalid %v", node)
	}

	return node, nil
}

func (n *nodes) update(node node) error {
	return node.update(n.entries)
}
