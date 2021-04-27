package pkg

import (
	"fmt"
	"log"

	rectangle "golangbot/pkg/package"
)

var _ = fmt.Println
var _ = rectangle.Area //error silencer

var rectLen, rectWidth float64 = 6, 7

func init() {
	// println("main package initialized")
	if rectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func Geometry() {
	fmt.Println("\n+++ package +++++")
	fmt.Println("Geometrical shape properties")
	fmt.Printf("Area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("Diagonal of the rectangle %.2f \n", rectangle.Diagonal(rectLen, rectWidth))
	fmt.Printf("Area of circle %.2f\n", CircleArea(5)) //circleArea() defined in 5-package-func.go

}
