package pkg

import (
	"fmt"
	"reflect"
	"unsafe"
)

type order_ref struct {
	ordId      int
	customerId int
}

func createQuery_ref(q interface{}) {
	t := reflect.TypeOf(q)
	v := reflect.ValueOf(q)
	k := t.Kind()

	fmt.Println("Type ", t)
	fmt.Println("Value ", v)
	fmt.Println("Kind", k)
	fmt.Println("Number of fields", v.NumField())

	for i := 0; i < v.NumField(); i++ {
		fmt.Printf("%T => %v\n", v.Field(i), v.Field(i))
	}

	a := 56
	x := reflect.ValueOf(a).Int()
	fmt.Printf("type:%T value:%v\n", x, x)
	b := "Manoj"
	y := reflect.ValueOf(b).String()
	fmt.Printf("type:%T value:%v\n", y, y)

	var a1 int = 456
	fmt.Printf("%T =>  %d\n", a1, unsafe.Sizeof(a1))

}
func Reflection() {
	fmt.Println("\n+++ reflection  +++++")
	o := order_ref{
		ordId:      456,
		customerId: 56,
	}
	createQuery_ref(o)

}
