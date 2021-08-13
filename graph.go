package klubok

type Graph struct {
	biloops biloops
}

func NewGraph(persister persister) *Graph {
	entries := newEntries(persister)
	entries.append(newVoidEntry())
	holes := newHoles(entries, void)
	vertices := newVertices(entries, holes, void)
	edges := newEdges(entries, holes)
	positiveLoop := newHeadLoop(vertices, edges)
	negativeLoop := newTailLoop(vertices, edges)

	return &Graph{
		biloops: newBiloops(
			vertices,
			edges,
			positiveLoop,
			negativeLoop,
		),
	}
}

func (g *Graph) Create() uint {
	return g.biloops.create().getPosition().toInteger()
}

func (g *Graph) ReadHeads(tail uint) []uint {
	positions := g.biloops.read(position(tail)).readHeads()

	heads := make([]uint, len(positions))
	for i, position := range positions {
		heads[i] = position.toInteger()
	}

	return heads
}

func (g *Graph) ReadTails(head uint) []uint {
	positions := g.biloops.read(position(head)).readTails()

	tails := make([]uint, len(positions))
	for i, position := range positions {
		tails[i] = position.toInteger()
	}

	return tails
}

func (g *Graph) Connect(tail uint, head uint) {
	g.biloops.read(position(tail)).addHead(position(head))
}

func (g *Graph) Disconnect(tail uint, head uint) {
	g.biloops.read(position(tail)).removeHead(position(head))
}

func (g *Graph) Delete(tail uint) {
	g.biloops.read(position(tail)).delete()
	//tailPosition, headPosition := position(message[0]), position(message[1])
	//
	//tailVertex := g.vertices.read(tailPosition)
	//
	//if !tailVertex.hasFirstEdgeTail() {
	//	return
	//}
	//
	//nextEdge := tailVertex.getFirstEdgeTail(g.edges)
	//if nextEdge.atPosition(headPosition) {
	//	g.holes.produce(nextEdge)
	//	// reconnect ring
	//}
	//
	//for nextEdge.hasNextPositiveEdge() {
	//	nextEdge = nextEdge.getNextEdgeHead(g.edges)
	//	if nextEdge.atPosition(headPosition) {
	//		g.holes.produce(nextEdge)
	//		// reconnect ring
	//	}
	//}
}

func (g *Graph) PositiveVertexDeleteStream(done <-chan struct{}) chan<- uint {
	tails := make(chan uint)

	go func() {
		defer close(tails)

		for {
			select {
			case <-done:
				return
			case tail := <-tails:
				tailVertex := g.vertices.read(position(tail))
				g.vertices.produceHole(tailVertex)

				if tailVertex.hasFirstPositiveEdge() {
					nextEdge := tailVertex.getFirstPositiveEdge(g.edges)
					g.edges.produceHole(nextEdge)

					for !nextEdge.hasNextPositiveEdge() {
						nextEdge = nextEdge.getNextPositiveEdge(g.edges)
						g.edges.produceHole(nextEdge)
					}
				}
			}
		}
	}()

	return tails
}

func (g *Graph) NegativeVertexDeleteStream(done <-chan struct{}) chan<- uint {
	tails := make(chan uint)

	go func() {
		defer close(tails)

		for {
			select {
			case <-done:
				return
			case tail := <-tails:
				tailVertex := g.vertices.read(position(tail))
				g.vertices.produceHole(tailVertex)

				if tailVertex.hasFirstNegativeEdge() {
					nextEdge := tailVertex.getFirstNegativeEdge(g.edges)
					g.edges.produceHole(nextEdge)

					for !nextEdge.hasNextNegativeEdge() {
						nextEdge = nextEdge.getNextNegativeEdge(g.edges)
						g.edges.produceHole(nextEdge)
					}
				}
			}
		}
	}()

	return tails
}

func (g *Graph) TotalVertexDeleteStream(done <-chan struct{}) chan<- uint {
	tails := make(chan uint)
	positive := g.PositiveVertexDeleteStream(done)
	negative := g.NegativeVertexDeleteStream(done)

	go func() {
		defer close(tails)
		defer close(positive)
		defer close(negative)

		for {
			select {
			case <-done:
				return
			case tail := <-tails:
				go func() {
					positive <- tail
				}()
				go func() {
					negative <- tail
				}()
			}
		}
	}()

	return tails
}
