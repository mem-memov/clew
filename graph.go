package klubok

const (
	entrySize = 6
	void uint = 0

	identifier uint = 0
	previousVertex uint = 1
	firstPositive uint = 2
	lastPositive uint = 3
	firstNegative uint = 4
	lastNegative uint = 5

	positiveDirection uint = 0
	positivePrevious uint = 1
	positiveNext uint = 2
	negativeDirection uint = 3
	negativePrevious uint = 4
	negativeNext uint = 5
)

type entry [entrySize]uint

type Graph struct {
	nextEntry uint
	lastVertex uint
	lastHole uint
	entries []entry
}

func NewGraph() *Graph {
	// void entry to make 0 a special value, it may contain graph metadata
	voidEntry := entry {uint(0), uint(0), uint(0), uint(0), uint(0), uint(0)}
	return &Graph{
		nextEntry: 1,
		lastVertex: void,
		lastHole: void,
		entries: []entry{voidEntry},
	}
}

func (g *Graph) Create() uint {
	tail := g.nextEntry

	g.entries = append(g.entries, entry{
		identifier: tail,
		previousVertex: g.lastVertex,
		firstPositive: void,
		lastPositive: void,
		firstNegative: void,
		lastNegative: void,
	})
	g.nextEntry++

	g.lastVertex = tail

	return tail
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

	tailVertex := g.entries[tail]
	headVertex := g.entries[head]

	edge := entry{
		positiveDirection: head,
		positivePrevious: void,
		positiveNext: void,
		negativeDirection: tail,
		negativePrevious: void,
		negativeNext: void,
	}
	
	if tailVertex[positiveNext] == void {
		tailVertex[positiveNext] = g.nextEntry
	}
	
	if tailVertex[positivePrevious] != void {
		positivePreviousEdge := g.entries[tailVertex[positivePrevious]]
		positivePreviousEdge[positiveNext] = g.nextEntry
		g.entries[tailVertex[positivePrevious]] = positivePreviousEdge
		tailVertex[lastPositive] = g.nextEntry
	}

	if headVertex[negativeNext] == void {
		headVertex[negativeNext] = g.nextEntry
	}

	if headVertex[negativePrevious] != void {
		negativePreviousEdge := g.entries[headVertex[negativePrevious]]
		negativePreviousEdge[negativeNext] = g.nextEntry
		g.entries[headVertex[negativePrevious]] = negativePreviousEdge
		headVertex[negativePrevious] = g.nextEntry
	}

	tailVertex[positivePrevious] = g.nextEntry
	headVertex[negativePrevious] = g.nextEntry
	g.entries[tail] = tailVertex
	g.entries[head] = headVertex

	g.entries = append(g.entries, edge)
	g.nextEntry++
}

func (g *Graph) Delete(tail uint) {
	tailVertex := g.entries[tail]
	g.entries[tail] = entry{}

	if tailVertex[firstPositive] != void {
		nextEdge := g.entries[tailVertex[firstPositive]]
		g.entries[tailVertex[firstPositive]] = entry{}

		for {
			if nextEdge[positiveNext] == void {
				break
			}
			nextEdge = g.entries[nextEdge[positiveNext]]
			g.entries[nextEdge[positiveNext]] = entry{}
		}
	}

	if tailVertex[firstNegative] != void {
		nextEdge := g.entries[tailVertex[firstNegative]]
		g.entries[tailVertex[firstNegative]] = entry{}

		for {
			if nextEdge[negativeNext] == void {
				break
			}
			nextEdge = g.entries[nextEdge[negativeNext]]
			g.entries[nextEdge[negativeNext]] = entry{}
		}
	}
}
