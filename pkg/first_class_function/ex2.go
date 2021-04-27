package myfirstclassfunction
import "fmt"

func iMap(s []int, f func(int) int) []int {  
    var r []int
    for _, v := range s {
        r = append(r, f(v))
    }
    return r
}

func Ex2() {  
fmt.Println("\n+++ first_class_function Ex2 +++++") 
    a := []int{5, 6, 7, 8, 9}
    r := iMap(a, func(n int) int {
        return n * 5
    })
    fmt.Println(r)
}
