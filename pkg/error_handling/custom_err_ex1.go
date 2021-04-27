package myerrorhandling

import (  
    "fmt"
    "math"
)

func circleArea1(radius float64) (float64, error) {  
    if radius < 0 {
        return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
    }
    return math.Pi * radius * radius, nil
}

func CustomErr1() {  
fmt.Println("\n+++ CustomErr 1 +++++")
    radius := 20.0
    area, err := circleArea1(radius)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("Area of circle %0.2f\n", area)
}

 