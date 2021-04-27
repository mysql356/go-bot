package myinterface

 
import "fmt"

type Describer_type interface {  
    Describe_type()
}
type Person_type struct {  
    name string
    age  int
}

func (p Person_type) Describe_type() {  
    fmt.Printf("%s is %d years old", p.name, p.age)
}

func findType2(i interface{}) {  
    switch v := i.(type) {
    case Describer_type:
        v.Describe_type()
    default:
        fmt.Printf("unknown type\n")
    }
}

func FindType2() {  
 fmt.Println("\n+++ FindType2 +++++")	
    findType2("Manoj")
    p := Person_type{
        name: "Manoj K",
        age:  25,
    }
    findType2(p)
}