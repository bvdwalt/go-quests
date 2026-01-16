package functions

import (
	"errors"
	"strings"
)

// Divide returns a / b or an error if b == 0
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// SumAll returns the sum of all provided integers
func SumAll(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

// MaxMin returns the max and min of all provided integers
// Returns an error if no numbers are provided
func MaxMin(nums ...int) (int, int, error) {
	if len(nums) == 0 {
		return 0, 0, errors.New("no numbers provided")
	}
	max, min := nums[0], nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return max, min, nil
}

// ConcatAll joins all strings using the provided separator
func ConcatAll(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	var builder strings.Builder
	for i, s := range strs {
		if i > 0 {
			builder.WriteString(sep)
		}
		builder.WriteString(s)
	}
	return builder.String()
}
