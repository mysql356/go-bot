/*
What is a map?

A map is a builtin type in Go which associates a value to a key. The value can be retrieved using the corresponding key.
How to create a map?

A map can be created by passing the type of key and value to the make function. make(map[type of key]type of value) is the syntax to create a map.
*/

package maps

import (  
    "fmt"
)

func Map_create() {  
    var personSalary map[string]int
    if personSalary == nil {
        fmt.Println("map is nil. Going to make one.")
        personSalary = make(map[string]int)
    }
	
	//==============
	var a map[string]int 
	a = map[string]int{}
	a["aa"] = 101
	fmt.Println(a)

	a1:= make(map[string]int)
	a1["aa"] = 10
	fmt.Println(a1)

	var a2 = map[string]int {
		"aa" : 10,
		"bb" : 20,
	}
	a2["cc"] = 30
	fmt.Println(a2)

	var data1 = map[string]string{}
	// var data1 = make(map[string]string)

	data1["a"] = "aa"
	data1["b"] = "bb"
	data1["c"] = "cc"
	fmt.Println(data1)
	//========
}

func Map_assign() {  
    personSalary := make(map[string]int)
    personSalary["steve"] = 12000
    personSalary["jamie"] = 15000
    personSalary["mike"] = 9000
    fmt.Println("personSalary map contents:", personSalary)
}

func Map_assign2() {  
    personSalary := map[string]int {
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("personSalary map contents:", personSalary)
}

func Map_fetch() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    employee := "jamie"
    fmt.Println("Salary of", employee, "is", personSalary[employee])
	fmt.Println("Salary of joe is", personSalary["joe"])
}

func In_map() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    newEmp := "joe"
    value, ok := personSalary[newEmp]
    if ok == true {
        fmt.Println("Salary of", newEmp, "is", value)
    } else {
        fmt.Println(newEmp,"not found")
    }

}

func Map_range() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("All items of a map")
    for key, value := range personSalary {
        fmt.Printf("personSalary[%s] = %d\n", key, value)
    }

}

func Map_delete_element() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("map before deletion", personSalary)
    delete(personSalary, "steve")
    fmt.Println("map after deletion", personSalary)

}

func Map_length() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("length is", len(personSalary))

}

func Map_reference_types() {  
    personSalary := map[string]int{
        "steve": 12000,
        "jamie": 15000,
    }
    personSalary["mike"] = 9000
    fmt.Println("Original person salary", personSalary)
    newPersonSalary := personSalary
    newPersonSalary["mike"] = 18000
    fmt.Println("Person salary changed", personSalary)
}

func Map_equality() {  
    map1 := map[string]int{
        "one": 1,
        "two": 2,
    }

	/*invalid operation: map1 == map3 (map can only be compared to nil)
	map3 := map1
    if map1 == map3 {
	}
	*/
	
	if map1 != nil {
		 fmt.Println("map1 is not nil")
    }
 
	var map2 map[string]int
	if map2 == nil {
		 fmt.Println("map2 is nil")
    }
	
}


 

