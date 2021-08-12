package klubok

type biloops struct {
	vertices     vertices
	edges        edges
	positiveLoop headLoop
	negativeLoop tailLoop
}

func newBiloops(vertices vertices, edges edges, positiveLoop headLoop, negativeLoop tailLoop) biloops {
	return biloops{vertices: vertices, edges: edges, positiveLoop: positiveLoop, negativeLoop: negativeLoop}
}

func (b biloops) create() biloop {
	vertex := b.vertices.create()
	b.vertices.update(vertex)
	return newBiloop(vertex, b.vertices, b.edges, b.positiveLoop, b.negativeLoop)
}

func (b biloops) read(position position) biloop {
	vertex := b.vertices.read(position)
	return newBiloop(vertex, b.vertices, b.edges, b.positiveLoop, b.negativeLoop)
}
