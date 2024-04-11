package set

// Set is the primary interface provided by the mapset package.  It
// represents an unordered set of data and a large number of
// operations that can be applied to that set.
type Set[T comparable] interface {
	// Add adds an element to the set. Returns whether
	// the item was added.
	Add(val T) bool

	// Append multiple elements to the set. Returns
	// the number of elements added.
	Append(val ...T) int

	// Cardinality returns the number of elements in the set.
	Cardinality() int

	// Clear removes all elements from the set, leaving
	// the empty set.
	Clear()

	// Clone returns a clone of the set using the same
	// implementation, duplicating all keys.
	Clone() Set[T]

	// Contains returns whether the given items
	// are all in the set.
	Contains(val ...T) bool

	// ContainsOne returns whether the given item
	// is in the set.
	//
	// Contains may cause the argument to escape to the heap.
	// See: https://github.com/deckarep/golang-set/issues/118
	ContainsOne(val T) bool

	// ContainsAny returns whether at least one of the
	// given items are in the set.
	ContainsAny(val ...T) bool

	// Difference returns the difference between this set
	// and other. The returned set will contain
	// all elements of this set that are not also
	// elements of other.
	//
	// Note that the argument to Difference
	// must be of the same type as the receiver
	// of the method. Otherwise, Difference will
	// panic.
	Difference(other Set[T]) Set[T]

	// Equal determines if two sets are equal to each
	// other. If they have the same cardinality
	// and contain the same elements, they are
	// considered equal. The order in which
	// the elements were added is irrelevant.
	//
	// Note that the argument to Equal must be
	// of the same type as the receiver of the
	// method. Otherwise, Equal will panic.
	Equal(other Set[T]) bool

	// Intersect returns a new set containing only the elements
	// that exist only in both sets.
	//
	// Note that the argument to Intersect
	// must be of the same type as the receiver
	// of the method. Otherwise, Intersect will
	// panic.
	Intersect(other Set[T]) Set[T]

	// IsEmpty determines if there are elements in the set.
	IsEmpty() bool

	// IsProperSubset determines if every element in this set is in
	// the other set but the two sets are not equal.
	//
	// Note that the argument to IsProperSubset
	// must be of the same type as the receiver
	// of the method. Otherwise, IsProperSubset
	// will panic.
	IsProperSubset(other Set[T]) bool

	// IsProperSuperset determines if every element in the other set
	// is in this set but the two sets are not
	// equal.
	//
	// Note that the argument to IsSuperset
	// must be of the same type as the receiver
	// of the method. Otherwise, IsSuperset will
	// panic.
	IsProperSuperset(other Set[T]) bool

	// IsSubset determines if every element in this set is in
	// the other set.
	//
	// Note that the argument to IsSubset
	// must be of the same type as the receiver
	// of the method. Otherwise, IsSubset will
	// panic.
	IsSubset(other Set[T]) bool

	// IsSuperset determines if every element in the other set
	// is in this set.
	//
	// Note that the argument to IsSuperset
	// must be of the same type as the receiver
	// of the method. Otherwise, IsSuperset will
	// panic.
	IsSuperset(other Set[T]) bool

	// Each iterates over elements and executes the passed func against each element.
	// If passed func returns true, stop iteration at the time.
	Each(func(T) bool)

	// Iter returns a channel of elements that you can
	// range over.
	Iter() <-chan T

	// Iterator returns an Iterator object that you can
	// use to range over the set.
	Iterator() *Iterator[T]

	// Remove removes a single element from the set.
	Remove(i T)

	// RemoveAll removes multiple elements from the set.
	RemoveAll(i ...T)

	// String provides a convenient string representation
	// of the current state of the set.
	String() string

	// SymmetricDifference returns a new set with all elements which are
	// in either this set or the other set but not in both.
	//
	// Note that the argument to SymmetricDifference
	// must be of the same type as the receiver
	// of the method. Otherwise, SymmetricDifference
	// will panic.
	SymmetricDifference(other Set[T]) Set[T]

	// Union returns a new set with all elements in both sets.
	//
	// Note that the argument to Union must be of the
	// same type as the receiver of the method.
	// Otherwise, Union will panic.
	Union(other Set[T]) Set[T]

	// Pop removes and returns an arbitrary item from the set.
	Pop() (T, bool)

	// ToSlice returns the members of the set as a slice.
	ToSlice() []T

	// MarshalJSON will marshal the set into a JSON-based representation.
	MarshalJSON() ([]byte, error)

	// UnmarshalJSON will unmarshal a JSON-based byte slice into a full Set datastructure.
	// For this to work, set subtypes must implemented the Marshal/Unmarshal interface.
	UnmarshalJSON(b []byte) error
}

// NewSet creates and returns a new set with the given elements.
// Operations on the resulting set are thread-safe.
func NewSet[T comparable](vals ...T) Set[T] {
	s := newThreadSafeSetWithSize[T](len(vals))
	for _, item := range vals {
		s.Add(item)
	}
	return s
}

// NewSetWithSize creates and returns a reference to an empty set with a specified
// capacity. Operations on the resulting set are thread-safe.
func NewSetWithSize[T comparable](cardinality int) Set[T] {
	s := newThreadSafeSetWithSize[T](cardinality)
	return s
}

// NewThreadUnsafeSet creates and returns a new set with the given elements.
// Operations on the resulting set are not thread-safe.
func NewThreadUnsafeSet[T comparable](vals ...T) Set[T] {
	s := newThreadUnsafeSetWithSize[T](len(vals))
	for _, item := range vals {
		s.Add(item)
	}
	return s
}

// NewThreadUnsafeSetWithSize creates and returns a reference to an empty set with
// a specified capacity. Operations on the resulting set are not thread-safe.
func NewThreadUnsafeSetWithSize[T comparable](cardinality int) Set[T] {
	s := newThreadUnsafeSetWithSize[T](cardinality)
	return s
}

// NewSetFromMapKeys creates and returns a new set with the given keys of the map.
// Operations on the resulting set are thread-safe.
func NewSetFromMapKeys[T comparable, V any](val map[T]V) Set[T] {
	s := NewSetWithSize[T](len(val))

	for k := range val {
		s.Add(k)
	}

	return s
}

// NewThreadUnsafeSetFromMapKeys creates and returns a new set with the given keys of the map.
// Operations on the resulting set are not thread-safe.
func NewThreadUnsafeSetFromMapKeys[T comparable, V any](val map[T]V) Set[T] {
	s := NewThreadUnsafeSetWithSize[T](len(val))

	for k := range val {
		s.Add(k)
	}

	return s
}