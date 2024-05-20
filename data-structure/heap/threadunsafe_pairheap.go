package heap

func newThreadUnsafeHeap() *pairHeap {
	// New returns an initialized PairHeap.
	return new(pairHeap).init()
}

type pairHeap struct {
	root *node
}

func (s *pairHeap) Insert(pitem Item) Item {
	s.root = merge(s.root, &node{item: pitem})
	return pitem
}

// IsEmpty returns true if pairHeap is empty, the complexity is O(1)
func (s *pairHeap) IsEmpty() bool {
	return s.root.item == nil
}

// FindMin finds the smallest item in the heap, the complexity is O(1)
func (s *pairHeap) FindMin() Item {
	if s.IsEmpty() {
		return nil
	}

	return s.root.item
}

// DeleteMin removes the top most value from the PairHeap and returns it
// The complexity is O(log n) amortized.
func (s *pairHeap) DeleteMin() Item {
	return s.deleteItem(nil, removeMin)
}

// Delete deletes a node from the heap and returns the item
// The complexity is O(log n) amortized.
func (s *pairHeap) Delete(pitem Item) Item {
	return s.deleteItem(pitem, removeItem)
}

// Clear resets the current PairHeap
func (s *pairHeap) Clear() {
	s.init()
}

// Adjust adjusts the value to the node item and returns it
// The complexity is O(n) amortized.
func (s *pairHeap) Adjust(pitem, pnew Item) Item {
	n := s.root.findNode(pitem)
	if n == nil {
		return nil
	}

	if n == s.root {
		s.DeleteMin()
		return s.Insert(pnew)
	} else {
		children := n.detach()
		s.Insert(pnew)
		s.root.children = append(s.root.children, children...)
		return n.item
	}
}

// Find exhausting search of the element that matches item and returns it
// The complexity is O(n) amortized.
func (s *pairHeap) Find(pitem Item) Item {
	if s.IsEmpty() {
		return nil
	}
	var found Item
	s.root.iterItem(func(i Item) bool {
		if pitem.Compare(i) == 0 {
			found = i
			return false
		} else {
			return true
		}
	})

	return found
}

// Do calls function cb on each element of the PairingHeap, in order of appearance.
// The behavior of Do is undefined if cb changes *p.
func (s *pairHeap) Do(pit ItemIterator) {
	if s.IsEmpty() {
		return
	}
	s.root.iterItem(pit)
}

func (s *pairHeap) init() *pairHeap {
	s.root = new(node)
	return s
}

func (s *pairHeap) deleteItem(pitem Item, ptype toDelete) Item {
	var result node

	if len(s.root.children) == 0 {
		result = *s.root
		s.root.item = nil
	} else {
		switch ptype {
		case removeMin:
			result = *s.root
			s.root = mergePairs(s.root, s.root.children)
		case removeItem:
			n := s.root.findNode(pitem)
			if n == nil {
				return nil
			} else {
				children := n.detach()
				s.root.children = append(s.root.children, children...)
				result = *n
			}
		default:
			panic("invalid type")
		}
	}

	return result.item
}

type node struct {
	// For use by client, untouched by this library
	item Item

	// A reference to the parent Heap node
	parent *node

	// List of children nodes all containing values less than the Top of the heap
	children []*node
}

func (s *node) detach() []*node {
	if s.parent == nil {
		return nil // avoid detaching root
	}

	for _, n := range s.children {
		n.parent = nil
	}

	var idx int
	for i, n := range s.parent.children {
		if n == s {
			idx = i
			break
		}
	}

	s.parent.children = append(s.parent.children[:idx], s.parent.children[idx+1:]...)
	s.parent = nil
	return s.children
}

func (s *node) iterItem(piter ItemIterator) {
	if !piter(s.item) {
		return
	}
	s.iterChildren(s.children, piter)
}

func (s *node) iterChildren(pchildren []*node, piter ItemIterator) {
	if len(pchildren) == 0 {
		return
	}

	for _, n := range pchildren {
		if !piter(n.item) {
			return
		}
		s.iterChildren(n.children, piter)
	}
}

func (s *node) findNode(pitem Item) *node {
	if s.item.Compare(pitem) == 0 {
		return s
	} else {
		return s.findInChilren(s.children, pitem)
	}
}

func (s *node) findInChilren(pchildren []*node, pitem Item) *node {
	if len(pchildren) == 0 {
		return nil
	}

	var n *node
loop:
	for _, child := range pchildren {
		n = child.findNode(pitem)
		if n != nil {
			break loop
		}
	}

	return n
}

func merge(pna, pnb *node) *node {
	if pna.item == nil { // case when root is empty
		pna = pnb
		return pna
	}

	if pna.item.Compare(pnb.item) < 0 {
		// Put second as the first child of first and update the parent
		pna.children = append([]*node{pnb}, pna.children...)
		pnb.parent = pna
		return pna
	} else {
		// Put first as the first child of second and update the parent
		pnb.children = append([]*node{pna}, pnb.children...)
		pna.parent = pnb
		return pnb
	}
}

func mergePairs(proot *node, pheaps []*node) *node {
	if len(pheaps) == 1 {
		proot = pheaps[0]
		pheaps[0].parent = nil
		return proot
	}

	var merged *node
	for { // iteratively merge heaps
		if len(pheaps) == 0 {
			break
		}
		if merged == nil {
			merged = merge(pheaps[0], pheaps[1])
			pheaps = pheaps[2:]
		} else {
			merged = merge(merged, pheaps[0])
			pheaps = pheaps[1:]
		}
	}
	proot = merged
	merged.parent = nil

	return proot
}
