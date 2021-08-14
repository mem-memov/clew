package clew

type sliceStorage struct {
	entries [][6]uint
}

func NewSliceStorage() *sliceStorage {
	return &sliceStorage{entries: [][6]uint{}}
}

func (s *sliceStorage) next() (uint, error) {
	return uint(len(s.entries)), nil
}

func (s *sliceStorage) read(position uint) ([6]uint, error) {
	return s.entries[int(position)], nil
}

func (s *sliceStorage) update(position uint, entry [6]uint) error {
	s.entries[int(position)] = entry
	return nil
}

func (s *sliceStorage) append(entry [6]uint) error {
	s.entries = append(s.entries, entry)
	return nil
}
