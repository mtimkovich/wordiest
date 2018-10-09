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
	file, err := os.Open("TWL06.txt")
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

func (s *Solution) Total() int {
	return s.Word1.Score + s.Word2.Score
}

func (s *Solution) String() string {
	return fmt.Sprintf("%v: %v + %v", s.Total(), s.Word1, s.Word2)
}

func Solve(tiles Tiles) *Solution {
	blacklist := map[string]bool{}
	solution := &Solution{}
	attempts := 0

	for {
		word1, remainder := highestWord(tiles, blacklist)
		word2, _ := highestWord(remainder, blacklist)

		if word1.Score == 0 {
			// Oh no
			log.Fatal("I could literally not make any words with these tiles.")
		}

		if word2.Score == 0 {
			// Our first word is too good! Blacklist it and try to find a slightly worse word.
			blacklist[word1.Word.String()] = true
		} else {
			newSolution := &Solution{word1, word2}

			if newSolution.Total() > solution.Total() {
				solution = newSolution
			}

			// Sometimes if the words are close in value, it's better to have 2 solid words than
			// 1 really good word. Try 20 more times to see if we can find something better.
			if word1.Score <= word2.Score+40 && attempts < 20 {
				blacklist[word1.Word.String()] = true
				attempts++
			} else {
				return solution
			}
		}
	}
}

// Find the highest scoring word we can make with our tiles
// Return the leftover tiles. There is also a blacklist of words to ignore for various scoring
// reasons.
func highestWord(tiles Tiles, blacklist map[string]bool) (highScore WordAndScore, remaining Tiles) {
	for _, word := range dictionary {
		if _, ok := blacklist[word]; ok {
			continue
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
		fmt.Println("Usage:", os.Args[0], "[a a2l a3w]+")
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
