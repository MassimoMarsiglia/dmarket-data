package utils

import "fmt"

var ErrFailedToParse = fmt.Errorf("failed to parse")

func ParseIotas[T ~int](s, name, idx string) (T, error) {
	for i := 0; i < len(idx)-1; i++ {
		if s == name[idx[i]:idx[i+1]] {
			return T(i), nil
		}
	}
	var zero T
	return zero, fmt.Errorf("%w invalid enum value: %q", ErrFailedToParse, s)
}
