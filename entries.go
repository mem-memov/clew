package klubok

type entries struct {
	persister persister
}

func newEntries(persister persister) entries {
	return entries{
		persister: persister,
	}
}

func (e entries) next() position {
	return newPosition(e.persister.next())
}

func (e entries) read(position position) entry {
	return newEntry(e.persister.read(position.toInteger()))
}

func (e entries) update(position position, entry entry) {
	e.persister.update(position.toInteger(), entry.toArray())
}

func (e entries) append(entry entry) {
	e.persister.append(entry.toArray())
}
