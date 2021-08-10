package klubok

type position uint

func newPosition(integer uint) position {
	return position(integer)
}

func (p position) toInteger() uint {
	return uint(p)
}
