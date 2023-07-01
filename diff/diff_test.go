package diff

import (
	"reflect"
	"testing"
)

func Test_DisjunctiveUnion(t *testing.T) {
	tests := []struct {
		a      []int
		b      []int
		notInA []int
		notInB []int
	}{
		{
			[]int{1, 2, 3, 8, 9, 10},
			[]int{3, 4, 5},
			[]int{4, 5},
			[]int{1, 2, 8, 9, 10},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3},
			[]int{},
			[]int{},
		},
		{
			[]int{1, 2, 3},
			[]int{1, 2, 3, 4},
			[]int{4},
			[]int{},
		},
	}

	for _, test := range tests {
		notInA, notInB := DisjunctiveUnion(test.a, test.b)

		if !reflect.DeepEqual(notInA, test.notInA) {
			t.Errorf("Expected notInA to be %v, got %v", test.notInA, notInA)
		}
		if !reflect.DeepEqual(notInB, test.notInB) {
			t.Errorf("Expected notInB to be %v, got %v", test.notInB, notInB)
		}
	}
}
