package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 5; j <= 10; j++ {
		fmt.Println(j)
	}

	cond := 0
	for {
		fmt.Println("eternal loop", cond)
		cond++
		if cond >= 10 {
			break
		}
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
