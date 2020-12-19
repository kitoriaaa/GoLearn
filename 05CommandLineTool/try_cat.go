package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

// このディレクトリ以下の指定されたファイルを出力する
const dir = "."

func main() {
	var (
		n = flag.Bool("n", false, "Whether Line number is needed")
	)
	flag.Parse()
	args := flag.Args()

	var lineNum int = 1
	for _, file := range args {
		path := filepath.Join(dir, file)
		printFile(*n, path, &lineNum)
	}
}

func printFile(n bool, path string, lineNum *int) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s is not exist\n", path)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if n {
			fmt.Println(strconv.Itoa(*lineNum)+":", scanner.Text())
		} else {
			fmt.Println(scanner.Text())
		}
		*lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed to load", err)
	}
}
