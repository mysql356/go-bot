/////////////////
/*
What is select?
The select statement is used to choose from multiple send/receive channel operations. 
The select statement blocks until one of the send/receive operation is ready. 
If multiple operations are ready, one of them is chosen at random. 
The syntax is similar to switch except that each of the case statement will be a channel operation. 
Lets dive right into some code for better understanding. 
*/

package mychannels

import (  
    "fmt"
    "time"
)


func server1(ch chan string) {  
    time.Sleep(4 * time.Second) 
    ch <- "from server1"
}
func server2(ch chan string) {  
    time.Sleep(2 * time.Second)
    ch <- "from server2"

}
func Channel_select() {  
	fmt.Println("\n+++ select channels +++++")	   
    output1 := make(chan string)
    output2 := make(chan string)
    go server1(output1)
    go server2(output2)
    select {
    case s1 := <-output1:
        fmt.Println(s1)
    case s2 := <-output2:
        fmt.Println(s2)
    }
}

/////////////////
func channel_select_process(ch chan string) {  
    time.Sleep(5500 * time.Millisecond)
    ch <- "process successful"
}

func Channel_select1() {  
    ch := make(chan string)
    go channel_select_process(ch)
    for {
        time.Sleep(1000 * time.Millisecond)
        select {
        case v := <-ch:
            fmt.Println("received value: ", v)
            return
        default:
            fmt.Println("no value received")
        }
    }

}

//Deadlock and default case
/*
The select statement will block forever since no other Goroutine is writing to this channel and hence will result in deadlock
*/

func hello_select_deadlock(done chan string) {  
    fmt.Println("hello go routine is going to sleep")
    done <- "done"
}

func Channel_select_deadlock() {  
    ch := make(chan string)
	go hello_select_deadlock(ch) //fatal error: all goroutines are asleep - deadlock!

    select {
    case <-ch:
    }
}

func Channel_select_deadlock_handling() {  
    ch := make(chan string)
    select {
    case <-ch:
    default:
        fmt.Println("default case executed")
    }
}

func Channel_nil_select() {  
    var ch chan string
    select {
    case v := <-ch:
        fmt.Println("received value", v)
    default:
        fmt.Println("default case executed")

    }
}

/*
func channel_select_nil() {
	select {}
}
fatal error: all goroutines are asleep - deadlock! 
no cases
*/

