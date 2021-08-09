package klubok

type entry [entrySize]position

func newEntry() entry {
	return entry{void, void, void, void, void, void}
}
