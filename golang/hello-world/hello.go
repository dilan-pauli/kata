package main

import "fmt"

func main() {
	fmt.Println("hello world")

	fmt.Println("go" + "lang")
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// https://gobyexample.com/variables
	var a = "inital"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)
}
