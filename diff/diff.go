package diff

// Assumes that a is a superset of b and that they are sorted.
// TODO need a better name.
func Diff[C comparable](a []C, b []C) []C {
	result := make([]C, 0, len(a)-len(b))
	bIdx := 0
	for _, i := range a {
		if bIdx < len(b) {
			if i == b[bIdx] {
				bIdx++
				continue
			}
		}
		result = append(result, i)
	}

	return result
}
