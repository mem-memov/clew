package klubok

type storage interface {
	read(p position) entry
	update(p position, e entry)
	// returns the position after the appended entry
	append(e entry) position
}

type updater interface {
	update(s storage, p position)
}