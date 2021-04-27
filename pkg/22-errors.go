package pkg

import myerrorhandling "golangbot/pkg/error_handling"

//D:\go_test\src\golangbot\error_handling

func Errors() {

	myerrorhandling.Defer_function()
	myerrorhandling.Defer_method()
	myerrorhandling.Defer_reverse_str()
	myerrorhandling.Usecase()
	myerrorhandling.Usecase_defer()

	myerrorhandling.IsExistFile1()
	myerrorhandling.IsExistFile2()
	myerrorhandling.IsHostExist()
	myerrorhandling.IsExistFilePath()
	myerrorhandling.CustomErr()
	myerrorhandling.CustomErr1()
	myerrorhandling.CustomErr2()
	myerrorhandling.CustomErr3()

	//myerrorhandling.Panic1();
	//myerrorhandling.Panic2();
	//myerrorhandling.PanicRecover();
	//myerrorhandling.PanicGoroutine();

	//myerrorhandling.PanicRuntime();
	myerrorhandling.PanicRuntimeRecover()

}
