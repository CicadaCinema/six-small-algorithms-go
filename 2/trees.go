package main

import (
	"fmt"
	"strings"
)

type Cell struct {
	data *Node
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

func (list *DoublyLinkedList) push(value *Node) {
	// create a new cell to hold the new item
	newCell := Cell{data: value}
	// use add_after to add the new cell after the top sentinel
	sentinel := list.top_sentinel
	sentinel.add_after(&newCell)
}

func (list *DoublyLinkedList) enqueue(value *Node) {
	list.push(value)
}

func (list *DoublyLinkedList) dequeue() *Node {
	// remove the item before the bottom sentinel
	return list.bottom_sentinel.prev.delete().data
}

type Node struct {
	data  string
	left  *Node
	right *Node
}

func build_tree() *Node {
	a := Node{"A", nil, nil}
	b := Node{"B", nil, nil}
	c := Node{"C", nil, nil}
	d := Node{"D", nil, nil}
	e := Node{"E", nil, nil}
	f := Node{"F", nil, nil}
	g := Node{"G", nil, nil}
	h := Node{"H", nil, nil}
	i := Node{"I", nil, nil}
	j := Node{"J", nil, nil}

	a.left = &b
	a.right = &c

	b.left = &d
	b.right = &e

	e.left = &g

	c.right = &f
	f.left = &h

	h.left = &i
	h.right = &j

	// return the root node
	return &a
}

func (node *Node) display_indented(indent string, depth int) string {
	result := ""

	// display the given node
	result += strings.Repeat(indent, depth)
	result += node.data
	result += "\n"

	// display the children
	if node.left != nil {
		result += node.left.display_indented(indent, depth+1)
	}
	if node.right != nil {
		result += node.right.display_indented(indent, depth+1)
	}

	return result

}

func (node *Node) preorder() string {
	result := ""

	// display the given node
	result += node.data

	// display the children
	if node.left != nil {
		result += " " + node.left.preorder()
	}
	if node.right != nil {
		result += " " + node.right.preorder()
	}

	return result
}

func (node *Node) inorder() string {
	result := ""

	if node.left != nil {
		result += node.left.inorder() + " "
	}

	result += node.data

	if node.right != nil {
		result += " " + node.right.inorder()
	}

	return result
}

func (node *Node) postorder() string {
	result := ""

	if node.left != nil {
		result += node.left.postorder() + " "
	}

	if node.right != nil {
		result += node.right.postorder() + " "
	}

	result += node.data

	return result
}

func (node *Node) breadth_first() string {
	result := ""
	queue := make_doubly_linked_list()

	queue.enqueue(node)

	for !queue.is_empty() {
		next_node_pointer := queue.dequeue()
		result += next_node_pointer.data
		if next_node_pointer.left != nil {
			queue.enqueue(next_node_pointer.left)
		}
		if next_node_pointer.right != nil {
			queue.enqueue(next_node_pointer.right)
		}
		if !queue.is_empty() {
			result += " "
		}
	}

	return result
}

func main() {
	// Build a tree.
	a_node := build_tree()

	// Display with indentation.
	fmt.Println(a_node.display_indented("  ", 0))

	// Display traversals.
	fmt.Println("Preorder:     ", a_node.preorder())
	fmt.Println("Inorder:      ", a_node.inorder())
	fmt.Println("Postorder:    ", a_node.postorder())
	fmt.Println("Breadth first:", a_node.breadth_first())
}
