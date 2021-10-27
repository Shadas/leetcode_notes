package _13_Roman_to_Integer

import "testing"

func TestRomanToInt(t *testing.T) {
	var (
		roman string
	)
	roman = "III"
	if num := romanToInt(roman); num != 3 {
		t.Errorf("wrong num with %d", num)
	}
}
