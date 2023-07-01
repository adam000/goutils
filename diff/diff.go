package diff

import "golang.org/x/exp/constraints"

// DisjunctiveUnion returns 2 slices containing the values not shared by both.
// DisjunctiveUnion assumes that there are not duplicate values in the input
// slices and that the input slices are sorted.
func DisjunctiveUnion[C constraints.Ordered](a []C, b []C) ([]C, []C) {
	notInA := make([]C, 0)
	notInB := make([]C, 0)
	aIdx := 0
	bIdx := 0

	for aIdx < len(a) && bIdx < len(b) {
		if a[aIdx] < b[bIdx] {
			notInB = append(notInB, a[aIdx])
			aIdx++
		} else if a[aIdx] == b[bIdx] {
			aIdx++
			bIdx++
		} else {
			notInA = append(notInA, b[bIdx])
			bIdx++
		}
	}

	for aIdx < len(a) {
		notInB = append(notInB, a[aIdx])
		aIdx++
	}

	for bIdx < len(b) {
		notInA = append(notInA, b[bIdx])
		bIdx++
	}

	return notInA, notInB
}
