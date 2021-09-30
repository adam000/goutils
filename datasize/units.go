package datasize

import (
	"fmt"
	"strings"
)

type Unit int

const gap Unit = 100

const (
	Invalid Unit = iota
	Byte
	Kilobyte
	Megabyte
	Gigabyte
	Terabyte
	Petabyte
	Exabyte

	Kibibyte = iota + gap - Exabyte + 1
	Mebibyte
	Gibibyte
	Tebibyte
	Pebibyte
	Exbibyte
)

func (u Unit) IsSi() bool {
	return u > gap
}

func (u Unit) Base() int {
	if u.IsSi() {
		return 1000
	}

	return 1024
}

func (u Unit) String() string {
	switch u {
	case Byte:
		return "Byte"
	case Kilobyte:
		return "Kilobyte"
	case Kibibyte:
		return "Kibibyte"
	case Megabyte:
		return "Megabyte"
	case Mebibyte:
		return "Mebibyte"
	case Gigabyte:
		return "Gigabyte"
	case Gibibyte:
		return "Gibibyte"
	case Terabyte:
		return "Terabyte"
	case Tebibyte:
		return "Tebibyte"
	case Petabyte:
		return "Petabyte"
	case Pebibyte:
		return "Pebibyte"
	case Exabyte:
		return "Exabyte"
	case Exbibyte:
		return "Exbibyte"
	default:
		return "Unknown"
	}
}

func (u Unit) ShortString() string {
	switch u {
	case Byte:
		return "B"
	case Kilobyte:
		return "KB"
	case Kibibyte:
		return "KiB"
	case Megabyte:
		return "MB"
	case Mebibyte:
		return "MiB"
	case Gigabyte:
		return "GB"
	case Gibibyte:
		return "GiB"
	case Terabyte:
		return "TB"
	case Tebibyte:
		return "TiB"
	case Petabyte:
		return "PB"
	case Pebibyte:
		return "PiB"
	case Exabyte:
		return "EB"
	case Exbibyte:
		return "EiB"
	default:
		return "??"
	}
}

func UnitFromString(input string, assumeSi bool) (Unit, error) {
	// TODO this is kind of lazy and assumes you've checked the input
	// already with a regex like `([bBkKmMgGtTpPeE](iB|ib|B|b)?)`.
	if len(input) == 0 {
		return Invalid, fmt.Errorf("Can't pass empty string to UnitFromString")
	}
	lowerInput := strings.ToLower(input)

	u := Invalid
	switch lowerInput[0] {
	case 'b':
		u = Byte
	case 'k':
		u = Kilobyte
	case 'm':
		u = Megabyte
	case 'g':
		u = Gigabyte
	case 't':
		u = Terabyte
	case 'p':
		u = Petabyte
	case 'e':
		u = Exabyte
	}

	if (assumeSi && len(input) == 1) || (len(input) > 1 && lowerInput[1:] == "ib") {
		u += gap
	}

	return u, nil
}

func (u Unit) Previous() (Unit, bool) {
	if u == Byte {
		return u, false
	}
	if u == Kibibyte {
		return Byte, true
	}
	return u - 1, true
}

func (u Unit) Next() (Unit, bool) {
	if u == Exabyte || u == Exbibyte {
		return u, false
	}
	return u + 1, true
}
