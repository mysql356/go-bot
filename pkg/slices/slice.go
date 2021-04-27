/*Slices
A slice is a convenient, flexible and powerful wrapper on top of an array. Slices do not own any data on their own. They are the just references to existing arrays.
*/

package slice

import (  
    "fmt"
)

func Slice1() {  
    a := [5]int{76, 77, 78, 79, 80}
    var b []int = a[1:4] //creates a slice from a[1] to a[3]
    fmt.Println(b)
	
    c := []int{6, 7, 8} //creates and returns a slice reference
    fmt.Println(c)	
	
	/*A slice does not own any data of its own. It is just a representation of the underlying array. 
	Any modifications done to the slice will be reflected in the underlying array.
	*/
    darr := [...]int{1,2,3,4,5,6,7,8,9}
    dslice := darr[2:5]
    fmt.Println("array before",darr)
    for i := range dslice {
        dslice[i]++
    }
    fmt.Println("array after",darr) 	
}

func Slice2() {  
    numa := [3]int{78, 79 ,80}
    nums1 := numa[:] //creates a slice which contains all elements of the array
    nums2 := numa[:]
    fmt.Println("array before change 1",numa)
    nums1[0] = 100
    fmt.Println("array after modification to slice nums1", numa)
    nums2[1] = 101
    fmt.Println("array after modification to slice nums2", numa)
}

/*
length and capacity of a slice
The length of the slice is the number of elements in the slice. 
The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
*/
func Slice3() { 
	fmt.Println()
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "banana"}
    fruitslice := fruitarray[2:4]
    fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice)) //length of is 2(grape,mango) and capacity is 3(grape,mango,banana)
}

func Slice4() { 
	fmt.Println()
    fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
    fruitslice := fruitarray[1:3]
    fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of is 2 and capacity is 6
    fruitslice = fruitslice[:cap(fruitslice)] //re-slicing furitslice till its capacity
    fmt.Println("After re-slicing length is",len(fruitslice), "and capacity is",cap(fruitslice))
}

/*
creating a slice using make
func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity. The capacity parameter is optional and defaults to the length. The make function creates an array and returns a slice reference to it.
*/

func Slice_make() {  
    i := make([]int, 5, 5)
    fmt.Println(i)
}
func Slice_append() {  
    cars := []string{"Ferrari", "Honda", "Ford"}
    fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
    cars = append(cars, "Toyota")
    fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	
	for k,v := range cars {
        fmt.Println(k,v);
    }	
	
    veggies := []string{"potatoes","tomatoes","brinjal"}
    fruits := []string{"oranges","apples"}
    food := append(veggies, fruits...)
    fmt.Println("food:",food)	
}

func Slice_nil() {  
    var names []string //zero value of a slice is nil
    if names == nil {
        fmt.Println("slice is nil going to append")
        names = append(names, "John", "Sebastian", "Vinay")
        fmt.Println("names contents:",names)
    }
}

func subtactOne(numbers []int) {  
    for i := range numbers {
        numbers[i] -= 2
    }

}
func Slice_in_function() {  
    nos := []int{8, 7, 6}
    fmt.Println("slice before function call", nos)
    subtactOne(nos)                               //function modifies the slice
    fmt.Println("slice after function call", nos) //modifications are visible outside
}

func Slice_multidimensional() {  
     pls := [][]string {
            {"C", "C++"},
            {"JavaScript"},
            {"Go", "Rust"},
            }
    for _, v1 := range pls {
        for _, v2 := range v1 {
            fmt.Printf("%s ", v2)
        }
        fmt.Printf("\n")
    }
}

func countries() []string {  
    countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
    neededCountries := countries[:len(countries)-2]
    countriesCpy := make([]string, len(neededCountries))
    copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
    return countriesCpy
}
func Slice_memory_optimisation() {  
    countriesNeeded := countries()
    fmt.Println(countriesNeeded)
}
