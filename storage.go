package klubok

type sliceStorage struct {
	entries []entry
}

func (s *sliceStorage) read(p position) entry {
	return s.entries[p]
}

func (s *sliceStorage) update(p position, e entry) {
	s.entries[p] = e
}

func (s *sliceStorage) append(e entry) position {
	s.entries = append(s.entries, e)
	return position(len(s.entries))
}