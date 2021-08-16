package clew

type count uint

func newCount(integer uint) count {
	return count(integer)
}

func (c count) toInteger() uint {
	return uint(c)
}
