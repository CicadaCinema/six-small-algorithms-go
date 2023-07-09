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
	name  string
	phone string
}

type LinearProbingHashTable struct {
	capacity  int
	employees []*Employee
}

// Initialize a LinearProbingHashTable and return a pointer to it.
func NewLinearProbingHashTable(capacity int) *LinearProbingHashTable {
	return &LinearProbingHashTable{
		capacity:  capacity,
		employees: make([]*Employee, capacity),
	}
}

// Display the hash table's contents.
func (hash_table *LinearProbingHashTable) dump() {
	for i, employee := range hash_table.employees {
		fmt.Printf(" %d: ", i)
		if employee == nil {
			fmt.Println("---")
			continue
		}
		// if we have reached this point, there is employee data at this index
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
func (hash_table *LinearProbingHashTable) find(name string) (int, int) {
	target_index := hash(name) % hash_table.capacity

	// enter a loop looking for a value with the matching key
	// we start with i=0 and end with i=hash_table.capacity-1 because any hash_table.capacity%hash_table.capacity==0
	for i := 0; i < hash_table.capacity; i++ {
		probe_index := (target_index + i) % hash_table.capacity
		// if there is a different key in this spot, move on
		if hash_table.employees[probe_index] != nil && hash_table.employees[probe_index].name != name {
			continue
		}
		// now we know this spot either contains the correct key or is empty
		return probe_index, i + 1
	}
	// they key is not present and the table does not have any empty spots (it is full)
	return -1, hash_table.capacity
}

// Add an item to the hash table.
func (hash_table *LinearProbingHashTable) set(name string, phone string) {
	index, _ := hash_table.find(name)

	// the key is not in the table and there is no room for new keys
	if index == -1 {
		panic("not enough space in hash table")
	}

	// now we know the key either exists at the index or it is an empty spot

	if hash_table.employees[index] == nil {
		// if it is empty, create a new struct
		hash_table.employees[index] = &Employee{name: name, phone: phone}
	} else {
		// otherwise, just update the phone number
		hash_table.employees[index].phone = phone
	}
}

// Return an item from the hash table.
func (hash_table *LinearProbingHashTable) get(name string) string {
	index, _ := hash_table.find(name)

	// the key is not in the table
	if index == -1 || hash_table.employees[index] == nil {
		return ""
	}

	return hash_table.employees[index].phone
}

// Return true if the person is in the hash table.
func (hash_table *LinearProbingHashTable) contains(name string) bool {
	index, _ := hash_table.find(name)
	// the person is in the hash table if the index is not -1 and if the pointer at the index is not nil
	return index != -1 && hash_table.employees[index] != nil
}

// Make a display showing whether each slice entry is nil.
func (hash_table *LinearProbingHashTable) dump_concise() {
	// Loop through the slice.
	for i, employee := range hash_table.employees {
		if employee == nil {
			// This spot is empty.
			fmt.Printf(".")
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
func (hash_table *LinearProbingHashTable) ave_probe_sequence_length() float32 {
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
		Employee{"Ann Archer", "202-555-0101"},
		Employee{"Bob Baker", "202-555-0102"},
		Employee{"Cindy Cant", "202-555-0103"},
		Employee{"Dan Deever", "202-555-0104"},
		Employee{"Edwina Eager", "202-555-0105"},
		Employee{"Fred Franklin", "202-555-0106"},
		Employee{"Gina Gable", "202-555-0107"},
	}

	hash_table := NewLinearProbingHashTable(10)
	for _, employee := range employees {
		hash_table.set(employee.name, employee.phone)
	}
	hash_table.dump()

	fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	// fmt.Println("Deleting Dan Deever")
	// hash_table.delete("Dan Deever")
	// fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hash_table.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))

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
