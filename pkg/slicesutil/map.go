package slicesutil

// Map transform a slice of A into a new slice of B
func Map[A, B any](s []A, f func(A) B) []B {
	r := make([]B, 0, len(s))
	for _, v := range s {
		r = append(r, f(v))
	}

	return r
}
