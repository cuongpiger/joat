package queue

type Queue[T comparable] interface {
	Push(pval T) bool
	Pop() (T, bool)
	Size() int
	Empty() bool
	Top() (T, bool)
}

func NewQueue[T comparable](pvals ...T) Queue[T] {
	queue := newThreadSafeQueueWithSize[T](len(pvals))
	for _, item := range pvals {
		queue.Push(item)
	}

	return queue
}

func NewQueueWithSize[T comparable](psize int) Queue[T] {
	return newThreadSafeQueueWithSize[T](psize)
}
