/*What are first class functions?
A language which supports first class functions allows functions to be assigned to variables, 
passed as arguments to other functions and returned from other functions. 
Go has support for first class functions.

anonymous functions - without function name 
*/
package myfirstclassfunction

import "fmt"
    
func Basic() {  
fmt.Println("\n+++ Basic +++++")
    a := func() {
        fmt.Println("Hello world : anonymous functions - 1 ")
    }
    a()
    fmt.Printf("%T", a)
}
////////////////
func Basic1() {  
fmt.Println("\n+++ Basic1 +++++")
    func() {
        fmt.Println("Hello world : anonymous functions - 2 ")
    }()
}
////////////////
func Basic2() {  
fmt.Println("\n+++ Basic2 +++++")
    func(n string) {
        fmt.Println("Welcome", n)
    }("Gophers")
}
////////////////////////////////////////////////
/*
Closures
Closures are a special case of anonymous functions. 
Closures are anonymous functions which access the variables defined outside the body of the function. 
*/

func Closures() {  
fmt.Println("\n+++ Closures +++++")
    a := 5
    func() {
        fmt.Println("a =", a)
    }()
}

//===========================
func appendStr() func(string) string {  
    t := "Hello"
    c := func(b string) string {
        t = t + " " + b
        return t
    }
    return c
}

func ClosuresEx() {  
fmt.Println("\n+++ ClosuresEx +++++")
    a := appendStr()
    b := appendStr()
    fmt.Println(a("World"))
    fmt.Println(b("Everyone"))

    fmt.Println(a("Gopher"))
    fmt.Println(b("!"))
}
