/*
Arrays
An array is a collection of elements that belong to the same type. For example the collection of integers 5, 8, 9, 79, 76 form an array. Mixing values of different types, for example an array that contains both strings and integers is not allowed in Go.
*/

package array

import (  
    "fmt"
)


func Array1() {  
    fmt.Println()
    var a [3]int //int array with length 3
    fmt.Println(a)
	

    a[0] = 12 // array index starts at 0
    a[1] = 78
    a[2] = 50
    fmt.Println(a)	
	
    b := [3]int{12, 78, 50} // short hand declaration to create array
    fmt.Println(b)	
	
    c := [3]int{12} 
    fmt.Println(c)	
	
    d := [...]int{12, 78, 50} // ... makes the compiler determine the length
    fmt.Println(d)	
	
  
    e := [3]int{5, 78, 8}
    //var f [5]int   //not working - The size of the array is a part of the type
	var f [3]int
    f = e //not possible since [3]int and [5]int are distinct types
    fmt.Println(e,f)	

	//array length
	g := [...]float64{67.7, 89.8, 21, 78}
    fmt.Println("length of g is",len(g))
}

func Array2() {  
    a := [...]string{"USA", "China", "India"}
    b := a // a copy of a is assigned to b
    b[0] = "UK"
    fmt.Println("a is ", a)
    fmt.Println("b is ", b) 
}

func changeLocal(num [5]int) {  
    num[0] = 27
    fmt.Println("inside function ", num)

}

func Array3() {  
    num := [...]int{5, 6, 7, 8, 8}
    fmt.Println("before passing to function ", num)
    changeLocal(num) //num is passed by value
    fmt.Println("after passing to function ", num)
}

func Array_iteration_for() {  
    a := [...]float64{67.7, 89.8, 21, 78}
    for i := 0; i < len(a); i++ { //looping from 0 to the length of the array
        fmt.Printf("%d th element of a is %.2f\n", i, a[i])
    }
}

func Array_iteration_range() {    
    a := [...]float64{67.7, 89.8, 21, 78}
    sum := float64(0)
    for i, v := range a {//range returns both the index and value
        fmt.Printf("%d the element of a is %.2f\n", i, v)
        sum += v
    }
    fmt.Println("\nsum of all elements of a",sum)	
	
	for _, v := range a { //ignores index  
	  fmt.Printf("element : %.2f\n", v)
	}		
}	

func printarray(a [3][2]string) {  
    for _, v1 := range a {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func Array_multidimensional() {  
    a := [3][2]string{
        {"lion", "tiger"},
        {"cat", "dog"},
        {"pigeon", "peacock"}, //this comma is necessary. The compiler will complain if you omit this comma
    }
    printarray(a)
    var b [3][2]string
    b[0][0] = "apple"
    b[0][1] = "samsung"
    b[1][0] = "microsoft"
    b[1][1] = "google"
    b[2][0] = "AT&T"
    b[2][1] = "T-Mobile"
    fmt.Printf("\n")
    printarray(b)
}

