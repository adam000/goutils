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
	var dateBuffer bytes.Buffer

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
	if hasHigherPlace || d.Years != 0 || alwaysShowYear {
		stringifyTimeUnit(&dateBuffer, d.Years, "year", false)
		if d.Years != 0 {
			hasHigherPlace = true
		}
	}
	if granularity <= Month {
		if hasHigherPlace || d.Months != 0 || alwaysShowMonth {
			stringifyTimeUnit(&dateBuffer, int(d.Months), "month", granularity == Month)
			if d.Months != 0 {
				hasHigherPlace = true
			}
		}
	}
	if granularity <= Day {
		if hasHigherPlace || d.Days != 0 || alwaysShowDay {
			stringifyTimeUnit(&dateBuffer, d.Days, "day", granularity == Day)
			if d.Days != 0 {
				hasHigherPlace = true
			}
		}
	}
	if granularity <= Hour {
		if hasHigherPlace || d.Hours != 0 || alwaysShowHour {
			stringifyTimeUnit(&dateBuffer, d.Hours, "hour", granularity == Hour)
		}
	}
	if granularity <= Minute {
		stringifyTimeUnit(&dateBuffer, d.Minutes, "minute", granularity == Minute)
	}

	return dateBuffer.String(), nil
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

func stringifyTimeUnit(dateBuffer *bytes.Buffer, count int, name string, isLast bool) {
	if isLast {
		dateBuffer.WriteString("and ")
	}
	dateBuffer.WriteString(strconv.Itoa(count))
	dateBuffer.WriteString(" ")
	dateBuffer.WriteString(name)
	if count != 1 {
		dateBuffer.WriteString("s")
	}
	if !isLast {
		dateBuffer.WriteString(", ")
	}
}
