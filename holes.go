package klubok

type holes struct {
	lastHole position
	entries  entries
}

func newHoles(entries entries, lastHole position) holes {
	return holes{
		lastHole: lastHole,
		entries:  entries,
	}
}

func (h holes) exist() bool {
	return h.lastHole == void
}

func (h holes) last() position {
	return h.lastHole
}

func (h holes) produceHole(u updater) {
	lastHole := newHole(h.lastHole, h.entries.read(h.lastHole))
}

func (h holes) consumeHole(u updater) {
	lastHole := newHole(h.lastHole, h.entries.read(h.lastHole))
	h.lastHole = lastHole.consume(h.entries, u, h.lastHole)
}
