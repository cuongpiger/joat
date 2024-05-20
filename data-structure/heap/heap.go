package heap

func NewHeap(pvals ...Item) Heap {
	heap := newThreadSafeHeap()
	for _, item := range pvals {
		heap.Insert(item)
	}

	return heap
}

type Heap interface {
	Insert(pval Item) Item
	IsEmpty() bool
	FindMin() Item
	Find(pitem Item) Item
	DeleteMin() Item
	Delete(pitem Item) Item
	Adjust(pitem, pnew Item) Item
	Do(pit ItemIterator)
	Clear()
}

// Interface is basic interface that all Heaps implement.
type Interface interface {
	// Insert add an element to the heap and returns it
	Insert(v Item) Item

	// DeleteMin deletes and returns the smallest element
	DeleteMin() Item

	// FindMin returns the minimum element
	FindMin() Item

	// Clear removes all items
	Clear()
}

// toDelete details what item to remove in a node call.
type toDelete int

const (
	removeItem toDelete = iota // removes the given item
	removeMin                  // removes min item in the heap
)
