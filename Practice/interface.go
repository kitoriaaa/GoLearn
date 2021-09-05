package main

import "fmt"

type Eater interface {
	PutIn()
	Chew()
	Swallow()
}

type Human struct {
	Height int
}

type Turtle struct {
	Kind string
}

func (h Human) PutIn() {
	fmt.Println("道具を使って丁寧に口に運ぶ")
}

func (h Human) Chew() {
	fmt.Println("歯でしっかり噛む")
}

func (h Human) Swallow() {
	fmt.Println("よく噛んだら飲み込む")
}

func (t Turtle) PutIn() {
	fmt.Println("獲物を見つけたら首を素早く伸ばして噛む")
}

func (t Turtle) Chew() {
	fmt.Println("くちばしでかみ砕く")
}

func (t Turtle) Swallow() {
	fmt.Println("小さくかみ砕いたら飲み込む")
}

func EatAll(e Eater) {
	e.PutIn()
	e.Chew()
	e.Swallow()
}

func main() {
	var man Human = Human{Height: 170}
	var cheloniaMydas = Turtle{Kind: "アオウミガメ"}
	var eat Eater
	fmt.Println("<人間が食べる>")
	eat = man
	EatAll(eat)
	fmt.Println("<カメが食べる>")
	eat = cheloniaMydas
	EatAll(eat)
}
