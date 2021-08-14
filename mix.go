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

func (m mix) readSources() []position {
	return m.heads.readHeads(m.node.toTarget())
}

func (m mix) readTargets() []position {
	return m.tails.readTails(m.node.toSource())
}

func (m mix) addTarget(position position) {

	source := m.node.toSource()
	target := m.nodes.read(position).toTarget()

	for _, present := range m.tails.readTails(source) {
		if present == position {
			return
		}
	}

	arrow := m.arrows.create(source, target)

	arrow = m.tails.addTail(source, arrow.toTail())
	arrow = m.heads.addHead(target, arrow.toHead())

	m.arrows.update(arrow)
}

func (m mix) removeTarget(position position) {

	source := m.node.toSource()

	if !source.hasFirstTail() {
		return
	}

	target := m.nodes.read(position).toTarget()

	if !target.hasFirstHead() {
		return
	}

	first := source.getFirstTail(m.arrows)
	tail := first

	if tail.toHead().hasTarget(target) {
		m.heads.removeHead(target, tail.toHead())
		m.tails.removeTail(source, tail)
		m.arrows.produceHole(tail.toArrow())
	}

	for {
		tail = tail.getNext(m.arrows)
		if tail.isSame(first) {
			return
		}
		if !tail.toHead().hasTarget(target) {
			continue
		}
		m.heads.removeHead(target, tail.toHead())
		m.tails.removeTail(source, tail)
		m.arrows.produceHole(tail.toArrow())
	}
}

func (m mix) delete() {
	m.tails.deleteSource(m.node.toSource())
	m.heads.deleteTarget(m.node.toTarget())
	m.nodes.produceHole(m.node)
}
