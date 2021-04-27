package myfilehandling

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func FileReadingChars() {  
	fmt.Println("\n+++ File reading  : Chars +++++")
    fptr := flag.String("fpath", "/Users/manoj/workspace/dev/collection/1-golangbot1/pkg/file_handling/test_read_chars.txt", "file path to read from")
    flag.Parse()
	fmt.Println("value of fpath is", *fptr) //file path

    f, err := os.Open(*fptr)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = f.Close(); err != nil {
            log.Fatal(err)
        }
    }()
 ////////////////////////////
    r := bufio.NewReader(f)
    b := make([]byte, 1)
    for {
        _, err := r.Read(b)
        if err != nil {
            fmt.Println("Error reading file:", err)
            break
        }
        fmt.Print(string(b), " - ")
    }
}
func FileReadingLine() {  
	fmt.Println("\n+++ File reading  : Line +++++")
    fptr := flag.String("fpathline", "/Users/manoj/workspace/dev/collection/1-golangbot1/pkg/file_handling/test_read.txt", "file path to read from")
    flag.Parse()

	fmt.Println("value of fpathline is", *fptr) //file path

    f, err := os.Open(*fptr)
    if err != nil {
        log.Fatal(err)
    }
    defer func() {
        if err = f.Close(); err != nil {
            log.Fatal(err)
        }
    }()
////////////////////////
s := bufio.NewScanner(f)
    for s.Scan() {
        fmt.Println(s.Text())
    }
    err = s.Err()
    if err != nil {
        log.Fatal(err)
    }
}