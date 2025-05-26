package utils

func MapArray[T any, K any](data []K, cb func(index int, item K) T) []T {
	result := make([]T, len(data))

	for i, v := range data {
		result[i] = cb(i, v)
	}

	return result
}
