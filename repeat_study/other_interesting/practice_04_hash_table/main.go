package main

import "fmt"

// ArraySize size of hash table
const ArraySize = 7

// HashTable struct - insert, search, delete
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket struct - insert, search, delete
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// Search take key and return true if key in hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// Delete take key and delete from hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)

}

// insert take key create a node with the key and store in the bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Printf("%s already exist", k)
	}
}

// search take key and find in from the bucket
func (b *bucket) search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete take key and remove from the bucket
func (b *bucket) delete(k string) {

	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash func
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize // hash table woth size from 0 to 99
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}
func main() {
	hashTable := Init()
	list := []string{
		"c",
		"python",
		"golang",
		"ruby",
		"java",
		"swift",
		"token",
	}
	for _, v := range list {
		hashTable.Insert(v)
	}
	hashTable.Delete("token")
}
