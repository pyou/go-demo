package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	for j := 5; j < 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}