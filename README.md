# listbench
Benchmarking slice, map, and double linked list in Go.

## Usage

```bash
go test -bench=. -benchtime 10ms
```

Without the benchtime flag, it will run fine but will take a long time. Setting benchtime to 10ms significantly shortens benchmarking time and still yields fairly stable results.

## Detail

This benchmarks three types of activities for each list type:
1. Adding an item
2. Traversing & accessing
3. Traversing & updating

Each activity is also performed on two types of objects:
1. A small object
2. A large object

## Sample Output

```bash
Computer:listbench user$ go test -bench=. -benchtime 10ms
goos: darwin
goarch: amd64
pkg: listbench
Benchmark_1_LargeObject_SliceAdd-8                           	 1000000	        21.3 ns/op
Benchmark_1_LargeObject_SlicePreSizedAdd-8                   	 1000000	        10.9 ns/op
Benchmark_1_LargeObject_MapAdd-8                             	  200000	       145 ns/op
Benchmark_1_LargeObject_MapPreSizedAdd-8                     	  500000	        76.8 ns/op
Benchmark_1_LargeObject_LinkedListAdd-8                      	 1000000	        10.2 ns/op
Benchmark_1_SmallObject_SliceAdd-8                           	 1000000	        45.6 ns/op
Benchmark_1_SmallObject_SlicePreSizedAdd-8                   	 1000000	        72.9 ns/op
Benchmark_1_SmallObject_MapAdd-8                             	  200000	        91.4 ns/op
Benchmark_1_SmallObject_MapPreSizedAdd-8                     	  500000	        80.0 ns/op
Benchmark_1_SmallObject_LinkedListAdd-8                      	10000000	         2.77 ns/op
Benchmark_2_LargeObject_SliceTraversalAndAccess-8            	50000000	         0.26 ns/op
Benchmark_2_LargeObject_SliceTraversalAndAccessWithRange-8   	50000000	         0.23 ns/op
Benchmark_2_LargeObject_MapTraversalAndAccess-8              	 1000000	        52.1 ns/op
Benchmark_2_LargeObject_MapTraversalAndAccessWithRange-8     	 1000000	        11.5 ns/op
Benchmark_2_LargeObject_LinkedListTraversalAndAccess-8       	 1000000	        32.9 ns/op
Benchmark_2_SmallObject_SliceTraversalAndAccess-8            	30000000	         0.46 ns/op
Benchmark_2_SmallObject_SliceTraversalAndAccessWithRange-8   	50000000	         0.23 ns/op
Benchmark_2_SmallObject_MapTraversalAndAccess-8              	 1000000	        50.9 ns/op
Benchmark_2_SmallObject_MapTraversalAndAccessWithRange-8     	 1000000	        11.1 ns/op
Benchmark_2_SmallObject_LinkedListTraversalAndAccess-8       	10000000	         2.33 ns/op
Benchmark_3_LargeObject_SliceTraversalAndUpdate-8            	 2000000	         8.92 ns/op
Benchmark_3_LargeObject_MapTraversalAndUpdate-8              	  500000	        46.5 ns/op
Benchmark_3_LargeObject_LinkedListTraversalAndUpdate-8       	 1000000	        28.0 ns/op
Benchmark_3_SmallObject_SliceTraversalAndUpdate-8            	10000000	         2.49 ns/op
Benchmark_3_SmallObject_MapTraversalAndUpdate-8              	 1000000	        42.2 ns/op
Benchmark_3_SmallObject_LinkedListTraversalAndUpdate-8       	10000000	         2.99 ns/op
PASS
ok  	listbench	48.367s
```

## Tags

#go #golang #map #slice #linkedlist
