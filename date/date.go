package date

import (
	"bytes"
	"errors"
	"strconv"
	"time"
)

func Diff(start, end time.Time) (string, error) {
	var dateBuffer bytes.Buffer

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

	if numMonths <= 0 {
		numYears--
		numMonths += 12
	}

	// Stringify
	if numYears != 0 {
		stringifyTimeUnit(&dateBuffer, numYears, "year", false)
	}
	stringifyTimeUnit(&dateBuffer, int(numMonths), "month", false)
	stringifyTimeUnit(&dateBuffer, numDays, "day", false)
	stringifyTimeUnit(&dateBuffer, numHours, "hour", false)
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
