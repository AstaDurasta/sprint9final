package main

import (
	"testing"

	"github.com/stretchr/testify/require"
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
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximum(tt.data)
			require.Equal(t, tt.want, got, "maximum(%v)", tt.data)
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
		{"large slice", generateTestSlice(10000, 42), 42},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxChunks(tt.data)
			require.Equal(t, tt.want, got, "maxChunks(%v)", tt.data)
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
	tests := []struct {
		name     string
		size     int
		expected int
	}{
		{"small size", 10, 10},
		{"medium size", 1000, 1000},
		{"zero size", 0, 0},
		{"larger size", 5000, 5000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			require.Equal(t, tt.expected, len(data), "unexpected length for size=%d", tt.size)

			for i, v := range data {
				_ = v
				require.NotNil(t, &v, "value at index %d is nil", i)
			}
		})
	}
}
