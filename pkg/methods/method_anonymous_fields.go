package mymethod

import (  
    "fmt"
)

type address struct {  
    city  string
    state string
}

func (a address) fullAddress() {  
    fmt.Printf("Full address: %s, %s", a.city, a.state)
}

type person struct {  
    firstName string
    lastName  string
    address
}

func Method_anonymous_fields() {  
	fmt.Println("\n+++  method_anonymous_fields  +++++")	 
    p := person{
        firstName: "Ram",
        lastName:  "Singh",
        address: address {
            city:  "New Delhi",
            state: "Delhi",
        },
    }

    p.fullAddress() //accessing fullAddress method of address struct
	fmt.Println()
    fmt.Println(p.firstName, p.address.city, p.city)

}