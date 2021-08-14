package clew

type entry [6]position


func newVoidEntry() entry {
	return entry{void, void, void, void, void, void}
}

func newEntry(integers [6]uint) entry {
	return entry{position(integers[0]), position(integers[1]), position(integers[2]), position(integers[3]), position(integers[4]), position(integers[5])}
}

func (e entry) toArray() [6]uint {
	return [6]uint{uint(e[0]), uint(e[1]), uint(e[2]), uint(e[3]), uint(e[4]), uint(e[5])}
}
