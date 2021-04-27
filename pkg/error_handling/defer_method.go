
package myerrorhandling
 

import (  
    "fmt"
)

type person_1005 struct {  
    firstName string
    lastName string
}

func (p person_1005) fullName() {  
    fmt.Printf("%s %s",p.firstName,p.lastName)
}

func Defer_method() {  

fmt.Println("\n+++ Defer_method +++++")

    p := person_1005 {
        firstName: "John",
        lastName: "Smith",
    }
    defer p.fullName()
    fmt.Printf("Welcome ")  
}