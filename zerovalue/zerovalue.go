package zerovalue

func New[T any]() T {
	var v T
	return v
}
