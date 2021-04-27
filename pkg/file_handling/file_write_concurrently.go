package myfilehandling

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func produce(data chan int, wg *sync.WaitGroup) {  
    n := rand.Intn(999)
    data <- n
    wg.Done()
}

func consume(data chan int, done chan bool) {  
    f, err := os.Create("/Users/manoj/workspace/dev/collection/1-golangbot1/pkg/file_handling/test_w_concurrent.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    for d := range data {
        //_, err = fmt.Fprintln(f, d)
		_, err = fmt.Fprintln(f, fmt.Sprintf("%d \n",d))
        if err != nil {
            fmt.Println(err)
            f.Close()
            done <- false
            return
        }
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        done <- false
        return
    }
    done <- true
}

func FileConcurrently() { 
fmt.Println("\n+++ File Write : File Concurrently +++++")   
    data := make(chan int)
    done := make(chan bool)
    wg := sync.WaitGroup{}
    for i := 0; i < 100; i++ {
        wg.Add(1)
        go produce(data, &wg)
    }
    go consume(data, done)
    go func() {
        wg.Wait()
        close(data)
    }()
    d := <-done
    if d == true {
        fmt.Println("File written successfully")
    } else {
        fmt.Println("File writing failed")
    }
}