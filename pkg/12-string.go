/*
A string in Go is a slice of bytes. Strings can be created by enclosing their contents inside " ".
Strings in Go are Unicode compliant and are UTF-8 Encoded.
*/

package pkg

import (
	"fmt"
	"unicode/utf8"
) //string_lengh()

func String_type() {
	name := "Hello World"
	fmt.Println(name)
}

func printBytes(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}
func printChars(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

//Since a string is a slice of bytes, it's possible to access each byte of a string.
func String_bytes() {
	name := "Hello World"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
	fmt.Println()

	name = "Golang"
	printBytes(name)
	fmt.Printf("\n")
	printChars(name)
}

/*rune

A rune is a builtin type in Go and it's the alias of int32. rune represents a Unicode code point in Go. It does not matter how many bytes the code point occupies, it can be represented by a rune. Lets modify the above program to print characters using a rune.

*/
func printBytesRune(s string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}
func printCharsRune(s string) {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}
func String_bytes_rune() {
	fmt.Println()
	name := "Hello World"
	printBytesRune(name)
	fmt.Printf("\n")
	printCharsRune(name)
	fmt.Println()

	name = "Golang"
	printBytesRune(name)
	fmt.Printf("\n")
	printCharsRune(name)
}

func printCharsAndBytes(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d %x\n", rune, index, rune)
	}
}

func For_range_string() {
	fmt.Println()
	name := "Golang"
	printCharsAndBytes(name)
}

//constructing string from slice of bytes
func String_from_byte() {
	byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	//byteSlice := []byte{67, 97, 102, 195, 169}
	str := string(byteSlice)
	fmt.Println(str)

	//byteSlice1 := []byte{0x47, 0x6f, 0x6c, 0x61, 0x6e, 0x67}
	byteSlice1 := []byte{71, 111, 108, 97, 110, 103}

	str1 := string(byteSlice1)
	fmt.Println(str1)
}

func String_from_rune() {
	runeSlice := []rune{0x0047, 0x006f}
	str := string(runeSlice)
	fmt.Println(str)
}

func length(s string) {
	fmt.Printf("length of %s is %d\n", s, utf8.RuneCountInString(s))
}
func String_lengh() {

	word1 := "Señor"
	length(word1)
	word2 := "Pets"
	length(word2)
}

func immutable(s string) string {
	// s[0] = 'a'//any valid unicode character within single quote is a rune
	return s
}

func String_immutable() {
	h := "hello"
	fmt.Println(immutable(h))
}

/*
To workaround this string immutability, strings are converted to a slice of runes.
Then that slice is mutated with whatever changes needed and converted back to a new string.
*/

func mutate_rune(s []rune) string {
	s[0] = 'a'
	return string(s)
}
func String_mutable() {
	h := "hello"
	fmt.Println(mutate_rune([]rune(h)))
}
