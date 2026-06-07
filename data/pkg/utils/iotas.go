package utils

import "fmt"

var ErrFailedToParse = fmt.Errorf("failed to parse")

func ParseIotas[T ~int](name string, table string, index []uint8) (T, error) {
	for i := 0; i < len(index)-1; i++ {
		s := table[index[i]:index[i+1]]
		if s == name {
			return T(i), nil
		}
	}
	var zero T
	return zero, fmt.Errorf("invalid value: %q", name)
}
