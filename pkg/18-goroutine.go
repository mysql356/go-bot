package pkg

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello world goroutine")
}
func Goroutine_basic() {
	go hello()
	time.Sleep(1 * time.Second)
	fmt.Println("goroutine_basic")
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
	fmt.Println("\nPress Enter key...")
}

func Goroutine_basic_scan() {
	f("F - 1")
	go f("GR - 2")

	//inline function
	go func() {
		fmt.Println("GR - 3")
	}() //func calling

	go func(msg string) {
		fmt.Println(msg)
	}("GR - 4") //func calling and params passing

	a := func(msg string) {
		fmt.Println(msg)
	}
	a("GR - 5") //func calling and params passing

	go numbers()
	go alphabets()
	fmt.Scanln()
	fmt.Println("Finish - 6")
}
