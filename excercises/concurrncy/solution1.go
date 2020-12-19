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

func readChannel(c chan FreqMap, chanel_length int) FreqMap {
	result := FreqMap{}
	for values := range c {
		for k, v := range values {
			result[k] = result[k] + v
		}
	}
	return result
}

func ConcurrentFrequency(s []string) FreqMap {
	c := make(chan FreqMap, len(s))
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

	return readChannel(c, len(s))
}

func main() {
	input := []string{"Maybe don't is a nic song"}
	ConcurrentFrequency(input)

}
