package calc

import (
	"fmt"
	"testing"
)

// TestAdd
// Sub tests for tests different scenarios
func TestAdd(t *testing.T) {
	c := &Calc{}
	t.Run("1 + 2", func(t *testing.T) {
		if c.Add(1, 2) != 3 {
			t.Fatal("fail!")
		}
	})

	t.Run("3 + 2", func(t *testing.T) {
		if c.Add(3, 2) != 5 {
			t.Fatal("fail!")
		}
	})
}

// TestSubs
// Create table driven tests
func TestSubs(t *testing.T) {
	c := &Calc{}

	cases := []struct{ A, B, Expected float64 }{
		{1, 2, -1},
		{5, 3, 2},
		{9, 9, 0},
	}
	for _, tc := range cases {
		t.Run(fmt.Sprintf("Subs %v - %v", tc.A, tc.B), func(t *testing.T) {
			actual := c.Subs(tc.A, tc.B)
			if actual != tc.Expected {
				t.Fatal("fail!")
			}
		})
	}
}
