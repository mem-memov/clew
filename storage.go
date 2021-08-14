package clew

type storage interface {
	create() (uint, error)
	read(uint) ([6]uint, error)
	update(uint, [6]uint) error
}
