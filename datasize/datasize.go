package datasize

import (
	"fmt"
	"math/big"
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

func (d DataSize) String() string {
	return fmt.Sprintf("%.2f%s", d.magnitude, d.unit.ShortString())
}
