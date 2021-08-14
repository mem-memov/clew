package clew

type heads struct {
	nodes  *nodes
	arrows *arrows
}

func newHeads(nodes *nodes, arrows *arrows) *heads {
	return &heads{nodes: nodes, arrows: arrows}
}

func (h *heads) readHeads(target target) ([]position, error) {
	heads := make([]position, 0)

	if !target.hasFirstHead() {
		return heads, nil
	}

	first, err := target.getFirstHead(h.arrows)
	if err != nil {
		return heads, err
	}

	next := first

	heads = append(heads, next.toTail().getSourcePosition())

	for {
		if !next.hasNext() {
			return heads, nil
		}
		next, err = next.getNext(h.arrows)
		if err != nil {
			return heads, err
		}
		if next.isSame(first) {
			return heads, nil
		}
		heads = append(heads, next.toTail().getSourcePosition())
	}
}

func (h *heads) addHead(target target, new head) (arrow, error) {

	if !target.hasFirstHead() {

		target = target.setFirstHead(new)
		err := h.nodes.update(target.toNode())
		if err != nil {
			return arrow{}, err
		}
	} else {

		first, err := target.getFirstHead(h.arrows)
		if err != nil {
			return arrow{}, err
		}

		if first.hasPrevious() {

			last, err := first.getPrevious(h.arrows)
			if err != nil {
				return arrow{}, err
			}
			last = last.setNext(new)
			new = new.setPrevious(last)
			new = new.setNext(first)
			first = first.setPrevious(new)

			err = h.arrows.update(last.toArrow())
			if err != nil {
				return arrow{}, err
			}
		} else {

			first = first.setPrevious(new)
			first = first.setNext(new)
			new = new.setPrevious(first)
			new = new.setNext(first)
		}

		err = h.arrows.update(first.toArrow())
		if err != nil {
			return arrow{}, err
		}
	}

	return new.toArrow(), nil
}

func (h *heads) removeHead(target target, removed head) error {

	first, err := target.getFirstHead(h.arrows)
	if err != nil {
		return err
	}

	if first.isSame(removed) {
		if first.isSurrounded() {
			previous, err := first.getPrevious(h.arrows)
			if err != nil {
				return err
			}

			next, err := first.getNext(h.arrows)
			if err != nil {
				return err
			}

			next, err = previous.bindNext(next, h.arrows)
			if err != nil {
				return err
			}

			target = target.setFirstHead(next)
		} else if first.isPaired() {
			second, err := first.getNext(h.arrows)
			if err != nil {
				return err
			}

			second = second.deletePrevious()
			second = second.deleteNext()
			target = target.setFirstHead(second)
		} else if first.isAlone() {
			target = target.deleteFirstHead()
		}
		return h.nodes.update(target.toNode())
	}

	previous := first

	for {
		current, err := previous.getNext(h.arrows)
		if err != nil {
			return err
		}

		if current.isSame(first) {
			return nil
		}
		if current.isSame(removed) {
			next, err := current.getNext(h.arrows)
			if err != nil {
				return err
			}

			_, err = previous.bindNext(next, h.arrows)
			if err != nil {
				return err
			}

			return nil
		}
		previous = current
	}
}

func (h *heads) deleteTarget(target target) error {
	if !target.hasFirstHead() {
		return nil
	}

	first, err := target.getFirstHead(h.arrows)
	if err != nil {
		return err
	}

	next := first
	next.toTail()
	err = h.arrows.produceHole(next.toArrow())
	if err != nil {
		return err
	}

	for {
		next, err = next.getNext(h.arrows)
		if err != nil {
			return err
		}

		if next.isSame(first) {
			return nil
		}

		err = h.arrows.produceHole(next.toArrow())
		if err != nil {
			return err
		}
	}
}
