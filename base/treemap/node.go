package treemap

import "golang.org/x/exp/constraints"

type node[K constraints.Ordered, V any] struct {
	right *node[K, V]
	left  *node[K, V]
	black bool
	key   K
	value V
}

func (p *node[K, V]) flip() {
	p.black = !p.black
	p.left.black = !p.left.black
	p.right.black = !p.right.black
}

func (p *node[K, V]) moveRedLeft() *node[K, V] {
	p.flip()

	if isRed(p.right.left) {
		p.right = p.right.rotateRight()
		elem := p.rotateLeft()
		elem.flip()

		return elem
	}

	return p
}

func (p *node[K, V]) rotateLeft() *node[K, V] {
	right := p.right
	// if right.black {
	// 	panic(ErrBlack)
	// }
	p.right = right.left
	right.left = p
	right.black = p.black
	p.black = false

	return right
}

func (p *node[K, V]) rotateRight() *node[K, V] {
	left := p.left
	// if left.black {
	// 	panic(ErrBlack)
	// }
	p.left = left.right
	left.right = p
	left.black = p.black
	p.black = false

	return left
}

func (p *node[K, V]) moveRedRight() *node[K, V] {
	p.flip()

	if isRed(p.left.left) {
		elem := p.rotateRight()
		elem.flip()

		return elem
	}

	return p
}

func (p *node[K, V]) fixUp() *node[K, V] {
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

func (p *node[K, V]) walkUpRot23() *node[K, V] {
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

func isRed[K constraints.Ordered, V any](elem *node[K, V]) bool {
	if elem == nil {
		return false
	}

	return !elem.black
}
