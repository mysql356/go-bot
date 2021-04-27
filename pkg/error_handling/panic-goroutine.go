package myerrorhandling

import (  
    "fmt"
    "time"
)

func recovery() {  
    if r := recover(); r != nil {
        fmt.Println("recovered:", r)
    }
}

func a() {  
    defer recovery()
    fmt.Println("Inside A")	
    //go b() //goroutine - not recovered	
    b()  //function - recovered
	/*because the recovery function is present in a different gouroutine and the panic is happening in a different goroutine.*/
    time.Sleep(1 * time.Second)
}

func b() {  
    fmt.Println("Inside B")
    panic("oh! B panicked")
}

func PanicGoroutine() {  
    a()
    fmt.Println("normally returned from main")
}
