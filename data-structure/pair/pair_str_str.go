package pair

type PairStringString struct {
	Key   string
	Value string
}

func NewPairStringString(pkey, pvalue string) PairStringString {
	return PairStringString{
		Key:   pkey,
		Value: pvalue,
	}
}
