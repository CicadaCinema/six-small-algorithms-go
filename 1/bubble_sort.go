package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func bubble_sort(arr []int) {
	// We require at most len(arr)-1 passes.
	for i := 0; i < len(arr)-1; i++ {
		// The last i elements are in their final positions by this point.
		swapped := false
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				index_j := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = index_j
				swapped = true
			}
		}
		if !swapped {
			// The elements are already sorted.
			break
		}
	}
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
	bubble_sort(arr)
	print_array(arr, 40)

	// Verify that it's sorted.
	check_sorted(arr)
}
