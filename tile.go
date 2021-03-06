package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Tile struct {
	Letter    byte
	LetterVal int // how much this letter is worth
	LetterMul int // letter multiplier
	WordMul   int // word multiplier
}

func (t *Tile) TileScore() int {
	return t.LetterVal * t.LetterMul
}

func (t *Tile) String() string {
	if t.LetterMul != 1 {
		return fmt.Sprintf("%c%v%v", t.Letter, t.LetterMul, "l")
	} else if t.WordMul != 1 {
		return fmt.Sprintf("%c%v%v", t.Letter, t.WordMul, "w")
	} else {
		return fmt.Sprintf("%c", t.Letter)
	}
}

func letterValue(letter byte) int {
	values := map[byte]int{
		'd': 2, 'n': 2, 'u': 2, 'l': 2,
		'g': 3, 'c': 3, 'p': 3, 'h': 3,
		'm': 4, 'b': 4, 'f': 4, 'w': 4, 'y': 4,
		'k': 5,
		'v': 6,
		'x': 8,
		'j': 10, 'q': 10, 'z': 10,
	}

	if val, ok := values[letter]; ok {
		return val
	} else {
		return 1
	}
}

func NewTile(input string) (*Tile, error) {
	t := &Tile{}
	t.LetterMul = 1
	t.WordMul = 1

	tileErr := fmt.Errorf("Invalid tile: %v", input)

	input = strings.ToLower(input)

	// Parse
	if len(input) == 1 {
		t.Letter = input[0]
	} else if len(input) == 3 {
		// input[0] is letter
		// input[1] is multiplier value
		// input[2] is type of multiplier
		t.Letter = input[0]
		mul, err := strconv.Atoi(string(input[1]))
		if err != nil {
			return nil, err
		}

		if input[2] == 'l' {
			t.LetterMul = mul
		} else if input[2] == 'w' {
			t.WordMul = mul
		} else {
			return nil, tileErr
		}
	} else {
		return nil, tileErr
	}

	t.LetterVal = letterValue(t.Letter)
	return t, nil
}

type Tiles []*Tile

// See if we can construct the given word with our tiles. Return the used tiles, the
// leftover tiles, and if we had a match.
func (tiles Tiles) Contains(word string) (used, remaining Tiles, match bool) {
	if len(word) > len(tiles) {
		return
	}

	remaining = append(remaining, tiles...)

	for _, c := range word {
		letterMatch := false
		for i, t := range remaining {
			if t.Letter == byte(c) {
				used = append(used, remaining[i])
				remaining = append(remaining[:i], remaining[i+1:]...)
				letterMatch = true
				break
			}
		}

		if !letterMatch {
			return
		}
	}

	match = true
	return
}

func (t Tiles) Score() int {
	score := 0
	mul := 1

	for _, tile := range t {
		score += tile.TileScore()
		mul *= tile.WordMul
	}

	return score * mul
}

func (t Tiles) String() (output string) {
	for _, tile := range t {
		output += fmt.Sprintf("%c", tile.Letter)
	}

	return
}

func (t Tiles) Debug() (output string) {
	for _, tile := range t {
		output += tile.String() + " "
	}

	return
}

func makeTiles(inputs []string) Tiles {
	var tiles Tiles

	for _, s := range inputs {
		tile, err := NewTile(s)
		if err != nil {
			fatal(err)
		}

		tiles = append(tiles, tile)
	}

	if len(tiles) != 14 {
		fatalf("Expected 14 tiles, got %v.\n", len(tiles))
	}

	// Sort tiles so more powerful tiles are sorted first.
	sort.Slice(tiles, func(i, j int) bool {
		if tiles[i].WordMul > tiles[j].WordMul {
			return true
		} else if tiles[i].WordMul < tiles[j].WordMul {
			return false
		} else {
			return tiles[i].LetterMul > tiles[j].LetterMul
		}
	})

	return tiles
}
