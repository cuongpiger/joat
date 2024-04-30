package queue

type threadUnsafeQueue[T comparable] struct {
	queue []T
}

var _ Queue[string] = &threadUnsafeQueue[string]{}

func newThreadUnsafeQueue[T comparable]() *threadUnsafeQueue[T] {
	return &threadUnsafeQueue[T]{
		queue: make([]T, 0),
	}
}

func newThreadUnsafeQueueWithSize[T comparable](psize int) *threadUnsafeQueue[T] {
	return &threadUnsafeQueue[T]{
		queue: make([]T, 0, psize),
	}
}

func (s *threadUnsafeQueue[T]) Push(pval T) bool {
	s.queue = append(s.queue, pval)
	return true
}

func (s *threadUnsafeQueue[T]) Pop() (T, bool) {
	if len(s.queue) == 0 {
		var empty T
		return empty, false
	}

	ret := s.queue[0]
	s.queue = s.queue[1:]
	return ret, true
}

func (s *threadUnsafeQueue[T]) Size() int {
	if s.queue == nil {
		return 0
	}

	return len(s.queue)
}

func (s *threadUnsafeQueue[T]) Empty() bool {
	return s.Size() == 0
}

func (s *threadUnsafeQueue[T]) Top() (T, bool) {
	if s.Empty() {
		var empty T
		return empty, false
	}

	return s.queue[0], true
}
