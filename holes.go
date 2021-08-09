package klubok

type holes struct {
	lastHole position
	entries storage
}

func newHoles(s storage, lastHole position) holes {
	return holes{
		lastHole: lastHole,
		entries: s,
	}
}

func (h holes) exist() bool {
	return h.lastHole == void
}

func (h holes) read(p position) hole {
	return newHole(h.entries.read(p))
}

func (h holes) last() position {
	return h.lastHole
}

func (h holes) produce(u updater) {

}

func (h holes) consume(u updater) {
	lastHole := newHole(h.entries.read(h.lastHole))
	h.lastHole = lastHole.moveNext(h.entries, u, h.lastHole)
}