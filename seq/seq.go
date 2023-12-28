package seq

func Last[T any](ts []T) T {
	return ts[len(ts)-1]
}
