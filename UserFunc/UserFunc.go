package main

import (
	"fmt"
	"runtime"
)

// define function named total that accept int values and returns int
func total(x int, y int) int {
	return x + y
}

func go_version() string {
	var go_version string
	go_version = runtime.Version()
	return go_version
}

func go_version_walk() {
	var go_version string
	go_version = runtime.Version()
	fmt.Print("Go version (index walk)\n")
	for index, s := range go_version {
		fmt.Printf("The index number of %c is %d\n", s, index)
	}
}

func main() {
	// call our function and store result in the answer variable
	answer := total(10, 20)
	fmt.Println("10 + 20  = ", answer)
	fmt.Println("100 + 500 = ", total(100, 500))
	fmt.Printf("Go version: %s\n", runtime.Version())
	fmt.Printf("Go version (via function): %s\n", go_version())
	go_version_walk()
}
