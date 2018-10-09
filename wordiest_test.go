package main

import (
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	loadWords()
	retCode := m.Run()
	os.Exit(retCode)
}

func getSolution(input string) *Solution {
	tiles := makeTiles(strings.Fields(input))
	return Solve(tiles)
}

func TestMultiplierPriority(t *testing.T) {
	solution := getSolution("c a i o o g w5w e u r2l r i2l u2w l")

	if solution.Total() < 100 {
		t.Error("Solution score too low:", solution)
	}
}

func TestSecondWord(t *testing.T) {
	solution := getSolution("e4w v e t5l n q i d2l f u i w s3w t")

	if solution.Word2.Score == 0 {
		t.Error("No word for word2.")
	}
}

func TestCloseWords(t *testing.T) {
	solution := getSolution("n e2l z t s a s n l t3l e e e f")

	if solution.Total() < 30 {
		t.Error("Did not find optimal solution:", solution)
	}
}
