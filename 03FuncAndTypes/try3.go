package main

func swap(x, y int) (int, int) {
	return y, x
}

func swap2(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	n, m := swap(10, 20)
	println(n, m)

	k, l := 10, 20
	swap2(&k, &l)
	println(k, l)
}
