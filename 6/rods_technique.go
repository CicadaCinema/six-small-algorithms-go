package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const num_items = 45

const min_value = 1
const max_value = 10
const min_weight = 4
const max_weight = 10

var allowed_weight int

type Item struct {
	id, blocked_by int
	i_block        []int // Other items that this one blocks.
	value, weight  int
	is_selected    bool
}

// Make some random items.
func make_items(num_items, min_value, max_value, min_weight, max_weight int) []Item {
	// Initialize a pseudorandom number generator.
	//random := rand.New(rand.NewSource(time.Now().UnixNano())) // Initialize with a changing seed
	random := rand.New(rand.NewSource(1337)) // Initialize with a fixed seed

	items := make([]Item, num_items)
	for i := 0; i < num_items; i++ {
		items[i] = Item{
			value:       random.Intn(max_value-min_value+1) + min_value,
			weight:      random.Intn(max_weight-min_weight+1) + min_weight,
			is_selected: false,
			id:          i,
			blocked_by:  -1,
			i_block:     nil,
		}
	}
	return items
}

// Return a copy of the items slice.
func copy_items(items []Item) []Item {
	new_items := make([]Item, len(items))
	copy(new_items, items)
	return new_items
}

// Return the total value of the items.
// If add_all is false, only add up the selected items.
func sum_values(items []Item, add_all bool) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if add_all || items[i].is_selected {
			total += items[i].value
		}
	}
	return total
}

// Return the total weight of the items.
// If add_all is false, only add up the selected items.
func sum_weights(items []Item, add_all bool) int {
	total := 0
	for i := 0; i < len(items); i++ {
		if add_all || items[i].is_selected {
			total += items[i].weight
		}
	}
	return total
}

// Return the value of this solution.
// If the solution is too heavy, return -1 so we prefer an empty solution.
func solution_value(items []Item, allowed_weight int) int {
	// If the solution's total weight > allowed_weight,
	// return 0 so we won't use this solution.
	if sum_weights(items, false) > allowed_weight {
		return -1
	}

	// Return the sum of the selected values.
	return sum_values(items, false)
}

// Print the selected items.
func print_selected(items []Item) {
	num_printed := 0
	for i, item := range items {
		if item.is_selected {
			fmt.Printf("%d(%d, %d) ", i, item.value, item.weight)
		}
		num_printed += 1
		if num_printed > 100 {
			fmt.Println("...")
			return
		}
	}
	fmt.Println()
}

func run_algorithm(alg func([]Item, int) ([]Item, int, int), items []Item, allowed_weight int) {
	// Copy the items so the run isn't influenced by a previous run.
	test_items := copy_items(items)

	start := time.Now()

	// Run the algorithm.
	solution, total_value, function_calls := alg(test_items, allowed_weight)

	elapsed := time.Since(start)

	fmt.Printf("Elapsed: %f\n", elapsed.Seconds())
	print_selected(solution)
	fmt.Printf("Value: %d, Weight: %d, Calls: %d\n",
		total_value, sum_weights(solution, false), function_calls)
	fmt.Println()
}

// Recursively assign values in or out of the solution.
// Return the best assignment, value of that assignment,
// and the number of function calls we made.
func exhaustive_search(items []Item, allowed_weight int) ([]Item, int, int) {
	return do_exhaustive_search(items, allowed_weight, 0)
}

func do_exhaustive_search(items []Item, allowed_weight, next_index int) ([]Item, int, int) {
	// base case - the 'next' index is one which does not exist
	if next_index == len(items) {
		// in this case, we already have a fully built solution from our previous function calls, so return it
		solution := copy_items(items)
		total_value := solution_value(items, allowed_weight)
		return solution, total_value, 1
	}

	// otherwise, we need to find the best solution from each of the two branches

	// try including this item
	items[next_index].is_selected = true
	incl_solution, incl_value, incl_calls_count := do_exhaustive_search(items, allowed_weight, next_index+1)

	// try excluding this item
	items[next_index].is_selected = false
	excl_solution, excl_value, excl_calls_count := do_exhaustive_search(items, allowed_weight, next_index+1)

	total_calls := incl_calls_count + excl_calls_count + 1
	if incl_value >= excl_value {
		// including is better
		return incl_solution, incl_value, total_calls
	} else {
		// excluding is better
		return excl_solution, excl_value, total_calls
	}
}

func branch_and_bound(items []Item, allowed_weight int) ([]Item, int, int) {
	best_value := 0
	current_value := 0
	current_weight := 0
	remaining_value := sum_values(items, true)

	return do_branch_and_bound(items, allowed_weight, 0, best_value, current_value, current_weight, remaining_value)
}

func do_branch_and_bound(items []Item, allowed_weight, next_index, best_value, current_value, current_weight, remaining_value int) ([]Item, int, int) {
	// base case - the 'next' index is one which does not exist
	// this is a full assignment
	if next_index == len(items) {
		// in this case, we already have a fully built solution from our previous function calls, so return it
		solution := copy_items(items)

		// assert that this is a new best solution
		if current_value <= best_value {
			panic("this is not better than the previous best solution")
		}

		return solution, current_value, 1
	}

	// we cannot do any better than the previous best solution, even if we added all the remaining items
	if current_value+remaining_value <= best_value {
		return nil, 0, 1
	}

	// we can try including this item
	var incl_solution []Item
	incl_solution, incl_value, incl_calls_count := nil, 0, 1
	if current_weight+items[next_index].weight <= allowed_weight {
		items[next_index].is_selected = true
		incl_solution, incl_value, incl_calls_count = do_branch_and_bound(items, allowed_weight, next_index+1, best_value, current_value+items[next_index].value, current_weight+items[next_index].weight, remaining_value-items[next_index].value)

		// if this solution has obtained a better value than the previously-known best value, then update this value
		if incl_value > best_value {
			best_value = incl_value
		}
	}

	// try excluding this item, only in the case that we have a shot at beating our current-best value
	var excl_solution []Item
	excl_solution, excl_value, excl_calls_count := nil, 0, 1
	if current_value+remaining_value-items[next_index].value > best_value {
		items[next_index].is_selected = false
		excl_solution, excl_value, excl_calls_count = do_branch_and_bound(items, allowed_weight, next_index+1, best_value, current_value, current_weight, remaining_value-items[next_index].value)
		// there is no need to update best_value because there are no more recursive function calls
	}

	total_calls := incl_calls_count + excl_calls_count + 1
	if incl_value >= excl_value {
		// including is better
		return incl_solution, incl_value, total_calls
	} else {
		// excluding is better
		return excl_solution, excl_value, total_calls
	}
}

// Build the items' block lists.
func make_block_lists(items []Item) {
	for i := range items {
		var s []int
		// check to see if i blocks j
		// this would mean that i is at least as good a choice as j (i must have greater than or equal value and smaller than or equal weight compared to j)
		for j := range items {
			if i != j && items[i].value >= items[j].value && items[i].weight <= items[j].weight {
				s = append(s, j)
			}
		}
		items[i].i_block = s
	}
}

// Block items on this item's blocks list.
func block_items(source Item, items []Item) {
	for _, v := range source.i_block {
		// ensure that this item is not already blocked
		if items[v].blocked_by == -1 {
			// mark this item as blocked by source
			items[v].blocked_by = source.id
		}
	}
}

// Unblock items on this item's blocks list.
func unblock_items(source Item, items []Item) {
	for _, v := range source.i_block {
		// ensure this item is blocked by source
		if items[v].blocked_by == source.id {
			// reset the blocked marking to -1
			items[v].blocked_by = -1
		}
	}
}

func rods_technique(items []Item, allowed_weight int) ([]Item, int, int) {
	best_value := 0
	current_value := 0
	current_weight := 0
	remaining_value := sum_values(items, true)

	make_block_lists(items)

	return do_rods_technique(items, allowed_weight, 0, best_value, current_value, current_weight, remaining_value)
}

type byBlockListLength []Item

func (s byBlockListLength) Len() int {
	return len(s)
}
func (s byBlockListLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byBlockListLength) Less(i, j int) bool {
	return len(s[i].i_block) > len(s[j].i_block)
}

func rods_technique_sorted(items []Item, allowed_weight int) ([]Item, int, int) {
	best_value := 0
	current_value := 0
	current_weight := 0
	remaining_value := sum_values(items, true)

	make_block_lists(items)

	sort.Sort(byBlockListLength(items))

	for i := range items {
		items[i].id = i
	}

	make_block_lists(items)

	return do_rods_technique(items, allowed_weight, 0, best_value, current_value, current_weight, remaining_value)
}

func do_rods_technique(items []Item, allowed_weight, next_index, best_value, current_value, current_weight, remaining_value int) ([]Item, int, int) {
	// base case - the 'next' index is one which does not exist
	// this is a full assignment
	if next_index == len(items) {
		// in this case, we already have a fully built solution from our previous function calls, so return it
		solution := copy_items(items)

		// assert that this is a new best solution
		if current_value <= best_value {
			panic("this is not better than the previous best solution")
		}

		return solution, current_value, 1
	}

	// we cannot do any better than the previous best solution, even if we added all the remaining items
	if current_value+remaining_value <= best_value {
		return nil, 0, 1
	}

	// we can try including this item (only if not blocked)
	var incl_solution []Item
	incl_solution, incl_value, incl_calls_count := nil, 0, 1
	if items[next_index].blocked_by == -1 && current_weight+items[next_index].weight <= allowed_weight {
		items[next_index].is_selected = true
		incl_solution, incl_value, incl_calls_count = do_rods_technique(items, allowed_weight, next_index+1, best_value, current_value+items[next_index].value, current_weight+items[next_index].weight, remaining_value-items[next_index].value)

		// if this solution has obtained a better value than the previously-known best value, then update this value
		if incl_value > best_value {
			best_value = incl_value
		}
	}

	// try excluding this item, only in the case that we have a shot at beating our current-best value
	var excl_solution []Item
	excl_solution, excl_value, excl_calls_count := nil, 0, 1
	if current_value+remaining_value-items[next_index].value > best_value {
		block_items(items[next_index], items)

		items[next_index].is_selected = false
		excl_solution, excl_value, excl_calls_count = do_rods_technique(items, allowed_weight, next_index+1, best_value, current_value, current_weight, remaining_value-items[next_index].value)
		// there is no need to update best_value because there are no more recursive function calls

		unblock_items(items[next_index], items)
	}

	total_calls := incl_calls_count + excl_calls_count + 1
	if incl_value >= excl_value {
		// including is better
		return incl_solution, incl_value, total_calls
	} else {
		// excluding is better
		return excl_solution, excl_value, total_calls
	}
}

func main() {
	items := make_items(num_items, min_value, max_value, min_weight, max_weight)
	allowed_weight = sum_weights(items, true) / 2

	// Display basic parameters.
	fmt.Println("*** Parameters ***")
	fmt.Printf("# items: %d\n", num_items)
	fmt.Printf("Total value: %d\n", sum_values(items, true))
	fmt.Printf("Total weight: %d\n", sum_weights(items, true))
	fmt.Printf("Allowed weight: %d\n", allowed_weight)
	fmt.Println()

	// Rod's technique
	if num_items > 350 {
		fmt.Println("Too many items for Rod's technique")
	} else {
		fmt.Println("*** Rod's technique ***")
		run_algorithm(rods_technique_sorted, items, allowed_weight)
	}
}
