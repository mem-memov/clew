package clew

type initializer struct {
	isInitialized bool
	entries       *entries
}

func newInitializer(entries *entries) *initializer {
	return &initializer{
		isInitialized: false,
		entries:       entries,
	}
}

func (i *initializer) initialize() error {
	if i.isInitialized {
		return nil
	}

	_, err := i.entries.read(void)
	// storage not empty
	if err == nil {
		i.isInitialized = true
		return nil
	}

	// a voidEntry which makes 0 to a special value, that means no position has been set, it may contain graph metadata
	_, err = i.entries.create()
	if err != nil {
		return err
	}

	i.isInitialized = true

	return nil
}
