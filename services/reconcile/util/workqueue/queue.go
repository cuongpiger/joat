package workqueue

type Interface TypedInterface[any]

type TypedInterface[T comparable] interface {
	ShutDown()
}
