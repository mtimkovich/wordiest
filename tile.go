package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Tile struct {
	letter    string
	letterVal int // how much this letter is worth
	letterMul int // letter multiplier
	wordMul   int // word multiplier
}

func (t *Tile) Print() {
	if t.letterMul != 1 {
		fmt.Printf("%v%v%v\n", t.letter, t.letterMul, "l")
	} else if t.wordMul != 1 {
		fmt.Printf("%v%v%v\n", t.letter, t.wordMul, "w")
	} else {
		fmt.Println(t.letter)
	}
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

func NewTile(input string) (*Tile, error) {
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
			return nil, err
		}

		if input[2] == 'l' {
			t.letterMul = mul
		} else if input[2] == 'w' {
			t.wordMul = mul
		} else {
			return nil, fmt.Errorf("Invalid input: %v", input)
		}
	} else {
		return nil, fmt.Errorf("Invalid input: %v", input)
	}

	t.letterVal = letterValue(t.letter)
	return t, nil
}

func makeTiles(input string) []*Tile {
	var tiles []*Tile
	split := strings.Split(input, " ")
	for _, s := range split {
		tile, err := NewTile(s)
		if err != nil {
			log.Fatal(err)
		}
		tiles = append(tiles, tile)
	}

	return tiles
}
