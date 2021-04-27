package main //1
import (
	"fmt" //2
	pkg "golangbot/pkg"
)

// import "os"
// import "strconv"

//Compile	go install index
//Run 		C:\Users\manoj.kumar\go\bin\index

func main() { //3

	v1 := []string{"1 variables ", " 2 type ", " 3 constants ", " 4 fucntions ", " 5 package ", " 6 if & switch ", " 7 loop ", " 8 array ", " 9 slice ", " 10 variadic function ", " 11 map ", " 12 string ", " 13 pointer ", " 14 structures ", " 15 method ", " 16 interface ", " 17 interface adv ", " 18 goroutine", " 19 channels  ", " 20 mutex ", " 21 oops ", " 22 error handling ", " 23 firstclassfunctions  ", " 24 reflection ", " 25 file handling ", " 26 ego ", " 27 pg ", " 28 mysql"}

	for _, v2 := range v1 {
		fmt.Printf("%s ", v2)
		fmt.Printf("\n")
	}

	fmt.Print("Enter chapter: ")
	var ch int
	fmt.Scanln(&ch)

	//ch,_ := strconv.Atoi(os.Args[1])

	switch ch {
	case 1: //variables
		fmt.Println("+++Variables - 1")
		pkg.Var1()
		pkg.Var2()
		pkg.Var3()
		pkg.Var4()
		pkg.Var5()
		pkg.Var6()
		pkg.Var7()
		pkg.Var8()
		pkg.Var9()

	case 2: //type
		fmt.Println("+++Type - 2")
		pkg.Type1_bool()
		pkg.Type2_signed_integers()
		pkg.Type3_signed_integers()
		pkg.Type4_unsigned_integers()
		pkg.Type5_floating_point()
		pkg.Type6_complex()
		pkg.Type7_string()
		pkg.Type8_conversion()

	case 3: //constants
		fmt.Println("+++constants - 3")
		pkg.Const1()
		pkg.Const2()
		pkg.Const3_string()
		pkg.Const4_boolean()
		pkg.Const5_numeric()
		pkg.Const6_numeric_expressions()

	case 4: //fucntions
		fmt.Println("+++fucntions - 4")
		pkg.Funct1_basic()
		pkg.Func1_multiple_return()

	case 5: //package : geometry
		fmt.Println("+++package - 5")
		pkg.Geometry()

	case 6: //if
		fmt.Println("+++if & switch - 6")
		pkg.If_stmt()
		pkg.Switch_stmt()
		pkg.Switch_fallthrough()

	case 7: //loop
		fmt.Println("+++loop - 7")
		pkg.Loop()

	case 8: //array
		fmt.Println("+++array - 8")
		pkg.Myarray()

	case 9: //slice
		fmt.Println("+++array & slice - 9")
		pkg.Myslice()

	case 10: //variadic function
		fmt.Println("+++variadic function - 10")
		pkg.Variadic_function1()
		pkg.Variadic_function2()
		pkg.Variadic_string1()
		pkg.Variadic_string2()

	case 11: //map
		fmt.Println("+++map - 11")
		pkg.Mymap()

	case 12: //string
		fmt.Println("+++string - 12")
		pkg.String_type()
		pkg.String_bytes()
		pkg.String_bytes_rune()
		pkg.For_range_string()
		pkg.String_from_byte()
		pkg.String_from_rune()
		pkg.String_lengh()
		pkg.String_immutable()
		pkg.String_mutable()

	case 13: //pointer
		fmt.Println("+++pointer - 13")
		pkg.Pointer()
		pkg.Pointer_nil()
		pkg.Pointer_new()
		pkg.Pointer_dereferencing()
		pkg.Pointer_dereferencing1()
		pkg.Function_passing_pointer()
		pkg.Function_returning_pointer()
		pkg.Function_passing_array()
		pkg.Function_passing_slices()
		pkg.Pointer_arithmetic()

	case 14: //structures
		fmt.Println("+++structures - 14")
		pkg.Struct_named()
		pkg.Struct_anonymous()
		pkg.Struct_nil()
		pkg.Struct_pointer()
		pkg.Struct_anonymous_fields()
		pkg.Struct_nested()
		pkg.Struct_promoted()
		pkg.Struct_package()
		pkg.Struct_equality()
		pkg.Struct_equality1()

	case 15: //method
		fmt.Println("+++method - 15")
		pkg.Methods()

	case 16: //interface
		fmt.Println("+++interface - 16")
		pkg.Interfaces()

	case 17: //interface
		fmt.Println("+++interface adv - 17")
		pkg.Interfaces_adv()

	case 18: //interface & goroutine
		fmt.Println("+++goroutine - 18")
		pkg.Goroutine_basic()
		pkg.Goroutine_basic_scan()

	case 19: //channels
		fmt.Println("+++channels - 19")
		pkg.Channels()

	case 20: //mutex
		fmt.Println("+++mutex - 20")
		pkg.Mutex_race_condition_main()
		pkg.Mutex_race_condition_solving_by_mutex()
		pkg.Mutex_race_condition_solving_by_channel()

	case 21: //oops
		fmt.Println("+++oops - 21")
		pkg.Oops()

	case 22: //error handling
		fmt.Println("+++error handling - 22")
		pkg.Errors()

	case 23: //firstclassfunctions
		fmt.Println("+++error firstclassfunctions - 23")
		pkg.Firstclassfunctions()

	case 24: //reflection
		fmt.Println("+++error reflection - 24")
		pkg.Reflection()
		pkg.Reflectionex()

	case 25: //file handling
		fmt.Println("+++file handling - 25")
		pkg.Files()

	case 26: //ego
		fmt.Println("+++ego - 26")
		pkg.Ego()

	case 27: //pg
		fmt.Println("+++pg - 27")
		pkg.Pg()

	case 28: //mysql
		fmt.Println("+++mysql - 28")
		pkg.Mysql_connect()
		pkg.Mysql()

	} //@end of switch
}
