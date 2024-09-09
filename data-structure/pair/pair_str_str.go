package pair

type PairStringString struct {
	First  string
	Second string
}

func NewPairStringString(pfirst, psecond string) PairStringString {
	return PairStringString{
		First:  pfirst,
		Second: psecond,
	}
}
