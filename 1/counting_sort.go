package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	id            string
	num_purchases int
}

func make_random_array(num_items, max int) []Customer {
	array := make([]Customer, num_items)

	for i, _ := range array {
		id := fmt.Sprintf("C%d", i)
		num_purchases := rand.Intn(max)
		array[i] = Customer{id: id, num_purchases: num_purchases}
	}

	return array
}

func print_array(arr []Customer, num_items int) {
	if num_items > len(arr) {
		num_items = len(arr)
	}

	for i := 0; i < num_items; i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println()
}

func check_sorted(arr []Customer) {
	// An array with 0 or 1 elements is trivially sorted.
	if len(arr) <= 1 {
		fmt.Println("The array is sorted")
		return
	}

	// Check every adjacent element.
	for i := 0; i < len(arr)-1; i++ {
		// If the array is sorted, we expect arr[i]<=arr[i+1].
		if arr[i].num_purchases > arr[i+1].num_purchases {
			fmt.Println("The array is NOT sorted!")
			return
		}
	}

	fmt.Println("The array is sorted")
	return
}

func counting_sort(arr []Customer, max int) []Customer {
	counts := make([]int, max)

	for _, v := range arr {
		// v.num_purchases is an integer in the range [0, max-1].
		counts[v.num_purchases] += 1
	}

	for i := 1; i < max; i++ {
		counts[i] += counts[i-1]
	}

	sorted := make([]Customer, len(arr))

	for _, v := range arr {
		k := v.num_purchases
		// Place this item in counts[k] - 1.
		sorted[counts[k]-1] = v
		// Decrement counts[k].
		counts[k] -= 1
	}

	return sorted
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
	sorted := counting_sort(arr, max)
	print_array(sorted, 40)

	// Verify that it's sorted.
	check_sorted(sorted)
}
