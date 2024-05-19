package workqueue

type RateLimitingInterface TypedRateLimitingInterface[any]
type TypedRateLimitingInterface[T comparable] interface {
	TypedDeplayingInterface[T]
}
