package diff

import "testing"

func Test_Diff(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}
	b := []int{2, 4, 5}

	allOfA := Diff[int](a, []int{})
	if len(allOfA) != len(a) {
		t.Errorf("Expected empty b to return full list, got %v", allOfA)
	}

	empty := Diff([]int{}, []int{})
	if len(empty) != 0 {
		t.Errorf("Expected empty from empty / empty")
	}

	result := Diff(a, b)
	if len(result) != len(a)-len(b) {
		t.Errorf("Expected diff of %v and %v to be %d long, but got %v", a, b, len(a)-len(b), result)
	}

	same := Diff(a, a)
	if len(same) != 0 {
		t.Errorf("Expected diffing the same thing against itself to come up empty, got %v", same)
	}
}
