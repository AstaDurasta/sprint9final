package main

import (
	"testing"
)

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"all equal", []int{7, 7, 7, 7}, 7},
		{"mixed values", []int{1, 3, 5, 2, 4}, 5},
		{"negative and positive", []int{-10, 0, 5, -3}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			if got != tt.want {
				t.Errorf("maximum(%v) = %d; want %d", tt.data, got, tt.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"empty slice", []int{}, 0},
		{"single element", []int{5}, 5},
		{"all equal", []int{7, 7, 7, 7}, 7},
		{"mixed values", []int{1, 3, 5, 2, 4}, 5},
		{"negative and positive", []int{-10, 0, 5, -3}, 5},
		{"large slice", generateTestSlice(10000, 42), 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			if got != tt.want {
				t.Errorf("maxChunks(%v) = %d; want %d", tt.data, got, tt.want)
			}
		})
	}
}

func generateTestSlice(size int, maxVal int) []int {
	s := make([]int, size)
	for i := range s {
		s[i] = i % maxVal
	}

	s[size-1] = maxVal
	return s
}

func TestGenerateRandomElements(t *testing.T) {
	size := 1000
	data := generateRandomElements(size)

	if len(data) != size {
		t.Errorf("generateRandomElements(%d) returned slice of length %d, want %d", size, len(data), size)
	}

	for _, v := range data {
		if v < 1 || v > SIZE {
			t.Errorf("generateRandomElements returned value out of range: %d", v)
		}
	}
}
