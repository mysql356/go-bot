package mymethod

import (  
    "fmt"
)

type Employee struct {  
    name     string
    salary   int
    currency string
}

//method
func (e Employee) displaySalary() {  
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

//function
func displaySalary1(e Employee) {  
    fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
}

func Method_basic() { 
    fmt.Println("\n+++  method_basic  +++++")	  
    emp1 := Employee {
        name:     "Ram Singh",
        salary:   5000,
        currency: "$",
    }
    emp1.displaySalary() //method calling
	fmt.Println()
    displaySalary1(emp1) //function calling
}

/*
Go is not a pure object oriented programming language and it does not support classes. 
Hence methods on types is a way to achieve behaviour similar to classes.

Methods with same name can be defined on different types whereas functions with the same names are not allowed. 

*/
