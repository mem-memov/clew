package clew

type storage interface {
	next() (uint, error)
	read(uint) ([6]uint, error)
	update(uint, [6]uint) error
	append([6]uint) error
}
