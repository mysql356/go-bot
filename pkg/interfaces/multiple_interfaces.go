package myinterface
import "fmt"

type SalaryCalculator_1002 interface {  
    DisplaySalary_1002()
}

type LeaveCalculator_1002 interface {  
    CalculateLeavesLeft_1002() int
}

type Employee_1002 struct {  
    firstName string
    lastName string
    basicPay int
    pf int
    totalLeaves int
    leavesTaken int
}

func (e Employee_1002) DisplaySalary_1002() {  
    fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Employee_1002) CalculateLeavesLeft_1002() int {  
    return e.totalLeaves - e.leavesTaken
}

func Multiple_interface() { 
fmt.Println("\n+++ Multiple_interface +++++")	 
    e := Employee_1002 {
        firstName: "Manoj",
        lastName: "Kumar",
        basicPay: 5000,
        pf: 200,
        totalLeaves: 30,
        leavesTaken: 5,
    }
    var s SalaryCalculator_1002 = e
    s.DisplaySalary_1002()
    var l LeaveCalculator_1002 = e
    fmt.Println("\nLeaves left =", l.CalculateLeavesLeft_1002())
}