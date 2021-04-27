package myerrorhandling
import "fmt"
/*
type Error interface {  
    error
    // RuntimeError is a no-op function but
    // serves to distinguish types that are run time
    // errors from ordinary errors: a type is a
    // run time error if it has a RuntimeError method.
    RuntimeError()
}
*/

func a1() {  
    n := []int{5, 7, 4}
    fmt.Println(n[3])
    fmt.Println("normally returned from a")
}
func PanicRuntime() {  
	fmt.Println("\n+++ PanicRuntime +++++")
    a1()
    fmt.Println("normally returned from main")
}