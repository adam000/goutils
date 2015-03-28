package date

import (
	"bytes"
	"errors"
	"strconv"
	"time"
)

type DiffOptions struct {
	ShowLeading     bool
	AlwaysShowYear  bool
	AlwaysShowMonth bool
	AlwaysShowDay   bool
	AlwaysShowHour  bool
}

func Diff(start, end time.Time, options *DiffOptions) (string, error) {
	var dateBuffer bytes.Buffer

	if options == nil {
		options = &DiffOptions{}
	}
	showLeading := options.ShowLeading
	alwaysShowYear := options.AlwaysShowYear
	alwaysShowMonth := options.AlwaysShowMonth
	alwaysShowDay := options.AlwaysShowDay
	alwaysShowHour := options.AlwaysShowHour

	if start.After(end) {
		return "", errors.New("start must be before end")
	}

	// Calculate the diff
	numYears := end.Year() - start.Year()
	numMonths := end.Month() - start.Month()
	numDays := end.Day() - start.Day()
	numHours := end.Hour() - start.Hour()
	numMinutes := end.Minute() - start.Minute()

	// Correct for negative values
	if numMinutes < 0 {
		numHours--
		numMinutes += 60
	}

	if numHours < 0 {
		numDays--
		numHours += 24
	}

	if numDays < 0 {
		numMonths--
		priorMonth := end.Month() - 1
		priorYear := end.Year()
		if priorMonth <= 0 {
			priorMonth += 12
			priorYear--
		}
		numDays += numDaysInMonth(priorMonth, priorYear)
	}

	if numMonths < 0 {
		numYears--
		numMonths += 12
	}

	// Stringify

	// Has a higher place value been shown? Then ignore zeroes (we set this to
	// showLeading because they act in the same manner)
	hasHigherPlace := showLeading
	if hasHigherPlace || numYears != 0 || alwaysShowYear {
		stringifyTimeUnit(&dateBuffer, numYears, "year", false)
		if numYears != 0 {
			hasHigherPlace = true
		}
	}
	if hasHigherPlace || numMonths != 0 || alwaysShowMonth {
		stringifyTimeUnit(&dateBuffer, int(numMonths), "month", false)
		if numMonths != 0 {
			hasHigherPlace = true
		}
	}
	if hasHigherPlace || numDays != 0 || alwaysShowDay {
		stringifyTimeUnit(&dateBuffer, numDays, "day", false)
		if numDays != 0 {
			hasHigherPlace = true
		}
	}
	if hasHigherPlace || numHours != 0 || alwaysShowHour {
		stringifyTimeUnit(&dateBuffer, numHours, "hour", false)
	}
	stringifyTimeUnit(&dateBuffer, numMinutes, "minute", true)

	return dateBuffer.String(), nil
}

func numDaysInMonth(month time.Month, year int) int {
	if month == time.February && year%4 == 0 {
		return 29
	}

	return map[time.Month]int{
		time.January:   31,
		time.February:  28,
		time.March:     31,
		time.April:     30,
		time.May:       31,
		time.June:      30,
		time.July:      31,
		time.August:    31,
		time.September: 30,
		time.October:   31,
		time.November:  30,
		time.December:  31,
	}[month]
}

func stringifyTimeUnit(dateBuffer *bytes.Buffer, count int, name string, isLast bool) {
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
