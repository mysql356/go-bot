package myinterface

import (  
    "fmt"
)

type SalaryCalculator_1003 interface {  
    DisplaySalary_1003()
}

type LeaveCalculator_1003 interface {  
    CalculateLeavesLeft_1003() int
}

type EmployeeOperations_1003 interface {  
    SalaryCalculator_1003
    LeaveCalculator_1003
}

type Employee_1003 struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}

func (e Employee_1003) DisplaySalary_1003() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee_1003) CalculateLeavesLeft_1003() int {  
    return e.totalLeaves - e.leavesTaken
}

func Embedded_interface() {  
fmt.Println("\n+++ Embedded_interface +++++")	
    e := Employee_1003 {
        firstName: "Manoj",
        lastName: "kumar",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var empOp EmployeeOperations_1003 = e
    empOp.DisplaySalary_1003()
    fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft_1003())
}