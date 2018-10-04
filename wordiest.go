package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/derekparker/trie"
)

var wordTrie *trie.Trie

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

// Return the letters of a string in sorted order.
func sorted(s string) string {
	runes := []rune(s)
	sort.Sort(RuneSlice(runes))
	return string(runes)
}

// Read all English words into the global trie.
func loadWords() {
	file, err := os.Open("sowpods.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	wordTrie = trie.New()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		signature := sorted(word)
		// The trie stores the word at the leaf of its sorted letters.
		wordTrie.Add(signature, word)
	}
}

func main() {
	// if len(os.Args) < 2 {
	// 	fmt.Println("Usage:", os.Args[0], "[a] [a2l] [a3w]")
	// 	fmt.Println("  - a: [a-z]")
	// 	fmt.Println("  - [1-9]: multiplier")
	// 	fmt.Println("  - l: letter multiplier")
	// 	fmt.Println("  - w: word multiplier")
	// 	os.Exit(0)
	// }

	// args := os.Args[1:]
	argsStr := "t e g5w o o l2l t i p o2l u d a n2w"
	args := strings.Split(argsStr, " ")
	fmt.Println(args)

	tiles := makeTiles(args)

	sort.Slice(tiles, func(i, j int) bool {
		if tiles[i].wordMul == tiles[j].wordMul {
			return tiles[i].TileScore() > tiles[j].TileScore()
		}

		return tiles[i].wordMul > tiles[j].wordMul
	})

	for _, tile := range tiles {
		fmt.Println(tile)
	}

	// loadWords()
}
