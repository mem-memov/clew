package clew

type storage interface {
	next() uint
	read(uint) [6]uint
	update(uint, [6]uint)
	append([6]uint)
}
