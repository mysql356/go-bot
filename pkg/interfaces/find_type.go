package myinterface

 

import (  
    "fmt"
)

func findType(i interface{}) {  
    switch i.(type) {
    case string:
        fmt.Printf("I am a string and my value is %s\n", i.(string))
    case int:
        fmt.Printf("I am an int and my value is %d\n", i.(int))
    default:
        fmt.Printf("Unknown type\n")
    }
}
func FindType() {  

 fmt.Println("\n+++ FindType +++++")	
    findType("Manoj")
    findType(77)
    findType(89.98)
}
