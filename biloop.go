package klubok

type biloop struct {
	node   node
	nodes  nodes
	arrows arrows
	heads  heads
	tails  tails
}

func newBiloop(vertex node, vertices nodes, arrows arrows, positiveLoop heads, negativeLoop tails) biloop {
	return biloop{node: vertex, nodes: vertices, arrows: arrows, heads: positiveLoop, tails: negativeLoop}
}

func (b biloop) getPosition() position {
	return b.node.getPosition()
}

func (b biloop) readSources() []position {
	return b.heads.readHeads(b.node.toTarget())
}

func (b biloop) readTargets() []position {
	return b.tails.readTails(b.node.toSource())
}

func (b biloop) addTarget(position position) {

	for _, present := range b.heads.readHeads(b.node.toTarget()) {
		if present == position {
			return
		}
	}

	arrow := b.arrows.create()

	arrow.toTail().setSource(b.node.toSource())
	arrow.toHead().setTarget(b.node.toTarget())

	b.heads.addHead(b.node.toTarget(), arrow.toHead())
	b.tails.addTail(b.node.toSource(), arrow.toTail())
}

func (b biloop) removeTarget(position position) {

	source := b.node.toSource()

	if !source.hasFirstTail() {
		return
	}

	target := b.nodes.read(position).toTarget()

	if !target.hasFirstHead() {
		return
	}

	first := source.getFirstTail(b.arrows)
	tail := first

	if tail.toArrow().toHead().hasTarget(target) {
		b.heads.removeHead(target, tail.toArrow().toHead())
		b.tails.removeTail(source, tail)
		b.arrows.produceHole(tail.toArrow())
	}

	for {
		tail = tail.getNext(b.arrows)
		if tail.toArrow().getPosition() == first.toArrow().getPosition() {
			return
		}
		if tail.toArrow().toHead().hasTarget(target) {
			b.heads.removeHead(target, tail.toArrow().toHead())
			b.tails.removeTail(source, tail)
			b.arrows.produceHole(tail.toArrow())
		}
	}
}

func (b biloop) delete() {

}
