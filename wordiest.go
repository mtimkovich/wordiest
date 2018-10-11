package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var dictionary []string

// Read all English words into slice
func loadWords() {
	file, err := os.Open("TWL06.txt")
	if err != nil {
		fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		if word[0] == '#' {
			continue
		}
		dictionary = append(dictionary, word)
	}
}

func fatal(a ...interface{}) {
	fmt.Println(a...)
	os.Exit(1)
}

func usage() {
	fmt.Println("Usage:", os.Args[0], "[a a2l a3w]+")
	fmt.Println("  - a: [a-z]")
	fmt.Println("  - [1-9]: multiplier")
	fmt.Println("  - l: letter multiplier")
	fmt.Println("  - w: word multiplier")
	os.Exit(0)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}

	args := os.Args[1:]
	tiles := makeTiles(args)

	loadWords()
	solution := Solve(tiles)
	fmt.Println(solution)
}
