package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"
)

func binary_search(arr []int, target int) (index, num_tests int) {
	examined_items := 0

	l := 0
	r := len(arr) - 1

	for l <= r {
		// We are examining arr[m] here.
		examined_items += 1

		m := int(math.Floor(float64(l+r) / 2))
		if arr[m] < target {
			l = m + 1
		} else if arr[m] > target {
			r = m - 1
		} else {
			return m, examined_items
		}
	}

	// Unsuccessful.
	return -1, examined_items
}

func partition(A []int) int {
	lo := 0
	hi := len(A) - 1

	// Choose the last element as the pivot.
	pivot := A[hi]

	// Temporary pivot index.
	i := lo - 1

	for j := lo; j < hi; j++ {

		// If the current element is less than or equal to the pivot.
		if A[j] <= pivot {

			// Move the temporary pivot index forward.
			i = i + 1
			// Swap the current element with the element at the temporary pivot index.
			A[i], A[j] = A[j], A[i]
		}
	}

	// Move the pivot element to the correct pivot position (between the smaller and larger elements).
	i = i + 1
	A[i], A[hi] = A[hi], A[i]

	// The pivot index.
	return i
}

func quicksort(A []int) {
	//Slice is so small that it doesn’t need sorting.
	if len(A) <= 1 {
		return
	}

	// Partition array and get the pivot index.
	p := partition(A)

	// Sort the two partitions.
	quicksort(A[0:p])
	quicksort(A[p+1 : len(A)])
}

func make_random_array(num_items, max int) []int {
	array := make([]int, num_items)

	for i, _ := range array {
		array[i] = rand.Intn(max)
	}

	return array
}

func print_array(arr []int, num_items int) {
	if num_items > len(arr) {
		num_items = len(arr)
	}

	for i := 0; i < num_items; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Get the number of items and maximum item value.
	var num_items, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make, sort and display the array.
	arr := make_random_array(num_items, max)
	quicksort(arr)
	print_array(arr, 40)
	fmt.Println()

	for {
		var user_input string
		fmt.Printf("Target: ")
		fmt.Scanln(&user_input)

		if user_input == "" {
			break
		}

		target, err := strconv.Atoi(user_input)
		if err != nil {
			log.Fatal("Error parsing int from string.")
		}

		index, num_tests := binary_search(arr, target)
		fmt.Printf("Index: %d\nNum tests: %d\n", index, num_tests)
	}
}
