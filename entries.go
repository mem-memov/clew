package klubok

type entries struct {
	storage storage
}

func newEntries(storage storage) entries {
	return entries{
		storage: storage,
	}
}

func (e entries) next() position {
	return newPosition(e.storage.next())
}

func (e entries) read(position position) entry {
	return newEntry(e.storage.read(position.toInteger()))
}

func (e entries) update(position position, entry entry) {
	e.storage.update(position.toInteger(), entry.toArray())
}

func (e entries) append(entry entry) {
	e.storage.append(entry.toArray())
}
