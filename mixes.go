package clew

type mixes struct {
	nodes        *nodes
	arrows       *arrows
	heads *heads
	tails *tails
}

func newMixes(nodes *nodes, arrows *arrows, heads *heads, tails *tails) *mixes {
	return &mixes{nodes: nodes, arrows: arrows, heads: heads, tails: tails}
}

func (b *mixes) create() mix {
	node := b.nodes.create()
	b.nodes.update(node)
	return newMix(node, b.nodes, b.arrows, b.heads, b.tails)
}

func (b *mixes) read(position position) mix {
	node := b.nodes.read(position)
	return newMix(node, b.nodes, b.arrows, b.heads, b.tails)
}
