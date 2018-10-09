package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	loadWords()
	retCode := m.Run()
	os.Exit(retCode)
}

func TestMultiplierPriority(t *testing.T) {
	tiles := makeTiles([]string{"g4w", "i", "b2l", "d", "q", "n", "e", "a", "o", "n4w", "i2l", "a", "l3l", "v2w"})
	solution := Solve(tiles)

	if solution.Total() < 950 {
		t.Error("Solution score too low:", solution)
	}
}

func TestSecondWord(t *testing.T) {
	tiles := makeTiles([]string{"e4w", "v", "e", "t5l", "n", "q", "i", "d2l", "f", "u", "i", "w", "s3w", "t"})
	solution := Solve(tiles)

	if solution.Word2.Score == 0 {
		t.Error("No word for word2.")
	}
}

func TestCloseWords(t *testing.T) {
	tiles := makeTiles([]string{"n", "e2l", "z", "t", "s", "a", "s", "n", "l", "t3l", "e", "e", "e", "f"})
	solution := Solve(tiles)

	if solution.Total() < 30 {
		t.Error("Did not find optimal solution:", solution)
	}
}
