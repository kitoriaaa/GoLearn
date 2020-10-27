package main

func main() {
	arr := []int{19, 86, 1, 12}
	var sum int = 0
	for _, n := range arr {
		sum += n
	}
	println(sum)
}
