package clew

type entries struct {
	storage storage
}

func newEntries(storage storage) *entries {
	return &entries{
		storage: storage,
	}
}

func (e *entries) create() (position, error) {
	next, err := e.storage.create()
	if err != nil {
		return void, err
	}
	return newPosition(next), err
}

func (e *entries) read(position position) (entry, error) {
	integers, err := e.storage.read(position.toInteger())
	if err != nil {
		return newVoidEntry(), err
	}
	return newEntry(integers), nil
}

func (e *entries) update(position position, entry entry) error {
	return e.storage.update(position.toInteger(), entry.toArray())
}
