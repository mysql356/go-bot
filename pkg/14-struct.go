/*
A structure is a user defined type which represents a collection of fields. It can be used in places where it makes sense to group the data into a sin
*/
package pkg

import (
	"fmt"
	"golangbot/pkg/structpackage" //struct_package()
)

type Employee_named struct {
	firstName, lastName string
	age, salary         int
}

func Struct_named() {

	//creating structure using field names
	emp1 := Employee_named{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	//creating structure without using field names
	emp2 := Employee_named{"Thomas", "Paul", 29, 800} //all
	//emp21 := Employee_named{"Thomas", "Paul"} //invalid
	emp3 := Employee_named{firstName: "Scott", salary: 600} //only selected value

	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)
	fmt.Println("Employee 3", emp3)
	fmt.Println("Employee 3 ::  Name :", emp3.firstName) //Accessing individual fields of a struct
}

func Struct_anonymous() {
	emp3 := struct {
		firstName, lastName, city string
		age, salary               int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		city:      "Delhi",
		age:       31,
		salary:    5000,
	}

	fmt.Println("Employee 3", emp3)
}

type Employee_nil struct {
	firstName, lastName string
	age, salary         int
}

func Struct_nil() {
	var emp4 Employee_nil //zero valued structure
	var emp41 Employee_nil
	fmt.Println("Employee 4", emp4)

	emp41.firstName = "Jack"
	emp41.lastName = "Adams"

	fmt.Println("Employee 41", emp41)
}

type Employee_pointer struct {
	firstName, lastName string
	age, salary         int
}

func Struct_pointer() {
	emp8 := &Employee_pointer{"Sam", "Anderson", 55, 6000}
	fmt.Println((*emp8).firstName, (*emp8).age)

	/*The language gives us the option to use emp8.firstName instead of the explicit dereference (*emp8).firstName to access the firstName field.*/
	fmt.Println(emp8.firstName, emp8.age)
}

/*
Anonymous fields
It is possible to create structs with fields which contain only a type without the field name. These kind of fields are called anonymous fields.
*/
type Person struct {
	string
	int
}

func Struct_anonymous_fields() {
	p := Person{"Naveen", 50}
	fmt.Println(p)
	fmt.Println(p.string, p.int)

	var p1 Person
	p1.string = "Manoj"
	p1.int = 50
	fmt.Println(p1)
	fmt.Println(p1.string, p1.int)
}

/*Nested structs
It is possible that a struct contains a field which in turn is a struct. These kind of structs are called as nested structs.
*/

type Address_nested struct {
	city, state string
}
type Person_nested struct {
	name    string
	age     int
	address Address_nested
}

func Struct_nested() {
	var p Person_nested
	p.name = "Naveen"
	p.age = 50
	p.address = Address_nested{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)
}

/*
Promoted fields
Fields that belong to a anonymous struct field in a structure are called promoted fields since they can be accessed as if they belong to the structure which holds the anonymous struct field. I can understand that this definition is quite complex so lets dive right into some code to understand this :).
*/
type Address_promoted struct {
	city, state string
}
type Person_promoted struct {
	name string
	age  int
	Address_promoted
}

func Struct_promoted() {
	var p Person_promoted
	p.name = "Naveen"
	p.age = 50
	p.Address_promoted = Address_promoted{
		city:  "Chicago",
		state: "Illinois",
	}
	fmt.Println("Name:", p.name)
	fmt.Println("Age:", p.age)
	fmt.Println("City:", p.city)   //city is promoted field
	fmt.Println("State:", p.state) //state is promoted field
}

/*
Exported Structs and Fields
If a struct type starts with a capital letter, then it is a exported type and it can be accessed from other packages. Similarly if the fields of a structure start with caps, they can be accessed from other packages.

*/

func Struct_package() {
	var spec structpackage.Spec
	spec.Maker = "apple"
	spec.Price = 50000
	//spec.model = "Mac Mini" //spec.model undefined : cannot refer to unexported field
	fmt.Println("Spec:", spec)
}

/*
Structs Equality

Structs are value types and are comparable if each of their fields are comparable. Two struct variables are considered equal if their corresponding fields are equal.
*/
type name_equality struct {
	firstName string
	lastName  string
}

func Struct_equality() {
	name1 := name_equality{"Steve", "Jobs"}
	name2 := name_equality{"Steve", "Jobs"}
	if name1 == name2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name_equality{firstName: "Steve", lastName: "Jobs"}
	name4 := name_equality{}
	name4.firstName = "Steve"
	//name4.lastName = "Jobs"
	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}

/*Struct variables are not comparable if they contain fields which are not comparable
invalid operation: image1 == image2 (struct containing map[int]int cannot be compared)
type image struct {
    data map[int]int
}
func struct_equality1() {
    image1 := image{data: map[int]int{
        0: 155,
    }}
    image2 := image{data: map[int]int{
        0: 155,
    }}
    if image1 == image2 {
        fmt.Println("image1 and image2 are equal")
    }
}
*/

type image struct {
	data int
}

func Struct_equality1() {
	image1 := image{data: 155}

	image2 := image{data: 155}

	if image1 == image2 {
		fmt.Println("image1 and image2 are equal")
	}
}
