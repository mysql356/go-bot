package myoops

import (
	"fmt"
	employee "golangbot/pkg/oops/oops"
)

func Oops_basic() {  
    e := employee.Employee {
        FirstName: "Ram",
        LastName: "Singh",
        TotalLeaves: 30,
        LeavesTaken: 15,
    }
    e.LeavesRemaining()
}

func Oops_basic1() {  
	fmt.Println()
    var e employee.Employee
    e.LeavesRemaining()
}

func Oops_construct() { 
	fmt.Println()
    e := employee.New("Shyam", "Singh", 30, 20)
    e.LeavesRemaining()
}
