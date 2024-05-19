package workqueue

type TypedDeplayingInterface[T comparable] interface {
	TypedInterface[T]
}
