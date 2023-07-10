package main

import (
	"fmt"
	"math/rand"
	"strings"
)

// djb2 hash function. See http://www.cse.yorku.ca/~oz/hash.html.
func hash(value string) int {
	hash := 5381
	for _, ch := range value {
		hash = ((hash << 5) + hash) + int(ch)
	}

	// Make sure the result is non-negative.
	if hash < 0 {
		hash = -hash
	}
	return hash
}

type Employee struct {
	name    string
	phone   string
	deleted bool
}

type QuadraticProbingHashTable struct {
	capacity  int
	employees []*Employee
}

// Initialize a LinearProbingHashTable and return a pointer to it.
func NewLinearProbingHashTable(capacity int) *QuadraticProbingHashTable {
	return &QuadraticProbingHashTable{
		capacity:  capacity,
		employees: make([]*Employee, capacity),
	}
}

// Display the hash table's contents.
func (hash_table *QuadraticProbingHashTable) dump() {
	for i, employee := range hash_table.employees {
		fmt.Printf(" %d: ", i)
		if employee == nil {
			fmt.Println("---")
			continue
		}
		// if we have reached this point, there is employee data at this index
		// (but it may be deleted)
		if employee.deleted {
			fmt.Println("xxx")
			continue
		}
		num_spaces := 16 - len(employee.name)
		fmt.Printf(
			"%s%s%s\n",
			employee.name,
			strings.Repeat(" ", num_spaces),
			employee.phone,
		)
	}
}

// Return the key's index or where it would be if present and
// the probe sequence length.
// If the key is not present and the table is full, return -1 for the index.
func (hash_table *QuadraticProbingHashTable) find(name string) (int, int) {
	hash := hash(name) % hash_table.capacity
	// the index of the first deleted item we come across (if we find one)
	deleted_index := -1

	// enter a loop looking for a value with the matching key
	// we start with i=0 and end with i=hash_table.capacity-1 because any hash_table.capacity%hash_table.capacity==0
	for i := 0; i < hash_table.capacity; i++ {
		probe_index := (hash + i*i) % hash_table.capacity
		if hash_table.employees[probe_index] == nil {
			// great! we have found an empty spot
			// we know that the target is not in the table
			if deleted_index == -1 {
				// we have not passed any deleted entries in coming to this point
				return probe_index, i + 1
			} else {
				// we saw a deleted entry earlier
				return deleted_index, i + 1
			}
		}
		// now we know this spot is not empty
		// if this is our first time coming across a deleted spot, save the index and move on
		if hash_table.employees[probe_index].deleted {
			if deleted_index == -1 {
				deleted_index = probe_index
			}
			continue
		}
		// now we know this spot is not empty and is not deleted
		// if it contains our key, return it!
		if hash_table.employees[probe_index].name == name {
			return probe_index, i + 1
		}
	}
	if deleted_index != -1 {
		// we saw a deleted entry earlier
		return deleted_index, hash_table.capacity
	}
	// they key is not present and the table does not have any empty or deleted spots (it is full)
	return -1, hash_table.capacity
}

// Add an item to the hash table.
func (hash_table *QuadraticProbingHashTable) set(name string, phone string) {
	index, _ := hash_table.find(name)

	// the key is not in the table and there is no room for new keys
	if index == -1 {
		panic("not enough space in hash table")
	}

	// now we know the key either exists at the index or it is an empty/recycled spot

	if hash_table.employees[index] == nil || hash_table.employees[index].deleted {
		// if it is empty/recycled, create a new struct
		hash_table.employees[index] = &Employee{name: name, phone: phone}
	} else {
		// otherwise, just update the phone number
		hash_table.employees[index].phone = phone
	}
}

// Return an item from the hash table.
func (hash_table *QuadraticProbingHashTable) get(name string) string {
	index, _ := hash_table.find(name)

	// the key is not in the table and the table is full, or we have been given an empty/recycled spot
	if index == -1 || hash_table.employees[index] == nil || hash_table.employees[index].deleted {
		return ""
	}

	return hash_table.employees[index].phone
}

// Return true if the person is in the hash table.
func (hash_table *QuadraticProbingHashTable) contains(name string) bool {
	index, _ := hash_table.find(name)
	// the person is in the hash table if the index is not -1 and if the pointer at the index is not nil and if the spot if not recycled
	return index != -1 && hash_table.employees[index] != nil && !hash_table.employees[index].deleted
}

func (hash_table *QuadraticProbingHashTable) delete(name string) {
	index, _ := hash_table.find(name)
	// if the table contains this key, mark it as deleted
	if hash_table.contains(name) {
		hash_table.employees[index].deleted = true
	}
}

// Show this key's probe sequence.
func (hash_table *QuadraticProbingHashTable) probe(name string) int {
	// Hash the key.
	hash := hash(name) % hash_table.capacity
	fmt.Printf("Probing %s (%d)\n", name, hash)

	// Keep track of a deleted spot if we find one.
	deleted_index := -1

	// Probe up to hash_table.capacity times.
	for i := 0; i < hash_table.capacity; i++ {
		index := (hash + i*i) % hash_table.capacity

		fmt.Printf("    %d: ", index)
		if hash_table.employees[index] == nil {
			fmt.Printf("---\n")
		} else if hash_table.employees[index].deleted {
			fmt.Printf("xxx\n")
		} else {
			fmt.Printf("%s\n", hash_table.employees[index].name)
		}

		// If this spot is empty, the value isn't in the table.
		if hash_table.employees[index] == nil {
			// If we found a deleted spot, return its index.
			if deleted_index >= 0 {
				fmt.Printf("    Returning deleted index %d\n", deleted_index)
				return deleted_index
			}

			// Return this index, which holds nil.
			fmt.Printf("    Returning nil index %d\n", index)
			return index
		}

		// If this spot is deleted, remember where it is.
		if hash_table.employees[index].deleted {
			if deleted_index < 0 {
				deleted_index = index
			}
		} else if hash_table.employees[index].name == name {
			// If this cell holds the key, return its data.
			fmt.Printf("    Returning found index %d\n", index)
			return index
		}

		// Otherwise continue the loop.
	}

	// If we get here, then the key is not
	// in the table and the table is full.

	// If we found a deleted spot, return it.
	if deleted_index >= 0 {
		fmt.Printf("    Returning deleted index %d\n", deleted_index)
		return deleted_index
	}

	// There's nowhere to put a new entry.
	fmt.Printf("    Table is full\n")
	return -1
}

// Make a display showing whether each slice entry is nil.
func (hash_table *QuadraticProbingHashTable) dump_concise() {
	// Loop through the slice.
	for i, employee := range hash_table.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
		} else if employee.deleted {
			// This spot is deleted.
			fmt.Printf("x")
		} else {
			// Display this entry.
			fmt.Printf("O")
		}
		if i%50 == 49 {
			fmt.Println()
		}
	}
	fmt.Println()
}

// Return the average probe sequence length for the items in the table.
func (hash_table *QuadraticProbingHashTable) ave_probe_sequence_length() float32 {
	total_length := 0
	num_values := 0
	for _, employee := range hash_table.employees {
		if employee != nil {
			_, probe_length := hash_table.find(employee.name)
			total_length += probe_length
			num_values++
		}
	}
	return float32(total_length) / float32(num_values)
}

func main() {
	// Make some names.
	employees := []Employee{
		Employee{"Ann Archer", "202-555-0101", false},
		Employee{"Bob Baker", "202-555-0102", false},
		Employee{"Cindy Cant", "202-555-0103", false},
		Employee{"Dan Deever", "202-555-0104", false},
		Employee{"Edwina Eager", "202-555-0105", false},
		Employee{"Fred Franklin", "202-555-0106", false},
		Employee{"Gina Gable", "202-555-0107", false},
	}

	hash_table := NewLinearProbingHashTable(10)
	for _, employee := range employees {
		hash_table.set(employee.name, employee.phone)
	}
	hash_table.dump()

	hash_table.probe("Hank Hardy")
	fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hash_table.delete("Dan Deever")
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hash_table.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	hash_table.dump()

	hash_table.probe("Ann Archer")
	hash_table.probe("Bob Baker")
	hash_table.probe("Cindy Cant")
	hash_table.probe("Dan Deever")
	hash_table.probe("Edwina Eager")
	hash_table.probe("Fred Franklin")
	hash_table.probe("Gina Gable")
	hash_table.set("Hank Hardy", "202-555-0108")
	hash_table.probe("Hank Hardy")

	// Look at clustering.
	random := rand.New(rand.NewSource(12345)) // Initialize with an unchanging seed
	big_capacity := 1009
	big_hash_table := NewLinearProbingHashTable(big_capacity)
	num_items := int(float32(big_capacity) * 0.9)
	for i := 0; i < num_items; i++ {
		str := fmt.Sprintf("%d-%d", i, random.Intn(1000000))
		big_hash_table.set(str, str)
	}
	big_hash_table.dump_concise()
	fmt.Printf("Average probe sequence length: %f\n",
		big_hash_table.ave_probe_sequence_length())
}
