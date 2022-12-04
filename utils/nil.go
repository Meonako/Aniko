package utils

func GetNonNil[T string | int | float64](value any) T {
	if value != nil {
		return value.(T)
	}

	return GetZeroValue[T]()
}

func GetZeroValue[T string | int | float64]() T {
	var val T
	return val
}
