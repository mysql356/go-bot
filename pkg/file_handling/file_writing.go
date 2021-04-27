package myfilehandling

import (
	"fmt"
	"os"
)

func FileWritingString() {  
fmt.Println("\n+++ File Write : String +++++")

    f, err := os.Create("/Users/manoj/workspace/dev/collection/1-golangbot1/pkg/file_handling/test_w.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    l, err := f.WriteString("Hello World")
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(l, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func FileWritingBytes() {  
fmt.Println("\n+++ File Write : Bytes +++++") 
    f, err := os.Create("/Users/manoj/workspace/dev/collection/1-golangbot1/pkg/file_handling/test_w_byte.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
    n2, err := f.Write(d2)
    if err != nil {
        fmt.Println(err)
        f.Close()
        return
    }
    fmt.Println(n2, "bytes written successfully")
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
}

func FileWritingLinebyline() {  
fmt.Println("\n+++ File Write : Linebyline +++++") 
    f, err := os.Create("/Users/manoj/workspace/dev/collection/1-golangbot1/pkg/file_handling/test_w_line.txt")
    if err != nil {
        fmt.Println(err)
                f.Close()
        return
    }
    d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}

    for _, v := range d {
        fmt.Fprintln(f, v)
        if err != nil {
            fmt.Println(err)
            return
        }
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("file written successfully")
}