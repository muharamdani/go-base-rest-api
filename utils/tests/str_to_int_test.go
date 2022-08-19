//go:build unit || utils || all

package tests

import (
	"github.com/muharamdani/go-base-rest-api/utils"
	"testing"
)

// TestStrToInt tests StrToInt function, need to change later using testify
func TestStrToInt(t *testing.T) {
	testCases := []struct {
		input string
		want  int
	}{
		{"1", 1},
		{"2", 2},
		{"3", 3},
		{"a", 0},
	}

	for _, tc := range testCases {
		got := utils.StrToInt(tc.input)
		if got != tc.want {
			t.Errorf("StrToInt(%q) = %d, want %d", tc.input, got, tc.want)
		}
	}
}
