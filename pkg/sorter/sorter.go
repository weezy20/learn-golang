package sorter

import (
	"cmp"
)

// This package would offer almost the same constraint as cmp.Ordered
// import "golang.org/x/exp/constraints"

// Define an interface sorter that is implemented by each type of sorter.
// This interface has one method, sort() that takes in a vector of generic type T and sorts it
type Sorter[T cmp.Ordered] interface {
	Sort([]T)
}

// Define a struct that implements the Sorter interface
type BubbleSort[T cmp.Ordered] struct{}

// Implement the Sorter interface for BubbleSort that takes a vector of generic type T and sorts it
// Go interfaces are implemented implicitly, there's no explicit declaration of intent
func (b *BubbleSort[T]) Sort(v []T) {
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

type InsertionSort[T cmp.Ordered] struct{}

func (i *InsertionSort[T]) Sort(v []T) {
	n := len(v)
	for i := 1; i < n; i++ {
		key := v[i]
		j := i - 1
		for j >= 0 && v[j] > key {
			v[j+1] = v[j]
			j = j - 1
		}
		v[j+1] = key
	}
}
