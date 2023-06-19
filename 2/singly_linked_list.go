package main

import "fmt"

type Cell struct {
	data string
	next *Cell
}

type LinkedList struct {
	sentinel *Cell
}

func make_linked_list() LinkedList {
	sentinel := Cell{next: nil}
	return LinkedList{sentinel: &sentinel}
}

// Add a cell immadiately after me.
func (me *Cell) add_after(after *Cell) {
	after.next = me.next
	me.next = after
}

// Delete a cell immadiately after me.
func (me *Cell) delete_after() Cell {
	if me.next == nil {
		panic("no cell after me")
	}
	deleted := *me.next
	me.next = deleted.next
	return deleted
}

func (list *LinkedList) add_range(values []string) {
	// last_cell is a pointer to the last cell
	last_cell := list.sentinel
	for {
		if last_cell.next != nil {
			// if there is a cell after last_cell, make that the new last_cell
			last_cell = last_cell.next
		} else {
			// otherwise, last_cell.next == nil and we are done
			break
		}

	}

	// iterate through the values to be added
	for _, value := range values {
		// create a new cell
		newCell := Cell{data: value}
		// use add_after to add that cell after last_cell
		last_cell.add_after(&newCell)
		// make last_cell point to the new last cell
		last_cell = &newCell
	}
}

func (list *LinkedList) to_string(separator string) string {
	output := ""

	// top is a pointer to the first cell
	top := list.sentinel.next
	for cell := top; cell != nil; cell = cell.next {
		output += cell.data
		// only output a separator if this cell is not the last cell
		// this loop can be refactored because we are essentially checking the same thing twice (the first time was in the for loop definition)
		if cell.next != nil {
			output += separator
		}
	}

	return output
}

func (list *LinkedList) length() int {
	count := 0

	// re-use the code above, incrementing count instead of reading the cell data
	top := list.sentinel.next
	for cell := top; cell != nil; cell = cell.next {
		count++
	}

	return count
}

func (list *LinkedList) is_empty() bool {
	sentinel := *list.sentinel
	// the list is empty if the sentinel’s next pointer is nil
	return sentinel.next == nil
}

func (list *LinkedList) push(value string) {
	// create a new cell to hold the new item
	newCell := Cell{data: value}
	// use add_after to add the new cell after the sentinel
	sentinel := list.sentinel
	sentinel.add_after(&newCell)
}

func (list *LinkedList) pop() string {
	// use delete_after to remove the cell from the list
	sentinel := list.sentinel
	return sentinel.delete_after().data
}

func main() {
	// small_list_test()

	// Make a list from a slice of values.
	greek_letters := []string{
		"α", "β", "γ", "δ", "ε",
	}
	list := make_linked_list()
	list.add_range(greek_letters)
	fmt.Println(list.to_string(" "))
	fmt.Println()

	// Demonstrate a stack.
	stack := make_linked_list()
	stack.push("Apple")
	stack.push("Banana")
	stack.push("Coconut")
	stack.push("Date")
	for !stack.is_empty() {
		fmt.Printf("Popped: %-7s   Remaining %d: %s\n",
			stack.pop(),
			stack.length(),
			stack.to_string(" "))
	}
}
