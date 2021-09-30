package datasize

import (
	"fmt"
	"math/big"
	"regexp"
)

type DataSize struct {
	magnitude *big.Float
	unit      Unit
}

func New(magnitude *big.Float, unit Unit) DataSize {
	return DataSize{magnitude, unit}
}

func (d DataSize) ToHumanReadable() DataSize {
	newValue := DataSize{d.magnitude, d.unit}

	base := big.NewFloat(float64(d.unit.Base()))
	for newValue.magnitude.Cmp(base) == 1 {
		newUnit, hasNextUnit := newValue.unit.Next()
		if !hasNextUnit {
			break
		}
		newValue.magnitude.Quo(newValue.magnitude, base)
		newValue.unit = newUnit
	}

	return newValue
}

// Always parses using big and `RoundingMode.ToNearestAway`
func Parse(input string) (DataSize, error) {
	r, err := regexp.Compile(`^([0-9]*\.?[0-9]+) ?([bBkKmMgGtTpPeE](iB|ib|B|b)?)$`)
	if err != nil {
		return DataSize{}, err
	}
	matches := r.FindStringSubmatch(input)
	if !r.MatchString(input) || len(matches) < 3 {
		return DataSize{}, fmt.Errorf("'%s' is not a valid data size", input)
	}

	magnitudeStr := matches[1]
	unitStr := matches[2]

	magnitude, _, err := big.ParseFloat(magnitudeStr, 10, big.MaxPrec, big.ToNearestAway)
	if err != nil {
		return DataSize{}, fmt.Errorf("Parsing big float: %s", err)
	}

	unit, err := UnitFromString(unitStr, true)
	if err != nil {
		return DataSize{}, fmt.Errorf("Parsing unit: %s", err)
	}

	return DataSize{
		magnitude: magnitude,
		unit:      unit,
	}, nil
}

func (d DataSize) String() string {
	return fmt.Sprintf("%.2f%s", d.magnitude, d.unit.ShortString())
}
