package main

import "fmt"

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

type ChainingHashTable struct {
	num_buckets int
	buckets     [][]*Employee
}

// Initialize a ChainingHashTable and return a pointer to it.
func NewChainingHashTable(num_buckets int) *ChainingHashTable {
	return &ChainingHashTable{
		num_buckets: num_buckets,
		buckets:     make([][]*Employee, num_buckets),
	}
}

// Display the hash table's contents.
func (hash_table *ChainingHashTable) dump() {
	for i, bucket := range hash_table.buckets {
		fmt.Printf("Bucket %d:\n", i)
		for _, employee := range bucket {
			fmt.Printf("    %s: %s\n", employee.name, employee.phone)
		}
	}
}

// Find the bucket and Employee holding this key.
// Return the bucket number and Employee number in the bucket.
// If the key is not present, return the bucket number and -1.
func (hash_table *ChainingHashTable) find(name string) (int, int) {
	bucket_number := hash(name) % hash_table.num_buckets
	for i, test_employee := range hash_table.buckets[bucket_number] {
		if test_employee.name == name {
			return bucket_number, i
		}
	}
	return bucket_number, -1
}

// Add an item to the hash table.
func (hash_table *ChainingHashTable) set(name string, phone string) {
	bucket_number, employee_number := hash_table.find(name)

	// if the emplyee is already present in the hash table, update their phone number
	if employee_number >= 0 {
		hash_table.buckets[bucket_number][employee_number].phone = phone
		return
	}

	hash_table.buckets[bucket_number] = append(hash_table.buckets[bucket_number], &Employee{name: name, phone: phone})
}

// Return an item from the hash table.
func (hash_table *ChainingHashTable) get(name string) string {
	bucket_number, employee_number := hash_table.find(name)

	// if the employee is present in the hash table
	if employee_number >= 0 {
		return hash_table.buckets[bucket_number][employee_number].phone
	}

	// key is not in the hash table, so return an empty string as a default value
	return ""
}

// Return true if the person is in the hash table.
func (hash_table *ChainingHashTable) contains(name string) bool {
	_, employee_number := hash_table.find(name)
	// if employee_number is not -1, then the key exists
	return employee_number != -1
}

// Delete this key's entry.
func (hash_table *ChainingHashTable) delete(name string) {
	bucket_number, employee_number := hash_table.find(name)

	// the key was never present to begin with
	if employee_number == -1 {
		return
	}

	// cut that struct out of its bucket
	hash_table.buckets[bucket_number] =
		append(hash_table.buckets[bucket_number][:employee_number],
			hash_table.buckets[bucket_number][employee_number+1:]...)
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
		Employee{"Herb Henshaw", "202-555-0108"},
		Employee{"Ida Iverson", "202-555-0109"},
		Employee{"Jeb Jacobs", "202-555-0110"},
	}

	hash_table := NewChainingHashTable(10)
	for _, employee := range employees {
		hash_table.set(employee.name, employee.phone)
	}
	hash_table.dump()

	fmt.Printf("Table contains Sally Owens: %t\n", hash_table.contains("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Println("Deleting Dan Deever")
	hash_table.delete("Dan Deever")
	fmt.Printf("Sally Owens: %s\n", hash_table.get("Sally Owens"))
	fmt.Printf("Table contains Dan Deever: %t\n", hash_table.contains("Dan Deever"))
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
	fmt.Println("Changing Fred Franklin")
	hash_table.set("Fred Franklin", "202-555-0100")
	fmt.Printf("Fred Franklin: %s\n", hash_table.get("Fred Franklin"))
}
