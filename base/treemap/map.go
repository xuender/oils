package treemap

import (
	"bytes"
)

// TreeMap 左倾红黑树Map.
// left-leaning red-black tree map.
type TreeMap[V any] struct {
	root     *node[V]
	count    int
	notFound V
}

// New 新建 TreeMap.
// notFound 找不到时返回的值.
func New[V any](notFound V) *TreeMap[V] {
	return &TreeMap[V]{
		notFound: notFound,
	}
}

// Len 长度.
func (p *TreeMap[V]) Len() int {
	return p.count
}

// Get 根据 key 获取 Value.
func (p *TreeMap[V]) Get(key []byte) (V, bool) {
	elem := p.root

	for elem != nil {
		switch bytes.Compare(key, elem.key) {
		case -1:
			elem = elem.left
		case 1:
			elem = elem.right
		default:
			return elem.value, true
		}
	}

	return p.notFound, false
}

// Set 覆盖，返回是之前否存在.
func (p *TreeMap[V]) Set(key []byte, value V) bool {
	var add bool
	p.root, add = p.set(p.root, key, value)
	p.root.black = true

	if add {
		p.count++
	}

	return add
}

func (p *TreeMap[V]) set(elem *node[V], key []byte, value V) (*node[V], bool) {
	if elem == nil {
		return &node[V]{key: key, value: value}, true
	}
	// elem = walkDownRot23(elem)
	add := false

	switch bytes.Compare(key, elem.key) {
	case -1:
		elem.left, add = p.set(elem.left, key, value)
	case 1:
		elem.right, add = p.set(elem.right, key, value)
	default:
		elem.value = value

		return elem, false
	}

	return elem.walkUpRot23(), add
}

// Add 增加，如果存在则忽略，返回之前是否存在.
func (p *TreeMap[V]) Add(key []byte, value V) bool {
	add := false
	p.root, add = p.add(p.root, key, value)
	p.root.black = true

	if add {
		p.count++
	}

	return add
}

func (p *TreeMap[V]) add(elem *node[V], key []byte, value V) (*node[V], bool) {
	if elem == nil {
		return &node[V]{key: key, value: value}, true
	}
	// elem = walkDownRot23(elem)
	add := false

	switch bytes.Compare(key, elem.key) {
	case -1:
		elem.left, add = p.add(elem.left, key, value)
	case 1:
		elem.right, add = p.add(elem.right, key, value)
	default:
		return elem, false
	}

	return elem.walkUpRot23(), add
}

// Clear 清空.
func (p *TreeMap[V]) Clear() {
	p.count = 0
	p.root = nil
}

// Min 最小键.
func (p *TreeMap[V]) Min() (V, []byte) {
	elem := p.root

	if elem == nil {
		return p.notFound, nil
	}

	for elem.left != nil {
		elem = elem.left
	}

	return elem.value, elem.key
}

// Max 最大键.
func (p *TreeMap[V]) Max() (V, []byte) {
	elem := p.root

	if elem == nil {
		return p.notFound, nil
	}

	for elem.right != nil {
		elem = elem.right
	}

	return elem.value, elem.key
}

// Del 删除.
func (p *TreeMap[V]) Del(key []byte) bool {
	var deleted bool
	p.root, deleted = p.del(p.root, key)

	if p.root != nil {
		p.root.black = true
	}

	if deleted {
		p.count--
	}

	return deleted
}

func (p *TreeMap[V]) del(elem *node[V], key []byte) (*node[V], bool) {
	if elem == nil {
		return nil, false
	}

	if bytes.Compare(key, elem.key) < 0 {
		return p.delLeft(elem, key)
	}

	if isRed(elem.left) {
		elem = elem.rotateRight()
	}

	if bytes.Compare(elem.key, key) >= 0 && elem.right == nil {
		return nil, true
	}

	if elem.right != nil && !isRed(elem.right) && !isRed(elem.right.left) {
		elem = elem.moveRedRight()
	}

	deleted := false

	if bytes.Compare(elem.key, key) >= 0 {
		deleted = p.delRight(elem)
	} else {
		elem.right, deleted = p.del(elem.right, key)
	}

	return elem.fixUp(), deleted
}

func (p *TreeMap[V]) delRight(elem *node[V]) bool {
	var (
		subDeleted []byte
		value      V
	)

	elem.right, subDeleted, value = p.delMin(elem.right)

	elem.key = subDeleted
	elem.value = value

	return true
}

func (p *TreeMap[V]) delLeft(elem *node[V], key []byte) (*node[V], bool) {
	if elem.left == nil {
		return elem, false
	}

	if !isRed(elem.left) && !isRed(elem.left.left) {
		elem = elem.moveRedLeft()
	}

	deleted := false
	elem.left, deleted = p.del(elem.left, key)

	return elem.fixUp(), deleted
}

func (p *TreeMap[V]) delMin(elem *node[V]) (*node[V], []byte, V) {
	if elem == nil {
		return nil, nil, p.notFound
	}

	if elem.left == nil {
		return nil, elem.key, elem.value
	}

	if !isRed(elem.left) && !isRed(elem.left.left) {
		elem = elem.moveRedLeft()
	}

	var (
		deleted []byte
		value   V
	)

	elem.left, deleted, value = p.delMin(elem.left)

	return elem.fixUp(), deleted, value
}

// DelMin 删除最小键.
func (p *TreeMap[V]) DelMin() bool {
	var deleted []byte
	p.root, deleted, _ = p.delMin(p.root)

	if p.root != nil {
		p.root.black = true
	}

	if deleted != nil {
		p.count--

		return true
	}

	return false
}

// DelMin 删除最大键.
func (p *TreeMap[V]) DelMax() bool {
	deleted := false
	p.root, deleted = delMax(p.root)

	if p.root != nil {
		p.root.black = true
	}

	if deleted {
		p.count--
	}

	return deleted
}

func delMax[V any](elem *node[V]) (*node[V], bool) {
	if elem == nil {
		return nil, false
	}

	if isRed(elem.left) {
		elem = elem.rotateRight()
	}

	if elem.right == nil {
		return nil, true
	}

	if !isRed(elem.right) && !isRed(elem.right.left) {
		elem = elem.moveRedRight()
	}

	deleted := false
	elem.right, deleted = delMax(elem.right)

	return elem.fixUp(), deleted
}
