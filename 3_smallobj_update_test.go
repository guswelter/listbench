package listbench

import (
	"testing"
)

func Benchmark_3_SmallObject_SliceTraversalAndUpdate(b *testing.B) {

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

	for _, j := range s {
		j.id = j.id + 1
	}
}

func Benchmark_3_SmallObject_MapTraversalAndUpdate(b *testing.B) {

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

	for _, j := range m {
		j.id = j.id + 1
	}
}

func Benchmark_3_SmallObject_LinkedListTraversalAndUpdate(b *testing.B) {

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
		ord.id = ord.id + 1
	}
}
