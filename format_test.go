package nearapi

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseAmountString(t *testing.T) {
	tt := []struct {
		arg      string
		expected string
	}{
		{
			arg:      "0",
			expected: "0",
		},
		{
			arg:      "990000000000000000000000",
			expected: "990000000000000000000000",
		},
		{
			arg:      "1000000000000000000000000",
			expected: "1000000000000000000000000",
		},
		{
			arg:      "11000000000000000000000000",
			expected: "11000000000000000000000000",
		},
		{
			arg:      "99900000000000000000000000",
			expected: "99900000000000000000000000",
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(fmt.Sprintf("parse %s", tc.arg), func(t *testing.T) {
			assert.Equal(t, tc.expected, ParseAmount(tc.arg).String())
		})
	}
}

func TestParseAmountToNEARs(t *testing.T) {
	tt := []struct {
		arg      string
		expected string
	}{
		{
			arg:      "0",
			expected: "0",
		},
		{
			arg:      "1",
			expected: "0.000000000000000000000001",
		},
		{
			arg:      "990000000000000000000000",
			expected: "0.99",
		},
		{
			arg:      "1000000000000000000000000",
			expected: "1",
		},
		{
			arg:      "11000000000000000000000000",
			expected: "11",
		},
		{
			arg:      "99900000000000000000000000",
			expected: "99.9",
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(fmt.Sprintf("parse %s", tc.arg), func(t *testing.T) {
			assert.Equal(t, tc.expected, ParseAmount(tc.arg).ToNEARs())
		})
	}
}

func TestParseAmountToNEARsF64(t *testing.T) {
	tt := []struct {
		arg      string
		expected float64
	}{
		{
			arg:      "0",
			expected: 0,
		},
		{
			arg:      "1000000",
			expected: 0.000000000000000001,
		},
		{
			arg:      "990000000000000000000000",
			expected: 0.99,
		},
		{
			arg:      "1000000000000000000000000",
			expected: 1,
		},
		{
			arg:      "11000000000000000000000000",
			expected: 11,
		},
		{
			arg:      "99900000000000000000000000",
			expected: 99.9,
		},
		{
			arg:      "999900000000000000000000000",
			expected: 999.9,
		},
	}

	for _, tc := range tt {
		tc := tc
		t.Run(fmt.Sprintf("parse %s", tc.arg), func(t *testing.T) {
			f64, _ := ParseAmount(tc.arg).ToNEARsF64()
			assert.Equal(t, tc.expected, f64)
		})
	}
}
