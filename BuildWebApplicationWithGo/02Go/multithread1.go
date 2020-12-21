// package main

// import "fmt"

// func sum(a []int, c chan int) {
// 	total := 0
// 	for _, v := range a {
// 		total += v
// 	}
// 	c <- total
// }

// func main() {
// 	a := []int{7, 2, 8, -9, 4, 0}

// 	c := make(chan int)
// 	go sum(a[:len(a)/2], c)
// 	go sum(a[len(a)/2:], c)

// 	x, y := <-c, <-c

// 	fmt.Println(x, y, x+y)
// }

// package main

// import "fmt"

// func main() {
// 	c := make(chan int)

// 	go func(s chan<- int) {
// 		for i := 0; i < 5; i++ {
// 			s <- i
// 		}
// 		close(s)
// 	}(c)

// 	for {
// 		if value, ok := <-c; !ok {
// 			break
// 		} else {
// 			fmt.Println(value)
// 		}
// 	}
// }

package main

import "fmt"

func main() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}
