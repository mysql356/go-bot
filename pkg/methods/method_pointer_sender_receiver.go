/*Value receivers in methods vs value arguments in functions

func:: sender:value - receiver:value
func:: sender:ref - receiver:pointer

func:: sender:value - receiver:pointer  X - not allowed
func:: sender:ref - receiver:value      X - not allowed

method:: sender:value/ref - receiver:value
method:: sender:value/ref - receiver:pointer

*/

package mymethod
import "fmt"

type rectangle struct {  
    length int
    width  int
}

func area(r rectangle) {  
    fmt.Printf("Area - Function (sender:value, receive: value): %d\n", (r.length * r.width))
}

func area_func_pointer(r *rectangle) {  
  fmt.Printf("Area - Function (sender:ref, receive: pointer): %d\n", (r.length * r.width))
}
//---------------------
func (r rectangle) area() {  
    fmt.Printf("Area + Method (sender:value/ref, receive: value): %d\n", (r.length * r.width))
}

func (r *rectangle) area_method_pointer() {  
    fmt.Printf("Area + Method (sender:value/ref, receive: pointer): %d\n", (r.length * r.width))
}

func Method_function_call() {  
	fmt.Println("\n+++  method_function_call  +++++")
    r := rectangle{
        length: 10,
        width:  5,
    }
    p := &r  

    //func:: sender:value - receiver:value
    area(r)      //area(*p), function call by value  
    //area(&r)   //area(p), function call by reference - not allowed

	//func:: sender:ref - receiver:pointer
    //area_func_pointer(r) //function call by value - not allowed
    area_func_pointer(p)   //function call by reference  
 
    fmt.Println("+++  ---------  +++++")
   
    //method:: sender:value/ref - receiver:value
    r.area()     //method call by value 
    (*p).area()

    p.area()    //method call by reference 
    (&r).area()

    fmt.Println("+++  ---------  +++++")
     
   //method:: sender:value/ref - receiver:pointer
    r.area_method_pointer()     //method call by value 
    (*p).area_method_pointer()

    p.area_method_pointer()    //method call by reference 
    (&r).area_method_pointer()
    
}

///////////////////////////

type rectangle1 struct {  
    length int
    width  int
}

func perimeter(r *rectangle1) {  
    fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *rectangle1) perimeter() {  
    fmt.Println("perimeter method output:", 2*(r.length+r.width))
}

func Method_geometry() {  
    r := rectangle1{
        length: 10,
        width:  5,
    }
    p := &r //pointer to r
    perimeter(p)
    p.perimeter()

    /*
        cannot use r (type rectangle) as type *rectangle in argument to perimeter
    */
    //perimeter(r)

    r.perimeter()//calling pointer receiver with a value

}