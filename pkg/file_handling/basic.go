package myfilehandling

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/gobuffalo/packr"
)

func Basic() {
	fmt.Println("\n+++ File reading : Using absolute file path +++++")

	data, err := ioutil.ReadFile("/Users/manoj/workspace/dev/collection/1-golangbot1/pkg/file_handling/test.txt")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

////golangbot -fpath d:/test.txt
func CommandLine() {
	fmt.Println("\n+++ File reading  : Passing the file path as a command line flag +++++")

	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()
	fmt.Println("value of fpath is", *fptr) //file path

	data, err := ioutil.ReadFile(*fptr) //file read
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	fmt.Println("Contents of file:", string(data))
}

//need to work again : embedding not done
func EmbeddedFile() {
	fmt.Println("\n+++ File reading  :  Bundling the text file along with the binary +++++")

	box := packr.NewBox("d:/")
	data := box.String("test.txt")
	fmt.Println("Contents of file:", data)
}
