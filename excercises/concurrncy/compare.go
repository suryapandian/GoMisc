package main

import (
	"fmt"
	"sync"
)

type FreqMap map[rune]int

var (
	euro = `Freude schöner Götterfunken
Tochter aus Elysium,
Wir betreten feuertrunken,
Himmlische, dein Heiligtum!
Deine Zauber binden wieder
Was die Mode streng geteilt;
Alle Menschen werden Brüder,
Wo dein sanfter Flügel weilt.`

	dutch = `Wilhelmus van Nassouwe
ben ik, van Duitsen bloed,
den vaderland getrouwe
blijf ik tot in den dood.
Een Prinse van Oranje
ben ik, vrij, onverveerd,
den Koning van Hispanje
heb ik altijd geëerd.`

	us = `O say can you see by the dawn's early light,
What so proudly we hailed at the twilight's last gleaming,
Whose broad stripes and bright stars through the perilous fight,
O'er the ramparts we watched, were so gallantly streaming?
And the rockets' red glare, the bombs bursting in air,
Gave proof through the night that our flag was still there;
O say does that star-spangled banner yet wave,
O'er the land of the free and the home of the brave?`
)

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

func ConcurrentFrequencyLoop(s []string) FreqMap {
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

func ConcurrentFrequencyRoutine(s []string) FreqMap {

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
	input := []string{euro, dutch, us}
	routineAns := ConcurrentFrequencyRoutine(input)
	loopAns := ConcurrentFrequencyLoop(input)
	for k, v := range routineAns {
		if loopAns[k] != v {
			fmt.Println("incorrct")
		}

	}

}
