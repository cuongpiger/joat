package queue

type Queue[T comparable] interface {
	Push(pval T) bool
	Pop() (T, bool)
	Size() int
	Empty() bool
	Top() (T, bool)
}

func NewQueue[T comparable]() Queue[T] {
	return newThreadUnsafeQueue[T]()
}
