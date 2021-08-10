package klubok

type biloops struct {
	vertices vertices
	loops loops
}

func newBiloops(vertices vertices, loops loops) biloops {
	return biloops{vertices: vertices, loops: loops}
}

func (b biloops) createBiloop() biloop {
	vertex := b.vertices.create()
	b.vertices.update(vertex)
	return newBiloop(b.loops.createLoop(vertex), b.loops.createLoop(vertex))
}
