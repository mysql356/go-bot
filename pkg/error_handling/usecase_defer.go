package myerrorhandling

import (  
    "fmt"
    "sync"
)

type rect_1007 struct {  
    length int
    width  int
}

func (r rect_1007) area_1007(wg *sync.WaitGroup) {  
    defer wg.Done()
    if r.length < 0 {
        fmt.Printf("rect %v's length should be greater than zero\n", r)
        return
    }
    if r.width < 0 {
        fmt.Printf("rect %v's width should be greater than zero\n", r)
        return
    }
    area := r.length * r.width
    fmt.Printf("rect %v's area %d\n", r, area)
}

func Usecase_defer() {  
fmt.Println("\n+++ Defer usecase +++++")
    var wg sync.WaitGroup
    r1 := rect_1007{-67, 89}
    r2 := rect_1007{5, -67}
    r3 := rect_1007{8, 9}
    rects := []rect_1007{r1, r2, r3}
    for _, v := range rects {
        wg.Add(1)
        go v.area_1007(&wg)
    }
    wg.Wait()
    fmt.Println("All go routines finished executing")
}

