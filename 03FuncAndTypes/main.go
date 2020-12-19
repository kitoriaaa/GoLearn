package main

import "fmt"

// Game score struct
type Game struct {
	userName int
	point    int
	gameID   int
}

// Person ...(name, age)
type Person struct {
	name string
	age  int
}

func isEven(x int) bool {
	return x%2 == 0
}

func swap(x int, y int) (int, int) {
	return y, x
}

func swap2(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

// Hes is int type
type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}

func main() {
	// cast
	// var a int = 5
	// var b float64 = float64(a)
	// println(b)

	// var sum int
	// sum = 5 + 6 + 3
	// var avg = float64(sum) / 3
	// println(avg)
	// if avg > 4.5 {
	// 	println("good")
	// }

	// -------------------------------

	// var (
	// 	a = true
	// 	b = false
	// 	c = false
	// )

	// if a && b || !c {
	// 	println("true")
	// } else {
	// 	println("false")
	// }

	// 配列
	// var ns1 []int
	// var ns2 = [5]int{1, 2, 3, 4, 5}
	// 要素数を推論
	// ns3 := [...]int{1,2, 3}

	// 5番目が50. 10番目が100で他が0の要素11の配列
	// var ns4 = [...]int{5: 50, 10: 100}
	// for i, v := range ns4 {
	// 	fmt.Printf("index: %d, num: %d\n", i, v)
	// }

	// 配列の長さ

	// println(len(ns4))
	// fmt.Println(ns2[1:3])

	// var m map[string]int

	// 構造体
	// var Person struct {
	// 	name string
	// 	age  int
	// }

	// bob := Person{"Bob", 25}

	// p := struct {
	// 	name string
	// 	age  int
	// }{
	// 	name: "test",
	// 	age:  10,
	// }

	// var ns = [...]int{19, 86, 1, 12}
	// var sum = 0
	// for i := range ns {
	// 	sum += ns[i]
	// }
	// println(sum)

	// map
	// var m map[string]int
	// m = make(map[string]int)
	// m = make(map[string]int, 10)
	// ms := map[string]int{"x": 10, "y": 20}
	// n, ok := ms["x"]

	// println(n, ok)

	// 関数定義

	for i := 1; i <= 100; i++ {
		switch {
		case isEven(i):
			fmt.Printf("%d is Even\n", i)
		default:
			fmt.Printf("%d is Odd\n", i)
		}
	}

	x, y := swap(10, 20)
	fmt.Println(x, y)

	n, m := 10, 20
	swap2(&n, &m)
	fmt.Println(n, m)

	var hex Hex = 100
	f := Hex.String
	fmt.Println(f(hex))
	fmt.Printf("%T\n%s\n", f, f(hex))
}
