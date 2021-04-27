package myfilehandling

import (
	"fmt"
	"os"
)

func FileApend() { 
fmt.Println("\n+++ File Write : Apend +++++")  
    f, err := os.OpenFile("/Users/manoj/workspace/dev/collection/1-golangbot1/pkg/file_handling/test_w_apend.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println(err)
        return
    }
    newLine := "File handling is easy."
    _, err = fmt.Fprintln(f, newLine)
    if err != nil {
        fmt.Println(err)
                f.Close()
        return
    }
    err = f.Close()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("file appended successfully")
}