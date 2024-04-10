package string

func Truncate(pStr string, pLen int) string {
	if len(pStr) <= pLen {
		return pStr
	}

	return pStr[:pLen]
}
