package datasize

import "testing"

func TestUnitFromString(t *testing.T) {
	tests := []struct {
		input    string
		assumeSi bool
		unit     Unit
		err      error
	}{
		{
			input:    "KiB",
			assumeSi: false,
			unit:     Kibibyte,
			err:      nil,
		},
		{
			input:    "KB",
			assumeSi: false,
			unit:     Kilobyte,
			err:      nil,
		},
		{
			input:    "KB",
			assumeSi: true,
			unit:     Kilobyte,
			err:      nil,
		},
		{
			input:    "K",
			assumeSi: false,
			unit:     Kilobyte,
			err:      nil,
		},
		{
			input:    "K",
			assumeSi: true,
			unit:     Kibibyte,
			err:      nil,
		},
		{
			input:    "asdf",
			assumeSi: false,
			unit:     Invalid,
			err:      nil,
		},
	}

	for _, test := range tests {
		unit, err := UnitFromString(test.input, test.assumeSi)
		if err != nil {
			if err != test.err {
				t.Errorf("%s: expected error %s but got different error %s", test.input, test.err, err)
			}
		} else if test.err != nil {
			t.Errorf("%s: expected error %s but didn't get an error", test.input, test.err)
		}

		if unit != test.unit {
			t.Errorf("%s: expected unit %s (%d) but got different unit %s (%d)", test.input, test.unit, test.unit, unit, unit)
		}
	}
}
