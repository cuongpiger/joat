package set

import (
	"encoding/json"
	"fmt"
	"strings"
)

type threadUnsafeSet[T comparable] map[T]struct{}

// Assert concrete type:threadUnsafeSet adheres to Set interface.
var _ Set[string] = (threadUnsafeSet[string])(nil)

func newThreadUnsafeSet[T comparable]() threadUnsafeSet[T] {
	return make(threadUnsafeSet[T])
}

func newThreadUnsafeSetWithSize[T comparable](cardinality int) threadUnsafeSet[T] {
	return make(threadUnsafeSet[T], cardinality)
}

func (s threadUnsafeSet[T]) Add(v T) bool {
	prevLen := len(s)
	s[v] = struct{}{}
	return prevLen != len(s)
}

func (s threadUnsafeSet[T]) Append(v ...T) int {
	prevLen := len(s)
	for _, val := range v {
		(s)[val] = struct{}{}
	}
	return len(s) - prevLen
}

// private version of Add which doesn't return a value
func (s threadUnsafeSet[T]) add(v T) {
	s[v] = struct{}{}
}

func (s threadUnsafeSet[T]) Cardinality() int {
	return len(s)
}

func (s threadUnsafeSet[T]) Clear() {
	for key := range s {
		delete(s, key)
	}
}

func (s threadUnsafeSet[T]) Clone() Set[T] {
	clonedSet := newThreadUnsafeSetWithSize[T](s.Cardinality())
	for elem := range s {
		clonedSet.add(elem)
	}
	return clonedSet
}

func (s threadUnsafeSet[T]) Contains(v ...T) bool {
	for _, val := range v {
		if _, ok := s[val]; !ok {
			return false
		}
	}
	return true
}

func (s threadUnsafeSet[T]) ContainsOne(v T) bool {
	_, ok := s[v]
	return ok
}

func (s threadUnsafeSet[T]) ContainsAny(v ...T) bool {
	for _, val := range v {
		if _, ok := s[val]; ok {
			return true
		}
	}
	return false
}

// private version of Contains for a single element v
func (s threadUnsafeSet[T]) contains(v T) (ok bool) {
	_, ok = s[v]
	return ok
}

func (s threadUnsafeSet[T]) Difference(other Set[T]) Set[T] {
	o := other.(threadUnsafeSet[T])

	diff := newThreadUnsafeSet[T]()
	for elem := range s {
		if !o.contains(elem) {
			diff.add(elem)
		}
	}
	return diff
}

func (s threadUnsafeSet[T]) Each(cb func(T) bool) {
	for elem := range s {
		if cb(elem) {
			break
		}
	}
}

func (s threadUnsafeSet[T]) Equal(other Set[T]) bool {
	o := other.(threadUnsafeSet[T])

	if s.Cardinality() != other.Cardinality() {
		return false
	}
	for elem := range s {
		if !o.contains(elem) {
			return false
		}
	}
	return true
}

func (s threadUnsafeSet[T]) Intersect(other Set[T]) Set[T] {
	o := other.(threadUnsafeSet[T])

	intersection := newThreadUnsafeSet[T]()
	// loop over smaller set
	if s.Cardinality() < other.Cardinality() {
		for elem := range s {
			if o.contains(elem) {
				intersection.add(elem)
			}
		}
	} else {
		for elem := range o {
			if s.contains(elem) {
				intersection.add(elem)
			}
		}
	}
	return intersection
}

func (s threadUnsafeSet[T]) IsEmpty() bool {
	return s.Cardinality() == 0
}

func (s threadUnsafeSet[T]) IsProperSubset(other Set[T]) bool {
	return s.Cardinality() < other.Cardinality() && s.IsSubset(other)
}

func (s threadUnsafeSet[T]) IsProperSuperset(other Set[T]) bool {
	return s.Cardinality() > other.Cardinality() && s.IsSuperset(other)
}

func (s threadUnsafeSet[T]) IsSubset(other Set[T]) bool {
	o := other.(threadUnsafeSet[T])
	if s.Cardinality() > other.Cardinality() {
		return false
	}
	for elem := range s {
		if !o.contains(elem) {
			return false
		}
	}
	return true
}

func (s threadUnsafeSet[T]) IsSuperset(other Set[T]) bool {
	return other.IsSubset(s)
}

func (s threadUnsafeSet[T]) Iter() <-chan T {
	ch := make(chan T)
	go func() {
		for elem := range s {
			ch <- elem
		}
		close(ch)
	}()

	return ch
}

func (s threadUnsafeSet[T]) Iterator() *Iterator[T] {
	iterator, ch, stopCh := newIterator[T]()

	go func() {
	L:
		for elem := range s {
			select {
			case <-stopCh:
				break L
			case ch <- elem:
			}
		}
		close(ch)
	}()

	return iterator
}

// Pop returns a popped item in case set is not empty, or nil-value of T
// if set is already empty
func (s threadUnsafeSet[T]) Pop() (v T, ok bool) {
	for item := range s {
		delete(s, item)
		return item, true
	}
	return v, false
}

func (s threadUnsafeSet[T]) Remove(v T) {
	delete(s, v)
}

func (s threadUnsafeSet[T]) RemoveAll(i ...T) {
	for _, elem := range i {
		delete(s, elem)
	}
}

func (s threadUnsafeSet[T]) String() string {
	items := make([]string, 0, len(s))

	for elem := range s {
		items = append(items, fmt.Sprintf("%v", elem))
	}
	return fmt.Sprintf("Set{%s}", strings.Join(items, ", "))
}

func (s threadUnsafeSet[T]) SymmetricDifference(other Set[T]) Set[T] {
	o := other.(threadUnsafeSet[T])

	sd := newThreadUnsafeSet[T]()
	for elem := range s {
		if !o.contains(elem) {
			sd.add(elem)
		}
	}
	for elem := range o {
		if !s.contains(elem) {
			sd.add(elem)
		}
	}
	return sd
}

func (s threadUnsafeSet[T]) ToSlice() []T {
	keys := make([]T, 0, s.Cardinality())
	for elem := range s {
		keys = append(keys, elem)
	}

	return keys
}

func (s threadUnsafeSet[T]) Union(other Set[T]) Set[T] {
	o := other.(threadUnsafeSet[T])

	n := s.Cardinality()
	if o.Cardinality() > n {
		n = o.Cardinality()
	}
	unionedSet := make(threadUnsafeSet[T], n)

	for elem := range s {
		unionedSet.add(elem)
	}
	for elem := range o {
		unionedSet.add(elem)
	}
	return unionedSet
}

// MarshalJSON creates a JSON array from the set, it marshals all elements
func (s threadUnsafeSet[T]) MarshalJSON() ([]byte, error) {
	items := make([]string, 0, s.Cardinality())

	for elem := range s {
		b, err := json.Marshal(elem)
		if err != nil {
			return nil, err
		}

		items = append(items, string(b))
	}

	return []byte(fmt.Sprintf("[%s]", strings.Join(items, ","))), nil
}

// UnmarshalJSON recreates a set from a JSON array, it only decodes
// primitive types. Numbers are decoded as json.Number.
func (s threadUnsafeSet[T]) UnmarshalJSON(b []byte) error {
	var i []T
	err := json.Unmarshal(b, &i)
	if err != nil {
		return err
	}
	s.Append(i...)

	return nil
}
