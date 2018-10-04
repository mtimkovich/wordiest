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

func sortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

// Read valid words into global trie
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
		signature := sortString(word)
		wordTrie.Add(signature, word)
	}
}

func main() {
	loadWords()
	// argv := "t e g5w o o l2l t i p o2l u d a n2w"
	argv := "t e g5w"

	tiles := makeTiles(argv)
	for _, tile := range tiles {
		fmt.Println(tile)
	}
}
