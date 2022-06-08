package utils

import (
	"sync"
)

const (
	DefaultChunkSize = 250
)

func Chunk(data []string, size int) [][]string {
	ln := len(data)

	if ln == 0 {
		return nil
	}

	chunks := make([][]string, (ln+size-1)/size)
	prev := 0
	i := 0
	till := ln - size

	for prev < till {
		next := prev + size
		chunks[i] = data[prev:next]
		prev = next
		i++
	}
	chunks[i] = data[prev:]

	return chunks
}

type T int

func X() {
	var slice []T
	var wg sync.WaitGroup

	queue := make(chan T, 1)

	// Create our data and send it into the queue.
	wg.Add(150)
	for i := 0; i < 150; i++ {
		go func(i int) {
			// defer wg.Done()  <- will result in the last int to be missed in the receiving channel
			queue <- T(i)
		}(i)
	}

	go func() {
		// defer wg.Done() <- Never gets called since the 100 `Done()` calls are made above, resulting in the `Wait()` to continue on before this is executed
		for t := range queue {
			slice = append(slice, t)
			wg.Done() // ** move the `Done()` call here
		}
	}()

	wg.Wait()

	// now prints off all 10000 int values
}
