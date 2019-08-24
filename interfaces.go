package learn

import "fmt"

//Search by int
type Search interface {
	Contains(i int) bool
}

//Ints is a slice of ints
type Ints []int

//Contains returns true if a given int is in a slice
func (sliceOfInts Ints) Contains(i int) bool {
	for _, item := range sliceOfInts {
		fmt.Printf("Item=%v, i=%v\n", item, i)
		if item == i {
			return true
		}
	}
	return false
}
