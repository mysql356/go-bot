package pkg

import myfilehandling "golangbot/pkg/file_handling"

//D:\go_test\src\golangbot\error_handling

func Files() {

	myfilehandling.Basic()

	//golangbot -fpath d:/test.txt
	//myfilehandling.CommandLine();

	//myfilehandling.EmbeddedFile();
	myfilehandling.FileReadingChars()
	myfilehandling.FileReadingLine()

	myfilehandling.FileWritingString()
	myfilehandling.FileWritingBytes()
	myfilehandling.FileWritingLinebyline()
	myfilehandling.FileApend()

	myfilehandling.FileConcurrently()

}
