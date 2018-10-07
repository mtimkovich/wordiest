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

type Solution struct {
	Word1 WordAndScore
	Word2 WordAndScore
}

func (s Solution) Total() int {
	return s.Word1.Score + s.Word2.Score
}

func (s Solution) String() string {
	return fmt.Sprintf("%v: %v + %v", s.Total(), s.Word1, s.Word2)
}

func Solve(tiles Tiles) (solution Solution) {
	var blacklist []string

	for {
		word1, remainder := highestWord(tiles, blacklist)
		word2, _ := highestWord(remainder, blacklist)

		if word1.Score == 0 {
			// Oh no
			log.Fatal("I could literally not make any words with these tiles.")
		}

		if word2.Score == 0 {
			// Our first word is too good! Blacklist it and try to find a slightly worse word.
			blacklist = append(blacklist, word1.Word.String())
		} else {
			return Solution{word1, word2}
		}
	}
}

// Find the highest scoring word we can make with our tiles
// Return the leftover ones.
func highestWord(tiles Tiles, blacklist []string) (highScore WordAndScore, remaining Tiles) {
OUTER:
	for _, word := range dictionary {
		for _, badWord := range blacklist {
			if word == badWord {
				continue OUTER
			}
		}

		if used, remainder, ok := tiles.Contains(word); ok {
			if score := used.Score(); score > highScore.Score {
				highScore = WordAndScore{used, score}
				remaining = remainder
			}
		}
	}

	return highScore, remaining
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
	tiles := makeTiles(args)

	loadWords()
	solution := Solve(tiles)
	fmt.Println(solution)
}
