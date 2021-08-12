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

	edge := b.edges.create()
	edge.setPositiveVertex(head)
	edge.setNegativeVertex(b.vertex)

	b.positiveLoop.addHead(b.vertex, head, edge)
	b.negativeLoop.addTail(head, b.vertex, edge)

	b.edges.update(edge)
}

func (b biloop) removeHead(position position) {

	if !b.vertex.hasFirstPositiveEdge() {
		return
	}

	head := b.vertices.read(position)

	if !head.hasFirstNegativeEdge() {
		return
	}

	firstEdge := b.vertex.getFirstPositiveEdge(b.edges)
	nextEdge := firstEdge

	if nextEdge.hasPositiveVertex(head) {
		b.positiveLoop.removeHead(b.vertex, nextEdge)
		b.negativeLoop.removeTail(head, nextEdge)
		b.edges.produceHole(nextEdge)
	}

	for {
		nextEdge = nextEdge.getNextPositiveEdge(b.edges)
		if nextEdge.getPosition() == firstEdge.getPosition() {
			return
		}
		if nextEdge.hasPositiveVertex(head) {
			b.positiveLoop.removeHead(b.vertex, nextEdge)
			b.negativeLoop.removeTail(head, nextEdge)
			b.edges.produceHole(nextEdge)
		}
	}
}

func (b biloop) delete() {

}
