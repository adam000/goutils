package date

import (
	"testing"
	"time"
)

func TestDiff(t *testing.T) {
	location, _ := time.LoadLocation("America/Los_Angeles")

	tests := []struct {
		start   time.Time
		end     time.Time
		output  string
		options *DiffOptions
	}{
		// alwaysShowYear
		{
			time.Date(2014, time.January, 1, 12, 0, 0, 0, location),
			time.Date(2014, time.January, 1, 13, 0, 0, 0, location),
			"0 years, 1 hour, 0 minutes",
			&DiffOptions{
				AlwaysShowYear: true,
			},
		},
		// Leap day check - not a leap year
		{
			time.Date(2015, time.February, 26, 0, 0, 0, 0, location),
			time.Date(2015, time.March, 3, 0, 0, 0, 0, location),
			"5 days, 0 hours, 0 minutes",
			nil,
		},
		// Leap day check - a leap year
		{
			time.Date(2016, time.February, 26, 0, 0, 0, 0, location),
			time.Date(2016, time.March, 3, 0, 0, 0, 0, location),
			"6 days, 0 hours, 0 minutes",
			nil,
		},
	}

	for _, test := range tests {
		result, err := Diff(test.start, test.end, test.options)
		if err != nil {
			t.Errorf("Unexpected error '%s'", err.Error())
		}
		if result != test.output {
			t.Errorf("Expected '%s', got '%s'", test.output, result)
		}
	}
}
