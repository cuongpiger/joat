package trio

type Trio[T1, T2, T3 any] struct {
	First  T1
	Second T2
	Third  T3
}

func NewTrio[T1, T2, T3 any](pfirst T1, psecond T2, pthird T3) Trio[T1, T2, T3] {
	return Trio[T1, T2, T3]{
		First:  pfirst,
		Second: psecond,
		Third:  pthird,
	}
}
