package treemap

import (
	"golang.org/x/exp/constraints"
)

// TreeMap 左倾红黑树Map.
// left-leaning red-black tree map.
type TreeMap[K constraints.Ordered, V any] struct {
	root          *node[K, V]
	count         int
	notFoundKey   K
	notFoundValue V
}

// New 新建 TreeMap.
// notFound 找不到时返回的值.
func New[K constraints.Ordered, V any](noFoundKey K, notFoundValue V) *TreeMap[K, V] {
	return &TreeMap[K, V]{
		notFoundKey:   noFoundKey,
		notFoundValue: notFoundValue,
	}
}

// Len 长度.
func (p *TreeMap[K, V]) Len() int {
	return p.count
}

// Get 根据 key 获取 Value.
func (p *TreeMap[K, V]) Get(key K) (V, bool) {
	elem := p.root

	for elem != nil {
		if key == elem.key {
			return elem.value, true
		}

		if key < elem.key {
			elem = elem.left
		} else {
			elem = elem.right
		}
	}

	return p.notFoundValue, false
}

// Set 覆盖，返回是之前否存在.
func (p *TreeMap[K, V]) Set(key K, value V) bool {
	var add bool
	p.root, add = p.set(p.root, key, value)
	p.root.black = true

	if add {
		p.count++
	}

	return add
}

func (p *TreeMap[K, V]) set(elem *node[K, V], key K, value V) (*node[K, V], bool) {
	if elem == nil {
		return &node[K, V]{key: key, value: value}, true
	}
	// elem = walkDownRot23(elem)
	add := false

	if key == elem.key {
		elem.value = value

		return elem, false
	}

	if key < elem.key {
		elem.left, add = p.set(elem.left, key, value)
	} else {
		elem.right, add = p.set(elem.right, key, value)
	}

	return elem.walkUpRot23(), add
}

// Add 增加，如果存在则忽略，返回之前是否存在.
func (p *TreeMap[K, V]) Add(key K, value V) bool {
	add := false
	p.root, add = p.add(p.root, key, value)
	p.root.black = true

	if add {
		p.count++
	}

	return add
}

func (p *TreeMap[K, V]) add(elem *node[K, V], key K, value V) (*node[K, V], bool) {
	if elem == nil {
		return &node[K, V]{key: key, value: value}, true
	}
	// elem = walkDownRot23(elem)
	add := false

	if key == elem.key {
		return elem, false
	}

	if key < elem.key {
		elem.left, add = p.add(elem.left, key, value)
	} else {
		elem.right, add = p.add(elem.right, key, value)
	}

	return elem.walkUpRot23(), add
}

// Clear 清空.
func (p *TreeMap[K, V]) Clear() {
	p.count = 0
	p.root = nil
}

// Min 最小键.
func (p *TreeMap[K, V]) Min() (V, K) {
	elem := p.root

	if elem == nil {
		return p.notFoundValue, p.notFoundKey
	}

	for elem.left != nil {
		elem = elem.left
	}

	return elem.value, elem.key
}

// Max 最大键.
func (p *TreeMap[K, V]) Max() (V, K) {
	elem := p.root

	if elem == nil {
		return p.notFoundValue, p.notFoundKey
	}

	for elem.right != nil {
		elem = elem.right
	}

	return elem.value, elem.key
}

// Del 删除.
func (p *TreeMap[K, V]) Del(key K) bool {
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

func (p *TreeMap[K, V]) del(elem *node[K, V], key K) (*node[K, V], bool) {
	if elem == nil {
		return nil, false
	}

	if key < elem.key {
		return p.delLeft(elem, key)
	}

	if isRed(elem.left) {
		elem = elem.rotateRight()
	}

	if elem.key >= key && elem.right == nil {
		return nil, true
	}

	if elem.right != nil && !isRed(elem.right) && !isRed(elem.right.left) {
		elem = elem.moveRedRight()
	}

	deleted := false

	if elem.key >= key {
		deleted = p.delRight(elem)
	} else {
		elem.right, deleted = p.del(elem.right, key)
	}

	return elem.fixUp(), deleted
}

func (p *TreeMap[K, V]) delRight(elem *node[K, V]) bool {
	var (
		subDeleted K
		value      V
	)

	elem.right, subDeleted, value = p.delMin(elem.right)

	elem.key = subDeleted
	elem.value = value

	return true
}

func (p *TreeMap[K, V]) delLeft(elem *node[K, V], key K) (*node[K, V], bool) {
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

func (p *TreeMap[K, V]) delMin(elem *node[K, V]) (*node[K, V], K, V) {
	if elem == nil {
		return nil, p.notFoundKey, p.notFoundValue
	}

	if elem.left == nil {
		return nil, elem.key, elem.value
	}

	if !isRed(elem.left) && !isRed(elem.left.left) {
		elem = elem.moveRedLeft()
	}

	var (
		deleted K
		value   V
	)

	elem.left, deleted, value = p.delMin(elem.left)

	return elem.fixUp(), deleted, value
}

// DelMin 删除最小键.
func (p *TreeMap[K, V]) DelMin() bool {
	var deleted K
	p.root, deleted, _ = p.delMin(p.root)

	if p.root != nil {
		p.root.black = true
	}

	if deleted != p.notFoundKey {
		p.count--

		return true
	}

	return false
}

// DelMin 删除最大键.
func (p *TreeMap[K, V]) DelMax() bool {
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

func delMax[K constraints.Ordered, V any](elem *node[K, V]) (*node[K, V], bool) {
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
