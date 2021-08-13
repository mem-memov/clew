package klubok

type biloop struct {
	vertex       vertex
	vertices     vertices
	edges        edges
	headLoop headLoop
	tailLoop tailLoop
}

func newBiloop(vertex vertex, vertices vertices, edges edges, positiveLoop headLoop, negativeLoop tailLoop) biloop {
	return biloop{vertex: vertex, vertices: vertices, edges: edges, headLoop: positiveLoop, tailLoop: negativeLoop}
}

func (b biloop) getPosition() position {
	return b.vertex.getPosition()
}

func (b biloop) readHeads() []position {
	return b.headLoop.readHeads(b.vertex.toHead())
}

func (b biloop) readTails() []position {
	return b.tailLoop.readTails(b.vertex.toTail())
}

func (b biloop) addHead(position position) {

	for _, present := range b.headLoop.readHeads(b.vertex.toHead()) {
		if present == position {
			return
		}
	}

	edge := b.edges.create()

	edge.toTail().setTailVertex(b.vertex.toTail())
	edge.toHead().setHeadVertex(b.vertex.toHead())

	b.headLoop.addHead(b.vertex.toHead(), edge.toHead())
	b.tailLoop.addTail(b.vertex.toTail(), edge.toTail())
}

func (b biloop) removeHead(position position) {

	tailVertex := b.vertex.toTail()

	if !tailVertex.hasFirstEdgeTail() {
		return
	}

	headVertex := b.vertices.read(position).toHead()

	if !headVertex.hasFirstEdgeHead() {
		return
	}

	firstEdge := tailVertex.getFirstEdgeTail(b.edges)
	nextEdge := firstEdge

	if nextEdge.hasTailVertex(tailVertex) {
		b.headLoop.removeHead(tailVertex, nextEdge)
		b.tailLoop.removeTail(headVertex, nextEdge)
		b.edges.produceHole(nextEdge)
	}

	for {
		nextEdge = nextEdge.getNextEdgeTail(b.edges)
		if nextEdge.getPosition() == firstEdge.getPosition() {
			return
		}
		if nextEdge.hasPositiveVertex(headVertex.toVertex()) {
			b.headLoop.removeHead(tailVertex, nextEdge)
			b.tailLoop.removeTail(headVertex, nextEdge)
			b.edges.produceHole(nextEdge)
		}
	}
}

func (b biloop) delete() {

}
