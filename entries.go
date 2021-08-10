package klubok

type entries interface {
	read(p position) entry
	update(p position, e entry)
	// returns the position after the appended entry
	append(e entry) position
}
