package klubok

type updater interface {
	update(s entries, p position)
}