package main

import "fmt"

const n int = 2

var patterns [][n]int

func main() {
	var pattern [n]int
	var isAdded [n]bool
	makePatterns(0, pattern, isAdded)

	fmt.Println(patterns)
}

// makePatterns 配列のパターンを計算
func makePatterns(i int, pattern [n]int, isAdded [n]bool) {
	if i == n {
		patterns = append(patterns, pattern)
		return
	}
	for addNumber := 1; addNumber <= n; addNumber++ {
		if isAdded[addNumber-1] {
			continue
		}
		isAdded[addNumber-1] = true
		pattern[i] = addNumber
		makePatterns(i+1, pattern, isAdded)
		isAdded[addNumber-1] = false
	}
}
