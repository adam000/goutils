package date

import (
	"bytes"
	"errors"
	"strconv"
	"time"
)

type DiffOptions struct {
	showLeading     bool
	alwaysShowYear  bool
	alwaysShowMonth bool
	alwaysShowDay   bool
	alwaysShowHour  bool
}

func Diff(start, end time.Time, options *DiffOptions) (string, error) {
	var dateBuffer bytes.Buffer

	if options == nil {
		options = &DiffOptions{}
	}
	showLeading := options.showLeading
	alwaysShowYear := options.alwaysShowYear
	alwaysShowMonth := options.alwaysShowMonth
	alwaysShowDay := options.alwaysShowDay
	alwaysShowHour := options.alwaysShowHour

	// TODO make this a function because leap years
	numDaysInMonth := map[time.Month]int{
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
	}

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
		priorMonth := numMonths
		priorYear := end.Year()
		if priorMonth <= 0 {
			priorMonth += 12
			priorYear--
		}
		// TODO add in priorYear when this becomes a function instead of a map
		numDays += numDaysInMonth[priorMonth]
	}

	if numMonths < 0 {
		numYears--
		numMonths += 12
	}

	// Stringify

	// Has a higher place value been shown? Then ignore zeroes
	hasHigherPlace := false
	if showLeading || numYears != 0 || alwaysShowYear {
		stringifyTimeUnit(&dateBuffer, numYears, "year", false)
		if numYears != 0 {
			hasHigherPlace = true
		}
	}
	if showLeading || hasHigherPlace || numMonths != 0 || alwaysShowMonth {
		stringifyTimeUnit(&dateBuffer, int(numMonths), "month", false)
		if numMonths != 0 {
			hasHigherPlace = true
		}
	}
	if showLeading || hasHigherPlace || numDays != 0 || alwaysShowDay {
		stringifyTimeUnit(&dateBuffer, numDays, "day", false)
		if numDays != 0 {
			hasHigherPlace = true
		}
	}
	if showLeading || hasHigherPlace || numHours != 0 || alwaysShowHour {
		stringifyTimeUnit(&dateBuffer, numHours, "hour", false)
	}
	stringifyTimeUnit(&dateBuffer, numMinutes, "minute", true)

	return dateBuffer.String(), nil
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
