package utils

func Point[T any](val T) *T {
	return &val
}
