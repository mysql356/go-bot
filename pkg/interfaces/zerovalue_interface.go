package myinterface

import "fmt"

type Describer_1004 interface {
	Describe_1004()
}

func Zerovalue_interface() {
fmt.Println("\n+++ Zerovalue_interface +++++")
	var d1 Describer_1004
        //d1.Describe()
	if d1 == nil {
		fmt.Printf("d1 is nil and has type %T value %v\n", d1, d1)
	}
}

 
