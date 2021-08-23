package clew

type Graph struct {
	initializer *initializer
	mixes       *mixes
}

func NewGraph(storage storage) *Graph {
	entries := newEntries(storage)
	initializer := newInitializer(entries)
	holes := newHoles(entries, void)
	nodes := newNodes(entries, holes)
	arrows := newArrows(entries, holes)
	heads := newHeads(nodes, arrows)
	tails := newTails(nodes, arrows)

	return &Graph{
		initializer: initializer,
		mixes: newMixes(
			nodes,
			arrows,
			heads,
			tails,
		),
	}
}

func (g *Graph) Has(source uint) (bool, error) {
	_, err := g.mixes.read(newPosition(source))

	return err == nil, nil
}

func (g *Graph) Create() (uint, error) {
	err := g.initializer.initialize()
	if err != nil {
		return 0, err
	}

	mix, err := g.mixes.create()
	if err != nil {
		return 0, err
	}

	return mix.getPosition().toInteger(), nil
}

func (g *Graph) ReadSources(target uint) ([]uint, error) {
	err := g.initializer.initialize()
	if err != nil {
		return []uint{}, err
	}

	mix, err := g.mixes.read(newPosition(target))
	if err != nil {
		return []uint{}, err
	}

	positions, err := mix.readSources()
	if err != nil {
		return []uint{}, err
	}

	heads := make([]uint, len(positions))
	for i, position := range positions {
		heads[i] = position.toInteger()
	}

	return heads, nil
}

func (g *Graph) ReadTargets(source uint) ([]uint, error) {
	err := g.initializer.initialize()
	if err != nil {
		return []uint{}, err
	}

	mix, err := g.mixes.read(newPosition(source))
	if err != nil {
		return []uint{}, err
	}

	positions, err := mix.readTargets()
	if err != nil {
		return []uint{}, err
	}

	tails := make([]uint, len(positions))
	for i, position := range positions {
		tails[i] = position.toInteger()
	}

	return tails, nil
}

func (g *Graph) SetReference(source uint, reference uint) error {
	err := g.initializer.initialize()
	if err != nil {
		return err
	}

	mix, err := g.mixes.read(newPosition(source))
	if err != nil {
		return err
	}

	err = mix.setReference(newPosition(reference))
	if err != nil {
		return err
	}

	return nil
}

func (g *Graph) GetReference(source uint) (uint, uint, error) {
	err := g.initializer.initialize()
	if err != nil {
		return 0, 0, err
	}

	mix, err := g.mixes.read(newPosition(source))
	if err != nil {
		return 0, 0, err
	}

	previousPosition, nextPosition, err := mix.getReference()
	if err != nil {
		return 0, 0, err
	}

	return previousPosition.toInteger(), nextPosition.toInteger(), nil
}

func (g *Graph) Connect(source uint, target uint) error {
	err := g.initializer.initialize()
	if err != nil {
		return err
	}

	mix, err := g.mixes.read(newPosition(source))
	if err != nil {
		return err
	}

	err = mix.addTarget(newPosition(target))
	if err != nil {
		return err
	}

	return nil
}

func (g *Graph) Disconnect(source uint, target uint) error {
	err := g.initializer.initialize()
	if err != nil {
		return err
	}

	mix, err := g.mixes.read(newPosition(source))
	if err != nil {
		return err
	}

	err = mix.removeTarget(newPosition(target))
	if err != nil {
		return err
	}

	return nil
}

func (g *Graph) Delete(source uint) error {
	err := g.initializer.initialize()
	if err != nil {
		return err
	}

	mix, err := g.mixes.read(newPosition(source))
	if err != nil {
		return err
	}

	previousNode, nextNode, err := mix.getReference()
	if err != nil {
		return err
	}

	previousMix, err := g.mixes.read(previousNode)
	if err != nil {
		return err
	}

	nextMix, err := g.mixes.read(nextNode)
	if err != nil {
		return err
	}

	err = previousMix.deleteNextReference()
	if err != nil {
		return err
	}

	err = nextMix.deletePreviousReference()
	if err != nil {
		return err
	}

	err = mix.delete()
	if err != nil {
		return err
	}

	return nil
}
