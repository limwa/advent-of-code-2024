package cast

import (
	"strconv"
)

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func ToString(i int) string {
	return strconv.Itoa(i)
}

func ToIntSlice(arr []string) []int {
	result := make([]int, len(arr))

	for i, s := range arr {
		result[i] = ToInt(s)
	}

	return result
}
