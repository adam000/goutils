package datasize

import (
	"fmt"
	"math"
	"math/big"
	"testing"
)

func TestToHumanReadable(t *testing.T) {
	tests := []struct {
		input     DataSize
		expected  DataSize
		outputStr string
	}{
		{
			input: DataSize{
				magnitude: big.NewFloat(2420),
				unit:      Kibibyte,
			},
			expected: DataSize{
				magnitude: big.NewFloat(2.42),
				unit:      Mebibyte,
			},
			outputStr: "2.42MiB",
		},
	}

	for _, test := range tests {
		actual := test.input.ToHumanReadable()
		if test.expected.unit != actual.unit {
			t.Errorf("Units differ: expected result %s, got %s", test.expected, actual)
		}
		if test.expected.magnitude.Cmp(actual.magnitude) != 0 {
			t.Errorf("Magnitudes differ: expected result %s, got %s", test.expected, actual)
		}
		if test.expected.String() != test.outputStr {
			t.Errorf("String result differs: expected %s, got %s", test.outputStr, test.expected.String())
		}
	}
}

func TestToBytesSi(t *testing.T) {
	tests := []Unit{
		Kibibyte,
		Mebibyte,
		Gibibyte,
		Tebibyte,
		Pebibyte,
		Exbibyte,
	}

	mag := big.NewFloat(1)
	for i, test := range tests {
		result := DataSize{magnitude: mag, unit: test}.ToBytes().magnitude
		expected := math.Pow(1000, float64(i+1))
		resultStr := fmt.Sprintf("%.f", result)
		expectedStr := fmt.Sprintf("%.f", expected)
		if expectedStr != resultStr {
			t.Errorf("Expected %s to be %.0f bytes, was %s bytes", test, expected, resultStr)
		}
	}
}

func TestToBytesRandom(t *testing.T) {
	test := DataSize{
		big.NewFloat(3.14),
		Gigabyte,
	}
	bytes := 3371549327

	result := test.ToBytes()

	resultStr := fmt.Sprintf("%.f", result.magnitude)
	expectedStr := fmt.Sprintf("%d", bytes)

	if resultStr != expectedStr {
		t.Errorf("Expected %s, got %s", expectedStr, resultStr)
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		input     string
		expectErr bool
	}{
		{"1.21Q", true},
		{"garbage", true},
		{"5 MiB", false},
		{"1.21G", false},
	}

	for _, test := range tests {
		_, err := Parse(test.input)

		if test.expectErr && err == nil {
			t.Errorf("Expected error but there wasn't one")
		} else if !test.expectErr && err != nil {
			t.Errorf("Didn't expect error, but got one anyways: %s", err)
		}
	}
}

func TestToBytesStd(t *testing.T) {
	tests := []Unit{
		Kilobyte,
		Megabyte,
		Gigabyte,
		Terabyte,
		Petabyte,
		Exabyte,
	}

	mag := big.NewFloat(1)
	for i, test := range tests {
		result := DataSize{magnitude: mag, unit: test}.ToBytes().magnitude
		expected := math.Pow(1024, float64(i+1))
		resultStr := fmt.Sprintf("%.f", result)
		expectedStr := fmt.Sprintf("%.f", expected)
		if expectedStr != resultStr {
			t.Errorf("Expected %s to be %.0f bytes, was %s bytes", test, expected, resultStr)
		}
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		first    DataSize
		second   DataSize
		expected DataSize
	}{
		{
			first: DataSize{
				magnitude: big.NewFloat(42),
				unit:      Kibibyte,
			},
			second: DataSize{
				magnitude: big.NewFloat(1),
				unit:      Byte,
			},
			expected: DataSize{
				magnitude: big.NewFloat(42),
				unit:      Kibibyte,
			},
		},
		{
			first: DataSize{
				magnitude: big.NewFloat(42),
				unit:      Kibibyte,
			},
			second: DataSize{
				magnitude: big.NewFloat(1),
				unit:      Kibibyte,
			},
			expected: DataSize{
				magnitude: big.NewFloat(43),
				unit:      Kibibyte,
			},
		},
		{
			first: DataSize{
				magnitude: big.NewFloat(500),
				unit:      Kibibyte,
			},
			second: DataSize{
				magnitude: big.NewFloat(500),
				unit:      Kibibyte,
			},
			expected: DataSize{
				magnitude: big.NewFloat(1),
				unit:      Mebibyte,
			},
		},
		{
			first: DataSize{
				magnitude: big.NewFloat(768),
				unit:      Gigabyte,
			},
			second: DataSize{
				magnitude: big.NewFloat(257),
				unit:      Gigabyte,
			},
			expected: DataSize{
				magnitude: big.NewFloat(1),
				unit:      Terabyte,
			},
		},
		{
			first: DataSize{
				magnitude: big.NewFloat(512),
				unit:      Kilobyte,
			},
			second: DataSize{
				magnitude: big.NewFloat(512),
				unit:      Kilobyte,
			},
			expected: DataSize{
				magnitude: big.NewFloat(1),
				unit:      Megabyte,
			},
		},
	}

	for _, test := range tests {
		result := test.first.Add(test.second)
		//if result.magnitude != test.expected.magnitude || result.unit != test.expected.unit {
		//if result.String() != test.expected.String() {
		hrResult := result.ToHumanReadable().String()
		hrExpected := test.expected.ToHumanReadable().String()

		if hrResult != hrExpected {
			t.Errorf("Expected %s as the result of adding %s and %s, got %s", hrExpected, test.first, test.second, hrResult)
		}
	}
}
