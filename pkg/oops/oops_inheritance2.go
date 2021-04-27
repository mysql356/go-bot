package myoops

 
import (  
    "fmt"
)

type author_1 struct {  
    firstName string
    lastName  string
    bio       string
}

func (a author_1) fullName() string {  
    return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

type post_1 struct {  
    title   string
    content string
    author_1
}

func (p post_1) details() {  
    fmt.Println("Title: ", p.title)
    fmt.Println("Content: ", p.content)
    fmt.Println("Author: ", p.fullName())
    fmt.Println("Bio: ", p.bio)
}

type website struct {  
 posts []post_1
}
func (w website) contents() {  
    fmt.Println("Contents of Website")
    for _, v := range w.posts {
        v.details()
        fmt.Println()
    }
}

func Oops_inheritance2() {  

fmt.Println("\n+++ oops_inheritance-2 +++++")	  
    author1 := author_1{
        "Naveen",
        "Ramanathan",
        "Golang Enthusiast",
    }
    post1 := post_1{
        "Inheritance in Go",
        "Go supports composition instead of inheritance",
        author1,
    }
    post2 := post_1{
        "Struct instead of Classes in Go",
        "Go does not support classes but methods can be added to structs",
        author1,
    }
    post3 := post_1{
        "Concurrency",
        "Go is a concurrent language and not a parallel one",
        author1,
    }
    w := website{
        posts: []post_1{post1, post2, post3},
    }
    w.contents()
}