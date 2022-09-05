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
	// TODO should I call ToHumanReadable() here?
	return DataSize{magnitude, unit}
}

func (d DataSize) ToHumanReadableSi() DataSize {
	num := d.ToBytes()
	bytesInKiB := big.NewFloat(float64(Kibibyte.Base()))
	return DataSize{
		magnitude: num.magnitude.Quo(num.magnitude, bytesInKiB),
		unit:      Kibibyte,
	}.ToHumanReadable()
}

func (d DataSize) ToHumanReadable() DataSize {
	newValue := DataSize{d.magnitude, d.unit}

	base := big.NewFloat(float64(d.unit.Base()))
	for newValue.magnitude.Cmp(base) > -1 {
		newUnit, hasNextUnit := newValue.unit.Next()
		if !hasNextUnit {
			break
		}
		newValue.magnitude.Quo(newValue.magnitude, base)
		newValue.unit = newUnit
	}

	return newValue
}

// Parse always parses using big and `RoundingMode.ToNearestAway`.
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
		return DataSize{}, fmt.Errorf("parsing big float: %s", err)
	}

	unit, err := UnitFromString(unitStr, true)
	if err != nil {
		return DataSize{}, fmt.Errorf("parsing unit: %s", err)
	}

	return DataSize{
		magnitude: magnitude,
		unit:      unit,
	}, nil
}

func (d DataSize) ToBytes() DataSize {
	numBytes, _ := d.magnitude.Float64()
	newU, hasPrev := d.unit.Previous()
	for hasPrev {
		numBytes *= float64(d.unit.Base())
		newU, hasPrev = newU.Previous()
	}

	return DataSize{
		magnitude: big.NewFloat(numBytes),
		unit:      Byte,
	}
}

// Add sums two numbers together, converting to the highest unit
// with a magnitude above 1. The result will be in the unit system
// (SI or std) of d.
func (d DataSize) Add(o DataSize) DataSize {
	db := d.ToBytes()
	ob := o.ToBytes()

	ds := DataSize{
		magnitude: db.magnitude.Add(db.magnitude, ob.magnitude),
		unit:      Byte,
	}
	if d.unit.IsSi() {
		return ds.ToHumanReadableSi()
	}
	return ds.ToHumanReadable()
}

func (d DataSize) Sub(o DataSize) DataSize {
	db := d.ToBytes()
	ob := o.ToBytes()

	ds := DataSize{
		magnitude: db.magnitude.Sub(db.magnitude, ob.magnitude),
		unit:      Byte,
	}
	if d.unit.IsSi() {
		return ds.ToHumanReadableSi()
	}
	return ds.ToHumanReadable()
}

func (d DataSize) Gte(o DataSize) bool {
	db := d.ToBytes()
	ob := o.ToBytes()

	return db.magnitude.Cmp(ob.magnitude) >= 0
}

func (d DataSize) String() string {
	return fmt.Sprintf("%.2f%s", d.magnitude, d.unit.ShortString())
}
