package date

import (
	"testing"
	"time"
)

func TestGetDiff(t *testing.T) {
	location, _ := time.LoadLocation("America/Los_Angeles")

	tests := []struct {
		start  time.Time
		end    time.Time
		output Diff
	}{
		// alwaysShowYear
		{
			time.Date(2014, time.January, 1, 12, 0, 0, 0, location),
			time.Date(2014, time.January, 1, 13, 0, 0, 0, location),
			Diff{
				Hours: 1,
			},
		},
		// Leap day check - not a leap year
		{
			time.Date(2015, time.February, 26, 0, 0, 0, 0, location),
			time.Date(2015, time.March, 3, 0, 0, 0, 0, location),
			Diff{
				Days: 5,
			},
		},
		// Leap day check - a leap year
		{
			time.Date(2016, time.February, 26, 0, 0, 0, 0, location),
			time.Date(2016, time.March, 3, 0, 0, 0, 0, location),
			Diff{
				Days: 6,
			},
		},
	}

	for _, test := range tests {
		result, err := GetDiff(test.start, test.end)
		if err != nil {
			t.Errorf("Unexpected error '%s'", err.Error())
		}
		if result != test.output {
			t.Errorf("Expected '%v', got '%v'", test.output, result)
		}
	}
}

func TestDiffDisplay(t *testing.T) {
	tests := []struct {
		diff    Diff
		options *DiffDisplayOptions
		output  string
	}{
		{
			Diff{
				Days: 5,
			},
			nil,
			"5 days, 0 hours, and 0 minutes",
		},
		// alwaysShowYear
		{
			Diff{
				Hours: 1,
			},
			&DiffDisplayOptions{
				AlwaysShowYear: true,
			},
			"0 years, 1 hour, and 0 minutes",
		},
		// Granularity
		{
			Diff{
				Months:  2,
				Days:    10,
				Hours:   12,
				Minutes: 12,
			},
			&DiffDisplayOptions{
				Granularity: Hour,
			},
			"2 months, 10 days, and 12 hours",
		},
		// 2 places
		{
			Diff{
				Hours:   12,
				Minutes: 12,
			},
			&DiffDisplayOptions{},
			"12 hours and 12 minutes",
		},
		// 1 place
		{
			Diff{
				Minutes: 12,
			},
			&DiffDisplayOptions{},
			"12 minutes",
		},
	}

	for _, test := range tests {
		result, err := test.diff.Display(test.options)
		if err != nil {
			t.Errorf("Unexpected error '%s'", err.Error())
		}
		if result != test.output {
			t.Errorf("Expected '%s', got '%s'", test.output, result)
		}
	}
}
