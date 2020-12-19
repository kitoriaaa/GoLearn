package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var msg = flag.String("msg", "デフォルト値", "説明")
var n int

func init() {
	flag.IntVar(&n, "n", 1, "回数")
}

func main() {
	flag.Parse()
	fmt.Println(strings.Repeat(*msg, n))
	scanner := bufio.NewScanner(os.Stdin)
	var str string
	for scanner.Scan() {
		str = scanner.Text()
		fmt.Println(scanner.Text())
		break
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "読み込みに失敗しました:", err)
	}

	path := filepath.Join("dir", str)
	fmt.Println(path)
	fmt.Println(filepath.Ext(path))
	fmt.Println(filepath.Base(path))
	fmt.Println(filepath.Dir(path))

	fmt.Println(strings.Repeat("--", 10))
	fmt.Println("file Walk")

	err := filepath.Walk(".",
		func(path string, info os.FileInfo, erro error) error {
			if filepath.Ext(path) == ".go" {
				fmt.Println(path)
			}
			return nil
		})

	if err != nil {
		fmt.Println(err)
	}
}
