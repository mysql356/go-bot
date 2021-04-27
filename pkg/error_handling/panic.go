package myerrorhandling
import "fmt"

func fullName(firstName *string, lastName *string) {  
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullName")
}

func Panic1() {  
    firstName := "Ram"
	//lastName := "Singh"
    //fullName(&firstName, &lastName)

    fullName(&firstName, nil)
    fmt.Println("returned normally from main")
}

//2nd ex

func fullName2(firstName *string, lastName *string) {  
    defer fmt.Println("deferred call in fullName")
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullName")
}

func Panic2() {  
    defer fmt.Println("deferred call in main")
    firstName := "Ram"
	//lastName := "Singh"
    //fullName2(&firstName, &lastName)

	fullName2(&firstName, nil)
    fmt.Println("returned normally from main")
}