package main

import "fmt"

func variables() {
	var a = "testVariable a"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	// Init a variable without a value, sets null value (for int 0)
	var e int
	fmt.Println(e)

	f := "shorthand initialization"
	fmt.Println(f)
}
