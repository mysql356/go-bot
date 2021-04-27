package myfirstclassfunction
import "fmt"

//User defined function types=======================
type add func(a int, b int) int

func UserDefinedFunction() {  
fmt.Println("\n+++ UserDefinedFunction +++++")
    var a add = func(a int, b int) int {
        return a + b
    }
    s := a(5, 6)
    fmt.Println("Sum", s)
}

//User defined function types 1=======================
func UserDefinedFunction1() {  
fmt.Println("\n+++ UserDefinedFunction1 +++++")
    a1 := func(a int, b int) int {
        return a + b
    }
    s := a1(5, 6)
    fmt.Println("Sum", s)
}

/////////////////
/*
Higher-order functions
The definition of Higher-order function from wiki is a function which does at least one of the following
-takes one or more functions as arguments
-returns a function as its result
*/

//Passing functions as arguments to other functions
func simple_ho1(a func(a, b int) int) {  
    fmt.Println(a(60, 7))
}

func HigherOrderFunc1() {  
fmt.Println("\n+++ HigherOrderFunc1 +++++")
    f := func(a, b int) int {
        return a * b
    }
    simple_ho1(f)
}
//Returning functions from other functions
func simple_ho2() func(a, b int) int {  
    f := func(a, b int) int {
        return a + b
    }
    return f
}

func HigherOrderFunc2() { 
fmt.Println("\n+++ HigherOrderFunc2 +++++") 
    s := simple_ho2()
    fmt.Println(s(60, 7))
}
////////////////////////////////////////////////
/*
Closures
Closures are a special case of anonymous functions. 
Closures are anonymous functions which access the variables defined outside the body of the function. 
*/

func main() {  
    a := 5
    func() {
        fmt.Println("a =", a)
    }()
}

