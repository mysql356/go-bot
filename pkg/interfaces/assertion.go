package myinterface
 

import (  
    "fmt"
)

func assert(i interface{}) {  
    s := i.(int) //get the underlying int value from i
    fmt.Println(s) 
}
func Assertion() {  
    var s interface{} = 56 // "ss" - panic 
    assert(s)
}

//////////////////

func assert2(i interface{}) {
	v, ok := i.(int)
	fmt.Println(v, ok)
}
func Assertion2() {
	var s interface{} = 56
	assert2(s)
	var i interface{} = "Steven Paul"
	assert2(i)
}
