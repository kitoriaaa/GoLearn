package main

import "fmt"

func isEvenIf() {
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d-偶数\n", i)
		} else {
			fmt.Printf("%d-奇数\n", i)
		}
	}
}

func isEvenSwitch() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%2 == 0:
			fmt.Printf("%d-偶数\n", i)
		default:
			fmt.Printf("%d-奇数\n", i)
		}
	}
}

func main() {
	isEvenIf()
	// isEvenSwitch()
}
