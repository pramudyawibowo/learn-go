package main

import (
	"fmt"
	"math/rand"
	"time"
)

// O(n^2)
func search1(numbers []int, target int) []int {
	// find what number that sum up to target
	// return the numbers in an array
	// if not found, return an empty array
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{numbers[i], numbers[j]}
			}
		}
	}
	return []int{}
}

// O(n)
func search2(numbers []int, target int) []int {
	numbersMap := make(map[int]int)
	for i := 0; i < len(numbers); i++ {
		if _, ok := numbersMap[target-numbers[i]]; ok {
			return []int{numbers[i], target - numbers[i]}
		}
		numbersMap[numbers[i]] = i
	}
	return []int{}
}

func main() {
	// make a slice of random numbers
	numbers := make([]int, 100000)
	for i := 0; i < 100000; i++ {
		numbers[i] = rand.Intn(100000)
	}

	target := 500 // find two numbers that sum up to 100

	// check performance
	start := time.Now()

	// result := search1(numbers, target)
	result := search2(numbers, target)

	// elapsed in milliseconds
	elapsed := time.Since(start).Milliseconds()

	fmt.Println(numbers)
	fmt.Println(result)
	fmt.Printf("Elapsed time: %d ms\n", elapsed)
}
