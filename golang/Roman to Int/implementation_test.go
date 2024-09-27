package roman

import (
	"fmt"
	"testing"
)

func TestCases(t *testing.T) {
	var cases = []struct {
		input string
		want  int
	}{
		{"III", 3},
		{"IV", 4}, // Added some variety to the test cases
		{"IX", 9},
		{"XII", 12},
		{"XL", 40},
		{"CD", 400},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for i, tt := range cases {
		t.Run(fmt.Sprintf("Test %d", i), func(t *testing.T) {
			num, err := convertRomanNumeralToInt(tt.input)
			if num != tt.want || err != nil {
				t.Fatalf(`convertRomanNumeralToInt(%v) = %d should be %d Error:%v`, tt.input, num, tt.want, err)
			}
		})
	}
}
