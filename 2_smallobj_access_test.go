package listbench

import (
	"testing"
)

func Benchmark_2_SmallObject_SliceTraversalAndAccess(b *testing.B) {

	// Create the objects
	o := make([]*obj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &obj{id: i}
	}

	// Setup the slice
	s := make([]*obj, b.N)

	// Add orders to the slice
	for i := 0; i < b.N; i++ {
		s = append(s, o[i])
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s[i]
	}
}

func Benchmark_2_SmallObject_SliceTraversalAndAccessWithRange(b *testing.B) {

	// Create the objects
	o := make([]*obj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &obj{id: i}
	}

	// Setup the slice
	s := make([]*obj, 0)

	// Add orders to the slice
	for i := 0; i < b.N; i++ {
		s = append(s, o[i])
	}

	b.ResetTimer()

	for i, j := range s {
		_, _ = i, j
	}
}

func Benchmark_2_SmallObject_MapTraversalAndAccess(b *testing.B) {

	// Create the objects
	o := make([]*obj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &obj{id: i}
	}

	// Setup the map
	m := make(map[int]*obj, b.N)

	// Add orders to the map
	for i := 0; i < b.N; i++ {
		m[i] = o[i]
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m[i]
	}
}

func Benchmark_2_SmallObject_MapTraversalAndAccessWithRange(b *testing.B) {

	// Create the objects
	o := make([]*obj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &obj{id: i}
	}

	// Setup the map
	m := make(map[int]*obj, b.N)

	// Add orders to the map
	for i := 0; i < b.N; i++ {
		m[i] = o[i]
	}

	b.ResetTimer()

	for i, j := range m {
		_, _ = i, j
	}
}

func Benchmark_2_SmallObject_LinkedListTraversalAndAccess(b *testing.B) {

	// Create the objects
	o := make([]*obj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &obj{id: i}
	}

	// Setup the linked list params
	var llhead, lltail *obj

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
