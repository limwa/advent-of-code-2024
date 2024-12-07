package lists

func Contains[T comparable](list []T, item T) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}
	return false
}

func Sum[T int | float64](list []T) T {
	sum := T(0)
	for _, i := range list {
		sum += i
	}
	return sum
}