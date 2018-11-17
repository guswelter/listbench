package listbench

import "time"

type obj struct {
	id   int
	next *obj
	prev *obj
}

type largeobj struct {
	id int

	A uint64
	B float64
	C uint64
	D uint64
	E uint8
	F float64
	G float64
	H float64
	I string
	J uint64
	K uint64
	L uint64
	M time.Time
	N time.Time
	O uint64
	P uint64
	Q uint64
	R float64
	S float64
	T float64
	U float64
	V float64
	W float64
	X bool
	Y int
	Z *largeobj

	next *largeobj
	prev *largeobj
}
