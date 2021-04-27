package mychannels
import ("fmt";"time";"sync")


func Channel_buffered_basic() {  
	fmt.Println("\n+++ buffered channels +++++")	 
    ch := make(chan string, 2)
    ch <- "naveen"
    ch <- "paul"
    fmt.Println(<- ch)
    fmt.Println(<- ch)
}

//////////////////////////
func write(ch chan int) {  
    for i := 0; i < 5; i++ {
        ch <- i
        fmt.Println("wrote-----", i, "to ch")
    }
    close(ch)
}
func Channel_buffered_close() {  
    ch := make(chan int, 2)
    go write(ch)
    time.Sleep(2 * time.Second)
    for v := range ch {
        fmt.Println("read value", v,"from ch")
        time.Sleep(2 * time.Second)

    }
}
////////////Deadlock
/*
When the control reaches the third write (ch <- "steve"), the write is blocked since the channel has exceeded its capacity. 
Now some Goroutine must read from the channel in order for the write to proceed, 
but in this case there is no concurrent routine reading from this channel. 
Hence there will be a deadlock and the program will panic at run time with the following message,
*/
func Channel_deadlock() {   
	fmt.Println("\n+++ buffered channels : deadlock+++++")	 
    ch := make(chan string, 2)
    ch <- "naveen"
    ch <- "paul"
    //ch <- "steve"	//fatal error: all goroutines are asleep - deadlock!
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
///////////////////
func Channel_length_capacity() {
	fmt.Println("\n+++ buffered channels : length-capacity+++++")	   
    ch := make(chan string, 3)
    ch <- "naveen"
    ch <- "paul"
    fmt.Println("capacity is", cap(ch))
    fmt.Println("length is", len(ch))
    fmt.Println("read value", <-ch)
    fmt.Println("new length is", len(ch))
}
////////////////////////
func channel_process(i int, wg *sync.WaitGroup) {  
    fmt.Println("started Goroutine ", i)
    time.Sleep(2 * time.Second)
    fmt.Printf("Goroutine %d ended\n", i)
    wg.Done()
}

func Channel_waitGroup() {  
	fmt.Println("\n+++ buffered channels : channel_waitGroup+++++")
    no := 3
    var wg sync.WaitGroup
    for i := 0; i < no; i++ {
        wg.Add(1)
        go channel_process(i, &wg)
    }
    wg.Wait()
    fmt.Println("All go routines finished executing")
}