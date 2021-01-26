package main

import (
	"fmt"
	"sync"
)

// Sort uses insertion sort. startIdx is inclusive, endIdx is not
func Sort(numbers []int, startIdx, endIdx int) {
	for i := startIdx; i < endIdx; i++ {
		valToInsert := numbers[i]
		holePos := i

		for holePos > 0 && numbers[holePos-1] > valToInsert {
			numbers[holePos] = numbers[holePos-1]
			holePos--
		}

		numbers[holePos] = valToInsert
	}
}

func SortRoutine(nums []int, n int, wg *sync.WaitGroup) {
	defer wg.Done()

	Sort(nums, 0, len(nums))

	fmt.Printf("Routine %d: %v\n", n, nums)
}

// Merge assumes that the len(finalArr) = len(arrOne) + len(arrTwo)
func Merge(arrOne, arrTwo, finalArr []int) {
	var i, j, k int

	lOne := len(arrOne)
	lTwo := len(arrTwo)

	for i < lOne && j < lTwo {
		if arrOne[i] < arrTwo[j] {
			finalArr[k] = arrOne[i]
			i++
		} else {
			finalArr[k] = arrTwo[j]
			j++
		}

		k++
	}

	// Only one of the two loops below will run because of the exit condition
	// of the loop above
	for i < lOne {
		finalArr[k] = arrOne[i]
		i++
		k++
	}

	for j < lTwo {
		finalArr[k] = arrTwo[j]
		j++
		k++
	}
}

func main() {
	MAX_NUMS := 5000

	// User Input Part
	fmt.Printf("Only going to input %d total numbers...\n\n", MAX_NUMS)
	numbers := getUserInput(MAX_NUMS)

	sortedTmp := make([]int, len(numbers))[0:] // hold numbers for final merge
	sortedNums := make([]int, len(numbers))[0:]

	fmt.Println()

	fmt.Printf("Unsorted: %v\n\n", numbers)

	l := len(numbers)
	diff := (l + 2) / 4
	idx := 0
	var wg sync.WaitGroup

	wg.Add(4)
	for i := 0; i < 3; i++ {
		go SortRoutine(numbers[idx:idx+diff], i, &wg)
		idx += diff
	}
	go SortRoutine(numbers[idx:l], 3, &wg)

	wg.Wait()

	// Merge 1st and 2nd quarter
	Merge(numbers[:diff], numbers[diff:2*diff], sortedTmp[:2*diff])
	// Merge 3rd and 4th quarter
	Merge(numbers[2*diff:3*diff], numbers[3*diff:], sortedTmp[2*diff:])

	// Merge into final array
	Merge(sortedTmp[:2*diff], sortedTmp[2*diff:], sortedNums)

	fmt.Printf("\nSorted: %v\n", sortedNums)
}

func getUserInput(maxNumbers int) []int {
	numbers := make([]int, maxNumbers)[0:]

	var tmp int
	for i := 0; i < maxNumbers; i++ {
		fmt.Print("Enter number (or 'x' and sort) -> ")
		_, err := fmt.Scanf("%d", &tmp)
		if err != nil {
			return numbers[:i]
		}

		numbers[i] = tmp
	}

	return numbers
}
