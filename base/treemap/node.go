package treemap

type node[V any] struct {
	right *node[V]
	left  *node[V]
	black bool
	key   []byte
	value V
}

func (p *node[V]) flip() {
	p.black = !p.black
	p.left.black = !p.left.black
	p.right.black = !p.right.black
}

func (p *node[V]) moveRedLeft() *node[V] {
	p.flip()

	if isRed(p.right.left) {
		p.right = p.right.rotateRight()
		elem := p.rotateLeft()
		elem.flip()

		return elem
	}

	return p
}

func (p *node[V]) rotateLeft() *node[V] {
	right := p.right
	if right.black {
		panic(ErrBlack)
	}

	p.right = right.left
	right.left = p
	right.black = p.black
	p.black = false

	return right
}

func (p *node[V]) rotateRight() *node[V] {
	left := p.left
	if left.black {
		panic(ErrBlack)
	}

	p.left = left.right
	left.right = p
	left.black = p.black
	p.black = false

	return left
}

func (p *node[V]) moveRedRight() *node[V] {
	p.flip()

	if isRed(p.left.left) {
		elem := p.rotateRight()
		elem.flip()

		return elem
	}

	return p
}

func (p *node[V]) fixUp() *node[V] {
	elem := p

	if isRed(elem.right) {
		elem = elem.rotateLeft()
	}

	if isRed(elem.left) && isRed(elem.left.left) {
		elem = elem.rotateRight()
	}

	if isRed(elem.left) && isRed(elem.right) {
		elem.flip()
	}

	return elem
}

func (p *node[V]) walkUpRot23() *node[V] {
	elem := p

	if isRed(elem.right) && !isRed(elem.left) {
		elem = elem.rotateLeft()
	}

	if isRed(elem.left) && isRed(elem.left.left) {
		elem = elem.rotateRight()
	}

	if isRed(elem.left) && isRed(elem.right) {
		elem.flip()
	}

	return elem
}

func isRed[V any](elem *node[V]) bool {
	if elem == nil {
		return false
	}

	return !elem.black
}
