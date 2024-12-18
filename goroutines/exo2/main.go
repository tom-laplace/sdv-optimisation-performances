package main

import (
	"fmt"
	"sync"
)

type SearchResult struct {
	SegmentIndex int
	LocalIndex   int
	GlobalIndex  int
	Found        bool
}

func main() {
	tableau := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 49}
	elementCible := 49
	tailleSegment := 4

	var wg sync.WaitGroup
	numOfGoroutines := len(tableau) / tailleSegment

	ch := make(chan SearchResult, numOfGoroutines)

	for i := 0; i < numOfGoroutines; i++ {
		wg.Add(1)
		debut := i * tailleSegment
		fin := debut + tailleSegment
		segment := tableau[debut:fin]

		go search(segment, elementCible, i, debut, &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	found := false
	for result := range ch {
		if result.Found {
			found = true
			fmt.Printf("Élément %d trouvé:\n", elementCible)
			fmt.Printf("- Segment: %d\n", result.SegmentIndex)
			fmt.Printf("- Index local dans le segment: %d\n", result.LocalIndex)
			fmt.Printf("- Index global dans le tableau: %d\n", result.GlobalIndex)
		}
	}

	if !found {
		fmt.Printf("Élément %d non trouvé dans le tableau\n", elementCible)
	}
}

func search(segment []int, elementCible int, segmentIndex int, baseIndex int, wg *sync.WaitGroup, ch chan SearchResult) {
	defer wg.Done()

	for i, val := range segment {
		if val == elementCible {
			ch <- SearchResult{
				SegmentIndex: segmentIndex,
				LocalIndex:   i,
				GlobalIndex:  baseIndex + i,
				Found:        true,
			}
			return
		}
	}

	ch <- SearchResult{Found: false}
}
