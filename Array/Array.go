package main

import "fmt"

func main() {
	// define array named domains as string type
	var domains [2]string
	fmt.Println("Array length : ", len(domains))
	// add value and print it
	fmt.Println("current values for array (non set):", domains)
	// get array length

	domains[0] = "cyberciti.biz"
	domains[1] = "nixcraft.com"
	fmt.Println("Set Array value : ", domains)

	// use for loop to print our array
	for i := 0; i < len(domains); i++ {
		fmt.Println("Get value for element ", i, " is ", domains[i])
	}
}
