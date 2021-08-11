package klubok

type positiveLoop struct {
	vertices vertices
	edges    edges
}

func newPositiveLoop(vertices vertices, edges edges) positiveLoop {
	return positiveLoop{vertices: vertices, edges: edges}
}

func (p positiveLoop) streamNextVertex(done <-chan struct{}, tail vertex, headStream chan<- uint) {
	if !tail.hasFirstPositiveEdge() {
		return
	}

	firstEdge := tail.getFirstPositiveEdge(p.edges)
	nextEdge := firstEdge

	select {
	case <-done:
		return
	default:
		nextEdge.sendPositiveVertex(headStream)
	}

	for {
		select {
		case <-done:
			return
		default:
			nextEdge = nextEdge.getNextPositiveEdge(p.edges)
			if nextEdge.getPosition() == firstEdge.getPosition() {
				return
			}
			nextEdge.sendPositiveVertex(headStream)
		}
	}
}

func (p positiveLoop) streamPreviousVertex(done <-chan struct{}, tail vertex, headStream chan<- uint) {
	if !tail.hasFirstPositiveEdge() {
		return
	}

	firstEdge := tail.getFirstPositiveEdge(p.edges)
	if firstEdge.hasDifferentPreviousPositiveEdge() {
		firstEdge = firstEdge.getPreviousPositiveEdge(p.edges)
	}
	previousEdge := firstEdge

	select {
	case <-done:
		return
	default:
		previousEdge.sendPositiveVertex(headStream)
	}

	for {
		select {
		case <-done:
			return
		default:
			previousEdge = previousEdge.getPreviousPositiveEdge(p.edges)
			if previousEdge.getPosition() == firstEdge.getPosition() {
				return
			}
			previousEdge.sendPositiveVertex(headStream)
		}
	}
}


