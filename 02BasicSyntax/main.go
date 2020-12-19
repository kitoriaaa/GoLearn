package main

import (
	"fmt"
	"unicode/utf8"
)

func f() int {
	return -2
}

func main() {
	// 変数定義方法たち
	// var n int = 100
	// var n int
	// n = 100

	// var n = 100

	// n := 100

	// var (
	// 	n = 100
	// 	m = 100
	// )

	// -------------------------------
	// 定数定義
	// const n int = 100

	// const m = 100

	// const s = "Hello, " + "世界"

	// const (
	// 	x = 100
	// 	y = 200
	// )

	// const (
	// 	a = iota
	// 	b
	// )

	// const (
	// 	c = 30 << iota
	// 	d
	// 	e
	// )
	// -------------------------------

	// アドレス演算子
	// & ポインタを取得 aPointer = &a
	// * ポインタがさす値を取得 b = *aPointer

	// チャネル演算
	// <-  チャネルへの送受信  ch<-100, <-ch

	// -------------------------------

	// 条件分岐

	// if x == 1 {
	// 	println("xは1")
	// } else if x == 2 {
	// 	println("xは2")
	// } else {
	// 	println("xは1でも2でもない")
	// }

	// if a := f(); a > 0 {
	// 	fmt.Println(a)
	// } else {
	// 	fmt.Println(2 * a)
	// }

	// var a int = 2323
	// switch a {
	// case 1, 2:
	// 	fmt.Println("a is 1 or 2")
	// default:
	// 	fmt.Println("default")
	// }

	// -------------------------------

	// 繰り返しは for しかない

	// for i := 0; i <= 100; i++ {
	// }
	// for i <= 100 {
	// }
	// 無限ループ
	// for {
	// }

	// rangeを使った繰り返し
	// iは index vは配列 いわゆる拡張for文
	// for i, v := range []int{2, 2, 3} {
	// 	println(i, v)
	// }

	// strArray := [5]string{"taro", "Jiro", "Saburo", "Shirou", "Gorou"}
	// for i, v := range strArray {
	// 	fmt.Printf("index is %d, name is %s\n", i, v)
	// }

	type myInteger int
	// var i myInteger = 12

	var str string = "Go言語"
	fmt.Println(str, len(str))
	fmt.Println(str, utf8.RuneCountInString(str))

	c := 'あ'
	fmt.Println(string(c) + " go")
}
