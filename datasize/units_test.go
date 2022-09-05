package datasize

import (
	"testing"
)

func TestUnitFromString(t *testing.T) {
	tests := []struct {
		input     string
		assumeSi  bool
		unit      Unit
		expectErr bool
	}{
		{
			input:     "KiB",
			assumeSi:  false,
			unit:      Kibibyte,
			expectErr: false,
		},
		{
			input:     "KB",
			assumeSi:  false,
			unit:      Kilobyte,
			expectErr: false,
		},
		{
			input:     "KB",
			assumeSi:  true,
			unit:      Kilobyte,
			expectErr: false,
		},
		{
			input:     "K",
			assumeSi:  false,
			unit:      Kilobyte,
			expectErr: false,
		},
		{
			input:     "K",
			assumeSi:  true,
			unit:      Kibibyte,
			expectErr: false,
		},
		{
			input:     "asdf",
			assumeSi:  false,
			unit:      Invalid,
			expectErr: true,
		},
	}

	for _, test := range tests {
		unit, err := UnitFromString(test.input, test.assumeSi)
		if err != nil {
			if !test.expectErr {
				t.Errorf("%s: unexpected error %s", test.input, err)
			}
		} else if test.expectErr {
			t.Errorf("%s: expected error but didn't get an error", test.input)
		}

		if unit != test.unit {
			t.Errorf("%s: expected unit %s (%d) but got different unit %s (%d)", test.input, test.unit, test.unit, unit, unit)
		}
	}
}

func TestToSi(t *testing.T) {
	tests := []struct {
		input    Unit
		expected Unit
	}{
		{
			input:    Kibibyte,
			expected: Kibibyte,
		},
		{
			input:    Byte,
			expected: Byte,
		},
		{
			input:    Kilobyte,
			expected: Kibibyte,
		},
	}

	for _, test := range tests {
		if test.expected != test.input.ToSi() {
			t.Errorf("Expected %s to be %s in SI", test.input, test.expected)
		}
	}
}

func TestToStd(t *testing.T) {
	tests := []struct {
		input    Unit
		expected Unit
	}{
		{
			input:    Kilobyte,
			expected: Kilobyte,
		},
		{
			input:    Byte,
			expected: Byte,
		},
		{
			input:    Kibibyte,
			expected: Kilobyte,
		},
	}

	for _, test := range tests {
		if test.expected != test.input.ToStd() {
			t.Errorf("Expected %s to be %s in SI", test.input, test.expected)
		}
	}
}
