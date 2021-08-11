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
	positiveLoop := newPositiveLoop(vertices, edges)
	negativeLoop := newNegativeLoop(vertices, edges)

	return &Graph{
		biloops: newBiloops(
			vertices,
			edges,
			positiveLoop,
			negativeLoop,
		),
	}
}

func (g *Graph) StreamNewVertices(done <-chan struct{}) (chan<- struct{}, <-chan uint) {
	tailStream := make(chan uint)
	requestStream := make(chan struct{})

	go func() {
		defer close(tailStream)
		defer close(requestStream)

		for {
			select {
			case <-done:
				return
			case <-requestStream:
				tailStream <- g.biloops.create().getPosition().toInteger()
			}
		}
	}()

	return requestStream, tailStream
}

func (g *Graph) StreamPositiveAdjacentVerticesForward(done <-chan struct{}) (chan<- uint, <-chan (<-chan uint)) {
	tailStream := make(chan uint)
	headStreams := make(chan (<-chan uint))

	go func() {
		defer close(tailStream)
		defer close(headStreams)

		for {
			select {
			case <-done:
				return
			case tail := <-tailStream:

				tailBiloop := g.biloops.read(position(tail))
				headStream := tailBiloop.streamNextPositiveVertex(done)

				select {
				case <-done:
					return
				case headStreams <- headStream:
				}
			}
		}
	}()

	return tailStream, headStreams
}

func (g *Graph) StreamNegativeAdjacentVerticesForward(done <-chan struct{}) (chan<- uint, <-chan (<-chan uint)) {
	headStream := make(chan uint)
	tailStreams := make(chan (<-chan uint))

	go func() {
		defer close(headStream)
		defer close(tailStreams)

		for {
			select {
			case <-done:
				return
			case head := <-headStream:
				headBiloop := g.biloops.read(position(head))
				tailStream := headBiloop.streamNextNegativeVertex(done)

				select {
				case <-done:
					return
				case tailStreams <- tailStream:
				}
			}
		}
	}()

	return headStream, tailStreams
}


func (g *Graph) StreamPositiveAdjacentVerticesBackward(done <-chan struct{}) (chan<- uint, <-chan (<-chan uint)) {
	tailStream := make(chan uint)
	headStreams := make(chan (<-chan uint))

	go func() {
		defer close(tailStream)
		defer close(headStreams)

		for {
			select {
			case <-done:
				return
			case tail := <-tailStream:

				tailBiloop := g.biloops.read(position(tail))
				headStream := tailBiloop.streamPreviousPositiveVertex(done)

				select {
				case <-done:
					return
				case headStreams <- headStream:
				}
			}
		}
	}()

	return tailStream, headStreams
}

func (g *Graph) StreamNegativeAdjacentVerticesBackward(done <-chan struct{}) (chan<- uint, <-chan (<-chan uint)) {
	headStream := make(chan uint)
	tailStreams := make(chan (<-chan uint))

	go func() {
		defer close(headStream)
		defer close(tailStreams)

		for {
			select {
			case <-done:
				return
			case head := <-headStream:
				headBiloop := g.biloops.read(position(head))
				tailStream := headBiloop.streamPreviousNegativeVertex(done)

				select {
				case <-done:
					return
				case tailStreams <- tailStream:
				}
			}
		}
	}()

	return headStream, tailStreams
}

func (g *Graph) NewEdges(done <-chan struct{}) chan<- [2]uint {
	edges := make(chan [2]uint)

	go func() {
		defer close(edges)

		for {
			select {
			case <-done:
				return
			case messsage := <-edges:
				tailVertex := g.vertices.read(position(messsage[0]))
				headVertex := g.vertices.read(position(messsage[1]))

				edgePosition := g.nextEntry
				newEdge := newVoidEdge()
				newEdge.setPositiveVertex(tailVertex)
				newEdge.setNegativeVertex(headVertex)

				tailVertex.setNextPositiveEdgeIfEmpty(edgePosition)

				if tailVertex[positivePrevious] != void {
					newEdge[positivePrevious] = tailVertex[positivePrevious]
					positivePreviousEdge := g.entries[tailVertex[positivePrevious]]
					positivePreviousEdge[positiveNext] = edgePosition
					g.entries[tailVertex[positivePrevious]] = positivePreviousEdge
					tailVertex[lastPositive] = edgePosition
				}

				headVertex.setNextNegativeEdgeIfEmpty(edgePosition)

				if headVertex[negativePrevious] != void {
					newEdge[negativePrevious] = headVertex[negativePrevious]
					negativePreviousEdge := g.entries[headVertex[negativePrevious]]
					negativePreviousEdge[negativeNext] = edgePosition
					g.entries[headVertex[negativePrevious]] = negativePreviousEdge
					headVertex[negativePrevious] = edgePosition
				}

				tailVertex.setPreviousPositiveEdge(edgePosition)
				headVertex.setPreviousNegativeEdge(edgePosition)

				g.vertices.update(tailVertex)
				g.vertices.update(headVertex)

				if g.holes.exist() {
					g.holes.consume(newEdge)
				} else {
					g.nextEntry = g.edges.append(newEdge)
				}
			}
		}
	}()

	return edges
}

func (g *Graph) DeletePositiveEdges(done <-chan struct{}) chan<- [2]uint {
	edges := make(chan [2]uint)

	go func() {
		defer close(edges)

		for {
			select {
			case <-done:
				return
			case message := <-edges:

				go func() {
					tailPosition, headPosition := position(message[0]), position(message[1])

					tailVertex := g.vertices.read(tailPosition)

					if !tailVertex.hasFirstPositiveEdge() {
						return
					}

					nextEdge := tailVertex.getFirstPositiveEdge(g.edges)
					if nextEdge.atPosition(headPosition) {
						g.holes.produce(nextEdge)
						// reconnect ring
					}

					for nextEdge.hasNextPositiveEdge() {
						nextEdge = nextEdge.getNextPositiveEdge(g.edges)
						if nextEdge.atPosition(headPosition) {
							g.holes.produce(nextEdge)
							// reconnect ring
						}
					}
				}()
			}
		}
	}()

	return edges
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
