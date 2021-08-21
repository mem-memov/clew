package clew

type mix struct {
	node   node
	nodes  *nodes
	arrows *arrows
	heads  *heads
	tails  *tails
}

func newMix(node node, nodes *nodes, arrows *arrows, heads *heads, tails *tails) mix {
	return mix{node: node, nodes: nodes, arrows: arrows, heads: heads, tails: tails}
}

func (m mix) getPosition() position {
	return m.node.getPosition()
}

func (m mix) readSources() ([]position, error) {
	return m.heads.readHeads(m.node.toTarget())
}

func (m mix) readTargets() ([]position, error) {
	return m.tails.readTails(m.node.toSource())
}

func (m mix) setReference(position position) error {
	reference, err := m.nodes.read(position)
	if err != nil {
		return err
	}

	reference = reference.setPreviousNode(m.node)
	m.node = m.node.setNextNode(reference)

	err = m.nodes.update(reference)
	if err != nil {
		return err
	}

	err = m.nodes.update(m.node)
	if err != nil {
		return err
	}

	return nil
}

func (m mix) getReference() (position, position, error) {

	previousNode, nextNode, err := m.node.getReference(m.nodes)
	if err != nil {
		return previousNode.getPosition(), nextNode.getPosition(), err
	}

	return previousNode.getPosition(), nextNode.getPosition(), nil
}

func (m mix) addTarget(position position) error {

	source := m.node.toSource()

	node, err := m.nodes.read(position)
	if err != nil {
		return err
	}

	target := node.toTarget()

	// optimizing next step
	if target.isSmaller(source) {
		return newMix(node, m.nodes, m.arrows, m.heads, m.tails).addSource(source.toNode().getPosition())
	}

	tails, err := m.tails.readTails(source)
	if err != nil {
		return err
	}

	for _, present := range tails {
		if present == position {
			return nil
		}
	}

	arrow, err := m.arrows.create(source, target)
	if err != nil {
		return err
	}

	source, arrow, err = m.tails.addTail(source, arrow.toTail())
	if err != nil {
		return err
	}

	if source.toNode().isSame(target.toNode()) {
		target = source.toTarget()
	}

	arrow, err = m.heads.addHead(target, arrow.toHead())
	if err != nil {
		return err
	}

	return m.arrows.update(arrow)
}

func (m mix) addSource(position position) error {

	target := m.node.toTarget()

	node, err := m.nodes.read(position)
	if err != nil {
		return err
	}

	source := node.toSource()

	// optimizing next step
	if source.isSmaller(target) {
		return newMix(node, m.nodes, m.arrows, m.heads, m.tails).addTarget(target.toNode().getPosition())
	}

	heads, err := m.heads.readHeads(target)
	if err != nil {
		return err
	}

	for _, present := range heads {
		if present == position {
			return nil
		}
	}

	arrow, err := m.arrows.create(source, target)
	if err != nil {
		return err
	}

	source, arrow, err = m.tails.addTail(source, arrow.toTail())
	if err != nil {
		return err
	}

	if source.toNode().isSame(target.toNode()) {
		target = source.toTarget()
	}

	arrow, err = m.heads.addHead(target, arrow.toHead())
	if err != nil {
		return err
	}

	return m.arrows.update(arrow)
}

func (m mix) removeTarget(position position) error {

	source := m.node.toSource()

	if !source.hasFirstTail() {
		return nil
	}

	node, err := m.nodes.read(position)
	if err != nil {
		return err
	}

	target := node.toTarget()

	if !target.hasFirstHead() {
		return nil
	}

	first, err := source.getFirstTail(m.arrows)
	if err != nil {
		return err
	}

	tail := first

	if tail.toHead().hasTarget(target) {
		err := m.heads.removeHead(target, tail.toHead())
		if err != nil {
			return err
		}

		err = m.tails.removeTail(source, tail)
		if err != nil {
			return err
		}

		err = m.arrows.produceHole(tail.toArrow())
		if err != nil {
			return err
		}
	}

	for {
		if !tail.hasNext() {
			return nil
		}

		tail, err = tail.getNext(m.arrows)
		if err != nil {
			return err
		}

		if tail.isSame(first) {
			return nil
		}

		if !tail.toHead().hasTarget(target) {
			continue
		}

		err := m.heads.removeHead(target, tail.toHead())
		if err != nil {
			return err
		}

		err = m.tails.removeTail(source, tail)
		if err != nil {
			return err
		}

		err = m.arrows.produceHole(tail.toArrow())
		if err != nil {
			return err
		}
	}
}

func (m mix) delete() error {
	err := m.tails.deleteSource(m.node.toSource(), m.heads)
	if err != nil {
		return err
	}

	err = m.heads.deleteTarget(m.node.toTarget(), m.tails)
	if err != nil {
		return err
	}

	err = m.nodes.produceHole(m.node)
	if err != nil {
		return err
	}

	return nil
}
