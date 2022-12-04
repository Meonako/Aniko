package utils

// Return true if slices contain value.
func Contains[T comparable](slices []T, value T) bool {
	for _, v := range slices {
		if v == value {
			return true
		}
	}

	return false
}

// Remove an element that has value match the param.
func Remove[T comparable](slices []T, value T) []T {
	for i, v := range slices {
		if v == value {
			return append(slices[:i], slices[i+1:]...)
		}
	}
	return slices
}

// Return value of first elemnt.
func First[T any](slices []T) T {
	return slices[0]
}

// Return value of last element.
func Last[T any](slices []T) T {
	return slices[LastIndex(slices)]
}

// Return last index of slices
func LastIndex[T any](slices []T) int {
	return len(slices) - 1
}
