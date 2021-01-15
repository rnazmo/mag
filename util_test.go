package main

import "testing"

func TestIsValidOUI(t *testing.T) {
	testcases := []struct {
		arg  string
		want bool
	}{
		{"", false}, // Special case
		{"a", false},
		{"001122", true},
		{"012ABC", true},
		{"012abc", true},
		{"abcabc", true},
		{"aBcAbC", true}, // make this case false? (Mixed uppercase and lowercase letters)
		{"0 2abC", false},
		{"0011223", false},
		{"00:11:22", true},
		{"01:2A:BC", true},
		{"01:2a:bc", true},
		{"00-11-22", true},
		{"00_11_22", false},
		{"00:11-22", false},
	}
	for i, tt := range testcases {
		if got := isValidOUI(tt.arg); got != tt.want {
			t.Fatalf(
				"#%d: bad return value: arg = %v, got = %v, want = %v",
				i, tt.arg, got, tt.want,
			)
		}
	}
}
