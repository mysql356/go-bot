/*Pointer receivers vs value receivers*/
package mymethod

import (  
    "fmt"
)

type Employee1 struct {  
    name string
    age  int
}

//value receiver  
func (e Employee1) changeName(newName string) {  
    e.name = newName
}

//pointer receiver  
func (e *Employee1) changeAge(newAge int) {  
    e.age = newAge
}

func Method_pointer_receiver() {  
	e := Employee1{
		name: "Ram",
		age:  50,
	}
	fmt.Println("\n+++  method_pointer_receiver  +++++")	 

	fmt.Printf(e.name) //Ram
	e.changeName("Shyam") //no impact 
	fmt.Printf(e.name) //Ram

	fmt.Printf("\nage: %d", e.age) //50 - before change
	(&e).changeAge(30) //method calling (by pointer)
	//e.changeAge(30)      //method calling  
	fmt.Printf("\nage: %d", e.age) //30 - after change

}
