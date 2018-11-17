package listbench

import (
	"testing"
)

func Benchmark_2_LargeObject_SliceTraversalAndAccess(b *testing.B) {

	// Create the objects
	o := make([]*largeobj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &largeobj{id: i}
	}

	// Setup the slice
	s := make([]*largeobj, b.N)

	// Add orders to the slice
	for i := 0; i < b.N; i++ {
		s = append(s, o[i])
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s[i]
	}
}

func Benchmark_2_LargeObject_SliceTraversalAndAccessWithRange(b *testing.B) {

	// Create the objects
	o := make([]*largeobj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &largeobj{id: i}
	}

	// Setup the slice
	s := make([]*largeobj, 0)

	// Add orders to the slice
	for i := 0; i < b.N; i++ {
		s = append(s, o[i])
	}

	b.ResetTimer()

	for i, j := range s {
		_, _ = i, j
	}
}

func Benchmark_2_LargeObject_MapTraversalAndAccess(b *testing.B) {

	// Create the objects
	o := make([]*largeobj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &largeobj{id: i}
	}

	// Setup the map
	m := make(map[int]*largeobj, b.N)

	// Add orders to the map
	for i := 0; i < b.N; i++ {
		m[i] = o[i]
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m[i]
	}
}

func Benchmark_2_LargeObject_MapTraversalAndAccessWithRange(b *testing.B) {

	// Create the objects
	o := make([]*largeobj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &largeobj{id: i}
	}

	// Setup the map
	m := make(map[int]*largeobj, b.N)

	// Add orders to the map
	for i := 0; i < b.N; i++ {
		m[i] = o[i]
	}

	b.ResetTimer()

	for i, j := range m {
		_, _ = i, j
	}
}

func Benchmark_2_LargeObject_LinkedListTraversalAndAccess(b *testing.B) {

	// Create the objects
	o := make([]*largeobj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &largeobj{id: i}
	}

	// Setup the linked list params
	var llhead, lltail *largeobj

	// Add orders to the linked list
	for i := 0; i < b.N; i++ {
		if llhead == nil {
			llhead = o[i]
			lltail = o[i]
		} else {
			lltail.next = o[i]
			o[i].prev = lltail
			lltail = o[i]
		}
	}

	b.ResetTimer()

	for ord := llhead; ord != nil; ord = ord.next {
		_ = ord
	}
}
