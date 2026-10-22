package mathutil_test

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name        string
		a, b        float64
		expected    float64
		expectError bool
	}{
		{name: "regular division", a: 10, b: 2, expected: 5, expectError: false},
		{name: "division by zero", a: 5, b: 0, expected: 0, expectError: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if (err != nil) != tt.expectError {
				t.Fatalf("unexpected error status: %v", err)
			}
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
