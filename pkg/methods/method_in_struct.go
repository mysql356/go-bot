/*
To define a method on a type, the definition of the receiver type of the method and the definition of the method should be in the same package. 
So far, all the structs and the methods on structs we defined where all located in the same main package and hence they worked.

This method is not linked with structure
cannot define new methods on non-local type int
It cannot be compiled 

package main


func (a int) add(b int) {  
}

func method_in_struct() {

}
*/
///////////////////////////////////////////////

//The way to get this working is to create a type alias for the built-in type int and then create a method with this type alias as the receiver.


package mymethod

import "fmt"

type myInt int

func (a myInt) add(b myInt) myInt {  
    return a + b
}

func Method_in_struct() {  
    num1 := myInt(5)
    num2 := myInt(10)
    sum := num1.add(num2)
    fmt.Println("Sum is", sum)
}
