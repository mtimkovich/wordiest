package main

import (
	"container/heap"
	"fmt"
)

type WordAndScoreHeap []*WordAndScore

func (h WordAndScoreHeap) Len() int           { return len(h) }
func (h WordAndScoreHeap) Less(i, j int) bool { return h[i].Score < h[j].Score }
func (h WordAndScoreHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *WordAndScoreHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*WordAndScore))
}

func (h *WordAndScoreHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type WordAndScore struct {
	Word      Tiles
	Remaining Tiles
	Score     int
}

func (w WordAndScore) String() string {
	return fmt.Sprintf("%v (%v)", w.Word, w.Score)
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
	solution := &Solution{}
	const ATTEMPTS = 50

	words := highestWords(tiles, ATTEMPTS)
	queue := make(chan *Solution, words.Len())

	for _, word := range words {
		go func(word *WordAndScore) {
			second := highestWords(word.Remaining, 1)

			if second.Len() > 0 {
				queue <- &Solution{*word, *second[0]}
			} else {
				queue <- nil
			}
		}(word)
	}

	for i := 0; i < words.Len(); i++ {
		newSolution := <-queue
		if newSolution != nil && newSolution.Total() > solution.Total() {
			solution = newSolution
		}
	}

	return solution
}

// Get the highest scoring |attempt| amount of words.
func highestWords(tiles Tiles, attempts int) WordAndScoreHeap {
	pq := make(WordAndScoreHeap, 0)
	heap.Init(&pq)

	for _, word := range dictionary {
		if used, remaining, ok := tiles.Contains(word); ok {
			score := used.Score()
			if pq.Len() == 0 || score > pq[0].Score {
				heap.Push(&pq, &WordAndScore{used, remaining, score})
			}

			if pq.Len() > attempts {
				heap.Pop(&pq)
			}
		}
	}

	return pq
}
