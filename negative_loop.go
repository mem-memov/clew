package klubok

type negativeLoop struct {
	vertices vertices
	edges    edges
}

func newNegativeLoop(vertices vertices, edges edges) negativeLoop {
	return negativeLoop{vertices: vertices, edges: edges}
}

func (n negativeLoop) streamNextVertex(done <-chan struct{}, head vertex, tailStream chan<- uint) {
	if !head.hasFirstPositiveEdge() {
		return
	}

	firstEdge := head.getFirstNegativeEdge(n.edges)
	nextEdge := firstEdge

	select {
	case <-done:
		return
	default:
		nextEdge.sendNegativeVertex(tailStream)
	}

	for {
		select {
		case <-done:
			return
		default:
			nextEdge = nextEdge.getNextNegativeEdge(n.edges)
			if nextEdge.getPosition() == firstEdge.getPosition() {
				return
			}
			nextEdge.sendNegativeVertex(tailStream)
		}
	}
}

func (n negativeLoop) streamPreviousVertex(done <-chan struct{}, head vertex, tailStream chan<- uint) {
	if !head.hasFirstPositiveEdge() {
		return
	}

	firstEdge := head.getFirstNegativeEdge(n.edges)
	if firstEdge.hasDifferentPreviousNegativeEdge() {
		firstEdge = firstEdge.getPreviousNegativeEdge(n.edges)
	}
	previousEdge := firstEdge

	select {
	case <-done:
		return
	default:
		previousEdge.sendNegativeVertex(tailStream)
	}

	for {
		select {
		case <-done:
			return
		default:
			previousEdge = previousEdge.getPreviousNegativeEdge(n.edges)
			if previousEdge.getPosition() == firstEdge.getPosition() {
				return
			}
			previousEdge.sendNegativeVertex(tailStream)
		}
	}
}
