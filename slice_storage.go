package clew

import "fmt"

type sliceStorage struct {
	entries [][6]uint
}

func NewSliceStorage() *sliceStorage {
	return &sliceStorage{entries: [][6]uint{}}
}

func (s *sliceStorage) create() (uint, error) {
	s.entries = append(s.entries, [6]uint{})
	return uint(len(s.entries) - 1), nil
}

func (s *sliceStorage) read(position uint) ([6]uint, error) {
	return s.entries[int(position)], nil
}

func (s *sliceStorage) update(position uint, entry [6]uint) error {
	s.entries[int(position)] = entry
	return nil
}

// toString for using within tests
func (s *sliceStorage) String() string {
	result := "\n"
	for i, entry := range s.entries {
		result += fmt.Sprintf("%d: %2d %2d %2d %2d %2d %2d \n", i, entry[0], entry[1], entry[2], entry[3], entry[4], entry[5])
	}
	result += "\n"

	return result
}
