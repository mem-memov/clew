package klubok

type positions struct {
	nexter nexter
}

func newPositions(nexter nexter) positions {
	return positions{nexter: nexter}
}

func (p positions) next() position {
	return p.nexter.next()
}
