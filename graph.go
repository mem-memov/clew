package clew

type Graph struct {
	initializer *initializer
	mixes       *mixes
}

func NewGraph(storage storage) (*Graph, error) {
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
	}, nil
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

	mix, err := g.mixes.read(position(target))
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

	mix, err := g.mixes.read(position(source))
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

func (g *Graph) Connect(source uint, target uint) error {
	err := g.initializer.initialize()
	if err != nil {
		return err
	}

	mix, err := g.mixes.read(position(source))
	if err != nil {
		return err
	}

	err = mix.addTarget(position(target))
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

	mix, err := g.mixes.read(position(source))
	if err != nil {
		return err
	}

	err = mix.removeTarget(position(target))
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

	mix, err := g.mixes.read(position(source))
	if err != nil {
		return err
	}

	err = mix.delete()
	if err != nil {
		return err
	}

	return nil
}
