package myerrorhandling

import (  
    "fmt"
    "sync"
)

type rect_1006 struct {  
    length int
    width  int
}

func (r rect_1006) area_1006(wg *sync.WaitGroup) {  
    if r.length < 0 {
        fmt.Printf("rect %v's length should be greater than zero\n", r)
        wg.Done()
        return
    }
    if r.width < 0 {
        fmt.Printf("rect %v's width should be greater than zero\n", r)
        wg.Done()
        return
    }
    area := r.length * r.width
    fmt.Printf("rect %v's area %d\n", r, area)
    wg.Done()
}

func Usecase() {  
fmt.Println("\n\n+++ Usecase +++++")
    var wg sync.WaitGroup
    r1 := rect_1006{-67, 89}
    r2 := rect_1006{5, -67}
    r3 := rect_1006{8, 9}
    rects := []rect_1006{r1, r2, r3}
    for _, v := range rects {
        wg.Add(1)
        go v.area_1006(&wg)
    }
    wg.Wait()
    fmt.Println("All go routines finished executing")
}

