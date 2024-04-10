//go:build go1.21
// +build go1.21

package set

import (
	"cmp"
	"slices"
)

// Sorted returns a sorted slice of a set of any ordered type in ascending order.
// When sorting floating-point numbers, NaNs are ordered before other values.
func Sorted[E cmp.Ordered](set Set[E]) []E {
	s := set.ToSlice()
	slices.Sort(s)
	return s
}
