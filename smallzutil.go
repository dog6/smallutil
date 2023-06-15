package smallzutil

import (
	"fmt"
	"math"
)

func JoinStringArray(arr []string) string {
	var s string
	for a := range arr {
		s = fmt.Sprintf("%v%v", s, arr[a])
	}
	return s
}

func AddIntArray(arr []int) int {
	s := 0
	for a := range arr {
		s += arr[a]
	}
	return s
}

func MultiplyIntArray(arr []int) int {
	s := 1
	for a := range arr {
		s = s * arr[a]
	}
	return s
}

func FloatToInt(f float64) int {
	i := int(math.Round(f))
	return i
}
