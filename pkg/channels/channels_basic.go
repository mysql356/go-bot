package mychannels

import ("fmt";"time")


func Channel_nil() {  
    var a chan int
    if a == nil {
        fmt.Println("channel a is nil, going to define it")
        a = make(chan int)
        fmt.Printf("Type of a is %T", a)
    }
}

func Channel_basic() {  
    messages := make(chan string)
    go func() { messages <- "ping.............." }() //goroutine invoke
    msg := <-messages //o/p of goroutine
    fmt.Println(msg)
}

/*
data := <- a // read from channel a  
a <- data // write to channel a  

*/

func channel_hello(done chan bool) { 
	fmt.Println()
    fmt.Println("Hello goroutine world")
    done <- true //return true
}
func Channel_in_goroutine() {  
    done := make(chan bool)
    go channel_hello(done) 
	<-done // wait for response of goroutine or using time.Sleep(1 * time.Second)
    //fmt.Println(<-done) //true 
    fmt.Println("channel_in_goroutine")
}
////////////////
func channel_hello1(done chan bool) {  
    fmt.Println("hello go routine is going to sleep")
    time.Sleep(4 * time.Second)
    fmt.Println("hello go routine awake and going to write to done")
    done <- true
}
func Channel_in_goroutine1() {  
    fmt.Println("\n+++Advance ex of channels+++++")	
    done := make(chan bool)
    fmt.Println("Main going to call hello go goroutine")
    go channel_hello1(done)
    <-done
    fmt.Println("Main received data")
}


