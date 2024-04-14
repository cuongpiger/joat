package math

func MaxNumeric[T float32 | float64 | int | int8 | int16 | int32 | int64 | uint64 | uint | uint8 | uint32 | uint16](a, b T) T {
	if a > b {
		return a
	}

	return b
}

func MinNumeric[T float32 | float64 | int | int8 | int16 | int32 | int64 | uint64 | uint | uint8 | uint32 | uint16](a, b T) T {
	if a < b {
		return a
	}

	return b
}
