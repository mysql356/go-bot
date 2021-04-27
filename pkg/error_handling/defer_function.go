/*What is Defer?
Defer statement is used to execute a function call just before the function where the defer statement is present returns. 
The definition might seem complex but it's pretty simple to understand by means of an example.
*/

package myerrorhandling
 

import (  
    "fmt"
)

func finished() {  
    fmt.Println("Finished finding largest")
}

func largest(nums []int) {  
    defer finished()    
    fmt.Println("Started finding largest")
    max := nums[0]
    for _, v := range nums {
        if v > max {
            max = v
        }
    }
    fmt.Println("Largest number in", nums, "is", max)
}

func Defer_function() {  
fmt.Println("\n+++ Defer_function +++++")
    nums := []int{78, 109, 2, 563, 300}
    largest(nums)
}

////////////////////
/*
Stack of defers

When a function has multiple defer calls, they are added on to a stack and executed in Last In First Out (LIFO) order.

*/

func Defer_reverse_str() {  

fmt.Println("\n\n+++ Defer_reverse_str +++++")
    name := "Manoj"
    fmt.Printf("Original String: %s\n", string(name))
    fmt.Printf("Reversed String: ")
    for _, v := range []rune(name) {
        defer fmt.Printf("%c", v)
    }
}


