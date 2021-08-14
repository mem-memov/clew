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

func (b *mixes) create() (mix, error) {
	node, err := b.nodes.create()
	if err != nil {
		return mix{}, err
	}

	return newMix(node, b.nodes, b.arrows, b.heads, b.tails), nil
}

func (b *mixes) read(position position) (mix, error) {
	node, err := b.nodes.read(position)
	if err != nil {
		return mix{}, err
	}

	return newMix(node, b.nodes, b.arrows, b.heads, b.tails), nil
}
