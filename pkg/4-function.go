package pkg

import (
	"fmt"
)

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}
func Funct1_basic() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("\nTotal price is", totalPrice)
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectProps_namedReturnValues(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return //no explicit return value
}

func Func1_multiple_return() {
	area, perimeter := rectProps(10, 5)
	fmt.Printf("\nArea %f Perimeter %f", area, perimeter)

	area, perimeter = rectProps_namedReturnValues(20, 10)
	fmt.Printf("\nArea %f Perimeter %f", area, perimeter)

	area, _ = rectProps(10.8, 5.6) // perimeter is discarded
	fmt.Printf("\nArea %f ", area)

}
