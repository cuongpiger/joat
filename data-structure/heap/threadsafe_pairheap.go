package heap

import lsync "sync"

type threadSafeHeap struct {
	lsync.RWMutex
	heap *pairHeap
}

func newThreadSafeHeap() *threadSafeHeap {
	return &threadSafeHeap{
		heap: newThreadUnsafeHeap(),
	}
}

func (s *threadSafeHeap) Insert(pitem Item) Item {
	s.Lock()
	defer s.Unlock()
	return s.heap.Insert(pitem)
}

func (s *threadSafeHeap) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return s.heap.IsEmpty()
}

func (s *threadSafeHeap) FindMin() Item {
	s.RLock()
	defer s.RUnlock()
	return s.heap.FindMin()
}

func (s *threadSafeHeap) Find(pitem Item) Item {
	s.RLock()
	defer s.RUnlock()
	return s.heap.Find(pitem)
}

func (s *threadSafeHeap) DeleteMin() Item {
	s.Lock()
	defer s.Unlock()
	return s.heap.DeleteMin()
}

func (s *threadSafeHeap) Delete(pitem Item) Item {
	s.Lock()
	defer s.Unlock()
	return s.heap.Delete(pitem)
}

func (s *threadSafeHeap) Adjust(pitem, pnew Item) Item {
	s.Lock()
	defer s.Unlock()
	return s.heap.Adjust(pitem, pnew)
}

func (s *threadSafeHeap) Do(pit ItemIterator) {
	s.Lock()
	defer s.Unlock()
	s.heap.Do(pit)
}

func (s *threadSafeHeap) Clear() {
	s.Lock()
	defer s.Unlock()
	s.heap.Clear()
}
