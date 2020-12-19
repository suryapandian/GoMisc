package main

import (
	"sync"
)

type FreqMap map[rune]int

func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {

	c := make(chan FreqMap, len(s))
	result := FreqMap{}
	go func() {
		for fm := range c {
			for k, v := range fm {
				result[k] = result[k] + v
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(len(s))
	for _, w := range s {
		go func(s string) {
			defer wg.Done()
			c <- Frequency(s)
		}(w)
	}
	wg.Wait()
	close(c)
	return result
}

func main() {
	input := []string{"Maybe don't is a nic song"}
	ConcurrentFrequency(input)

}
