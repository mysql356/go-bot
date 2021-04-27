package pkg

import (
	"fmt"
)

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num, "not found in ", nums)
	}
	fmt.Printf("\n")
}
func Variadic_function1() {
	find(89, 89, 90, 95)
	find(45, 56, 67, 45, 90, 109)
	find(78, 38, 56, 98)
	find(87)
}

func Variadic_function2() {
	nums := []int{89, 90, 95}
	//find(89, nums) //compilation error : we are passing a slice to a function which expects a variable number of arguments
	find(89, nums...)
}

func change(s ...string) {
	s[0] = "Go"
}

func Variadic_string1() {
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}

func change2(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func Variadic_string2() {
	welcome := []string{"hello", "world"}
	fmt.Println(welcome)
	change(welcome...)
	fmt.Println(welcome)
}
