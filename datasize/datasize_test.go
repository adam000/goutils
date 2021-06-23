package datasize

import (
	"math/big"
	"testing"
)

func TestToHumanReadable(t *testing.T) {
	tests := []struct {
		input     DataSize
		expected  DataSize
		outputStr string
	}{
		{
			input: DataSize{
				magnitude: big.NewFloat(2420),
				unit:      Kibibyte,
			},
			expected: DataSize{
				magnitude: big.NewFloat(2.42),
				unit:      Mebibyte,
			},
			outputStr: "2.42MiB",
		},
	}

	for _, test := range tests {
		actual := test.input.ToHumanReadable()
		if test.expected.unit != actual.unit {
			t.Errorf("Units differ: expected result %s, got %s", test.expected, actual)
		}
		if test.expected.magnitude.Cmp(actual.magnitude) != 0 {
			t.Errorf("Magnitudes differ: expected result %s, got %s", test.expected, actual)
		}
		if test.expected.String() != test.outputStr {
			t.Errorf("String result differs: expected %s, got %s", test.outputStr, test.expected.String())
		}
	}
}
