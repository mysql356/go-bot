package pkg

import "math"

func init() {
	//fmt.Println("Area of circle function is defined in main package")
}

func CircleArea(radius float64) float64 {
	area := math.Pi * radius * radius
	return area
}
