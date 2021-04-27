package myerrorhandling
import "fmt"
//import "runtime/debug"

 

func r2() {  
    if r := recover(); r != nil {
        fmt.Println("Recovered", r)
        //debug.PrintStack() - There is a way to print the stack trace using the PrintStack function of the Debug package
    }
}

func a2() {  
    defer r2()
    n := []int{5, 7, 4}
    //fmt.Println(n[2]) //normally
	fmt.Println(n[3])   //panic
    fmt.Println("normally returned from a")
}

func PanicRuntimeRecover() {  
	fmt.Println("\n+++ PanicRuntimeRecover +++++")
    a2()
    fmt.Println("normally returned from main")
}

 /*
-normally (o/p)
+++ PanicRuntimeRecover +++++
4
normally returned from a
normally returned from main

-Panic
+++ PanicRuntimeRecover +++++
Recovered runtime error: index out of range
normally returned from main


*/