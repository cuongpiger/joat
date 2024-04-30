package queue

import (
	lsync "sync"
)

type threadSafeQueue[T comparable] struct {
	lsync.RWMutex
	usq *threadUnsafeQueue[T]
}

func newThreadSafeQueue[T comparable]() *threadSafeQueue[T] {
	return &threadSafeQueue[T]{
		usq: newThreadUnsafeQueue[T](),
	}
}

func newThreadSafeQueueWithSize[T comparable](psize int) *threadSafeQueue[T] {
	return &threadSafeQueue[T]{
		usq: newThreadUnsafeQueueWithSize[T](psize),
	}
}

func (s *threadSafeQueue[T]) Push(pval T) bool {
	s.Lock()
	defer s.Unlock()
	return s.usq.Push(pval)
}

func (s *threadSafeQueue[T]) Pop() (T, bool) {
	s.Lock()
	defer s.Unlock()
	return s.usq.Pop()
}

func (s *threadSafeQueue[T]) Size() int {
	s.RLock()
	defer s.RUnlock()
	return s.usq.Size()
}

func (s *threadSafeQueue[T]) Empty() bool {
	s.RLock()
	defer s.RUnlock()
	return s.usq.Empty()
}

func (s *threadSafeQueue[T]) Top() (T, bool) {
	s.RLock()
	defer s.RUnlock()
	return s.usq.Top()
}
