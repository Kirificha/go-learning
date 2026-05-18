package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name     string
		a, b     float64
		expected float64
		wantErr  bool
	}{
		{"error", 2, 0, 0, true},
		{"обычное деление", 10, 2, 5, false},
		{"деление на 1", 7, 1, 7, false},
		{"отрицательные", -6, 2, -3, false},
		{"дробный результат", 1, 4, 0.25, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Divide(tt.a, tt.b)
			if tt.wantErr && err == nil {
				t.Error("Ожидали ошибку, но её нет")
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
