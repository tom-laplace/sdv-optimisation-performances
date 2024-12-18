package main

import (
	"fmt"
	"sync"
)

func main() {
	listes := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	var wg sync.WaitGroup
	ch := make(chan int, len(listes))

	for i := range listes {
		wg.Add(1)
		go sum(listes[i], &wg, ch)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	total := 0
	for sum := range ch {
		total += sum
	}

	fmt.Printf("%v", listes)
	fmt.Println()
	fmt.Println("Toutes les sommes ont été calculées :", total)
}

func sum(slice []int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()

	sum := 0
	for _, val := range slice {
		sum += val
	}

	ch <- sum
}
