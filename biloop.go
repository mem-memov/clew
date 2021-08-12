package klubok

type biloop struct {
	vertex       vertex
	vertices     vertices
	edges        edges
	positiveLoop headLoop
	negativeLoop tailLoop
}

func newBiloop(vertex vertex, vertices vertices, edges edges, positiveLoop headLoop, negativeLoop tailLoop) biloop {
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

	for _, present := range b.positiveLoop.readHeads(b.vertex.toTail()) {
		if present == position {
			return
		}
	}

	edge := b.edges.create()

	edge.toTail().setTailVertex(b.vertex.toTail())
	edge.toHead().setHeadVertex(b.vertex.toHead())

	b.positiveLoop.addHead(b.vertex.toTail(), edge.toTail())
	b.negativeLoop.addTail(b.vertex.toHead(), edge.toHead())
}

func (b biloop) removeHead(position position) {

	tail := b.vertex.toTail()

	if !tail.hasFirstEdgeTail() {
		return
	}

	head := b.vertices.read(position).toHead()

	if !head.hasFirstEdgeHead() {
		return
	}

	firstEdge := tail.getFirstEdgeTail(b.edges)
	nextEdge := firstEdge

	if nextEdge.hasPositiveVertex(head.toVertex()) {
		b.positiveLoop.removeHead(tail, nextEdge)
		b.negativeLoop.removeTail(head, nextEdge)
		b.edges.produceHole(nextEdge)
	}

	for {
		nextEdge = nextEdge.getNextPositiveEdge(b.edges)
		if nextEdge.getPosition() == firstEdge.getPosition() {
			return
		}
		if nextEdge.hasPositiveVertex(head.toVertex()) {
			b.positiveLoop.removeHead(tail, nextEdge)
			b.negativeLoop.removeTail(head, nextEdge)
			b.edges.produceHole(nextEdge)
		}
	}
}

func (b biloop) delete() {

}
