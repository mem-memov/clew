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

func (b biloop) streamNextPositiveVertex(done <-chan struct{}) <-chan uint {
	headStream := make(chan uint)

	go func() {
		defer close(headStream)
		b.positiveLoop.streamNextVertex(done, b.vertex, headStream)
	}()

	return headStream
}

func (b biloop) streamNextNegativeVertex(done <-chan struct{}) <-chan uint {
	tailStream := make(chan uint)

	go func() {
		defer close(tailStream)
		b.negativeLoop.streamNextVertex(done, b.vertex, tailStream)
	}()

	return tailStream
}

func (b biloop) streamPreviousPositiveVertex(done <-chan struct{}) <-chan uint {
	headStream := make(chan uint)

	go func() {
		defer close(headStream)
		b.positiveLoop.streamPreviousVertex(done, b.vertex, headStream)
	}()

	return headStream
}

func (b biloop) streamPreviousNegativeVertex(done <-chan struct{}) <-chan uint {
	tailStream := make(chan uint)

	go func() {
		defer close(tailStream)
		b.negativeLoop.streamPreviousVertex(done, b.vertex, tailStream)
	}()

	return tailStream
}
