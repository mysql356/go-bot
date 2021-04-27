package mychannels

import (  
    "fmt"
)

/////////////
func calcSquares(number int, squareop chan int) {  
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    squareop <- sum
}

func calcCubes(number int, cubeop chan int) {  
    sum := 0 
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }
    cubeop <- sum
} 

func Channel_advance_calc() { 
    fmt.Println("\n+++Advance ex of channels+++++")	 
    number := 5
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares(number, sqrch)
    go calcCubes(number, cubech)
    squares, cubes := <-sqrch, <-cubech
    fmt.Println("Final output", squares + cubes)
}

/*
Deadlock
One important factor to consider while using channels is deadlock. If a Goroutine is sending data on a channel, 
then it is expected that some other Goroutine should be receiving the data. If this does not happen, then the program will panic at runtime with Deadlock.

Similarly if a Goroutine is waiting to receive data from a channel, 
then some other Goroutine is expected to write data on that channel, else the program will panic.

func channel_deadlock() {  
    ch := make(chan int)
    ch <- 5
}
*/

/*
Unidirectional channels

All the channels we discussed so far are bidirectional channels, that is data can be both sent and received on them. 
It is also possible to create unidirectional channels, that is channels that only send or receive data.
*/
func sendData_unidirectional(sendch chan<- int) {  
    sendch <- 10
}

func Channel_advance_unidirectional() {  

	fmt.Println("\n+++Advance ex of unidirectional+++++")	 
    //sendch := make(chan<- int) //invalid operation: <-sendch (receive from send-only type chan<- int)
	
	sendch := make(chan int)
    go sendData_unidirectional(sendch)
    fmt.Println(<-sendch)
		
}

/*pings-pongs channel
<-chan : Receive from channel
chan<- : send to channel
*/

func ping(pings chan<- string, msg string) {
    pings <- msg
}

// The `pong` function accepts one channel for receives
// (`pings`) and a second for sends (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}
func Channel_pings_pongs() {
	fmt.Println("\n+++Advance ex of pings-pongs+++++")	
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}

////////////////////////
func producer(chnl chan int) {  
    for i := 0; i < 5; i++ {
        chnl <- i
    }
    close(chnl)
}
func Channel_close_by_for() { 
 	fmt.Println("\n+++Advance ex of channel close - using for loop +++++")	
    ch := make(chan int)
    go producer(ch)
    for {
        v, ok := <-ch
        if ok == false {
            break
        }
        fmt.Println("Received ", v, ok)
    }
}
func Channel_close_by_range() {  
 	fmt.Println("\n+++Advance ex of channel close - using range +++++")	
    ch := make(chan int)
    go producer(ch)
    for v := range ch {
        fmt.Println("Received ",v)
    }
}

////////////////////
func digits_range(number int, dchnl chan int) {  
    for number != 0 {
        digit := number % 10
        dchnl <- digit
        number /= 10
    }
    close(dchnl)
}
func calcSquares_range(number int, squareop chan int) {  
    sum := 0
    dch := make(chan int)
    go digits_range(number, dch)
    for digit := range dch {
        sum += digit * digit
    }
    squareop <- sum
}

func calcCubes_range(number int, cubeop chan int) {  
    sum := 0
    dch := make(chan int)
    go digits_range(number, dch)
    for digit := range dch {
        sum += digit * digit * digit
    }
    cubeop <- sum
}

func Channel_advance_calc_using_range() {  
    number := 5
    sqrch := make(chan int)
    cubech := make(chan int)
    go calcSquares_range(number, sqrch)
    go calcCubes_range(number, cubech)
    squares, cubes := <-sqrch, <-cubech
    fmt.Println("Final output", squares+cubes)
}

