package itertools

import (
	"sync"
)

// Product generate cartesian product of input iterables.
func Product(params ...[]int) <-chan []int {
	c := make(chan []int)
	var wg sync.WaitGroup
	wg.Add(1)

	iterate(&wg, c, []int{}, params...)
	go func() { wg.Wait(); close(c) }()

	return c
}

func iterate(wg *sync.WaitGroup, channel chan<- []int, result []int, params ...[]int) {
	defer wg.Done()
	if len(params) == 0 {
		channel <- result
		return
	}
	p, params := params[0], params[1:]
	for i := 0; i < len(p); i++ {
		wg.Add(1)
		go iterate(wg, channel, append(result, p[i]), params...)
	}
}

// Combinations return l length subsequences of elements from the input alphabet.
func Combinations(alphabet []interface{}, l int) <-chan []interface{} {
	c := make(chan []interface{})
	var wg sync.WaitGroup

	wg.Add(1)

	if l == 0 {
		close(c)
		return c
	}

	comb(&wg, c, []interface{}{}, alphabet, l)
	go func() { wg.Wait(); close(c) }()

	return c
}

func comb(wg *sync.WaitGroup, c chan<- []interface{}, result []interface{}, alphabet []interface{}, l int) {
	defer wg.Done()

	if l == 0 {
		c <- result
		return
	}

	for i := 0; i < len(alphabet); i++ {
		wg.Add(1)
		go comb(wg, c, append(result, alphabet[i]), alphabet[i+1:], l-1)
	}
}
