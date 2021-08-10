package klubok

const (
	nextHole position = 0
)

type hole struct {
	position position
	entry    entry
}

func newHole(position position, entry entry) hole {
	return hole{position: position, entry: entry}
}

func (h hole) produce() {

}

func (h hole) consume(s entries, u updater, p position) position {
	u.update(s, p)
	return h.entry[nextHole]
}
