package main

import "fmt"

/*This is a comment in Golang */
// this is also a comment

func main() {
//	var x float64 = 20.0
//	y := 42
//	var a, b, c = 3, 4, "foo"
//   fmt.Println(a)
//	fmt.Println(b)
//	fmt.Println(c)
//	fmt.Printf("a is of the type %T\n", a)
//	fmt.Printf("b is of the type %T\n", b)
//	fmt.Printf("c is of the type %T\n", c)
	const LENGTH int = 10
	const WIDTH int = 5
	var area int

	area = LENGTH * WIDTH
	fmt.Printf("value of are : %d", area)
}
