package main

import (
	"fmt"

	. "iter"
)

func main() {

	println("\nFibonnaci:")
	for x := range Fib(100) {
		fmt.Printf("i = %6v\n", x)
	}

	println("\n - - - ")
	println("Filter:")
	fibs := Fib(100)
	for i := range Filter(fibs) {
		fmt.Printf("i = %6v\n", i)
	}

	println("\n - - - ")
	println("Map, Filter:")
	fibs = Fib(100)
	for i := range Map(Filter(fibs)) {
		fmt.Printf("i = %6v\n", i)
	}

	println("\n - - - ")
	println("Filter, Map, Reduce [1]")
	fibs = Fib(100)
	filtered := Filter(fibs)
	mapped := Map(filtered)
	reduced := Reduce(mapped)
	fmt.Printf("i = %v\n", reduced)

	println("\n - - - ")
	println("Filter, Map, Reduce [2]")
	result := Reduce(Map(Filter(Fib(100))))
	fmt.Printf("i = %v\n", result)

	println("\n - - - ")
	println("Reduce")
	fmt.Printf("i = %v\n", reduced)

	println("\n - - - ")
	println("Take")
	fibs = Fib(10_000_000)
	j := 0
	for i := range Take(fibs, 10) {
		j++
		fmt.Printf("%v : i = %v\n", j, i)
	}

	println("\n - - - ")
	println("Enumerate")
	fibs = Fib(100)
	for item := range Enumerate(fibs) {
		fmt.Printf("index :%6v,  value = %6v\n", item.Index, item.Value)
	}

	println("\n - - - ")
	println("Range")
	for item := range Range(10) {
		squared := item * item
		fmt.Printf("index = %6v, squared = %6v  \n", item, squared)
	}

}
