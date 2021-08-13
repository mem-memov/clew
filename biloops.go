package klubok

type biloops struct {
	nodes        nodes
	arrows       arrows
	positiveLoop heads
	negativeLoop tails
}

func newBiloops(vertices nodes, arrows arrows, positiveLoop heads, negativeLoop tails) biloops {
	return biloops{nodes: vertices, arrows: arrows, positiveLoop: positiveLoop, negativeLoop: negativeLoop}
}

func (b biloops) create() biloop {
	vertex := b.nodes.create()
	b.nodes.update(vertex)
	return newBiloop(vertex, b.nodes, b.arrows, b.positiveLoop, b.negativeLoop)
}

func (b biloops) read(position position) biloop {
	vertex := b.nodes.read(position)
	return newBiloop(vertex, b.nodes, b.arrows, b.positiveLoop, b.negativeLoop)
}
