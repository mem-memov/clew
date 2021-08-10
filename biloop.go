package klubok

type biloop struct {
	positive loop
	negative loop
}

func newBiloop(positive loop, negative loop) biloop {
	return biloop{positive: positive, negative: negative}
}

func (b biloop) getPosition() position {
	return b.positive.
}
