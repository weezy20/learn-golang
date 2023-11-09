package sorter

import "golang.org/x/exp/constraints"

// Define an interface sorter that is implemented by each type of sorter.
// This interface has one method, sort() that takes in a vector of generic type T and sorts it
type Sorter[T constraints.Ordered] interface {
	Sort([]T)
}

// Define a struct that implements the Sorter interface
type BinarySort[T constraints.Ordered] struct{}

// Implement the Sorter interface for BinarySort that takes a vector of generic type T and sorts it
// Go interfaces are implemented implicitly, there's no explicit declaration of intent
func (b *BinarySort[T]) Sort(v []T) {
	// implement binary sort on v
	n := len(v)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if v[j] > v[j+1] {
				v[j], v[j+1] = v[j+1], v[j]
			}
		}
	}
}
