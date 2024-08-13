package main

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend      float32
		divider       float32
		expected      float32
		expectedError bool
	}{
		{100, 10, 10, false},
		{9, 3, 3, false},
		{0, 0, 0, true},
	}

	for _, tt := range tests {
		result, err := Divide(tt.dividend, tt.divider)
		if tt.expectedError && err == nil {
			t.Errorf("expected an error but got none for dividend %f and divider %f", tt.dividend, tt.divider)
		}
		if !tt.expectedError && err != nil {
			t.Errorf("expected no error but got one for dividend %f and divider %f", tt.dividend, tt.divider)
		}
		if result != tt.expected {
			t.Errorf("expected result %f but got result is %f for dividend %f and divider %f", tt.expected, result, tt.dividend, tt.divider)
		}
	}
}
