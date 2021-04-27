package myinterface

import "fmt"

type Describer_1001 interface {  
    Describe_1001()
}

type Person_1001 struct {  
    name string
    age  int
}

func (p *Person_1001) Describe_1001() { 
    fmt.Printf("%s is %d years old\n", p.name, p.age)
}
 
func Interface_pointer_struct_method() { 

fmt.Println("\n+++ Interface_pointer_struct_method +++++")

    //struct + method + interface
    var d1 Describer_1001
    p1 := Person_1001{"Ram", 25}
    d1 = &p1
    d1.Describe_1001()

    //struct + method
    p2 := Person_1001{"Shyam", 25}
    p2.Describe_1001()
    (&p2).Describe_1001()

}


/*
package main

import "fmt"

type person struct {
	first string
}

type human interface {
	speak()
}

func (p *person) speak() {
	fmt.Println("Hello")
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "James",
	}

	saySomething(&p1)

	// does not work
	// saySomething(p1)

	p1.speak()
}


*/