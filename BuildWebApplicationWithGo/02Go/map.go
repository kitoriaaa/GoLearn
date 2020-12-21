package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func sumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func add(a *int) int {
	*a = *a + 1
	return *a
}

func main() {
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["tow"] = 2

	fmt.Printf("one %d\n", numbers["one"])
	// fmt.Print("Hello")

	rating := map[string]float32{"C#": 5, "GO": 4.5, "Python": 4.5, "C++": 2}
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}

	delete(rating, "C")

	x := 3
	y := 4
	s, p := sumAndProduct(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, s)
	fmt.Printf("%d * %d = %d\n", x, y, p)

	x1 := 3
	x1Inc := add(&x1)
	fmt.Printf("x1, x1_inc = %d, %d", x1, x1Inc)

}
