package klubok

type biloop struct {
	vertex       vertex
	vertices     vertices
	edges        edges
	positiveLoop positiveLoop
	negativeLoop negativeLoop
}

func newBiloop(vertex vertex, vertices vertices, edges edges, positiveLoop positiveLoop, negativeLoop negativeLoop) biloop {
	return biloop{vertex: vertex, vertices: vertices, edges: edges, positiveLoop: positiveLoop, negativeLoop: negativeLoop}
}

func (b biloop) getPosition() position {
	return b.vertex.getPosition()
}

func (b biloop) readHeads() []position {
	return b.positiveLoop.readHeads(b.vertex)
}

func (b biloop) readTails() []position {
	return b.negativeLoop.readTails(b.vertex)
}

func (b biloop) addHead(position position) {
	for _, present := range b.positiveLoop.readHeads(b.vertex) {
		if present == position {
			return
		}
	}
	head := b.vertices.read(position)
	b.positiveLoop.addHead(b.vertex, head)
}

func (b biloop) removeHead(position position) {
	head := b.vertices.read(position)
	b.positiveLoop.removeHead(b.vertex, head)
}

func (b biloop) delete() {

}
