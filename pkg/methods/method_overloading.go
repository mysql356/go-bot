package mymethod

import (  
    "fmt"
    "math"
)

type Rectangle struct {  
    length int
    width  int
}

type Circle struct {  
    radius float64
}

func (r Rectangle) Area() int {  
    return r.length * r.width
}

func (c Circle) Area() float64 {  
    return math.Pi * c.radius * c.radius
}

func Method_overloading() { 
	fmt.Println("\n+++  method_overloading  +++++")	   
    r := Rectangle{
        length: 10,
        width:  5,
    }
    fmt.Printf("Area of rectangle %d\n", r.Area())
    c := Circle{
        radius: 12,
    }
    fmt.Printf("Area of circle %f", c.Area())
}
