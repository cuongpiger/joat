package heap

type Item interface {
	// Should return a number:
	//    negative , if a < b
	//    zero     , if a == b
	//    positive , if a > b
	Compare(pthan Item) int
}

// ItemIterator allows callers of Do to iterate in-order over portions of the tree. When this function return false,
// iteration will stop and the function will immediately return.
type ItemIterator func(pitem Item) bool

// String implements the Item interface
type String string

// Integer implements the Item interface
type Integer int

func (a String) Compare(b Item) int {
	s1 := a
	s2 := b.(String)
	minVal := len(s2)
	if len(s1) < len(s2) {
		minVal = len(s1)
	}
	diff := 0
	for i := 0; i < minVal && diff == 0; i++ {
		diff = int(s1[i]) - int(s2[i])
	}
	if diff == 0 {
		diff = len(s1) - len(s2)
	}
	if diff < 0 {
		return -1
	}
	if diff > 0 {
		return 1
	}
	return 0
}

func (a Integer) Compare(b Item) int {
	a1 := a
	a2 := b.(Integer)
	switch {
	case a1 > a2:
		return 1
	case a1 < a2:
		return -1
	default:
		return 0
	}
}
