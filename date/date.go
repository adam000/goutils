package date

import (
	"bytes"
	"errors"
	"strconv"
	"time"
)

type TimeUnit int

const (
	Minute TimeUnit = iota
	Hour
	Day
	Month
	Year
)

type DiffDisplayOptions struct {
	ShowLeading     bool
	AlwaysShowYear  bool
	AlwaysShowMonth bool
	AlwaysShowDay   bool
	AlwaysShowHour  bool
	Granularity     TimeUnit
	// The string buffer used when building the string to display
	stringBuffer bytes.Buffer
}

type Diff struct {
	Years, Months, Days, Hours, Minutes int
}

func GetDiffString(start, end time.Time, options *DiffDisplayOptions) (string, error) {
	diff, err := GetDiff(start, end)
	if err != nil {
		return "", err
	}

	str, err := diff.Display(options)
	return str, err
}

func GetDiff(start, end time.Time) (d Diff, err error) {
	if start.After(end) {
		return Diff{}, errors.New("start must be before end")
	}

	// Calculate the diff
	d.Years = end.Year() - start.Year()
	d.Months = int(end.Month() - start.Month())
	d.Days = end.Day() - start.Day()
	d.Hours = end.Hour() - start.Hour()
	d.Minutes = end.Minute() - start.Minute()

	// Correct for negative values
	if d.Minutes < 0 {
		d.Hours--
		d.Minutes += 60
	}

	if d.Hours < 0 {
		d.Days--
		d.Hours += 24
	}

	if d.Days < 0 {
		d.Months--
		priorMonth := end.Month() - 1
		priorYear := end.Year()
		if priorMonth <= 0 {
			priorMonth += 12
			priorYear--
		}
		d.Days += numDaysInMonth(priorMonth, priorYear)
	}

	if d.Months < 0 {
		d.Years--
		d.Months += 12
	}

	return
}

func NumDaysDiff(start, end time.Time) (numDays int, err error) {
	if start.After(end) {
		return 0, errors.New("start must be before end")
	}

	year := start.Year()

	numDays -= start.YearDay()

	for year < end.Year() {
		numDays += numDaysInYear(year)
		year++
	}

	numDays += end.YearDay() + 1

	return
}

func (d Diff) Display(options *DiffDisplayOptions) (string, error) {
	if options == nil {
		options = &DiffDisplayOptions{}
	}
	showLeading := options.ShowLeading
	alwaysShowYear := options.AlwaysShowYear
	alwaysShowMonth := options.AlwaysShowMonth
	alwaysShowDay := options.AlwaysShowDay
	alwaysShowHour := options.AlwaysShowHour
	granularity := options.Granularity

	// Has a higher place value been shown? Then ignore zeroes (we set this to
	// showLeading because they act in the same manner)
	hasHigherPlace := showLeading
	numPlaces := d.getNumPlaces()
	if hasHigherPlace || d.Years != 0 || alwaysShowYear {
		options.stringifyTimeUnit(d.Years, "year", hasHigherPlace, false, numPlaces == 2)
		if d.Years != 0 {
			hasHigherPlace = true
		}
	}
	if granularity <= Month {
		if hasHigherPlace || d.Months != 0 || alwaysShowMonth {
			options.stringifyTimeUnit(d.Months, "month", hasHigherPlace, granularity == Month, numPlaces == 2)
			if d.Months != 0 {
				hasHigherPlace = true
			}
		}
	}
	if granularity <= Day {
		if hasHigherPlace || d.Days != 0 || alwaysShowDay {
			options.stringifyTimeUnit(d.Days, "day", hasHigherPlace, granularity == Day, numPlaces == 2)
			if d.Days != 0 {
				hasHigherPlace = true
			}
		}
	}
	if granularity <= Hour {
		if hasHigherPlace || d.Hours != 0 || alwaysShowHour {
			options.stringifyTimeUnit(d.Hours, "hour", hasHigherPlace, granularity == Hour, numPlaces == 2)
			if d.Hours != 0 {
				hasHigherPlace = true
			}
		}
	}
	if granularity <= Minute {
		options.stringifyTimeUnit(d.Minutes, "minute", hasHigherPlace, granularity == Minute, numPlaces == 2)
	}

	return options.stringBuffer.String(), nil
}

func (d Diff) getNumPlaces() (numPlaces int) {
	if d.Years != 0 {
		numPlaces++
	}
	if d.Months != 0 {
		numPlaces++
	}
	if d.Days != 0 {
		numPlaces++
	}
	if d.Hours != 0 {
		numPlaces++
	}
	if d.Minutes != 0 {
		numPlaces++
	}
	return
}

func numDaysInYear(year int) int {
	if year%4 == 0 {
		return 366
	} else {
		return 365
	}
}

func numDaysInMonth(month time.Month, year int) int {
	if month == time.February && year%4 == 0 {
		return 29
	}

	switch month {
	case time.February:
		return 28
	case time.January, time.March, time.May, time.July, time.August, time.October, time.December:
		return 31
	case time.April, time.June, time.September, time.November:
		return 30
	}

	// Can't reach this...
	return 0
}

func (d *DiffDisplayOptions) stringifyTimeUnit(count int, name string, hasHigherPlace, isLast, hideComma bool) {
	if isLast && hasHigherPlace {
		if hideComma {
			d.stringBuffer.WriteString(" ")
		}

		d.stringBuffer.WriteString("and ")
	}
	d.stringBuffer.WriteString(strconv.Itoa(count))
	d.stringBuffer.WriteString(" ")
	d.stringBuffer.WriteString(name)
	if count != 1 {
		d.stringBuffer.WriteString("s")
	}
	if !isLast && !hideComma {
		d.stringBuffer.WriteString(", ")
	}
}
