package klubok

const (
	nextHole position = 0
)

type hole entry

func newHole(e entry) hole {
	return hole(e)
}

func (h hole) moveNext(s storage, u updater, p position) position {
	u.update(s, p)
	return h[nextHole]
}
