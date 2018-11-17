package listbench

import (
	"testing"
)

func Benchmark_1_SmallObject_SliceAdd(b *testing.B) {

	// Create the objects
	o := make([]*obj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &obj{id: i}
	}

	// Setup the slice
	s := make([]*obj, 0)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = append(s, o[i])
	}
}

func Benchmark_1_SmallObject_SlicePreSizedAdd(b *testing.B) {

	// Create the objects
	o := make([]*obj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &obj{id: i}
	}

	// Setup the slice (pre-sized)
	s := make([]*obj, b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s = append(s, o[i])
	}
}

func Benchmark_1_SmallObject_MapAdd(b *testing.B) {

	// Create the objects
	o := make([]*obj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &obj{id: i}
	}

	// Setup the map
	m := make(map[int]*obj, 0)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m[i] = o[i]
	}
}

func Benchmark_1_SmallObject_MapPreSizedAdd(b *testing.B) {

	// Create the objects
	o := make([]*obj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &obj{id: i}
	}

	// Setup the map
	m := make(map[int]*obj, b.N)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m[i] = o[i]
	}
}

func Benchmark_1_SmallObject_LinkedListAdd(b *testing.B) {

	// Create the objects
	o := make([]*obj, b.N)
	for i := 0; i < b.N; i++ {
		o[i] = &obj{id: i}
	}

	// Setup the linked list params
	var llhead, lltail *obj

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
