package main

import (
	"fmt"
	"sync"
)

func worker(a int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- a * a
}

func closers(ch chan int, wg *sync.WaitGroup) {
	wg.Wait()
	close(ch)
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var result int
	var wg sync.WaitGroup
	ch := make(chan int, len(numbers))

	for _, a := range numbers {
		wg.Add(1)
		go worker(a, ch, &wg)
	}

	go closers(ch, &wg)

	for a := range ch {
		result += a
	}

	fmt.Println(result)
}
