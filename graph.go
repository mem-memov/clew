package klubok

type Graph struct {
	biloops biloops
}

func NewGraph(persister persister) *Graph {
	entries := newEntries(persister)
	entries.append(newVoidEntry())
	holes := newHoles(entries, void)
	nodes := newNodes(entries, holes, void)
	arrows := newArrows(entries, holes)
	heads := newHeads(nodes, arrows)
	tails := newTails(nodes, arrows)

	return &Graph{
		biloops: newBiloops(
			nodes,
			arrows,
			heads,
			tails,
		),
	}
}

func (g *Graph) Create() uint {
	return g.biloops.create().getPosition().toInteger()
}

func (g *Graph) ReadSources(target uint) []uint {
	positions := g.biloops.read(position(target)).readSources()

	heads := make([]uint, len(positions))
	for i, position := range positions {
		heads[i] = position.toInteger()
	}

	return heads
}

func (g *Graph) ReadTargets(source uint) []uint {
	positions := g.biloops.read(position(source)).readTargets()

	tails := make([]uint, len(positions))
	for i, position := range positions {
		tails[i] = position.toInteger()
	}

	return tails
}

func (g *Graph) Connect(source uint, target uint) {
	g.biloops.read(position(source)).addTarget(position(target))
}

func (g *Graph) Disconnect(tail uint, head uint) {
	g.biloops.read(position(tail)).removeTarget(position(head))
}

func (g *Graph) Delete(tail uint) {
	g.biloops.read(position(tail)).delete()
}
