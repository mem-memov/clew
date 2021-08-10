package klubok

const (
	entrySize          = 6
	void      position = 0
)

type position uint

type Graph struct {
	nextEntry position
	vertices  vertices
	edges     edges
	holes     holes
}

func NewGraph() *Graph {
	// void entry to make 0 a special value, it may contain graph metadata
	voidEntry := newEntry()
	entries := &sliceStorage{entries: []entry{voidEntry}}
	return &Graph{
		nextEntry: 1,
		vertices:  newVertices(entries, void),
		edges:     newEdges(entries),
		holes:     newHoles(entries, void),
	}
}

func (g *Graph) Create() uint {

	if g.holes.exist() {
		tailVertex := g.vertices.create(g.holes.last())
		g.holes.consume(tailVertex)
		hole := g.entries[g.lastHole][nextHole]
		tail := g.lastHole
		g.lastHole = hole
		g.entries[tail] = entry{
			identifier:     tail,
			previousVertex: g.lastVertex,
			firstPositive:  void,
			lastPositive:   void,
			firstNegative:  void,
			lastNegative:   void,
		}
		return tail
	} else {
		tail := g.nextEntry
		g.entries = append(g.entries, entry{
			identifier:     tail,
			previousVertex: g.lastVertex,
			firstPositive:  void,
			lastPositive:   void,
			firstNegative:  void,
			lastNegative:   void,
		})
		g.nextEntry++

		g.lastVertex = tail
		return tail
	}
}

func (g *Graph) ReadPositive(tail uint) []uint {
	heads := make([]uint, 0)

	tailVertex := g.entries[tail]

	if tailVertex[firstPositive] == void {
		return heads
	}

	nextEdge := g.entries[tailVertex[firstPositive]]
	heads = append(heads, nextEdge[positiveDirection])

	for {
		if nextEdge[positiveNext] == void {
			break
		}
		nextEdge = g.entries[nextEdge[positiveNext]]
		heads = append(heads, nextEdge[positiveDirection])
	}

	return heads
}

func (g *Graph) ReadNegative(tail uint) []uint {
	heads := make([]uint, 0)

	tailVertex := g.entries[tail]

	if tailVertex[firstNegative] == void {
		return heads
	}

	nextEdge := g.entries[tailVertex[firstNegative]]
	heads = append(heads, nextEdge[negativeDirection])

	for {
		if nextEdge[negativeNext] == void {
			break
		}
		nextEdge = g.entries[nextEdge[negativeNext]]
		heads = append(heads, nextEdge[negativeDirection])
	}

	return heads
}

func (g *Graph) Update(tail uint, head uint) {

	tailVertex := g.vertices.read(position(tail))
	headVertex := g.vertices.read(position(head))

	edgePosition := g.nextEntry
	newEdge := newEmptyEdge()
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

func (g *Graph) Delete(tail uint) {
	tailVertex := g.vertices.read(position(tail))
	g.holes.produce(tailVertex)

	if tailVertex.hasFirstPositiveEdge() {
		nextEdge := tailVertex.getFirstPositiveEdge(g.edges)
		g.holes.produce(nextEdge)

		for !nextEdge.hasNextPositiveEdge() {
			nextEdge = nextEdge.getNextPositiveEdge(g.edges)
			g.holes.produce(nextEdge)
		}
	}

	if tailVertex.hasFirstNegativeEdge() {
		nextEdge := tailVertex.getFirstNegativeEdge(g.edges)
		g.holes.produce(nextEdge)

		for !nextEdge.hasNextNegativeEdge() {
			nextEdge = nextEdge.getNextNegativeEdge(g.edges)
			g.holes.produce(nextEdge)
		}
	}
}
