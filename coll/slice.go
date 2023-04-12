package coll

func Map[S any, T any](source []S, fn func(S) T) []T {
	var target []T
	for _, s := range source {
		target = append(target, fn(s))
	}
	return target
}
