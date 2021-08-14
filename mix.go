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

func (m mix) addTarget(position position) error {

	source := m.node.toSource()

	node, err := m.nodes.read(position)
	if err != nil {
		return err
	}

	target := node.toTarget()

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

	arrow, err = m.tails.addTail(source, arrow.toTail())
	if err != nil {
		return err
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

	node, err  := m.nodes.read(position)
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

		m.arrows.produceHole(tail.toArrow())
	}

	for {
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

		m.arrows.produceHole(tail.toArrow())
	}
}

func (m mix) delete() error {
	err := m.tails.deleteSource(m.node.toSource())
	if err != nil {
		return err
	}

	err = m.heads.deleteTarget(m.node.toTarget())
	if err != nil {
		return err
	}

	m.nodes.produceHole(m.node)

	return nil
}
