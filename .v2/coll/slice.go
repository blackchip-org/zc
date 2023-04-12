package coll

func Reverse[T any](source []T) []T {
	var target []T
	for _, v := range source {
		target = append([]T{v}, target...)
	}
	return target
}
