package pkg

import (
	"fmt"
)

//3 - Just like our own struct types, it is possible to define our own function types.
type add func(a int, b int) int

//4 - Higher-order functions
func simple(a func(a, b int) int) {
	fmt.Println(a(60, 7))
}

func simple1() func(a, b int) int {
	f := func(a, b int) int {
		return a + b
	}
	return f
}

//6
func appendStr() func(string) string {
	t := "Hello"
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}

func Firstclassfunctions() {
	//1
	a := func() {
		fmt.Println("hello world first class function")
	}
	a()
	fmt.Printf("%T\n", a)

	//2
	func() {
		fmt.Println("anonymous function")
	}()

	func(n string) {
		fmt.Println("Welcome", n)
	}("Gophers")

	//3
	var a1 add = func(a int, b int) int {
		return a + b
	}
	s := a1(5, 6)
	fmt.Println("Sum", s)

	//4 - Higher-order functions
	//functions as arguments
	f := func(a, b int) int {
		return a + b
	}
	simple(f)

	//Returning functions
	s1 := simple1()
	fmt.Println(s1(60, 7))

	//5 - Closures : a anonymous functions which can access the variables defined outside of function
	a5 := 5
	func() {
		fmt.Println("a5 =", a5)
	}()

	//6
	a6 := appendStr()
	b6 := appendStr()
	fmt.Println(a6("World"))
	fmt.Println(b6("Everyone"))

}
