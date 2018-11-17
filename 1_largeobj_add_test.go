package listbench

import (
	"testing"
)

func Benchmark_1_LargeObject_SliceAdd(b *testing.B) {

	// Create the orders
	o := make([]*largeobj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &largeobj{id: i}
	}

	// Setup the slice
	s := make([]*largeobj, 0)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = append(s, o[i])
	}
}

func Benchmark_1_LargeObject_SlicePreSizedAdd(b *testing.B) {

	// Create the orders
	o := make([]*largeobj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &largeobj{id: i}
	}

	// Setup the slice (pre-sized)
	s := make([]*largeobj, b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = append(s, o[i])
	}
}

func Benchmark_1_LargeObject_MapAdd(b *testing.B) {

	// Create the orders
	o := make([]*largeobj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &largeobj{id: i}
	}

	// Setup the map
	m := make(map[int]*largeobj, 0)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m[i] = o[i]
	}
}

func Benchmark_1_LargeObject_MapPreSizedAdd(b *testing.B) {

	// Create the orders
	o := make([]*largeobj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &largeobj{id: i}
	}

	// Setup the map
	m := make(map[int]*largeobj, b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m[i] = o[i]
	}
}

func Benchmark_1_LargeObject_LinkedListAdd(b *testing.B) {

	// Create the orders
	o := make([]*largeobj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &largeobj{id: i}
	}

	// Setup the linked list params
	var llhead, lltail *largeobj

	b.ResetTimer()

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
}
