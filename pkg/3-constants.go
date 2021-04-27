package pkg

import (
	"fmt"
	"math"
)


func Const1() {  
    const a = 55 //allowed
    //a = 89 //reassignment not allowed
	fmt.Println(a)
}

func Const2() {  
    var a = math.Sqrt(4)//allowed
    //const b = math.Sqrt(4)//not allowed
	fmt.Println(a)
}

func Const3_string() {  

	const hello = "Hello World"  
	fmt.Printf(hello)
	//hello = "modify"

    var defaultName = "Sam" //allowed
    fmt.Printf("\ntype %T value %v", defaultName, defaultName)
	
	type myString string
	var customName myString = "Sam" //allowed
	customName = "defaultName" //not allowed	
	 
	fmt.Printf("\ntype %T value %v", customName, customName)
}

func Const4_boolean() { 

	fmt.Println()
    const trueConst = true
    type myBool bool
    var defaultBool = trueConst //allowed
    var customBool myBool = trueConst //allowed
    //defaultBool = customBool //not allowed
	fmt.Println(trueConst,defaultBool,customBool)
}

func Const5_numeric() { 

    const a = 5
    var intVar int = a
    var int32Var int32 = a
    var float64Var float64 = a
    var complex64Var complex64 = a
    fmt.Println("intVar",intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var",complex64Var)
	
	
    var i = 5
    var f = 5.6
    var c = 5 + 6i
    fmt.Printf("i's type %T, f's type %T, c's type %T", i, f, c)
}

func Const6_numeric_expressions() { 

    var a = 5.9/8
    fmt.Printf("\na's type %T value %v",a, a)
}
