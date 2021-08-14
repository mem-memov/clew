package klubok

type Graph struct {
	mixes mixes
}

func NewGraph(storage storage) *Graph {
	entries := newEntries(storage)
	entries.append(newVoidEntry())
	holes := newHoles(entries, void)
	nodes := newNodes(entries, holes, void)
	arrows := newArrows(entries, holes)
	heads := newHeads(nodes, arrows)
	tails := newTails(nodes, arrows)

	return &Graph{
		mixes: newMixes(
			nodes,
			arrows,
			heads,
			tails,
		),
	}
}

func (g *Graph) Create() uint {
	return g.mixes.create().getPosition().toInteger()
}

func (g *Graph) ReadSources(target uint) []uint {
	positions := g.mixes.read(position(target)).readSources()

	heads := make([]uint, len(positions))
	for i, position := range positions {
		heads[i] = position.toInteger()
	}

	return heads
}

func (g *Graph) ReadTargets(source uint) []uint {
	positions := g.mixes.read(position(source)).readTargets()

	tails := make([]uint, len(positions))
	for i, position := range positions {
		tails[i] = position.toInteger()
	}

	return tails
}

func (g *Graph) Connect(source uint, target uint) {
	g.mixes.read(position(source)).addTarget(position(target))
}

func (g *Graph) Disconnect(source uint, target uint) {
	g.mixes.read(position(source)).removeTarget(position(target))
}

func (g *Graph) Delete(source uint) {
	g.mixes.read(position(source)).delete()
}
