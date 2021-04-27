package myinterface

import (  
    "fmt"
)

func describe_empty(i interface{}) {  
    fmt.Printf("Type = %T, value = %v\n", i, i)
}

func Interface_empty() {  
 fmt.Println("\n+++ Interface_empty +++++")
    s := "Hello World"
    describe_empty(s)
    i := 55
    describe_empty(i)
    strt := struct {
        name string
    }{
        name: "Manoj k",
    }
    describe_empty(strt)
}