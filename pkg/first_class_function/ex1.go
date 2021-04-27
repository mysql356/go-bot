package myfirstclassfunction
import "fmt"

type student struct {  
    firstName string
    lastName  string
    grade     string
    country   string
}

func filter(s []student, f func(student) bool) []student {  
    var r []student
    for _, v := range s {
        if f(v) == true {
            r = append(r, v)
        }
    }
    return r
}

func Ex1() {  
fmt.Println("\n+++ first_class_function Ex1 +++++")
    s1 := student{
        firstName: "Manoj",
        lastName:  "Kumar",
        grade:     "A",
        country:   "India",
    }
    s2 := student{
        firstName: "Ram",
        lastName:  "Singh",
        grade:     "B",
        country:   "UK",
    }
    s := []student{s1, s2}
    f := filter(s, func(s student) bool {
        if s.grade == "A" {
            return true
        }
        return false
    })
	fmt.Println(s)
    fmt.Println(f)
}