package letter

import (
	"sync"
)

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}

func ConcurrentFrequency(s []string) FreqMap {
	m := FreqMap{}
	mut := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, batch := range s {
		wg.Add(1)
		go func(batch string) {
			batchResults := Frequency(batch)

			mut.Lock()
			for k,v := range batchResults {
				m[k] = m[k] + v
			}
			mut.Unlock()
			wg.Done()
		}(batch)
	}
	wg.Wait()
	return m
}
