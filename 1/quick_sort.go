package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func check_sorted(arr []int) {
	// An array with 0 or 1 elements is trivially sorted.
	if len(arr) <= 1 {
		fmt.Println("The array is sorted")
		return
	}

	// Check every adjacent element.
	for i := 0; i < len(arr)-1; i++ {
		// If the array is sorted, we expect arr[i]<=arr[i+1].
		if arr[i] > arr[i+1] {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}

	fmt.Println("The array is sorted")
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Get the number of items and maximum item value.
	var num_items, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&num_items)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display the unsorted array.
	arr := make_random_array(num_items, max)
	print_array(arr, 40)
	fmt.Println()

	// Sort and display the result.
	quicksort(arr)
	print_array(arr, 40)

	// Verify that it's sorted.
	check_sorted(arr)
}
