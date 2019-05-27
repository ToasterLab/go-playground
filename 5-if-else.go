package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even (it's not)")
	} else {
		fmt.Println("7 is odd (yes)")
	}

	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative (it's not)")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit only")
	} else {
		fmt.Println(num, "has multiple digits")
	}
}
