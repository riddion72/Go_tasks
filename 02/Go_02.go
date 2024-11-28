package main

import (
	"fmt"
)

type multiValue struct {
	id    int
	value int
}

func worker(id int, a int, ch chan multiValue) {
	ch <- multiValue{id, a * a}
}

func main() {
	var array [5]int = [5]int{2, 4, 6, 8, 10}
	var resArray [5]int
	ch := make(chan multiValue, len(array))

	for i, a := range array {
		go worker(i, a, ch)
	}

	for i := 0; i < len(array); i++ {
		result := <-ch
		resArray[result.id] = result.value
	}

	fmt.Println(resArray)
}
