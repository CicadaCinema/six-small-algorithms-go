package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func linear_search(arr []int, target int) (index, num_tests int) {
	// If the value is in the list.
	for i, v := range arr {
		if v == target {
			return i, i + 1
		}
	}
	// Otherwise.
	return -1, len(arr)
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

	// Make and display the unsorted array.
	arr := make_random_array(num_items, max)
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

		index, num_tests := linear_search(arr, target)
		fmt.Printf("Index: %d\nNum tests: %d\n", index, num_tests)
	}
}
