package clew

type tails struct {
	nodes  *nodes
	arrows *arrows
}

func newTails(nodes *nodes, arrows *arrows) *tails {
	return &tails{nodes: nodes, arrows: arrows}
}

func (t *tails) readTails(source source) ([]position, error) {
	tails := make([]position, 0)

	if !source.hasFirstTail() {
		return tails, nil
	}

	first, err := source.getFirstTail(t.arrows)
	if err != nil {
		return tails, err
	}

	next := first

	tails = append(tails, next.toHead().getTargetPosition())

	for {
		if !next.hasNext() {
			return tails, nil
		}

		next, err = next.getNext(t.arrows)
		if err != nil {
			return tails, err
		}

		if next.isSame(first) {
			return tails, nil
		}
		tails = append(tails, next.toHead().getTargetPosition())
	}
}

func (t *tails) addTail(source source, new tail) (source, arrow, error) {

	if !source.hasFirstTail() {

		source = source.setFirstTail(new)

	} else {

		first, err := source.getFirstTail(t.arrows)
		if err != nil {
			return source, arrow{}, err
		}

		if first.hasPrevious() {
			last, err := first.getPrevious(t.arrows)
			if err != nil {
				return source, arrow{}, err
			}

			last = last.setNext(new)
			new = new.setPrevious(last)
			new = new.setNext(first)
			first = first.setPrevious(new)

			err = t.arrows.update(last.toArrow())
			if err != nil {
				return source, arrow{}, err
			}
		} else {
			first = first.setPrevious(new)
			first = first.setNext(new)
			new = new.setPrevious(first)
			new = new.setNext(first)
		}

		err = t.arrows.update(first.toArrow())
		if err != nil {
			return source, arrow{}, err
		}
	}

	source = source.incrementTailCount()

	err := t.nodes.update(source.toNode())
	if err != nil {
		return source, arrow{}, err
	}

	return source, new.toArrow(), nil
}

func (t *tails) removeTail(source source, removed tail) error {

	first, err := source.getFirstTail(t.arrows)
	if err != nil {
		return err
	}

	if first.isSame(removed) {
		if first.isSurrounded() {
			previous, err := first.getPrevious(t.arrows)
			if err != nil {
				return err
			}

			next, err := first.getNext(t.arrows)
			if err != nil {
				return err
			}

			next, err = previous.bindNext(next, t.arrows)
			if err != nil {
				return err
			}

			source.setFirstTail(next)
		} else if first.isPaired() {
			second, err := first.getNext(t.arrows)
			if err != nil {
				return err
			}

			second = second.deletePrevious()
			second = second.deleteNext()
			source.setFirstTail(second)
		} else if first.isAlone() {
			source = source.deleteFirstTail()
		}

		source = source.decrementTailCount()

		return t.nodes.update(source.toNode())
	}

	previous := first

	for {
		if !previous.hasNext() {
			return nil
		}

		current, err := previous.getNext(t.arrows)
		if err != nil {
			return err
		}

		if current.isSame(first) {
			return nil
		}

		if current.isSame(removed) {
			next, err := current.getNext(t.arrows)
			if err != nil {
				return err
			}

			_, err = previous.bindNext(next, t.arrows)
			if err != nil {
				return err
			}

			source = source.decrementTailCount()

			return t.nodes.update(source.toNode())
		}
		previous = current
	}
}

func (t *tails) deleteSource(source source, heads *heads) error {
	if !source.hasFirstTail() {
		return nil
	}

	first, err := source.getFirstTail(t.arrows)
	if err != nil {
		return err
	}

	next := first

	err = t.removeTail(source, next)
	if err != nil {
		return err
	}

	err = t.arrows.produceHole(next.toArrow())
	if err != nil {
		return err
	}

	head := next.toHead()
	target, err := head.getTarget(t.nodes)
	if err != nil {
		return err
	}

	err = heads.removeHead(target, head)
	if err != nil {
		return err
	}

	for {
		if !next.hasNext() {
			return nil
		}

		next, err = next.getNext(t.arrows)
		if err != nil {
			return err
		}

		if next.isSame(first) {
			return nil
		}

		err = t.removeTail(source, next)
		if err != nil {
			return err
		}

		err = t.arrows.produceHole(next.toArrow())
		if err != nil {
			return err
		}

		head := next.toHead()
		target, err := head.getTarget(t.nodes)
		if err != nil {
			return err
		}

		err = heads.removeHead(target, head)
		if err != nil {
			return err
		}
	}
}
