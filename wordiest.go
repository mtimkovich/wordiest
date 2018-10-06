package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var dictionary []string

// Read all English words into slice
func loadWords() {
	file, err := os.Open("sowpods.txt")
	if err != nil {
		log.Fatal(err)
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

// Find the highest scoring word we can make with our tiles
// Return the leftover ones.
// TODO: Edge cases, e.g. no matches
func solve(tiles Tiles) (WordAndScore, Tiles) {
	highScore := WordAndScore{Tiles{}, 0}
	var bestRemaining Tiles

	for _, word := range dictionary {
		if used, remaining, ok := tiles.Contains(word); ok {
			if score := used.Score(); score > highScore.Score {
				highScore = WordAndScore{used, score}
				bestRemaining = remaining
			}
		}
	}

	return highScore, bestRemaining
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:", os.Args[0], "[a] [a2l] [a3w]")
		fmt.Println("  - a: [a-z]")
		fmt.Println("  - [1-9]: multiplier")
		fmt.Println("  - l: letter multiplier")
		fmt.Println("  - w: word multiplier")
		os.Exit(0)
	}

	args := os.Args[1:]
	// argsStr := "t e g5w o o l2l t i p o2l u d a n2w"
	// args := strings.Split(argsStr, " ")
	// fmt.Println(args)

	tiles := makeTiles(args)

	loadWords()

	word1, remainder := solve(tiles)
	word2, _ := solve(remainder)

	fmt.Printf("%v: %v + %v\n", word1.Score+word2.Score, word1, word2)
}
