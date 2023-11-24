package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter your number: ")
	fmt.Scanf("%d", &n)
	fmt.Println(n)
	if n >= 0 {
		fmt.Println("Number is positive")
	} else {
		fmt.Println("Number is negative")
	}
}
