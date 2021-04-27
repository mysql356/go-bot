package pkg

import (
	"fmt"
)

func Loop() {
	fmt.Println()
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}

	//break
	fmt.Println("\nbreak")
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")

	//continue
	fmt.Println("\ncontinue")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	//Nested for loops
	fmt.Println("\nNested for loops")
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//Labels
outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}

	//initialisation and post are omitted
	fmt.Println("\ninitialisation and post are omitted")
	i := 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}

	//semicolons are ommitted and only condition is present
	fmt.Println("\nsemicolons are ommitted and only condition is present")
	i = 0
	for i <= 10 {
		fmt.Printf("%d ", i)
		i += 2
	}

	//multiple initialisation and increment
	fmt.Println("\nmultiple initialisation and increment")
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 {
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	//infinite loop
	for {
		fmt.Println("Hello World")
		break
	}
}
