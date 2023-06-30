package main

import (
	"fmt"
)

type Cell struct {
	data string
	prev *Cell
	next *Cell
}

type DoublyLinkedList struct {
	top_sentinel    *Cell
	bottom_sentinel *Cell
}

func make_doubly_linked_list() DoublyLinkedList {
	// Create the sentinels.
	top_sentinel := Cell{prev: nil, next: nil}
	bottom_sentinel := Cell{prev: nil, next: nil}

	// Make them point to each other.
	top_sentinel.next = &bottom_sentinel
	bottom_sentinel.prev = &top_sentinel

	return DoublyLinkedList{top_sentinel: &top_sentinel, bottom_sentinel: &bottom_sentinel}
}

// Add a cell immadiately after me.
func (me *Cell) add_after(after *Cell) {
	other := (*me).next

	// The ordering should now be: me, after, other

	after.next = other
	after.prev = me

	me.next = after
	other.prev = after
}

// Add a cell immediately before me.
func (me *Cell) add_before(before *Cell) {
	// This is equivalent to adding this cell immedaitely after my prev.
	me.prev.add_after(before)
}

// Delete me.
func (me *Cell) delete() Cell {
	if me.next == nil || me.prev == nil {
		panic("no cell after me, or no cell before me")
	}

	me.prev.next = me.next
	me.next.prev = me.prev

	return *me
}

func (list *DoublyLinkedList) add_range(values []string) {
	// iterate through the values to be added
	for _, value := range values {
		// create a new cell
		newCell := Cell{data: value}
		// use add_before to add that cell before the bottom sentinel
		list.bottom_sentinel.add_before(&newCell)
	}
}

func (list *DoublyLinkedList) to_string(separator string) string {
	output := ""

	// grab the first cell and add its data to the output
	cell := list.top_sentinel.next
	// if this cell is in fact the bottom sentinel, return nothing
	if cell.next == nil {
		return output
	}
	// otherwise, this is a real cell and has data
	output += cell.data
	for {
		// move onto the next cell
		cell = cell.next

		// if this cell is in fact the bottom sentinel, break
		if cell.next == nil {
			break
		}

		// otherwise, it is a real cell and has data
		output += separator
		output += cell.data
	}

	return output
}

func (list *DoublyLinkedList) length() int {
	count := 0

	// grab the first cell
	cell := list.top_sentinel.next
	for {
		// if this cell is in fact the bottom sentinel, break
		if cell.next == nil {
			break
		}

		// otherwise, it is a real cell
		count++
		cell = cell.next
	}

	return count
}

func (list *DoublyLinkedList) is_empty() bool {
	// the list is empty if the cell after the top sentinel is in fact the bottom sentinel
	return list.top_sentinel.next.next == nil
}

func (list *DoublyLinkedList) push(value string) {
	// create a new cell to hold the new item
	newCell := Cell{data: value}
	// use add_after to add the new cell after the top sentinel
	sentinel := list.top_sentinel
	sentinel.add_after(&newCell)
}

func (list *DoublyLinkedList) pop() string {
	return list.top_sentinel.next.delete().data
}

/*
func main() {
	// Make a list from a slice of values.
	list := make_doubly_linked_list()
	if list.to_string(" ") != "" || list.length() != 0 || list.is_empty() != true {
		log.Fatal()
	}

	animals := []string{
		"Ant",
		"Bat",
		"Cat",
		"Dog",
		"Elk",
		"Fox",
	}
	list.add_range(animals)
	if list.to_string(" ") != "Ant Bat Cat Dog Elk Fox" || list.length() != 6 || list.is_empty() != false {
		log.Fatal()
	}

	aaa := Cell{data: "aaa"}
	list.top_sentinel.add_after(&aaa)
	if list.to_string(" ") != "aaa Ant Bat Cat Dog Elk Fox" {
		log.Fatal()
	}
	aaa.add_before(&Cell{data: "bbb"})
	if list.to_string(" ") != "bbb aaa Ant Bat Cat Dog Elk Fox" {
		log.Fatal()
	}
	aaa_deleted := aaa.delete()
	if list.to_string(" ") != "bbb Ant Bat Cat Dog Elk Fox" || aaa_deleted.data != aaa.data {
		log.Fatal()
	}

	list.push("c")
	if list.to_string(" ") != "c bbb Ant Bat Cat Dog Elk Fox" {
		log.Fatal()
	}
	pop_result := list.pop()
	if list.to_string(" ") != "bbb Ant Bat Cat Dog Elk Fox" || pop_result != "c" {
		log.Fatal()
	}

	fmt.Println("succcess!!!")
}
*/

func (list *DoublyLinkedList) enqueue(value string) {
	list.push(value)
}

func (list *DoublyLinkedList) dequeue() string {
	// remove the item before the bottom sentinel
	return list.bottom_sentinel.prev.delete().data
}

func (list *DoublyLinkedList) push_bottom(value string) {
	// add an item to the bottom of the list just before the bottom sentinel
	list.bottom_sentinel.add_before(&Cell{data: value})
}

func (list *DoublyLinkedList) push_top(value string) {
	// add an item to the top of the list just after the top sentinel
	list.top_sentinel.add_after(&Cell{data: value})
}

func (list *DoublyLinkedList) pop_top() string {
	return list.top_sentinel.next.delete().data
}

func (list *DoublyLinkedList) pop_bottom() string {
	return list.bottom_sentinel.prev.delete().data
}

func main() {
	// Test queue functions.
	fmt.Printf("*** Queue Functions ***\n")
	queue := make_doubly_linked_list()
	queue.enqueue("Agate")
	queue.enqueue("Beryl")
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Citrine")
	fmt.Printf("%s ", queue.dequeue())
	fmt.Printf("%s ", queue.dequeue())
	queue.enqueue("Diamond")
	queue.enqueue("Emerald")
	for !queue.is_empty() {
		fmt.Printf("%s ", queue.dequeue())
	}
	fmt.Printf("\n\n")

	// Test deque functions. Names starting
	// with F have a fast pass.
	fmt.Printf("*** Deque Functions ***\n")
	deque := make_doubly_linked_list()
	deque.push_top("Ann")
	deque.push_top("Ben")
	fmt.Printf("%s ", deque.pop_bottom())
	deque.push_bottom("F-Cat")
	fmt.Printf("%s ", deque.pop_bottom())
	fmt.Printf("%s ", deque.pop_bottom())
	deque.push_bottom("F-Dan")
	deque.push_top("Eva")
	for !deque.is_empty() {
		fmt.Printf("%s ", deque.pop_bottom())
	}
	fmt.Printf("\n")
}
