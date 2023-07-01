package main

import (
	"fmt"
	"log"
	"strings"
)

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

func (node *Node) insert_value(value string) {
	new_node := Node{data: value}

	current_node := node
	for {
		if value < current_node.data {
			// go down the left branch
			// if there is no node on the left branch, we are done
			if current_node.left == nil {
				current_node.left = &new_node
				return
			}
			// otherwise, make the left node the new node and continue
			current_node = current_node.left
		} else if value > current_node.data {
			// go down the right branch
			// if there is no node on the right branch, we are done
			if current_node.right == nil {
				current_node.right = &new_node
				return
			}
			// otherwise, make the right node the new node and continue
			current_node = current_node.right
		} else {
			log.Panic("New node data equal to data of existing node.")
		}

	}
}

func (node *Node) find_value(value string) *Node {
	current_node := node
	for {
		if current_node == nil {
			// no match
			return nil
		}
		if value < current_node.data {
			// go down the left branch
			current_node = current_node.left
		} else if value > current_node.data {
			// go down the right branch
			current_node = current_node.right
		} else {
			// we now know that value == current_node.data, so return current_node
			return current_node
		}

	}
}

func main() {
	// Make a root node to act as sentinel.
	root := Node{"", nil, nil}

	// Add some values.
	root.insert_value("I")
	root.insert_value("G")
	root.insert_value("C")
	root.insert_value("E")
	root.insert_value("B")
	root.insert_value("K")
	root.insert_value("S")
	root.insert_value("Q")
	root.insert_value("M")

	// Add F.
	root.insert_value("F")

	// Display the values in sorted order.
	fmt.Printf("Sorted values: %s\n", root.right.inorder())

	// Let the user search for values.
	for {
		// Get the target value.
		target := ""
		fmt.Printf("String: ")
		fmt.Scanln(&target)
		if len(target) == 0 {
			break
		}

		// Find the value's node.
		node := root.find_value(target)
		if node == nil {
			fmt.Printf("%s not found\n", target)
		} else {
			fmt.Printf("Found value %s\n", target)
		}
	}
}
