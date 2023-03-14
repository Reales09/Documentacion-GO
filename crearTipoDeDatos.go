package main

import "fmt"

var d int

type dinero int

var e dinero

func main() {
	d = 42

	fmt.Println(d)
	fmt.Printf("%T\n", d)

	e = 1000
	fmt.Println(d)
	fmt.Printf("%T\n", e)

	d = e
	fmt.Println(a)
	fmt.Printf("%T\n", a)

}
