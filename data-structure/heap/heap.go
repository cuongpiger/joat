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

// toDelete details what item to remove in a node call.
type toDelete int

const (
	removeItem toDelete = iota // removes the given item
	removeMin                  // removes min item in the heap
)
