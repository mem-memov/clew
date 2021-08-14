package clew

type sliceStorage struct {
	entries [][6]uint
}

func NewSliceStorage() *sliceStorage {
	return &sliceStorage{entries: [][6]uint{}}
}

func (s *sliceStorage) next() uint {
	return uint(len(s.entries))
}

func (s *sliceStorage) read(position uint) [6]uint {
	return s.entries[int(position)]
}

func (s *sliceStorage) update(position uint, entry [6]uint) {
	s.entries[int(position)] = entry
}

func (s *sliceStorage) append(entry [6]uint) {
	s.entries = append(s.entries, entry)
}
