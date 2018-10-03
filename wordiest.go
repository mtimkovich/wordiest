package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var dictionary []string

type Tile struct {
	letter    string
	letterVal int // how much this letter is worth
	letterMul int // letter multiplier
	wordMul   int // word multiplier
}

func letterValue(letter string) int {
	values := map[string]int{
		"d": 2, "g": 2,
		"c": 3, "m": 3, "b": 3, "p": 3,
		"h": 4, "f": 4, "w": 4, "y": 4, "v": 4,
		"k": 5,
		"j": 8, "x": 8,
		"q": 10, "z": 10,
	}

	if val, ok := values[letter]; ok {
		return val
	} else {
		return 1
	}
}

func NewTile(input string) *Tile {
	t := &Tile{}
	t.letterMul = 1
	t.wordMul = 1

	// Parse
	if len(input) == 1 {
		t.letter = input
	} else if len(input) == 3 {
		// input[0] is letter
		// input[1] is multiplier value
		// input[2] is type of multiplier
		t.letter = string(input[0])
		mul, err := strconv.Atoi(string(input[1]))
		if err != nil {
			log.Fatal(err)
		}

		if input[2] == 'l' {
			t.letterMul = mul
		} else if input[2] == 'w' {
			t.wordMul = mul
		} else {
			log.Fatal("Invalid input: " + input)
		}
	} else {
		log.Fatal("Invalid input: " + input)
	}

	t.letterVal = letterValue(t.letter)
	return t
}

func makeTiles(input string) []*Tile {
	var tiles []*Tile
	split := strings.Split(input, " ")
	for _, s := range split {
		tiles = append(tiles, NewTile(s))
	}

	return tiles
}

// Read valid words into global dictionary
func loadDictionary() {
	file, err := os.Open("sowpods.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		dictionary = append(dictionary, word)
	}
}

func main() {
	// loadDictionary()
	argv := "t e g5w o o l2l t i p o2l u d a n2w"

	tiles := makeTiles(argv)
	for _, tile := range tiles {
		fmt.Println(tile)
	}
}
