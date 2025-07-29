package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {

	if size <= 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	sliceNumbers := make([]int, size)
	for i := 0; i < size; i++ {
		sliceNumbers[i] = rand.Int()
	}
	return sliceNumbers
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for _, v := range data {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) == 1 {
		return data[0]
	}
	chunkSize := len(data) / CHUNKS
	maxes := make([]int, CHUNKS)
	var wg sync.WaitGroup

	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		go func(data []int) {
			defer wg.Done()
			maxes[i] = maximum(data)

		}(data[start:end])
	}

	wg.Wait()

	finalMax := maximum(maxes)

	return finalMax
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	data := generateRandomElements(SIZE)
	fmt.Println("Ищем максимальное значение в один поток")
	// ваш код здесь
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	// ваш код здесь

	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
